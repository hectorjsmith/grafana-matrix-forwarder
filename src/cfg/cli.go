package cfg

import (
	"flag"
	"fmt"
)

const (
	versionFlagName         = "version"
	userFlagName            = "user"
	passwordFlagName        = "password"
	homeServerFlagName      = "homeserver"
	hostFlagName            = "host"
	portFlagName            = "port"
	metricRoundingFlagName  = "metricRounding"
	logPayloadFlagName      = "logPayload"
	resolveModeFlagName     = "resolveMode"
	envFlagName             = "env"
	persistAlertMapFlagName = "persistAlertMap"
	authSchemeFlagName      = "auth.scheme"
	authCredentialsFlagName = "auth.credentials"
)

func (settings *AppSettings) updateSettingsFromCommandLine() {
	versionFlag := flag.Bool(versionFlagName, false, "show version info and exit")
	userFlag := flag.String(userFlagName, "", "username used to login to matrix")
	passwordFlag := flag.String(passwordFlagName, "", "password used to login to matrix")
	homeServerFlag := flag.String(homeServerFlagName, defaultHomeServerUrl, "url of the homeserver to connect to")
	hostFlag := flag.String(hostFlagName, defaultServerHost, "host address the server connects to")
	portFlag := flag.Int(portFlagName, defaultServerPort, "port to run the webserver on")
	roundingFlag := flag.Int(metricRoundingFlagName, defaultMetricRounding, "round metric values to the specified decimal places (set -1 to disable rounding)")
	logPayloadFlag := flag.Bool(logPayloadFlagName, false, "print the contents of every alert request received from grafana")
	persistAlertMapFlag := flag.Bool(persistAlertMapFlagName, defaultPersistAlertMap,
		"persist the internal map between grafana alerts and matrix messages - this is used to support resolving alerts using replies")
	authSchemeFlag := flag.String(authSchemeFlagName, "", "set the scheme for required authentication")
	authCredentialsFlag := flag.String(authCredentialsFlagName, "", "credentials required to forward alerts")

	var resolveModeStr string
	flag.StringVar(&resolveModeStr, resolveModeFlagName, string(defaultResolveMode),
		fmt.Sprintf("set how to handle resolved alerts - valid options are: '%s', '%s', '%s'", ResolveWithMessage, ResolveWithReaction, ResolveWithReply))

	var envFlag bool
	flag.BoolVar(&envFlag, envFlagName, false, "ignore all other flags and read all configuration from environment variables")

	flag.Parse()
	if !envFlag {
		settings.VersionMode = *versionFlag
		if wasCliFlagProvided(userFlagName) {
			settings.UserID = *userFlag
		}
		if wasCliFlagProvided(passwordFlagName) {
			settings.UserPassword = *passwordFlag
		}
		if wasCliFlagProvided(homeServerFlagName) {
			settings.HomeserverURL = *homeServerFlag
		}
		if wasCliFlagProvided(hostFlagName) {
			settings.ServerHost = *hostFlag
		}
		if wasCliFlagProvided(portFlagName) {
			settings.ServerPort = *portFlag
		}
		if wasCliFlagProvided(logPayloadFlagName) {
			settings.LogPayload = *logPayloadFlag
		}
		if wasCliFlagProvided(metricRoundingFlagName) {
			settings.MetricRounding = *roundingFlag
		}
		if wasCliFlagProvided(resolveModeFlagName) {
			settings.setResolveMode(resolveModeStr)
		}
		if wasCliFlagProvided(persistAlertMapFlagName) {
			settings.PersistAlertMap = *persistAlertMapFlag
		}
		if wasCliFlagProvided(*authSchemeFlag) {
			settings.AuthScheme = *authSchemeFlag
		}
		if wasCliFlagProvided(*authCredentialsFlag) {
			settings.AuthCredentials = *authCredentialsFlag
		}
	}
}

func wasCliFlagProvided(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}
