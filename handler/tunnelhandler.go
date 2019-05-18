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
	if err := req.ParseForm(); err == nil {
		if action := req.PostFormValue("action"); action != "" {
			var err error
			switch action {
			case "start":
				if !t.SAMTunnel.Up() {
					fmt.Println("Starting tunnel", t.ID())
					if t.SAMTunnel, err = t.Load(); err == nil {
						t.Serve()
					}
					//return
				} else {
					fmt.Println(t.ID(), "already started")
					req.URL.Path = req.URL.Path + "/color"
				}
			case "stop":
				if t.SAMTunnel.Up() {
					fmt.Println("Stopping tunnel", t.ID())
					t.Close()
				} else {
					fmt.Println(t.ID(), "already stopped")
					req.URL.Path = req.URL.Path + "/color"
				}
			case "restart":
				if t.SAMTunnel.Up() {
					fmt.Println("Stopping tunnel", t.ID())
					t.Close()
					fmt.Println("Starting tunnel", t.ID())
					if t.SAMTunnel, err = t.Load(); err == nil {
						t.Serve()
					}
					return
				} else {
					fmt.Println(t.ID(), "stopped.")
					req.URL.Path = req.URL.Path + "/color"
				}
			default:
			}
		}
	}

	if strings.HasSuffix(req.URL.Path, "color") {
		fmt.Fprintf(rw, "  <div id=\"%s\" class=\"%s\" >", t.SAMTunnel.ID(), t.SAMTunnel.GetType())
	}

	t.Printdivf(t.SAMTunnel.ID(), "TunName", t.SAMTunnel.ID(), rw, req)
	fmt.Fprintf(rw, "  <span id=\"toggle%s\" class=\"control\">\n", t.SAMTunnel.ID())
	fmt.Fprintf(rw, "    <a href=\"#\" onclick=\"toggle_visibility_class('%s');\"> Show/Hide %s</a><br>\n", t.SAMTunnel.ID(), t.SAMTunnel.ID())
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
	fmt.Fprintf(rw, "    <form class=\"linkstyle\" name=\"start\" action=\"/%s\" method=\"post\">", t.SAMTunnel.ID())
	fmt.Fprintf(rw, "      <input class=\"linkstyle\" type=\"hidden\" value=\"start\" name=\"action\" />")
	fmt.Fprintf(rw, "      <input class=\"linkstyle\" type=\"submit\" value=\".[START]\">")
	fmt.Fprintf(rw, "    </form>")

	fmt.Fprintf(rw, "    <form class=\"linkstyle\" name=\"stop\" action=\"/%s\" method=\"post\">", t.SAMTunnel.ID())
	fmt.Fprintf(rw, "      <input class=\"linkstyle\" type=\"hidden\" value=\"stop\" name=\"action\" />")
	fmt.Fprintf(rw, "      <input class=\"linkstyle\" type=\"submit\" value=\".[STOP].\">")
	fmt.Fprintf(rw, "    </form>")

	fmt.Fprintf(rw, "    <form class=\"linkstyle\" name=\"restart\" action=\"/%s\" method=\"post\">", t.SAMTunnel.ID())
	fmt.Fprintf(rw, "      <input class=\"linkstyle\" type=\"hidden\" value=\"restart\" name=\"action\" />")
	fmt.Fprintf(rw, "      <input class=\"linkstyle\" type=\"submit\" value=\"[RESTART].\">")
	fmt.Fprintf(rw, "    </form>")

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
