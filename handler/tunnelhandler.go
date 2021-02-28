package samtunnelhandler

import (
	"fmt"
	"net/http"
	"sort"
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
	fmt.Fprintf(rw, t.ColorWrap(ID, t.SAMTunnel.ID(), key, t.SAMTunnel.GetType(), prop, value))
}

func (t *TunnelHandler) Printf(id, key, value string, rw http.ResponseWriter, req *http.Request) {
	if key == "" || value == "" {
		return
	}
	ID := t.SAMTunnel.ID()
	if id != "" {
		ID = t.SAMTunnel.ID() + "." + id
	}
	fmt.Fprintf(rw, "%s=%s\n", ID, t.SAMTunnel.ID())
}

func PropSort(props map[string]string) []string {
	var slice []string
	for k, v := range props {
		slice = append(slice, k+"="+v)
	}
	sort.Strings(slice)
	return slice
}

func (t *TunnelHandler) ControlForm(rw http.ResponseWriter, req *http.Request) {
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
}

func (t *TunnelHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	t.ControlForm(rw, req)
	if strings.HasSuffix(req.URL.Path, "color") {
		fmt.Fprintf(rw, t.ColorDiv(t.SAMTunnel.ID(), t.SAMTunnel.GetType()))
		t.Printdivf(t.SAMTunnel.ID(), "TunName", t.SAMTunnel.ID(), rw, req)
		fmt.Fprintf(rw, t.ColorSpan(t.SAMTunnel.ID()))
		for _, value := range PropSort(t.SAMTunnel.Props()) {
			key := strings.SplitN(value, "=", 2)[0]
			val := strings.SplitN(value, "=", 2)[1]
			if key != "TunName" {
				t.Printdivf(key, key, val, rw, req)
			}
		}
		fmt.Fprintf(rw, t.ColorForm(t.SAMTunnel.ID(), t.SAMTunnel.GetType()))
	} else {
		t.Printf(t.SAMTunnel.ID(), "TunName", t.SAMTunnel.ID(), rw, req)
		for _, value := range PropSort(t.SAMTunnel.Props()) {
			key := strings.SplitN(value, "=", 2)[0]
			val := strings.SplitN(value, "=", 2)[1]
			if key != "TunName" {
				t.Printf(key, key, val, rw, req)
			}
		}
	}
}

func NewTunnelHandler(ob samtunnel.SAMTunnel, err error) (*TunnelHandler, error) {
	var t TunnelHandler
	t.SAMTunnel = ob
	return &t, err
}
