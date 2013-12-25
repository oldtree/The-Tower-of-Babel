package utils

import (
	"github.com/Unknwon/goconfig"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "mysql"
	"os"
	"path"
)

var (
	AppName    string
	AppVersion string
	AppHost    string
	AppUrl     string
	AppLogo    string
	//EnforceRedirect     bool
	AvatarURL string
	//SecretKey           string
	IsProMode bool
	//MailUser            string
	//MailFrom            string
	ActiveCodeLives   int
	ResetPwdCodeLives int
	LoginRememberDays int

	DateFormat          string
	DateTimeFormat      string
	DateTimeShortFormat string
	RealtimeRenderMD    bool
	//ImageSizeSmall      int
	//ImageSizeMiddle     int
	//ImageLinkAlphabets  []byte
	//ImageXSend          bool
	//ImageXSendHeader    string
	//Langs               []string
	LoginMaxRetries   int
	LoginFailedBlocks int
)

const (
	AppConfigPath = "conf/app.ini"
)

var (
	Cfg *goconfig.ConfigFile
)

func loadConfigVar() {
	AppName = beego.AppName

	AppHost = Cfg.MustValue("app", "app_host")
	AppUrl = Cfg.MustValue("app", "app_url")
	AppLogo = Cfg.MustValue("app", "app_logo")
	AppVersion = Cfg.MustValue("app", "app_ver")
	AvatarURL = Cfg.MustValue("app", "avatar_url")

	ActiveCodeLives = Cfg.MustInt("app", "acitve_code_live_hours")

	if ActiveCodeLives <= 0 {
		ActiveCodeLives = 12
	}
	ResetPwdCodeLives = Cfg.MustInt("app", "resetpwd_code_live_hours")
	if ResetPwdCodeLives <= 0 {
		ResetPwdCodeLives = 12
	}

	LoginRememberDays = Cfg.MustInt("app", "login_remember_days")

	LoginMaxRetries = Cfg.MustInt("app", "login_max_retries")
	if LoginMaxRetries <= 0 {
		LoginMaxRetries = 1
	}

	LoginFailedBlocks = Cfg.MustInt("app", "login_failed_blocks")
	if LoginFailedBlocks <= 0 {
		LoginFailedBlocks = 1
	}

}

func LoadConfig() {

	if fh, _ := os.OpenFile(AppConfigPath, os.O_RDONLY|os.O_CREATE, 0600); fh != nil {
		fh.Close()
	}

	var err error

	Cfg, err = goconfig.LoadConfigFile(AppConfigPath)

	Cfg.BlockMode = false

	if err != nil {
		panic("Fail to load configuration file: " + err.Error())
	}
	beego.AppName = Cfg.MustValue("app", "app_name")
	beego.HttpPort = Cfg.MustInt("app", "http_port")
	beego.RunMode = Cfg.MustValue("app", "run_mode")

	IsProMode = (beego.RunMode == "pro")
	if IsProMode {
		beego.SetLevel(beego.LevelInfo)
	}

	//session
	beego.SessionOn = true
	beego.SessionProvider = Cfg.MustValue("session", "session_provider")
	beego.SessionSavePath = Cfg.MustValue("session", "session_path")
	beego.SessionName = Cfg.MustValue("session", "session_name")

	//database
	driverName := Cfg.MustValue("orm", "driver_name")
	dataSource := Cfg.MustValue("orm", "data_source")
	maxIdle := Cfg.MustInt("orm", "max_idle_conn")
	maxOpen := Cfg.MustInt("orm", "max_open_conn")

	//set default database
	if _, err := os.Open(dataSource); err != nil && os.IsNotExist(err) {
		os.MkdirAll(path.Dir(dataSource), os.ModePerm)
		os.Create(dataSource)
	}
	orm.RegisterDataBase("default", driverName, dataSource, maxIdle, maxOpen)
	orm.RunCommand()
	orm.Debug = true
	err = orm.RunSyncdb("default", false, true)

	orm.Debug = Cfg.MustBool("orm", "debug_log")

	loadConfigVar()

}
