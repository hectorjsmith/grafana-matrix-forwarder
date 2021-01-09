#/bin/sh

serverHostOpt="-host $GMF_SERVER_HOST"
if [ -z $GMF_SERVER_HOST ]; then
  serverHostOpt=""
fi
portHostOpt="-port $GMF_SERVER_PORT"
if [ -z $GMF_SERVER_PORT ]; then
  portHostOpt=""
fi
resolveModeOpt="-resolveMode $GMF_RESOLVE_MODE"
if [ -z $GMF_RESOLVE_MODE ]; then
  resolveModeOpt=""
fi
logPayloadOpt="-logPayload"
if [ -z $GMF_LOG_PAYLOAD ] || [ $GMF_LOG_PAYLOAD = "no" ] || [ $GMF_LOG_PAYLOAD = "false" ]; then
  logPayloadOpt=""
fi

# Print version to logs for debugging purposes
/app/grafana-matrix-forwarder -version

# Start the forwarder (use exec to support graceful shutdown)
# Inspired by: https://akomljen.com/stopping-docker-containers-gracefully/
exec /app/grafana-matrix-forwarder \
    -user $GMF_MATRIX_USER \
    -password $GMF_MATRIX_PASSWORD \
    -homeserver $GMF_MATRIX_HOMESERVER \
    $portHostOpt \
    $serverHostOpt \
    $resolveModeOpt \
    $logPayloadOpt
