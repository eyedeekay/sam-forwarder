// +build static

package sammanager

import (
	"log"
	"os"
	"os/signal"
	"time"
)

import (
	. "github.com/eyedeekay/sam-forwarder/gui"
	"github.com/zserge/lorca"
)

var view lorca.UI

func (s *SAMManager) RunUI() {
	view, err = LaunchUI(s)
	if err != nil {
		log.Println(err.Error())
	}
}

func (s *SAMManager) Serve() bool {
	log.Println("Starting Tunnels()")
	for _, element := range s.handlerMux.Tunnels() {
		log.Println("Starting service tunnel", element.ID())
		go element.Serve()
	}

	if s.UseWebUI() == true {
		go s.handlerMux.ListenAndServe()
		if view, err = LaunchUI(s); err != nil {
			log.Println(err.Error())
			return false
		} else {
			return Exit()
		}
	} else {
		return Exit()
	}
	return false
}

func Exit() bool {
	Close := false
	for !Close {
		time.Sleep(1 * time.Second)
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		go func() {
			for sig := range c {
				log.Println(sig)
				if view != nil {
					view.Close()
				}
				Close = true
			}
		}()
	}
	return Close
}
