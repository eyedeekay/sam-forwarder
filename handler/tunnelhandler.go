package samtunnelhandler

import (
	"html/template"
	"net/http"

	samtunnel "github.com/eyedeekay/sam-forwarder/interface"
)

type TunnelHandler struct {
	samtunnel.SAMTunnel
	template     string
	htmlTemplate *template.Template
}

func (t *TunnelHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	t.htmlTemplate.Execute(rw, t.SAMTunnel)
}

func NewTunnelHandler(ob samtunnel.SAMTunnel, err error) (*TunnelHandler, error) {
	var t TunnelHandler
	t.SAMTunnel = ob
	t.template = `<div class="samtunnel">
<div class="samtunnel-header">
<div class="samtunnel-header-title">
<span class="samtunnel-header-title-text">
<span class="samtunnel-header-title-text-id">{{.ID}}</span>
<span class="samtunnel-header-title-text-type">{{.Type}}</span>
</span>
</div>
<div class="samtunnel-header-controls">
<form method="post" action="{{.ID}}/control">
<input type="hidden" name="action" value="start">
<input type="submit" value="Start">
</form>
<form method="post" action="{{.ID}}/control">
<input type="hidden" name="action" value="stop">
<input type="submit" value="Stop">
</form>
<form method="post" action="{{.ID}}/control">
<input type="hidden" name="action" value="restart">
<input type="submit" value="Restart">
</form>
</div>
</div>
<div class="samtunnel-body">
{{range $key, $value := .Props }}
<div class="samtunnel-body-prop">
<span class="samtunnel-body-prop-key">{{$key}}</span>
<span class="samtunnel-body-prop-value">{{$value}}</span>
</div>
{{end}}
</div>
</div>
`
	t.htmlTemplate, err = template.New(ob.ID()).Parse(t.template)
	if err != nil {
		return nil, err
	}
	return &t, err
}
