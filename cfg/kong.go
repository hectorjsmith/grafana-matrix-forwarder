package cfg

import (
	"fmt"
	"os"
	"strings"

	"github.com/alecthomas/kong"
)

var cli struct {
	VersionMode     bool   `name:"version" short:"v" help:"Show version info and exit"`
	Host            string `name:"host" group:"server" env:"GMF_SERVER_HOST" help:"Host address the server connects to" default:"${default_host}"`
	Port            int    `name:"port" group:"server" env:"GMF_SERVER_PORT" help:"Port to run the webserver on" default:"${default_port}"`
	AuthScheme      string `name:"auth.scheme" group:"server" env:"GMF_AUTH_SCHEME" help:"Set the scheme for required authentication - valid options are: ${auth_scheme_options}"`
	AuthCredentials string `name:"auth.credentials" group:"server" env:"GMF_AUTH_CREDENTIALS" help:"Credentials required to forward alerts"`
	HomeserverURL   string `name:"homeserver" group:"matrix" env:"GMF_MATRIX_HOMESERVER" help:"URL of the homeserver to connect to" default:"${default_homeserver}"`
	User            string `name:"user" group:"matrix" env:"GMF_MATRIX_USER" help:"Username used to login to matrix"`
	Password        string `name:"password" group:"matrix" env:"GMF_MATRIX_PASSWORD" help:"Password used to login to matrix"`
	ResolveMode     string `name:"resolveMode" group:"alerts" env:"GMF_RESOLVE_MODE" help:"Set how to handle resolved alerts - valid options are: ${resolve_mode_options}" default:"${default_resolve_mode}"`
	PersistAlertMap bool   `name:"persistAlertMap" group:"alerts" env:"GMF_PERSIST_ALERT_MAP" help:"Persist the internal map between grafana alerts and matrix messages - this is used to support resolving alerts using replies" default:"${default_persist_alert_map}" negatable:"true"`
	MetricRounding  int    `name:"metricRounding" group:"alerts" env:"GMF_METRIC_ROUNDING" help:"Round metric values to the specified decimal places" default:"${default_metric_rounding}"`
	LogPayload      bool   `name:"logPayload" group:"debug" env:"GMF_LOG_PAYLOAD" help:"Print the contents of every alert request received from grafana"`

	Env bool `name:"env" group:"deprecated" help:"No longer has any effect"`
}

func Parse() AppSettings {
	ctx := kong.Parse(
		&cli,
		kong.Vars{
			"default_host":              "0.0.0.0",
			"default_port":              "6000",
			"default_homeserver":        "matrix.org",
			"default_metric_rounding":   "3",
			"default_persist_alert_map": "true",
			"default_resolve_mode":      string(ResolveWithMessage),
			"resolve_mode_options":      strings.Join(AvailableResolveModesStr(), ", "),
			"auth_scheme_options":       "bearer",
		},
		kong.Name("grafana_matrix_forwarder"),
		kong.Description("Forward alerts from Grafana to a Matrix room"),
		kong.UsageOnError(),
		kong.ExplicitGroups(
			[]kong.Group{
				{
					Key:         "server",
					Title:       "🔌 Server",
					Description: "Configure the server used to receive alerts from Grafana",
				},
				{
					Key:         "matrix",
					Title:       "💬 Matrix",
					Description: "How to connect to Matrix to forward alerts",
				},
				{
					Key:         "alerts",
					Title:       "❗ Alerts",
					Description: "Configuration for the alerts themselves",
				},
				{
					Key:         "debug",
					Title:       "❔ Debug",
					Description: "Options to help debugging issues",
				},
				{
					Key:         "deprecated",
					Title:       "🔻 Deprecated",
					Description: "Flags that have been deprecated and should no longer be used",
				},
			},
		),
	)

	flagsValid, messages := validateFlags()
	if len(messages) > 0 {
		if !flagsValid {
			ctx.PrintUsage(false)
		}
		fmt.Println()
		for i := 0; i < len(messages); i++ {
			fmt.Println(messages[i])
		}
	}
	if !flagsValid {
		os.Exit(1)
	}

	resolveMode, err := ToResolveMode(cli.ResolveMode)
	if err != nil {
		// This should have already been validated
		panic(err)
	}
	return AppSettings{
		VersionMode:     cli.VersionMode,
		ServerHost:      cli.Host,
		ServerPort:      cli.Port,
		AuthScheme:      cli.AuthScheme,
		AuthCredentials: cli.AuthCredentials,
		HomeserverURL:   cli.HomeserverURL,
		UserID:          cli.User,
		UserPassword:    cli.Password,
		ResolveMode:     resolveMode,
		LogPayload:      cli.LogPayload,
		PersistAlertMap: cli.PersistAlertMap,
		MetricRounding:  cli.MetricRounding,
	}
}
