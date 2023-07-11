---
title: Configuration
weight: 30
---

The tool can be configured with either CLI flags and environment variables. CLI flags will take priority.

## CLI Flags

```
Usage: grafana-matrix-forwarder

Forward alerts from Grafana to a Matrix room

Flags:
  -h, --help       Show context-sensitive help.
  -v, --version    Show version info and exit

🔌 Server
  Configure the server used to receive alerts from Grafana

  --host="0.0.0.0"             Host address the server connects to ($GMF_SERVER_HOST)
  --port=6000                  Port to run the webserver on ($GMF_SERVER_PORT)
  --auth.scheme=STRING         Set the scheme for required authentication - valid options are: bearer ($GMF_AUTH_SCHEME)
  --auth.credentials=STRING    Credentials required to forward alerts ($GMF_AUTH_CREDENTIALS)

💬 Matrix
  How to connect to Matrix to forward alerts

  --homeserver="matrix.org"    URL of the homeserver to connect to ($GMF_MATRIX_HOMESERVER)
  --user=STRING                Username used to login to matrix ($GMF_MATRIX_USER)
  --password=STRING            Password used to login to matrix ($GMF_MATRIX_PASSWORD)
  --token=STRING               Auth token used to authenticate with matrix ($GMF_MATRIX_TOKEN)

❗ Alerts
  Configuration for the alerts themselves

  --resolveMode="message"    Set how to handle resolved alerts - valid options are: message, reaction, reply ($GMF_RESOLVE_MODE)
  --[no-]persistAlertMap     Persist the internal map between grafana alerts and matrix messages - this is used to support resolving alerts using replies ($GMF_PERSIST_ALERT_MAP)
  --metricRounding=3         Round metric values to the specified decimal places ($GMF_METRIC_ROUNDING)

❔ Debug
  Options to help debugging issues

  --logPayload    Print the contents of every alert request received from grafana ($GMF_LOG_PAYLOAD)

🔻 Deprecated
  Flags that have been deprecated and should no longer be used

  --env    No longer has any effect
``` 

## Environment Variables

The following environment variables should be set to configure how the forwarder container runs.
These environment variables map directly to the CLI parameters of the application.

| Name | Required | Description |
|------|----------|-------------|
| `GMF_MATRIX_USER` | X | Username used to login to matrix |
| `GMF_MATRIX_PASSWORD` | X | Password used to login to matrix |
| `GMF_MATRIX_TOKEN` | X | Token to be used to authenticate against the matrix server |
| `GMF_MATRIX_HOMESERVER` | X | URL of the matrix homeserver to connect to |
| `GMF_SERVER_HOST` | | Host address the server connects to (defaults to "0.0.0.0") |
| `GMF_SERVER_PORT` | | Port to run the webserver on (default 6000) |
| `GMF_RESOLVE_MODE` | | Set how to handle resolved alerts - valid options are: 'message', 'reaction', and 'reply' |
| `GMF_LOG_PAYLOAD` | | Set to any value to print the contents of every alert request received from grafana (disabled if blank or set to "false") |
| `GMF_METRIC_ROUNDING` | | Set the number of decimal places to round metric values to (-1 to disable all rounding) |
| `GMF_PERSIST_ALERT_MAP` | | Persist the internal map between grafana alerts and matrix messages - this is used to support resolving alerts using replies (defaults to "true") |

{{< hint info >}}
Either the token or the password must be set, not both.
{{< /hint >}}

## Getting a Matrix Token

To get a matrix token you can run the following command:

```
curl -XPOST -d '{"type": "m.login.password", "identifier": {"user": "myusername", "type": "m.id.user"}, "password": "mypassword"}' "https://matrix.org/_matrix/client/r0/login"
```

{{< hint info >}}
Don't forget to change the `user`, `password`, and homeserver URL for the account you want to use.
{{< /hint >}}

The output will look something like:

```
{"user_id":"@mysername:matrix.org","access_token":"xxxxx","home_server":"matrix.org","device_id":"something","well_known":{"m.homeserver":{"base_url":"https://matrix-client.matrix.org/"}}}
```

Then just copy the value for `access_token` and use it.
