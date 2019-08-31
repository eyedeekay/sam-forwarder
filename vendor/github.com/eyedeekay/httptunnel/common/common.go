package proxycommon

import (
	"io"
	"log"
	"net/http"
)

var hopHeaders = []string{
	"Accept-Language",
	"Connection",
	"Keep-Alive",
	//	"Proxy-Authenticate",
	//	"Proxy-Authorization",
	"Proxy-Connection",
	"Trailers",
	"Upgrade",
	"X-Forwarded-For",
	"X-Real-IP",
}

func Transfer(destination io.WriteCloser, source io.ReadCloser) {
	defer destination.Close()
	defer source.Close()
	log.Println("connecting connection")
	io.Copy(destination, source)
}

func DelHopHeaders(header http.Header) {
	for _, h := range hopHeaders {
		header.Del(h)
	}
	if header.Get("User-Agent") != "MYOB/6.66 (AN/ON)" {
		header.Set("User-Agent", "MYOB/6.66 (AN/ON)")
	}
}

func CopyHeader(dst, src http.Header) {
	for k, vv := range src {
		for _, v := range vv {
			dst.Add(k, v)
		}
	}
}
