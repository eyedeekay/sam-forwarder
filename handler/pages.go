package samtunnelhandler

import (
	"fmt"
	"net/http"
	"strings"
)

func DefaultCSS() string {
	return `.server {
    width: 63%;
    min-height: 15%;
    background-color: #9DABD5;
    float: left;
    overflow-wrap: break-word;
}
.client {
    width: 63%;
    min-height: 15%;
    background-color: #2D4470;
    float: left;
    overflow-wrap: break-word;
}
.tcpclient {
    width: 63%;
    min-height: 15%;
    background-color: #2D4470;
    float: left;
    overflow-wrap: break-word;
}
.http {
    width: 63%;
    min-height: 15%;
    background-color: #00ffff;
    float: left;
    overflow-wrap: break-word;
}
.httpclient {
    width: 63%;
    min-height: 15%;
    background-color: #709fa6;
    float: left;
    overflow-wrap: break-word;
}
.udpserver {
    width: 63%;
    min-height: 15%;
    background-color: #265ea7;
    float: left;
    overflow-wrap: break-word;
}
.outproxy {
    width: 63%;
    min-height: 15%;
    background-color: #265ea7;
    float: left;
    overflow-wrap: break-word;
}
.outproxyhttp {
    width: 63%;
    min-height: 15%;
    background-color: #265ea7;
    float: left;
    overflow-wrap: break-word;
}
.udpclient {
    width: 63%;
    min-height: 15%;
    background-color: #222187;
    float: left;
    overflow-wrap: break-word;
}
.vpnserver {
    width: 63%;
    min-height: 15%;
    background-color: #265ea7;
    float: left;
    overflow-wrap: break-word;
}
.vpnclient {
    width: 63%;
    min-height: 15%;
    background-color: #222187;
    float: left;
    overflow-wrap: break-word;
}
.TunName {
    font-weight: bold;
}
.panel {
    width: 33%;
    float: right;
}
.prop {

}
.prop:hover {
    box-shadow: inset 0 0 100px 100px rgba(255, 255, 255, 0.1);
}
.global {
    background-color: #00ffff;
}
body {
    background-color: #9e9e9e;
    color: #070425;
    font-family: "monospace";
    font-size: 1rem;
}
a {
    color: #080808;
}
h1 {
    background-color: #9e9e9e;
}
span {
    float: left;
    display: inline;
}
textarea {
    display: inline-block;
    width: 100%;
    resize: none;
    height: 1rem;
    float: right;
    display: inline;
}
.linkstyle {
    align-items: normal;
    background-color: rgba(0,0,0,0);
    border-color: rgb(0, 0, 238);
    border-style: none;
    box-sizing: content-box;
    color: rgb(0, 0, 238);
    cursor: pointer;
    font: inherit;
    height: auto;
    display: inline;
    padding: 0;
    perspective-origin: 0 0;
    text-align: start;
    text-decoration: underline;
    transform-origin: 0 0;
    width: auto;
    -moz-appearance: none;
    -webkit-logical-height: 1em; /* Chrome ignores auto, so we have to use this hack to set the correct height  */
    -webkit-logical-width: auto; /* Chrome ignores auto, but here for completeness */
}

@supports (-moz-appearance:none) { /* Mozilla-only */
    .linkstyle::-moz-focus-inner { /* reset any predefined properties */
        border: none;
        padding: 0;
    }
    .linkstyle:focus { /* add outline to focus pseudo-class */
        outline-style: dotted;
        outline-width: 1px;
    }
}
`
}

func DefaultJS() string {
	return `function toggle_visibility_id(id) {
   var e = document.getElementById(id);
   if(e.style.display == 'block')
      e.style.display = 'none';
   else
      e.style.display = 'block';
}
function toggle_visibility_class(id) {
   var elist = document.getElementsByClassName(id)
   for (let e of elist) {
       if(e.style.display == 'block')
          e.style.display = 'none';
       else
          e.style.display = 'block';
   }
   var tlist = document.getElementsByClassName("TunName")
   for (let t of tlist) {
       t.style.display = 'block';
   }
   var clist = document.getElementsByClassName("control")
   for (let c of clist) {
       c.style.display = 'block';
   }
   var slist = document.getElementsByClassName("status")
   for (let s of slist) {
       s.style.display = 'block';
   }
}
toggle_visibility_class("prop")
`
}

func (t *TunnelHandler) ColorSpan(shortid string) string {
	r := fmt.Sprintf("  <span id=\"toggle%s\" class=\"control\">\n", t.SAMTunnel.ID())
	r += fmt.Sprintf("    <a href=\"#\" onclick=\"toggle_visibility_class('%s');\"> Show/Hide %s</a><br>\n", t.SAMTunnel.ID(), t.SAMTunnel.ID())
	r += fmt.Sprintf("    <a href=\"/%s/color\">Tunnel page</a>\n", t.SAMTunnel.ID())
	r += fmt.Sprintf("  </span>\n")
	return r
}

func (t *TunnelHandler) ColorForm(shortid, tuntype string) string {
	r := fmt.Sprintf("  </div>\n\n")
	r += fmt.Sprintf("  <div id=\"%s\" class=\"%s control panel\" >", shortid+".control", tuntype)

	r += fmt.Sprintf("    <form class=\"linkstyle\" name=\"start\" action=\"/%s\" method=\"post\">", shortid)
	r += fmt.Sprintf("      <input class=\"linkstyle\" type=\"hidden\" value=\"start\" name=\"action\" />")
	r += fmt.Sprintf("      <input class=\"linkstyle\" type=\"submit\" value=\".[START]\">")
	r += fmt.Sprintf("    </form>")

	r += fmt.Sprintf("    <form class=\"linkstyle\" name=\"stop\" action=\"/%s\" method=\"post\">", shortid)
	r += fmt.Sprintf("      <input class=\"linkstyle\" type=\"hidden\" value=\"stop\" name=\"action\" />")
	r += fmt.Sprintf("      <input class=\"linkstyle\" type=\"submit\" value=\".[STOP].\">")
	r += fmt.Sprintf("    </form>")

	r += fmt.Sprintf("    <form class=\"linkstyle\" name=\"restart\" action=\"/%s\" method=\"post\">", shortid)
	r += fmt.Sprintf("      <input class=\"linkstyle\" type=\"hidden\" value=\"restart\" name=\"action\" />")
	r += fmt.Sprintf("      <input class=\"linkstyle\" type=\"submit\" value=\"[RESTART].\">")
	r += fmt.Sprintf("    </form>")

	r += fmt.Sprintf("    <div id=\"%s.status\" class=\"%s status\">.[STATUS].</div>", shortid, shortid)
	r += fmt.Sprintf("  </div>\n\n")
	return r
}

func (t *TunnelHandler) ColorWrap(longid, shortid, key, tuntype, prop, value string) string {
	r := fmt.Sprintf("    <div id=\"%s\" class=\"%s %s %s %s\" >\n", longid, shortid, key, tuntype, prop)
	r += fmt.Sprintf("      <span id=\"%s\" class=\"key\">%s</span>=", longid, key)
	r += fmt.Sprintf("      <textarea id=\"%s\" rows=\"1\" class=\"value\">%s</textarea>\n", longid, value)
	r += fmt.Sprintf("    </div>\n\n")
	return r
}

func (t *TunnelHandler) ColorDiv(shortid, tuntype string) string {
	return fmt.Sprintf("  <div id=\"%s\" class=\"%s\" >", t.SAMTunnel.ID(), t.SAMTunnel.GetType())
}

func (m *TunnelHandlerMux) ColorHeader(h http.Handler, r *http.Request, w http.ResponseWriter) {
	if !strings.HasSuffix(r.URL.Path, "color") {
		h.ServeHTTP(w, r)
	} else {
		fmt.Fprintf(w, "<!DOCTYPE html>\n")
		fmt.Fprintf(w, "<html>\n")
		fmt.Fprintf(w, "<head>\n")
		fmt.Fprintf(w, "  <link rel=\"stylesheet\" href=\"/styles.css\">")
		fmt.Fprintf(w, "</head>\n")
		fmt.Fprintf(w, "<body>\n")
		fmt.Fprintf(w, "<h1>\n")
		w.Write([]byte(fmt.Sprintf("<a href=\"/index.html\">Welcome %s! you are serving %d tunnels. </a>\n", m.user, len(m.tunnels))))
		fmt.Fprintf(w, "</h1>\n")
		fmt.Fprintf(w, "  <div id=\"toggleall\" class=\"global control\">\n")
		fmt.Fprintf(w, "    <a href=\"#\" onclick=\"toggle_visibility_class('%s');\">Show/Hide %s</a>\n", "prop", "all")
		fmt.Fprintf(w, "  </div>\n")
		h.ServeHTTP(w, r)
		fmt.Fprintf(w, "  <script src=\"/scripts.js\"></script>\n")
		fmt.Fprintf(w, "</body>\n")
		fmt.Fprintf(w, "</html>\n")
	}
}
