package samforwarder

import (
	"io"
	"log"
	"net"
)

/*
Adapted from: https://gist.github.com/jbardin/821d08cb64c01c84b81a
*/

func Proxy(srvConn, cliConn net.Conn) {
	// channels to wait on the close event for each connection
	serverClosed := make(chan struct{}, 1)
	clientClosed := make(chan struct{}, 1)

	go broker(srvConn, cliConn, clientClosed)
	go broker(cliConn, srvConn, serverClosed)

	// wait for one half of the proxy to exit, then trigger a shutdown of the
	// other half by calling CloseRead(). This will break the read loop in the
	// broker and allow us to fully close the connection cleanly without a
	// "use of closed network connection" error.
	var waitFor chan struct{}
	select {
	case <-clientClosed:
		// the client closed first and any more packets from the server aren't
		// useful, so we can optionally SetLinger(0) here to recycle the port
		// faster.
		//		srvConn.SetLinger(0)
		//		srvConn.CloseRead()
		waitFor = serverClosed
	case <-serverClosed:
		//		cliConn.CloseRead()
		waitFor = clientClosed
	}

	// Wait for the other connection to close.
	// This "waitFor" pattern isn't required, but gives us a way to track the
	// connection and ensure all copies terminate correctly; we can trigger
	// stats on entry and deferred exit of this function.
	<-waitFor
}

// This does the actual data transfer.
// The broker only closes the Read side.
func broker(dst, src net.Conn, srcClosed chan struct{}) {
	// We can handle errors in a finer-grained manner by inlining io.Copy (it's
	// simple, and we drop the ReaderFrom or WriterTo checks for
	// net.Conn->net.Conn transfers, which aren't needed). This would also let
	// us adjust buffersize.
	_, err := io.Copy(dst, src)

	if err != nil {
		log.Printf("Copy error: %s", err)
	}
	if err := src.Close(); err != nil {
		log.Printf("Close error: %s", err)
	}
	srcClosed <- struct{}{}
}
