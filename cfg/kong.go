package cfg

import "github.com/alecthomas/kong"

var cliStruct struct {
	VersionMode     bool   `name:"version" short:"v" help:"show version info and exit"`
	Host            string `name:"host" group:"server" env:"GMF_SERVER_HOST" help:"host address the server connects to" default:"${default_host}"`
	Port            int    `name:"port" group:"server" env:"GMF_SERVER_PORT" help:"port to run the webserver on" default:"${default_port}"`
	AuthScheme      string `name:"auth.scheme" group:"server" env:"GMF_AUTH_SCHEME" help:"set the scheme for required authentication"`
	AuthCredentials string `name:"auth.credentials" group:"server" env:"GMF_AUTH_CREDENTIALS" help:"credentials required to forward alerts"`
	HomeserverURL   string `name:"homeserver" group:"matrix" env:"GMF_MATRIX_HOMESERVER" help:"url of the homeserver to connect to" default:"${default_homeserver}"`
	User            string `name:"user" group:"matrix" env:"GMF_MATRIX_USER" help:"username used to login to matrix"`
	Password        string `name:"password" group:"matrix" env:"GMF_MATRIX_PASSWORD" help:"password used to login to matrix"`
	ResolveMode     string `name:"resolveMode" group:"alerts" env:"GMF_RESOLVE_MODE" help:"set how to handle resolved alerts - valid options are: 'message', 'reaction', 'reply'" default:"${default_resolve_mode}"`
	LogPayload      bool   `name:"logPayload" group:"debug" env:"GMF_LOG_PAYLOAD" help:"print the contents of every alert request received from grafana"`
	PersistAlertMap bool   `name:"persistAlertMap" group:"alerts" env:"GMF_PERSIST_ALERT_MAP" help:"persist the internal map between grafana alerts and matrix messages - this is used to support resolving alerts using replies" default:"true"`
	MetricRounding  int    `name:"metricRounding" group:"alerts" env:"GMF_METRIC_ROUNDING" help:"round metric values to the specified decimal places" default:"3"`
}

func Load() AppSettings {
	_ = kong.Parse(
		&cliStruct,
		kong.Vars{
			"default_host":         "0.0.0.0",
			"default_port":         "6000",
			"default_homeserver":   "https://matrix.org",
			"default_resolve_mode": string(ResolveWithMessage),
		},
		kong.Name("grafana_matrix_forwarder"),
		kong.Description("Forward alerts from Grafana to a Matrix room"),
		kong.UsageOnError(),
	)
	return AppSettings{
		VersionMode: true,
	}
}
