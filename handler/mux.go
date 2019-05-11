package samtunnelhandler

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type TunnelHandlerMux struct {
	http.Server
	pagenames    []string
	tunnels      []*TunnelHandler
	user         string
	password     string
	sessionToken string
	cssString    string
	jsString     string
}

func (m *TunnelHandlerMux) ListenAndServe() {
	m.Server.ListenAndServe()
}

func (m *TunnelHandlerMux) PageCheck(path string) bool {
	for _, v := range m.pagenames {
		if strings.Contains(path, strings.Replace(v, "/", "", 0)) {
			return true
		}
	}
	return false
}

func (m *TunnelHandlerMux) CheckCookie(w http.ResponseWriter, r *http.Request) bool {
	if m.password != "" {
		if m.sessionToken == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return false
		}
		c, err := r.Cookie("session_token")
		if err != nil {
			if err == http.ErrNoCookie {
				w.WriteHeader(http.StatusUnauthorized)
				return false
			}
			w.WriteHeader(http.StatusBadRequest)
			return false
		}
		if m.sessionToken != c.Value {
			w.WriteHeader(http.StatusUnauthorized)
			return false
		}
	}
	return true

}

func (m *TunnelHandlerMux) HandlerWrapper(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if m.CheckCookie(w, r) == false {
			return
		}
		if m.PageCheck(r.URL.Path) {
			fmt.Fprintf(w, "<!DOCTYPE html>\n")
			fmt.Fprintf(w, "<html>\n")
			fmt.Fprintf(w, "<head>\n")
			fmt.Fprintf(w, "  <link rel=\"stylesheet\" href=\"/styles.css\">")
			fmt.Fprintf(w, "</head>\n")
			fmt.Fprintf(w, "<body>\n")
			h.ServeHTTP(w, r)
			fmt.Fprintf(w, "  <script src=\"/scripts.js\"></script>\n")
			fmt.Fprintf(w, "</body>\n")
			fmt.Fprintf(w, "</html>\n")
		} else if !strings.HasSuffix(r.URL.Path, "color") {
			h.ServeHTTP(w, r)
		} else {
			fmt.Fprintf(w, "<!DOCTYPE html>\n")
			fmt.Fprintf(w, "<html>\n")
			fmt.Fprintf(w, "<head>\n")
			fmt.Fprintf(w, "  <link rel=\"stylesheet\" href=\"/styles.css\">")
			fmt.Fprintf(w, "</head>\n")
			fmt.Fprintf(w, "<body>\n")
			h.ServeHTTP(w, r)
			fmt.Fprintf(w, "  <script src=\"/scripts.js\"></script>\n")
			fmt.Fprintf(w, "</body>\n")
			fmt.Fprintf(w, "</html>\n")
		}
	})
}

func (t *TunnelHandlerMux) Tunnels() []*TunnelHandler {
	return t.tunnels
}

func (m *TunnelHandlerMux) Append(v *TunnelHandler) *TunnelHandlerMux {
	for _, prev := range m.tunnels {
		if v.ID() == prev.ID() {
			log.Printf("v.ID() found, %s == %s", v.ID(), prev.ID())
			return m
		}
	}
	log.Printf("Adding tunnel ID: %s", v.ID())
	m.tunnels = append(m.tunnels, v)
	Handler := m.Handler.(*http.ServeMux)
	Handler.Handle(fmt.Sprintf("/%d", len(m.tunnels)), m.HandlerWrapper(v))
	Handler.Handle(fmt.Sprintf("/%s", v.ID()), m.HandlerWrapper(v))
	Handler.Handle(fmt.Sprintf("/%d/color", len(m.tunnels)), m.HandlerWrapper(v))
	Handler.Handle(fmt.Sprintf("/%s/color", v.ID()), m.HandlerWrapper(v))
	m.Handler = Handler
	return m
}

func ReadFile(filename string) (string, error) {
	r, e := ioutil.ReadFile(filename)
	return string(r), e
}

func NewTunnelHandlerMux(host, port, user, password, css, javascript string) *TunnelHandlerMux {
	var m TunnelHandlerMux
	m.Addr = host + ":" + port
	Handler := http.NewServeMux()
	m.pagenames = []string{"index.html", "/"}
	m.user = user
	m.password = password
	m.sessionToken = ""
	m.tunnels = []*TunnelHandler{}
	var err error
	m.cssString, err = ReadFile(css)
	if err != nil {
		m.cssString = DefaultCSS()
	}
	m.jsString, err = ReadFile(javascript)
	if err != nil {
		m.jsString = DefaultJS()
	}
	for _, v := range m.pagenames {
		Handler.HandleFunc(fmt.Sprintf("/%s", v), m.Home)
	}
	Handler.HandleFunc("/styles.css", m.CSS)
	Handler.HandleFunc("/scripts.js", m.JS)
	if m.password != "" {
		Handler.HandleFunc("/login", m.Signin)
	}
	m.Handler = Handler
	return &m
}
