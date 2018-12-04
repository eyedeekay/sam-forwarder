package sammanager

import (
	"fmt"
	"log"
	"strings"
)

import (
	"github.com/eyedeekay/sam-forwarder"
	"github.com/eyedeekay/sam-forwarder/config"
	"github.com/eyedeekay/sam-forwarder/csvpn"
	"github.com/eyedeekay/sam-forwarder/udp"
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
	WebHost    string
	WebPort    string

	forwarders          []*samforwarder.SAMForwarder
	clientforwarders    []*samforwarder.SAMClientForwarder
	udpforwarders       []*samforwarderudp.SAMSSUForwarder
	udpclientforwarders []*samforwarderudp.SAMSSUClientForwarder
	vpnforwarders       []*samforwardervpn.SAMClientServerVPN
	vpnclientforwarders []*samforwardervpn.SAMClientVPN
}

func stringify(s []string) string {
	var p string
	for _, x := range s {
		if x != "ntcpserver" && x != "httpserver" && x != "ssuserver" && x != "ntcpclient" && x != "ssuclient" {
			p += x + ","
		}
	}
	r := strings.Trim(strings.Trim(strings.Replace(p, ",,", ",", -1), " "), "\n")
	return r
}

func (s *SAMManager) List(search ...string) *[]string {
	var r []string
	if search == nil {
		for index, element := range s.forwarders {
			r = append(r, fmt.Sprintf("  %v. %s", index, element.Print()))
		}
		for index, element := range s.clientforwarders {
			r = append(r, fmt.Sprintf("  %v. %s", index, element.Print()))
		}
		for index, element := range s.udpforwarders {
			r = append(r, fmt.Sprintf("  %v. %s", index, element.Print()))
		}
		for index, element := range s.udpclientforwarders {
			r = append(r, fmt.Sprintf("  %v. %s", index, element.Print()))
		}
		return &r
	} else if len(search) > 0 {
		switch search[0] {
		case "":
			for index, element := range s.forwarders {
				r = append(r, fmt.Sprintf("  %v. %s", index, element.Print()))
			}
			for index, element := range s.clientforwarders {
				r = append(r, fmt.Sprintf("  %v. %s", index, element.Print()))
			}
			for index, element := range s.udpforwarders {
				r = append(r, fmt.Sprintf("  %v. %s", index, element.Print()))
			}
			for index, element := range s.udpclientforwarders {
				r = append(r, fmt.Sprintf("  %v. %s", index, element.Print()))
			}
			return &r
		case "ntcpserver":
			for index, element := range s.forwarders {
				r = append(r, fmt.Sprintf("  %v. %s", index, element.Search(stringify(search))))
			}
			return &r
		case "httpserver":
			for index, element := range s.forwarders {
				if element.Type == "http" {
					r = append(r, fmt.Sprintf("  %v. %s", index, element.Search(stringify(search))))
				}
			}
			return &r
		case "ntcpclient":
			for index, element := range s.clientforwarders {
				r = append(r, fmt.Sprintf("  %v. %s", index, element.Search(stringify(search))))
			}
			return &r
		case "ssuserver":
			for index, element := range s.udpforwarders {
				r = append(r, fmt.Sprintf("  %v. %s", index, element.Search(stringify(search))))
			}
			return &r
		case "ssuclient":
			for index, element := range s.udpclientforwarders {
				r = append(r, fmt.Sprintf("  %v. %s", index, element.Search(stringify(search))))
			}
			return &r
		default:
			for index, element := range s.forwarders {
				if element.Search(stringify(search)) != "" {
					r = append(r, fmt.Sprintf("  %v. %s", index, element.Search(stringify(search))))
				}
			}
			for index, element := range s.clientforwarders {
				if element.Search(stringify(search)) != "" {
					r = append(r, fmt.Sprintf("  %v. %s", index, element.Search(stringify(search))))
				}
			}
			for index, element := range s.udpforwarders {
				if element.Search(stringify(search)) != "" {
					r = append(r, fmt.Sprintf("  %v. %s", index, element.Search(stringify(search))))
				}
			}
			for index, element := range s.udpclientforwarders {
				if element.Search(stringify(search)) != "" {
					r = append(r, fmt.Sprintf("  %v. %s", index, element.Search(stringify(search))))
				}
			}
			return &r
		}
	}
	return &r
}

func (s *SAMManager) Cleanup() {
	for _, k := range s.forwarders {
		k.Cleanup()
	}
	for _, k := range s.clientforwarders {
		k.Cleanup()
	}
	for _, k := range s.udpforwarders {
		k.Cleanup()
	}
	for _, k := range s.udpclientforwarders {
		k.Cleanup()
	}
}

func (s *SAMManager) Serve() bool {
	log.Println("Starting tunnels")
	for _, element := range s.forwarders {
		log.Println("Starting NTCP service tunnel", element.TunName)
		go element.Serve()
	}
	for _, element := range s.clientforwarders {
		log.Println("Starting NTCP client tunnel", element.TunName)
		go element.Serve()
	}
	for _, element := range s.udpforwarders {
		log.Println("Starting SSU service tunnel", element.TunName)
		go element.Serve()
	}
	for _, element := range s.udpclientforwarders {
		log.Println("Starting SSU client tunnel", element.TunName)
		go element.Serve()
	}
	for true {
	}
	return false
}

func NewSAMManagerFromOptions(opts ...func(*SAMManager) error) (*SAMManager, error) {
	var s SAMManager
	s.FilePath = ""
	s.save = true
	s.start = false
	s.config = i2ptunconf.NewI2PBlankTunConf()
	s.ServerHost = "localhost"
	s.ServerPort = "8081"
	s.SamHost = "localhost"
	s.SamPort = "7656"
	s.WebHost = "localhost"
	s.WebPort = "7957"
	s.tunName = "samcatd-"
	for _, o := range opts {
		if err := o(&s); err != nil {
			return nil, err
		}
	}
	log.Println("tunnel settings", s.ServerHost, s.ServerPort, s.SamHost, s.SamPort)
	var err error
	if s.FilePath != "" {
		s.config, err = i2ptunconf.NewI2PTunConf(s.FilePath)
		s.config.TargetHost = s.config.GetHost(s.ServerHost, "127.0.0.1")
		s.config.TargetPort = s.config.GetPort(s.ServerPort, "8081")
		if err != nil {
			return nil, err
		}
	} else {
		if s.config == nil {
			return nil, fmt.Errorf("Configuration not found")
		}
		s.FilePath = s.config.FilePath
	}
	for _, label := range s.config.Labels {
		if t, e := s.config.Get("type", label); e {
			switch t {
			case "http":
				if f, e := i2ptunconf.NewSAMForwarderFromConfig(s.FilePath, s.SamHost, s.SamPort, label); e == nil {
					log.Println("found http under", label)
					s.forwarders = append(s.forwarders, f)
				} else {
					return nil, fmt.Errorf(e.Error())
				}
			case "server":
				if f, e := i2ptunconf.NewSAMForwarderFromConfig(s.FilePath, s.SamHost, s.SamPort, label); e == nil {
					log.Println("found server under", label)
					s.forwarders = append(s.forwarders, f)
				} else {
					return nil, fmt.Errorf(e.Error())
				}
			case "client":
				if f, e := i2ptunconf.NewSAMClientForwarderFromConfig(s.FilePath, s.SamHost, s.SamPort, label); e == nil {
					log.Println("found client under", label)
					s.clientforwarders = append(s.clientforwarders, f)
				} else {
					return nil, fmt.Errorf(e.Error())
				}
			case "udpserver":
				if f, e := i2ptunconf.NewSAMSSUForwarderFromConfig(s.FilePath, s.SamHost, s.SamPort, label); e == nil {
					log.Println("found udpserver under", label)
					s.udpforwarders = append(s.udpforwarders, f)
				} else {
					return nil, fmt.Errorf(e.Error())
				}
			case "udpclient":
				if f, e := i2ptunconf.NewSAMSSUClientForwarderFromConfig(s.FilePath, s.SamHost, s.SamPort, label); e == nil {
					log.Println("found udpclient under", label)
					s.udpclientforwarders = append(s.udpclientforwarders, f)
				} else {
					return nil, fmt.Errorf(e.Error())
				}
			case "vpnserver":
				if f, e := samforwardervpn.NewSAMVPNForwarderFromConfig(s.FilePath, s.SamHost, s.SamPort, label); e == nil {
					log.Println("found vpnclient under", label)
					s.vpnforwarders = append(s.vpnforwarders, f)
				} else {
					return nil, fmt.Errorf(e.Error())
				}
			case "vpnclient":
				if f, e := samforwardervpn.NewSAMVPNClientForwarderFromConfig(s.FilePath, s.SamHost, s.SamPort, label); e == nil {
					log.Println("found vpnclient under", label)
					s.vpnclientforwarders = append(s.vpnclientforwarders, f)
				} else {
					return nil, fmt.Errorf(e.Error())
				}
			default:
				if f, e := i2ptunconf.NewSAMForwarderFromConfig(s.FilePath, s.SamHost, s.SamPort, label); e == nil {
					log.Println("found server under", label)
					s.forwarders = append(s.forwarders, f)
				} else {
					return nil, fmt.Errorf(e.Error())
				}
			}
		}
	}
	if len(s.config.Labels) == 0 || s.start {
		t, b := s.config.Get("type")
		if !b {
			t = "client"
			//return nil, fmt.Errorf("samcat was instructed to start a tunnel with insufficient default settings information.")
		}
		switch t {
		case "http":
			if f, e := i2ptunconf.NewSAMForwarderFromConfig(s.FilePath, s.SamHost, s.SamPort); e == nil {
				log.Println("found default http")
				s.forwarders = append(s.forwarders, f)
			} else {
				return nil, fmt.Errorf(e.Error())
			}
		case "server":
			if f, e := i2ptunconf.NewSAMForwarderFromConfig(s.FilePath, s.SamHost, s.SamPort); e == nil {
				log.Println("found default server")
				s.forwarders = append(s.forwarders, f)
			} else {
				return nil, fmt.Errorf(e.Error())
			}
		case "client":
			if f, e := i2ptunconf.NewSAMClientForwarderFromConfig(s.FilePath, s.SamHost, s.SamPort); e == nil {
				log.Println("found default client")
				s.clientforwarders = append(s.clientforwarders, f)
			} else {
				return nil, fmt.Errorf(e.Error())
			}
		case "udpserver":
			if f, e := i2ptunconf.NewSAMSSUForwarderFromConfig(s.FilePath, s.SamHost, s.SamPort); e == nil {
				log.Println("found default udpserver")
				s.udpforwarders = append(s.udpforwarders, f)
			} else {
				return nil, fmt.Errorf(e.Error())
			}
		case "udpclient":
			if f, e := i2ptunconf.NewSAMSSUClientForwarderFromConfig(s.FilePath, s.SamHost, s.SamPort); e == nil {
				log.Println("found default udpclient")
				s.udpclientforwarders = append(s.udpclientforwarders, f)
			} else {
				return nil, fmt.Errorf(e.Error())
			}
		case "vpnserver":
			if f, e := samforwardervpn.NewSAMVPNForwarderFromConfig(s.FilePath, s.SamHost, s.SamPort); e == nil {
				log.Println("found default vpnserver")
				s.vpnforwarders = append(s.vpnforwarders, f)
			} else {
				return nil, fmt.Errorf(e.Error())
			}
		case "vpnclient":
			if f, e := samforwardervpn.NewSAMVPNClientForwarderFromConfig(s.FilePath, s.SamHost, s.SamPort); e == nil {
				log.Println("found default vpnclient")
				s.vpnclientforwarders = append(s.vpnclientforwarders, f)
			} else {
				return nil, fmt.Errorf(e.Error())
			}
		default:
			if f, e := i2ptunconf.NewSAMClientForwarderFromConfig(s.FilePath, s.SamHost, s.SamPort); e == nil {
				log.Println("found default client")
				s.clientforwarders = append(s.clientforwarders, f)
			} else {
				return nil, fmt.Errorf(e.Error())
			}
		}
	}
	return &s, nil
}

func NewSAMManager(inifile, servhost, servport, samhost, samport, webhost, webport string, start bool) (*SAMManager, error) {
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
	)
}

func NewSAMManagerFromConf(conf *i2ptunconf.Conf, servhost, servport, samhost, samport, webhost, webport string, start bool) (*SAMManager, error) {
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
	)
}
