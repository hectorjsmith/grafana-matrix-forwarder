package cfg

import "strings"

const (
	minServerPort = 1000
	maxServerPort = 65535
)

func validateFlags() (bool, []string) {
	var flagsValid = true
	var messages = []string{}
	if !cli.VersionMode {
		if cli.Env {
			messages = append(messages, "warn: the env flag has been deprecated and no longer has any function")
		}
		if cli.User == "" {
			messages = append(messages, "error: matrix username must not be blank")
			flagsValid = false
		}
		passwordSet := cli.Password != ""
		tokenSet := cli.Token != ""
		if passwordSet == tokenSet {
			messages = append(messages, "error: must set either password or token (only one, not both)")
			flagsValid = false
		}
		if cli.HomeserverURL == "" {
			messages = append(messages, "error: matrix homeserver url must not be blank")
			flagsValid = false
		}
		if cli.Port < minServerPort || cli.Port > maxServerPort {
			messages = append(messages, "error: invalid server port selected")
			flagsValid = false
		}
		if (cli.AuthScheme == "") != (cli.AuthCredentials == "") {
			messages = append(messages, "error: invalid auth setup - both scheme and credentials should be set")
			flagsValid = false
		}
		if strings.ToLower(cli.AuthScheme) != "" && strings.ToLower(cli.AuthScheme) != "bearer" {
			messages = append(messages, "error: unsupported auth scheme selected")
			flagsValid = false
		}
		_, err := ToResolveMode(cli.ResolveMode)
		if err != nil {
			messages = append(messages, "error: invalid resolve mode selected")
			flagsValid = false
		}
	}
	return flagsValid, messages
}
