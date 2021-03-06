//|------------------------------------------------------------------
//|        __
//|     __/  \
//|  __/  \__/_
//| /  \__/    \
//|/\__/CellGo /_
//|\/_/NetFW__/  \
//|  /\__ _/  \__/
//|  \/_/  \__/_/
//|    /\__/_/
//|    \/_/
//|------------------------------------------------------------------
//| Cellgo Framework conf file
//| All rights reserved: By cellgo.cn CopyRight
//| You are free to use the source code, but in the use of the process,
//| please keep the author information. Respect for the work of others
//| is respect for their own
//|-------------------------------------------------------------------
// Author:Tommy.Jin Dtime:2016-7-18

package cellgo

// Conf is the main struct for Config
type Conf struct {
	NetName    string //Net name
	ServerName string
	Listen
	SiteConfig
}

// Listen holds for http/https/websocket related config
type Listen struct {
	ServerTimeOut int64
	HTTPAddr      string //http  conn addr
	HTTPPort      int    //http  conn Port
	HTTPSAddr     string //https conn addr
	HTTPSPort     int    //https conn Port
	HTTPSCertFile string //https conn certfile
	HTTPSKeyFile  string //https conn keyfile
	WEBSOCKETAddr string //websocket conn Port
	WEBSOCKETPort int    //websocket conn Port
}

// SiteConfig holds Site related config
type SiteConfig struct {
	AutoDisplay      bool
	DefaultBeforeAct string
	DefaultAfterAct  string
	Dynamic          string
	StaticDir        string
	StaticRouter     []string
	LabLeft          string
	LabRight         string
	TemplateExt      string
	TemplatePath     string
	//Session
}

// Version number of the cellgo.
const (
	VERSION  = "0.0.5"
	LASTDATE = "July 20, 2016"
)

var (
	// CellConf is the default config for Cellgo
	CellConf *Conf
)

func init() {
	CellConf = &Conf{
		NetName:    "cellgo",
		ServerName: "CellgoService_" + VERSION,
		Listen: Listen{
			ServerTimeOut: 0,
			HTTPAddr:      "",
			HTTPPort:      80,
			HTTPSAddr:     "",
			HTTPSPort:     10443,
			HTTPSCertFile: "",
			HTTPSKeyFile:  "",
			WEBSOCKETAddr: "",
			WEBSOCKETPort: 8088,
		},
		SiteConfig: SiteConfig{
			AutoDisplay:      true,
			DefaultBeforeAct: "Before",
			DefaultAfterAct:  "After",
			Dynamic:          "/",
			StaticDir:        "static",
			StaticRouter:     []string{"/css/", "/js/", "/images/"},
			LabLeft:          "{{",
			LabRight:         "}}",
			TemplateExt:      "html",
			TemplatePath:     "template",
		},
	}

}
