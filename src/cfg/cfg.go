package cfg

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// ResolveMode determines how the application will handle resolved alerts
type ResolveMode string

// AppSettings includes all application parameters
type AppSettings struct {
	VersionMode   bool
	UserID        string
	UserPassword  string
	HomeserverURL string
	ServerHost    string
	ServerPort    int
	LogPayload    bool
	ResolveMode   ResolveMode
}

const (
	ResolveWithReaction ResolveMode = "reaction"
	ResolveWithMessage  ResolveMode = "message"
	ResolveWithReply    ResolveMode = "reply"
	minServerPort                   = 1000
	maxServerPort                   = 65535
)

// Parse the AppSettings data from the command line
func Parse() AppSettings {
	appSettings := &AppSettings{}
	appSettings.setDefaults()
	appSettings.updateSettingsFromEnvironment()
	appSettings.updateSettingsFromCommandLine()

	appSettings.validateConfiguration()
	return *appSettings
}

func (settings *AppSettings) setDefaults() {
	settings.ServerPort = 6000
	settings.ServerHost = "0.0.0.0"
	settings.ResolveMode = ResolveWithMessage
}

func (settings *AppSettings) updateSettingsFromEnvironment() {
	var envValue string
	var envExists bool

	if envValue, envExists = os.LookupEnv("GMF_MATRIX_USER"); envExists {
		settings.UserID = envValue
	}
	if envValue, envExists = os.LookupEnv("GMF_MATRIX_PASSWORD"); envExists {
		settings.UserPassword = envValue
	}
	if envValue, envExists = os.LookupEnv("GMF_MATRIX_HOMESERVER"); envExists {
		settings.HomeserverURL = envValue
	}
	if envValue, envExists = os.LookupEnv("GMF_SERVER_HOST"); envExists {
		settings.ServerHost = envValue
	}
	if envValue, envExists = os.LookupEnv("GMF_SERVER_PORT"); envExists {
		intValue, err := strconv.Atoi(envValue)
		if err != nil {
			log.Printf("ignoring invalid port number: %s", envValue)
		} else {
			settings.ServerPort = intValue
		}
	}
	if envValue, envExists = os.LookupEnv("GMF_RESOLVE_MODE"); envExists {
		settings.setResolveMode(envValue)
	}
	if envValue, envExists = os.LookupEnv("GMF_LOG_PAYLOAD"); envExists {
		lowerEnvValue := strings.ToLower(envValue)
		if envValue != "" && lowerEnvValue != "false" && lowerEnvValue != "no" {
			settings.LogPayload = true
		}
	}
}

func (settings *AppSettings) updateSettingsFromCommandLine() {
	versionFlag := flag.Bool("version", false, "show version info and exit")
	userFlag := flag.String("user", "", "username used to login to matrix")
	passwordFlag := flag.String("password", "", "password used to login to matrix")
	homeserverFlag := flag.String("homeserver", "matrix.org", "url of the homeserver to connect to")
	hostFlag := flag.String("host", "0.0.0.0", "host address the server connects to")
	portFlag := flag.Int("port", 6000, "port to run the webserver on")
	logPayloadFlag := flag.Bool("logPayload", false, "print the contents of every alert request received from grafana")

	var resolveModeStr string
	flag.StringVar(&resolveModeStr, "resolveMode", string(ResolveWithMessage),
		fmt.Sprintf("set how to handle resolved alerts - valid options are: '%s', '%s', '%s'", ResolveWithMessage, ResolveWithReaction, ResolveWithReply))

	var envFlag bool
	flag.BoolVar(&envFlag, "env", false, "ignore all other flags and read all configuration from environment variables")

	flag.Parse()
	if !envFlag {
		settings.VersionMode = *versionFlag
		settings.UserID = *userFlag
		settings.UserPassword = *passwordFlag
		settings.HomeserverURL = *homeserverFlag
		settings.ServerHost = *hostFlag
		settings.ServerPort = *portFlag
		settings.LogPayload = *logPayloadFlag
		settings.setResolveMode(resolveModeStr)
	}
}

func (settings *AppSettings) setResolveMode(resolveModeStr string) {
	resolveModeStrLower := strings.ToLower(resolveModeStr)
	if resolveModeStrLower == string(ResolveWithReaction) {
		settings.ResolveMode = ResolveWithReaction
	} else if resolveModeStrLower == string(ResolveWithMessage) {
		settings.ResolveMode = ResolveWithMessage
	} else if resolveModeStrLower == string(ResolveWithReply) {
		settings.ResolveMode = ResolveWithReply
	} else {
		log.Printf("invalid resolve mode provided (%s) - defaulting to %s", resolveModeStr, ResolveWithMessage)
		settings.ResolveMode = ResolveWithMessage
	}
}

func (settings *AppSettings) validateConfiguration() {
	var flagsValid = true
	if !settings.VersionMode {
		if settings.UserID == "" {
			fmt.Println("missing parameter '-user'")
			flagsValid = false
		}
		if settings.UserPassword == "" {
			fmt.Println("missing flag '-password'")
			flagsValid = false
		}
		if settings.HomeserverURL == "" {
			fmt.Println("missing flag '-homeserver'")
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
