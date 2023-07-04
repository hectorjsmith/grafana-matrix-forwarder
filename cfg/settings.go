package cfg

import (
	"fmt"
	"strings"
)

// ResolveMode determines how the application will handle resolved alerts
type ResolveMode string

const (
	ResolveWithReaction ResolveMode = "reaction"
	ResolveWithMessage  ResolveMode = "message"
	ResolveWithReply    ResolveMode = "reply"
)

// AppSettings includes all application parameters
type AppSettings struct {
	VersionMode     bool
	UserID          string
	UserPassword    string
	HomeserverURL   string
	ServerHost      string
	MetricRounding  int
	ServerPort      int
	LogPayload      bool
	ResolveMode     ResolveMode
	PersistAlertMap bool
	AuthScheme      string
	AuthCredentials string
}

func ToResolveMode(raw string) (ResolveMode, error) {
	resolveModeStrLower := strings.ToLower(raw)
	if resolveModeStrLower == string(ResolveWithReaction) {
		return ResolveWithReaction, nil
	} else if resolveModeStrLower == string(ResolveWithMessage) {
		return ResolveWithMessage, nil
	} else if resolveModeStrLower == string(ResolveWithReply) {
		return ResolveWithReply, nil
	}
	return ResolveWithMessage, fmt.Errorf("invalid resolve mode '%s'", raw)
}

func AvailableResolveModes() []ResolveMode {
	return []ResolveMode{
		ResolveWithMessage,
		ResolveWithReaction,
		ResolveWithReply,
	}
}

func AvailableResolveModesStr() []string {
	modes := AvailableResolveModes()
	modesStr := make([]string, len(modes))
	for i, m := range modes {
		modesStr[i] = string(m)
	}
	return modesStr
}
