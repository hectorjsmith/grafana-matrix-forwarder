package cfg

import (
	"flag"
	"fmt"
	"os"
)

var (
	VersionMode   bool
	UserId        string
	UserPassword  string
	HomeserverUrl string
)

func Parse() {
	flag.BoolVar(&VersionMode, "version", false, "show version info and exit")
	flag.StringVar(&UserId, "user", "", "username used to login to matrix")
	flag.StringVar(&UserPassword, "password", "", "password used to login to matrix")
	flag.StringVar(&HomeserverUrl, "homeserver", "matrix.org", "url of the homeserver to connect to")

	flag.Parse()
	validateFlags()
}

func validateFlags() {
	var flagsValid = true
	if !VersionMode {
		if UserId == "" {
			fmt.Println("missing flag 'user'")
			flagsValid = false
		}
		if UserPassword == "" {
			fmt.Println("missing flag 'password'")
			flagsValid = false
		}
	}
	if !flagsValid {
		flag.Usage()
		os.Exit(1)
	}
}
