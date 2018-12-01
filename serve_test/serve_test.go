package samforwardertest

import (
    "testing"
    "time"
    "net/http"
    "log"
)

func TestTCP(t *testing.T) {
    go serve()
    time.Sleep(time.Duration(60 * time.Second))
    go client()
    time.Sleep(time.Duration(60 * time.Second))
    resp, err := http.Get("http://127.0.0.1:"+cport+"/test.html")
    //defer forwarder.Close()
    //defer forwarderclient.Close()
    if err != nil {
        t.Fatal(err)
    }
    log.Println(resp)
}

func TestUDP(t *testing.T){

}
