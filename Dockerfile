FROM alpine

# Create main app folder to run from
WORKDIR /app

# Copy compiled binary to release image
# (must build the binary before running docker build)
COPY grafana-matrix-forwarder /app/grafana-matrix-forwarder

ENTRYPOINT ["/app/grafana-matrix-forwarder"]
