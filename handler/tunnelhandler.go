package samtunnelhandler

import (
	"fmt"
	"net/http"
	"strings"
)

import (
	"github.com/eyedeekay/sam-forwarder/interface"
)

type TunnelHandler struct {
	samtunnel.SAMTunnel
}

func (t *TunnelHandler) Printdivf(id, key, value string, rw http.ResponseWriter, req *http.Request) {
	if key == "" || value == "" {
		return
	}
	ID := t.SAMTunnel.ID()
	if id != "" {
		ID = t.SAMTunnel.ID() + "." + id
	}
	prop := ""
	if key != "TunName" {
		prop = "prop"
	}
	if strings.HasSuffix(req.URL.Path, "color") {
		fmt.Fprintf(rw, "    <div id=\"%s\" class=\"%s %s %s %s\" >\n", ID, t.SAMTunnel.ID(), key, t.SAMTunnel.GetType(), prop)
		fmt.Fprintf(rw, "      <span id=\"%s\" class=\"key\">%s</span>=", ID, key)
		fmt.Fprintf(rw, "      <span id=\"%s\" class=\"value\">%s</span>\n", ID, value)
		fmt.Fprintf(rw, "    </div>\n\n")
	} else {
		fmt.Fprintf(rw, "%s=%s\n", ID, t.SAMTunnel.ID())
	}
}

func (t *TunnelHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	if strings.HasSuffix(req.URL.Path, "color") {
		fmt.Fprintf(rw, "  <div id=\"%s\" class=\"%s\" >", t.SAMTunnel.ID(), t.SAMTunnel.GetType())
	}
	t.Printdivf(t.SAMTunnel.ID(), "TunName", t.SAMTunnel.ID(), rw, req)
	fmt.Fprintf(rw, "  <span id=\"toggle%s\" class=\"control\">\n", t.SAMTunnel.ID())
	fmt.Fprintf(rw, "    <a href=\"#\" onclick=\"toggle_visibility_class('%s');\"> Show/Hide %s</a>\n", t.SAMTunnel.ID(), t.SAMTunnel.ID())
	fmt.Fprintf(rw, "    <a href=\"/%s/color\">Tunnel page</a>\n", t.SAMTunnel.ID())
	fmt.Fprintf(rw, "  </span>\n")
	for key, value := range t.SAMTunnel.Props() {
		if key != "TunName" {
			t.Printdivf(key, key, value, rw, req)
		}
	}
	if strings.HasSuffix(req.URL.Path, "color") {
		fmt.Fprintf(rw, "  </div>\n\n")
	}
	if strings.HasSuffix(req.URL.Path, "color") {
		fmt.Fprintf(rw, "  <div id=\"%s\" class=\"%s control panel\" >", t.SAMTunnel.ID()+".control", t.SAMTunnel.GetType())
	}
	fmt.Fprintf(rw, "    <a href=\"/%s/start\">.[START]</a>", t.SAMTunnel.ID())
	fmt.Fprintf(rw, "    <a href=\"/%s/stop\">.[STOP].</a>", t.SAMTunnel.ID())
	fmt.Fprintf(rw, "    <a href=\"/%s/restart\">[RESTART].</a><br>", t.SAMTunnel.ID())
	fmt.Fprintf(rw, "    <div id=\"%s.status\" class=\"%s status\">.[STATUS].</div>", t.SAMTunnel.ID(), t.SAMTunnel.ID())
	if strings.HasSuffix(req.URL.Path, "color") {
		fmt.Fprintf(rw, "  </div>\n\n")
	}
}

func NewTunnelHandler(ob samtunnel.SAMTunnel, err error) (*TunnelHandler, error) {
	var t TunnelHandler
	t.SAMTunnel = ob
	return &t, err
}
