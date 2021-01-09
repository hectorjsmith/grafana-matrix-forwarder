FROM alpine

WORKDIR /app
COPY dist/grafana-matrix-forwarder_linux_amd64/grafana-matrix-forwarder /app
COPY docker/run.sh /app

RUN chmod +x /app/*

ENTRYPOINT /app/run.sh
