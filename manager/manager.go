package sammanager

import (
	//"fmt"
	"log"
	"os/exec"
	"os/user"
	"runtime"
	"strconv"
)

import (
	"github.com/eyedeekay/portcheck"
	"github.com/eyedeekay/sam-forwarder/config"
	"github.com/eyedeekay/sam-forwarder/handler"
	"github.com/justinas/nosurf"
)

type SAMManager struct {
	FilePath string
	save     bool
	start    bool
	config   *i2ptunconf.Conf

	tunName string

	ServerHost string
	ServerPort string
	SamHost    string
	SamPort    string
	UseWeb     bool
	WebHost    string
	WebPort    string

	cssFile string
	jsFile  string

	handlerMux *samtunnelhandler.TunnelHandlerMux
}

var err error

func (s *SAMManager) Cleanup() {
	for _, k := range s.handlerMux.Tunnels() {
		k.Cleanup()
	}
}

func (s *SAMManager) UseWebUI() bool {
	return s.UseWeb
}

func (s *SAMManager) Title() string {
	return s.config.UserName
}

func (s *SAMManager) Width() int {
	return 800
}

func (s *SAMManager) Height() int {
	return 600
}

func (s *SAMManager) Resizable() bool {
	return true
}

func (s *SAMManager) URL() string {
	return "http://" + s.WebHost + ":" + s.WebPort
}

func User() string {
	runningUser, _ := user.Current()
	if runtime.GOOS != "windows" {
		if runningUser.Uid == "0" {
			cmd := exec.Command("logname")
			out, err := cmd.Output()
			if err != nil {
				return err.Error()
			}
			return string(out)
		}
	}
	return runningUser.Name
}

var runningUser = User()

func NewSAMManagerFromOptions(opts ...func(*SAMManager) error) (*SAMManager, error) {
	var s SAMManager
	s.FilePath = ""
	s.save = true
	s.start = false
	s.UseWeb = true
	s.config = i2ptunconf.NewI2PBlankTunConf()
	s.ServerHost = "localhost"
	s.ServerPort = "8081"
	s.SamHost = "localhost"
	s.SamPort = "7656"
	s.WebHost = "localhost"
	s.WebPort = "7957"
	s.tunName = "samcatd-"
	s.config.UserName = "samcatd"
	s.config.Password = ""
	s.cssFile = ""
	s.jsFile = ""
	for _, o := range opts {
		if err := o(&s); err != nil {
			return nil, err
		}
	}
	if port, err := strconv.Atoi(s.WebPort); err != nil {
		log.Println("Error:", err)
		return nil, err
	} else {
		if pc.SCL(port) {
			log.Println("Service found, launching GUI")
			s.RunUI()
			return nil, nil
		}
	}
	s.handlerMux = samtunnelhandler.NewTunnelHandlerMux(s.WebHost, s.WebPort, s.config.UserName, s.config.Password, s.cssFile, s.jsFile)
	log.Println("tunnel settings", s.ServerHost, s.ServerPort, s.SamHost, s.SamPort)
	var err error
	if s.FilePath != "" {
		s.config, err = i2ptunconf.NewI2PTunConf(s.FilePath)
		s.config.TargetHost = s.config.GetHost(s.ServerHost, "127.0.0.1")
		s.config.TargetPort = s.config.GetPort(s.ServerPort, "8081")
		if err != nil {
			return nil, err
		}

		for _, label := range s.config.Labels {
			if t, e := s.config.Get("type", label); e {
				switch t {
				case "http":
					if f, e := samtunnelhandler.NewTunnelHandler(i2ptunconf.NewSAMForwarderFromConfig(s.FilePath, s.SamHost, s.SamPort, label)); e == nil {
						log.Println("found http under", label)
						s.handlerMux = s.handlerMux.Append(f)
					} else {
						return nil, e
					}
				case "httpclient":
					if f, e := samtunnelhandler.NewTunnelHandler(i2ptunconf.NewSAMHTTPClientFromConfig(s.FilePath, s.SamHost, s.SamPort, label)); e == nil {
						log.Println("found http under", label)
						s.handlerMux = s.handlerMux.Append(f)
					} else {
						return nil, e
					}
				case "server":
					if f, e := samtunnelhandler.NewTunnelHandler(i2ptunconf.NewSAMForwarderFromConfig(s.FilePath, s.SamHost, s.SamPort, label)); e == nil {
						log.Println("found server under", label)
						s.handlerMux = s.handlerMux.Append(f)
					} else {
						return nil, e
					}
				case "client":
					if f, e := samtunnelhandler.NewTunnelHandler(i2ptunconf.NewSAMClientForwarderFromConfig(s.FilePath, s.SamHost, s.SamPort, label)); e == nil {
						log.Println("found client under", label)
						s.handlerMux = s.handlerMux.Append(f)
					} else {
						return nil, e
					}
				case "udpserver":
					if f, e := samtunnelhandler.NewTunnelHandler(i2ptunconf.NewSAMSSUForwarderFromConfig(s.FilePath, s.SamHost, s.SamPort, label)); e == nil {
						log.Println("found udpserver under", label)
						s.handlerMux = s.handlerMux.Append(f)
					} else {
						return nil, e
					}
				case "udpclient":
					if f, e := samtunnelhandler.NewTunnelHandler(i2ptunconf.NewSAMSSUClientForwarderFromConfig(s.FilePath, s.SamHost, s.SamPort, label)); e == nil {
						log.Println("found udpclient under", label)
						s.handlerMux = s.handlerMux.Append(f)
					} else {
						return nil, e
					}
				/*case "vpnserver":
					if f, e := samtunnelhandler.NewTunnelHandler(samforwardervpnserver.NewSAMVPNForwarderFromConfig(s.FilePath, s.SamHost, s.SamPort, label)); e == nil {
						log.Println("found vpnserver under", label)
						s.handlerMux = s.handlerMux.Append(f)
					} else {
						return nil, e
					}
				case "vpnclient":
					if f, e := samtunnelhandler.NewTunnelHandler(samforwardervpn.NewSAMVPNClientForwarderFromConfig(s.FilePath, s.SamHost, s.SamPort, label)); e == nil {
						log.Println("found vpnclient under", label)
						s.handlerMux = s.handlerMux.Append(f)
					} else {
						return nil, e
					}*/
				default:
					if f, e := samtunnelhandler.NewTunnelHandler(i2ptunconf.NewSAMForwarderFromConfig(s.FilePath, s.SamHost, s.SamPort, label)); e == nil {
						log.Println("found server under", label)
						s.handlerMux = s.handlerMux.Append(f)
					} else {
						return nil, e
					}
				}
			}
		}
	} else {
		t, b := s.config.Get("type")
		if !b {
			t = "client"
			//return nil, fmt.Errorf("samcat was instructed to start a tunnel with insufficient default settings information.")
		}
		switch t {
		case "http":
			if f, e := samtunnelhandler.NewTunnelHandler(i2ptunconf.NewSAMForwarderFromConf(s.config)); e == nil {
				log.Println("found default http")
				s.handlerMux = s.handlerMux.Append(f)
			} else {
				return nil, e
			}
		case "httpclient":
			if f, e := samtunnelhandler.NewTunnelHandler(i2ptunconf.NewSAMHTTPClientFromConf(s.config)); e == nil {
				log.Println("found default httpclient")
				s.handlerMux = s.handlerMux.Append(f)
			} else {
				return nil, e
			}
		case "browserclient":
			if f, e := samtunnelhandler.NewTunnelHandler(i2ptunconf.NewSAMBrowserClientFromConf(s.config)); e == nil {
				log.Println("found default browserclient")
				s.handlerMux = s.handlerMux.Append(f)
			} else {
				return nil, e
			}
		case "server":
			if f, e := samtunnelhandler.NewTunnelHandler(i2ptunconf.NewSAMForwarderFromConf(s.config)); e == nil {
				log.Println("found default server")
				s.handlerMux = s.handlerMux.Append(f)
			} else {
				return nil, e
			}
		case "client":
			if f, e := samtunnelhandler.NewTunnelHandler(i2ptunconf.NewSAMClientForwarderFromConf(s.config)); e == nil {
				log.Println("found default client")
				s.handlerMux = s.handlerMux.Append(f)
			} else {
				return nil, e
			}
		case "udpserver":
			if f, e := samtunnelhandler.NewTunnelHandler(i2ptunconf.NewSAMSSUForwarderFromConf(s.config)); e == nil {
				log.Println("found default udpserver")
				s.handlerMux = s.handlerMux.Append(f)
			} else {
				return nil, e
			}
		case "udpclient":
			if f, e := samtunnelhandler.NewTunnelHandler(i2ptunconf.NewSAMSSUClientForwarderFromConf(s.config)); e == nil {
				log.Println("found default udpclient")
				s.handlerMux = s.handlerMux.Append(f)
			} else {
				return nil, e
			}
		/*case "vpnserver":
			if f, e := samtunnelhandler.NewTunnelHandler(samforwardervpnserver.NewSAMVPNForwarderFromConf(s.config)); e == nil {
				log.Println("found default vpnserver")
				s.handlerMux = s.handlerMux.Append(f)
			} else {
				return nil, e
			}
		case "vpnclient":
			if f, e := samtunnelhandler.NewTunnelHandler(samforwardervpn.NewSAMVPNClientForwarderFromConf(s.config)); e == nil {
				log.Println("found default vpnclient")
				s.handlerMux = s.handlerMux.Append(f)
			} else {
				return nil, e
			}*/
		default:
			if f, e := samtunnelhandler.NewTunnelHandler(i2ptunconf.NewSAMClientForwarderFromConf(s.config)); e == nil {
				log.Println("found default client")
				s.handlerMux = s.handlerMux.Append(f)
			} else {
				return nil, e
			}
		}
	}
	s.handlerMux.Handler = nosurf.New(s.handlerMux.Handler)
	return &s, nil
}

func NewSAMManager(inifile, servhost, servport, samhost, samport, webhost, webport, cssfile, jsfile string, start, web bool, webuser, webpass string) (*SAMManager, error) {
	log.Println("tunnel settings", servhost, servport, samhost, samport)
	return NewSAMManagerFromOptions(
		SetManagerFilePath(inifile),
		SetManagerHost(servhost),
		SetManagerPort(servport),
		SetManagerSAMHost(samhost),
		SetManagerSAMPort(samport),
		SetManagerWebHost(webhost),
		SetManagerWebPort(webport),
		SetManagerStart(start),
		SetManagerWebUser(webuser),
		SetManagerWebPass(webpass),
		SetManagerWeb(web),
	)
}

func NewSAMManagerFromConf(conf *i2ptunconf.Conf, servhost, servport, samhost, samport, webhost, webport, cssfile, jsfile string, start, web bool, webuser, webpass string) (*SAMManager, error) {
	log.Println("tunnel settings", servhost, servport, samhost, samport)
	return NewSAMManagerFromOptions(
		SetManagerConf(conf),
		SetManagerHost(servhost),
		SetManagerPort(servport),
		SetManagerSAMHost(samhost),
		SetManagerSAMPort(samport),
		SetManagerWebHost(webhost),
		SetManagerWebPort(webport),
		SetManagerStart(start),
		SetManagerWebUser(webuser),
		SetManagerWebPass(webpass),
		SetManagerWeb(web),
	)
}
