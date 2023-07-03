package server

import (
	"context"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"grafana-matrix-forwarder/cfg"
	"grafana-matrix-forwarder/matrix"
	"grafana-matrix-forwarder/server/metrics"
	"grafana-matrix-forwarder/server/v0"
	"grafana-matrix-forwarder/server/v1"
	"grafana-matrix-forwarder/service"

	"log"
	"net/http"
	"time"
)

// Server data structure that holds data necessary for the web server to function
type Server struct {
	ctx               context.Context
	matrixWriteCloser matrix.WriteCloser
	appSettings       cfg.AppSettings
	alertForwarder    service.Forwarder
	metricsCollector  *metrics.Collector
}

// BuildServer builds a Server instance based on the provided context.Context, a matrix.WriteCloser, and the cfg.AppSettings
func BuildServer(ctx context.Context, matrixWriteCloser matrix.WriteCloser, appSettings cfg.AppSettings) Server {
	return Server{
		ctx:               ctx,
		matrixWriteCloser: matrixWriteCloser,
		appSettings:       appSettings,
		alertForwarder:    service.NewForwarder(appSettings, matrixWriteCloser.GetWriter()),
		metricsCollector:  metrics.NewCollector(),
	}
}

// Start the web server and listen for incoming requests
func (server Server) Start() (err error) {
	log.Print("starting webserver ...")
	log.Printf("resolve mode set to: %s", server.appSettings.ResolveMode)
	mux := http.NewServeMux()
	mux.Handle("/api/v0/forward", http.HandlerFunc(
		func(response http.ResponseWriter, request *http.Request) {
			server.HandleGrafanaAlert(&v0.Handler{}, response, request)
		},
	))
	mux.Handle("/api/v1/standard", http.HandlerFunc(
		func(response http.ResponseWriter, request *http.Request) {
			server.HandleGrafanaAlert(&v0.Handler{}, response, request)
		},
	))
	mux.Handle("/api/v1/unified", http.HandlerFunc(
		func(response http.ResponseWriter, request *http.Request) {
			server.HandleGrafanaAlert(&v1.Handler{}, response, request)
		},
	))
	mux.Handle("/metrics", promhttp.Handler())

	prometheus.MustRegister(server.metricsCollector)
	serverAddr := fmt.Sprintf("%s:%d", server.appSettings.ServerHost, server.appSettings.ServerPort)
	srv := &http.Server{
		Addr:    serverAddr,
		Handler: mux,
	}

	go func() {
		if err = srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server error: %+s\n", err)
		}
	}()

	log.Printf("webserver listening at %s", serverAddr)
	log.Print("ready")

	<-server.ctx.Done()

	log.Printf("shutting down ...")

	ctxShutDown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()

	if err = server.matrixWriteCloser.Close(); err != nil {
		log.Fatalf("matrix client logout failed: %+s", err)
	}
	if err = srv.Shutdown(ctxShutDown); err != nil {
		log.Fatalf("server shutdown failed: %+s", err)
	}

	if err == http.ErrServerClosed {
		err = nil
	}
	return
}
