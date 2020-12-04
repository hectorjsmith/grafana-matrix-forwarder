package cfg

import (
	"flag"
	"fmt"
	"os"
)

// AppSettings includes all application parameters
type AppSettings struct {
	VersionMode   bool
	UserID        string
	UserPassword  string
	HomeserverURL string
	ServerHost    string
	ServerPort    int
	LogPayload    bool
}

const minServerPort = 1000
const maxServerPort = 65535

// Parse the AppSettings data from the command line
func Parse() AppSettings {
	appSettings := AppSettings{}
	flag.BoolVar(&appSettings.VersionMode, "version", false, "show version info and exit")
	flag.StringVar(&appSettings.UserID, "user", "", "username used to login to matrix")
	flag.StringVar(&appSettings.UserPassword, "password", "", "password used to login to matrix")
	flag.StringVar(&appSettings.HomeserverURL, "homeserver", "matrix.org", "url of the homeserver to connect to")
	flag.StringVar(&appSettings.ServerHost, "host", "0.0.0.0", "host address the server connects to")
	flag.IntVar(&appSettings.ServerPort, "port", 6000, "port to run the webserver on")
	flag.BoolVar(&appSettings.LogPayload, "logPayload", false, "print the contents of every alert request received from grafana")

	flag.Parse()
	appSettings.validateFlags()
	return appSettings
}

func (settings AppSettings) validateFlags() {
	var flagsValid = true
	if !settings.VersionMode {
		if settings.UserID == "" {
			fmt.Println("missing flag 'user'")
			flagsValid = false
		}
		if settings.UserPassword == "" {
			fmt.Println("missing flag 'password'")
			flagsValid = false
		}
		if settings.ServerPort < minServerPort || settings.ServerPort > maxServerPort {
			fmt.Printf("invalid server port, must be within %d and %d (found %d)\n",
				minServerPort, maxServerPort, settings.ServerPort)
			flagsValid = false
		}
	}
	if !flagsValid {
		flag.Usage()
		os.Exit(1)
	}
}
