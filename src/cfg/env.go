package cfg

import (
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	userEnvName            = "GMF_MATRIX_USER"
	passwordEnvName        = "GMF_MATRIX_PASSWORD"
	homeServerEnvName      = "GMF_MATRIX_HOMESERVER"
	hostEnvName            = "GMF_SERVER_HOST"
	portEnvName            = "GMF_SERVER_PORT"
	resolveModeEnvName     = "GMF_RESOLVE_MODE"
	metricRoundingEnvName  = "GMF_METRIC_ROUNDING"
	logPayloadEnvName      = "GMF_LOG_PAYLOAD"
	persistAlertMapEnvName = "GMF_PERSIST_ALERT_MAP"
	authSchemeEnvName      = "GMF_AUTH_SCHEME"
	authCredentialsEnvName = "GMF_AUTH_CREDENTIALS"
)

func (settings *AppSettings) updateSettingsFromEnvironment() {
	var envValue string
	var envExists bool

	if envValue, envExists = os.LookupEnv(userEnvName); envExists {
		settings.UserID = envValue
	}
	if envValue, envExists = os.LookupEnv(passwordEnvName); envExists {
		settings.UserPassword = envValue
	}
	if envValue, envExists = os.LookupEnv(homeServerEnvName); envExists {
		settings.HomeserverURL = envValue
	}
	if envValue, envExists = os.LookupEnv(hostEnvName); envExists {
		settings.ServerHost = envValue
	}
	if envValue, envExists = os.LookupEnv(portEnvName); envExists {
		intValue, err := strconv.Atoi(envValue)
		if err != nil {
			log.Printf("ignoring invalid port number: %s", envValue)
		} else {
			settings.ServerPort = intValue
		}
	}
	if envValue, envExists = os.LookupEnv(resolveModeEnvName); envExists {
		settings.setResolveMode(envValue)
	}
	if envValue, envExists = os.LookupEnv(metricRoundingEnvName); envExists {
		intValue, err := strconv.Atoi(envValue)
		if err != nil {
			log.Printf("ignoring invalid metric rounding number: %s", envValue)
		} else {
			settings.MetricRounding = intValue
		}
	}
	if envValue, envExists = os.LookupEnv(logPayloadEnvName); envExists {
		lowerEnvValue := strings.ToLower(envValue)
		if envValue != "" && lowerEnvValue != "false" && lowerEnvValue != "no" {
			settings.LogPayload = true
		}
	}
	if envValue, envExists = os.LookupEnv(persistAlertMapEnvName); envExists {
		lowerEnvValue := strings.ToLower(envValue)
		if envValue != "" && lowerEnvValue != "false" && lowerEnvValue != "no" {
			settings.PersistAlertMap = true
		}
	}
	if envValue, envExists = os.LookupEnv(authSchemeEnvName); envExists {
		settings.AuthScheme = envValue
	}
	if envValue, envExists = os.LookupEnv(authCredentialsEnvName); envExists {
		settings.AuthCredentials = envValue
	}
}
