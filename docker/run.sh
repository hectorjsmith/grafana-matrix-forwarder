#/bin/sh

# Print version to logs for debugging purposes
/app/grafana-matrix-forwarder -version

# Start the forwarder (use exec to support graceful shutdown)
# Inspired by: https://akomljen.com/stopping-docker-containers-gracefully/
exec /app/grafana-matrix-forwarder -env
