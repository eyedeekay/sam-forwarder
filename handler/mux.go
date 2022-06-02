package samtunnelhandler

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

type TunnelHandlerMux struct {
	http.Server
	pagenames    []string
	tunnels      []*TunnelHandler
	user         string
	password     string
	sessionToken string
	//cssString    string
	//jsString     string

	templateTop     string
	templateBot     string
	templateTopHTML *template.Template
	templateBotHTML *template.Template
}

func (m *TunnelHandlerMux) ListenAndServe() {
	m.Server.ListenAndServe()
}

func (m *TunnelHandlerMux) Count() int {
	return len(m.tunnels)
}

func (m *TunnelHandlerMux) User() string {
	if m.user == "" {
		return "anonymous"
	}
	return m.user
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
		h.ServeHTTP(w, r)
	})
}

func (t *TunnelHandlerMux) Tunnels() []*TunnelHandler {
	return t.tunnels
}

func (m *TunnelHandlerMux) Append(v *TunnelHandler) *TunnelHandlerMux {
	if m == nil {
		return m
	}
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
	m.pagenames = []string{"index.html", "index", ""}
	m.user = user
	m.password = password
	m.sessionToken = ""
	m.tunnels = []*TunnelHandler{}
	m.templateTop = `
<head>
<meta charset="utf-8">
<title>SAMTunnel</title>
<link rel="stylesheet" href="/styles.css">
<script src="/scripts.js"></script>
</head>
<body>
<h1>SAMTunnel</h1>
<a class="samtunnel-home-page" href="/index.html">Welcome {{.User}}! you are serving {{.Count}} tunnels. </a>
<div id="toggleall" class="global control">
<a class="samtunnel-global-toggle" href="#" onclick="toggle_visibility_class('prop');">Show/Hide All</a>
</div>
`
	m.templateBot = `</body>
</html>
`
	var err error
	m.templateTopHTML = template.Must(template.New("top").Parse(m.templateTop))
	if err != nil {
		log.Printf("Error parsing templateTop: %s", err)
	}
	m.templateBotHTML = template.Must(template.New("bot").Parse(m.templateBot))
	if err != nil {
		log.Printf("Error parsing templateBot: %s", err)
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

func (m *TunnelHandlerMux) Home(w http.ResponseWriter, r *http.Request) {
	if m.CheckCookie(w, r) == false {
		return
	}
	if r.URL.Path == "/" {
		http.Redirect(w, r, "/index.html", 301)
		fmt.Fprintf(w, "redirecting to index.html")
		return
	}
	r2, err := http.NewRequest("GET", r.URL.Path+"/color", r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	m.templateTopHTML.Execute(w, m)
	for _, tunnel := range m.Tunnels() {
		tunnel.ServeHTTP(w, r2)
	}
	m.templateBotHTML.Execute(w, m)
}
