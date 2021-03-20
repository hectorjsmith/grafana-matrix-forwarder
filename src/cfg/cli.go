package cfg

import (
	"flag"
	"fmt"
)

func (settings *AppSettings) updateSettingsFromCommandLine() {
	versionFlag := flag.Bool("version", false, "show version info and exit")
	userFlag := flag.String("user", "", "username used to login to matrix")
	passwordFlag := flag.String("password", "", "password used to login to matrix")
	homeserverFlag := flag.String("homeserver", "matrix.org", "url of the homeserver to connect to")
	hostFlag := flag.String("host", "0.0.0.0", "host address the server connects to")
	portFlag := flag.Int("port", 6000, "port to run the webserver on")
	roundingFlag := flag.Int("metricRounding", 3, "round metric values to the specified decimal places (set -1 to disable rounding)")
	logPayloadFlag := flag.Bool("logPayload", false, "print the contents of every alert request received from grafana")

	var resolveModeStr string
	flag.StringVar(&resolveModeStr, "resolveMode", string(ResolveWithMessage),
		fmt.Sprintf("set how to handle resolved alerts - valid options are: '%s', '%s', '%s'", ResolveWithMessage, ResolveWithReaction, ResolveWithReply))

	var envFlag bool
	flag.BoolVar(&envFlag, "env", false, "ignore all other flags and read all configuration from environment variables")

	flag.Parse()
	flag.CommandLine.Lookup("version")
	if !envFlag {
		settings.VersionMode = *versionFlag
		settings.UserID = *userFlag
		settings.UserPassword = *passwordFlag
		settings.HomeserverURL = *homeserverFlag
		settings.ServerHost = *hostFlag
		settings.ServerPort = *portFlag
		settings.LogPayload = *logPayloadFlag
		settings.MetricRounding = *roundingFlag
		settings.setResolveMode(resolveModeStr)
	}
}
