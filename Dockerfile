FROM alpine

# Create main app folder to run from
WORKDIR /app

# Copy compiled binary to release image
# (must build the binary before running docker build)
COPY grafana_matrix_forwarder /app/grafana_matrix_forwarder

ENTRYPOINT ["/app/grafana_matrix_forwarder"]
