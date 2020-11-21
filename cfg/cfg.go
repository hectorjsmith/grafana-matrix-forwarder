package cfg

import (
	"flag"
	"fmt"
	"os"
)

type AppSettings struct {
	VersionMode   bool
	UserId        string
	UserPassword  string
	HomeserverUrl string
}

func Parse() AppSettings {
	appSettings := AppSettings{}
	flag.BoolVar(&appSettings.VersionMode, "version", false, "show version info and exit")
	flag.StringVar(&appSettings.UserId, "user", "", "username used to login to matrix")
	flag.StringVar(&appSettings.UserPassword, "password", "", "password used to login to matrix")
	flag.StringVar(&appSettings.HomeserverUrl, "homeserver", "matrix.org", "url of the homeserver to connect to")

	flag.Parse()
	appSettings.validateFlags()
	return appSettings
}

func (settings AppSettings) validateFlags() {
	var flagsValid = true
	if !settings.VersionMode {
		if settings.UserId == "" {
			fmt.Println("missing flag 'user'")
			flagsValid = false
		}
		if settings.UserPassword == "" {
			fmt.Println("missing flag 'password'")
			flagsValid = false
		}
	}
	if !flagsValid {
		flag.Usage()
		os.Exit(1)
	}
}
