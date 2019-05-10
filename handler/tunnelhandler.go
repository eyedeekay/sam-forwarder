package samtunnelhandler

import (
	"fmt"
	"net/http"
)

import (
	"github.com/eyedeekay/sam-forwarder/interface"
)

type TunnelHandler struct {
	samtunnel.SAMTunnel
}

func (t *TunnelHandler) Printdivf(id, key, value string, rw http.ResponseWriter, req *http.Request) {
	if req.FormValue("color") == "true" {
		fmt.Fprintf(rw, "    <div id=\"%s\" class=\"%s\" >\n", t.SAMTunnel.ID()+"."+id, t.SAMTunnel.ID())
		fmt.Fprintf(rw, "      <div id=\"%s\" class=\"key\">%s</div>=", t.SAMTunnel.ID()+"."+id, key)
		fmt.Fprintf(rw, "      <div id=\"%s\" class=\"value\">%s</div>\n", t.SAMTunnel.ID()+"."+id, value)
		fmt.Fprintf(rw, "    </div>\n")
	} else {
		fmt.Fprintf(rw, "%s=%s", t.SAMTunnel.ID()+"."+id, t.SAMTunnel.ID())
	}
}

func (t *TunnelHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	if req.FormValue("color") == "true" {
		fmt.Fprintf(rw, "  <div id=\"%s\" class=\"%s\" >", t.SAMTunnel.ID(), t.SAMTunnel.GetType())
	}
	t.Printdivf(t.SAMTunnel.ID(), "TunName", t.SAMTunnel.ID(), rw, req)
	for key, value := range t.SAMTunnel.Props() {
		t.Printdivf(key, key, value, rw, req)
	}
	if req.FormValue("color") == "true" {
		fmt.Fprintf(rw, "  </div>")
	}
}

func NewTunnelHandler(ob samtunnel.SAMTunnel, err error) (*TunnelHandler, error) {
	var t TunnelHandler
	t.SAMTunnel = ob
	return &t, err
}
