package main

type completedConfig struct {
	Random          bool        `json:"random"`
	AuthPass        interface{} `json:"authPass"`
	UseOnlinePac    bool        `json:"useOnlinePac"`
	TTL             int         `json:"TTL"`
	Global          bool        `json:"global"`
	ReconnectTimes  int         `json:"reconnectTimes"`
	Index           int         `json:"index"`
	ProxyType       int         `json:"proxyType"`
	ProxyHost       interface{} `json:"proxyHost"`
	AuthUser        interface{} `json:"authUser"`
	ProxyAuthPass   interface{} `json:"proxyAuthPass"`
	IsDefault       bool        `json:"isDefault"`
	PacURL          interface{} `json:"pacUrl"`
	Configs         []Config    `json:"configs"`
	ProxyPort       int         `json:"proxyPort"`
	RandomAlgorithm int         `json:"randomAlgorithm"`
	ProxyEnable     bool        `json:"proxyEnable"`
	Enabled         bool        `json:"enabled"`
	Autoban         bool        `json:"autoban"`
	ProxyAuthUser   interface{} `json:"proxyAuthUser"`
	ShareOverLan    bool        `json:"shareOverLan"`
	LocalPort       int         `json:"localPort"`
}

type Config struct {
	Enable        bool   `json:"enable"`
	Password      string `json:"password"`
	Method        string `json:"method"`
	Remarks       string `json:"remarks"`
	Server        string `json:"server"`
	Obfs          string `json:"obfs"`
	Protocol      string `json:"protocol"`
	ServerPort    int    `json:"server_port"`
	RemarksBase64 string `json:"remarks_base64"`
}

func newCompletedConfig() *completedConfig {
	return &completedConfig{
		Random:         false,
		ReconnectTimes: 3,
		Enabled:        true,
		LocalPort:      1080,
	}
}
