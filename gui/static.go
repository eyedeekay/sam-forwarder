// +build !nostatic

package gui

import (
	"github.com/eyedeekay/sam-forwarder/interface"
	"github.com/zserge/lorca"
)

var USER = ""

func LaunchUI(s samtunnel.WebUI) (lorca.UI, error) {
	if s.UseWebUI() == true {
		if lorca.LocateChrome() != "" {
			return lorca.New(s.URL(), s.Title(), s.Width(), s.Height())
		}
	}
	return nil, nil
}
