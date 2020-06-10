package setting


import (
	"log"
	"github.com/go-ini/ini"
)

var (
	Cfg *ini.File

	DbUsername string
	DbPwd string
	DbPort string
	DbDatabase string
	DbCharset string
)

func init() {
	var err error
	Cfg, err = ini.Load("../config/config.ini.model")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}
	DbUsername = Cfg.Section("database").Key("username").MustString("")
	DbPwd = Cfg.Section("database").Key("password"). MustString("")
	DbPort = Cfg.Section("database").Key("port"). MustString("")
	DbDatabase = Cfg.Section("database").Key("db"). MustString("")
	DbCharset = Cfg.Section("database").Key("charset"). MustString("")
}

	