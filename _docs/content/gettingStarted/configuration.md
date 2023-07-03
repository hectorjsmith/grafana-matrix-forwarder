---
title: Configuration
weight: 30
---

The tool can be configured with either CLI flags and environment variables. CLI flags will take priority.

{{< hint info >}}
NOTE: all CLI flags can be ignored by providing the `-env` flag
{{< /hint >}}

## CLI Flags

```
$ grafana-matrix-forwarder -h

  -env
        ignore all other flags and read all configuration from environment variables
  -homeserver string
        url of the homeserver to connect to (default "matrix.org")
  -host string
        host address the server connects to (default "0.0.0.0")
  -logPayload
        print the contents of every alert request received from grafana
  -metricRounding int
        round metric values to the specified decimal places (set -1 to disable rounding) (default 3)
  -password string
        password used to login to matrix
  -persistAlertMap
        persist the internal map between grafana alerts and matrix messages - this is used to support resolving alerts using replies (default true)
  -port int
        port to run the webserver on (default 6000)
  -resolveMode string
        set how to handle resolved alerts - valid options are: 'message', 'reaction', 'reply' (default "message")
  -user string
        username used to login to matrix
  -version
        show version info and exit
``` 

## Environment Variables

The following environment variables should be set to configure how the forwarder container runs.
These environment variables map directly to the CLI parameters of the application.

| Name | Required | Description |
|------|----------|-------------|
| `GMF_MATRIX_USER` | X | Username used to login to matrix |
| `GMF_MATRIX_PASSWORD` | X | Password used to login to matrix |
| `GMF_MATRIX_HOMESERVER` | X | URL of the matrix homeserver to connect to |
| `GMF_SERVER_HOST` | | Host address the server connects to (defaults to "0.0.0.0") |
| `GMF_SERVER_PORT` | | Port to run the webserver on (default 6000) |
| `GMF_RESOLVE_MODE` | | Set how to handle resolved alerts - valid options are: 'message', 'reaction', and 'reply' |
| `GMF_LOG_PAYLOAD` | | Set to any value to print the contents of every alert request received from grafana (disabled if set to "no" or "false") |
| `GMF_METRIC_ROUNDING` | | Set the number of decimal places to round metric values to (-1 to disable all rounding) |
| `GMF_PERSIST_ALERT_MAP` | | Persist the internal map between grafana alerts and matrix messages - this is used to support resolving alerts using replies (defaults to "true") |
