//go:build nostatic
// +build nostatic

package gui

import (
	"github.com/eyedeekay/sam-forwarder/interface"
	"github.com/zserge/webview"
)

func LaunchUI(s samtunnel.WebUI) (webview.WebView, error) {
	if s.UseWebUI() == true {
		settings := webview.Settings{
			Title:     s.Title(),
			URL:       s.URL(),
			Height:    s.Height(),
			Width:     s.Width(),
			Resizable: s.Resizable(),
			Debug:     true,
		}
		view := webview.New(settings)
		return view, nil
	}
	return nil, nil
}
