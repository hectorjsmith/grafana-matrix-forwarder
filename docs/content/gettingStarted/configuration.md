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

- `GMF_MATRIX_USER` (required) - Username used to login to matrix
- `GMF_MATRIX_PASSWORD` (required) - Password used to login to matrix
- `GMF_MATRIX_HOMESERVER` (required) - URL of the matrix homeserver to connect to
- `GMF_SERVER_HOST` (optional) - Host address the server connects to (defaults to "0.0.0.0")
- `GMF_SERVER_PORT` (optional) - Port to run the webserver on (default 6000)
- `GMF_RESOLVE_MODE` (optional) - Set how to handle resolved alerts - valid options are: 'message', 'reaction', and 'reply'
- `GMF_LOG_PAYLOAD` (optional) - Set to any value to print the contents of every alert request received from grafana (disabled if set to "no" or "false")
- `GMF_METRIC_ROUNDING` (optional) - Set the number of decimal places to round metric values to (-1 to disable all rounding)
