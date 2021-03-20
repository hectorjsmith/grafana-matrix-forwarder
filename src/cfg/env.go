package cfg

import (
	"log"
	"os"
	"strconv"
	"strings"
)

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
	if envValue, envExists = os.LookupEnv("GMF_METRIC_ROUNDING"); envExists {
		intValue, err := strconv.Atoi(envValue)
		if err != nil {
			log.Printf("ignoring invalid metric rounding number: %s", envValue)
		} else {
			settings.MetricRounding = intValue
		}
	}
	if envValue, envExists = os.LookupEnv("GMF_LOG_PAYLOAD"); envExists {
		lowerEnvValue := strings.ToLower(envValue)
		if envValue != "" && lowerEnvValue != "false" && lowerEnvValue != "no" {
			settings.LogPayload = true
		}
	}
}
