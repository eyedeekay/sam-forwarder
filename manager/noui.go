//go:build cli && (!nostatic || !static)
// +build cli
// +build !nostatic !static

package sammanager

import (
	"log"
	"os"
	"os/signal"
	"time"
)

func (s *SAMManager) RunUI() {
}

func (s *SAMManager) Serve() bool {
	log.Println("Starting Tunnels()")
	for _, element := range s.handlerMux.Tunnels() {
		log.Println("Starting service tunnel", element.ID())
		go element.Serve()
	}

	return Exit()
}

func Exit() bool {
	Close := false
	for !Close {
		time.Sleep(1 * time.Second)
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		go func() {
			for sig := range c {
				log.Println(sig)

				Close = true
			}
		}()
	}
	return false
}
