package server

import (
	"context"
	"fmt"
	"grafana-matrix-forwarder/cfg"
	"grafana-matrix-forwarder/matrix"
	"log"
	"net/http"
	"time"
)

// Server data structure that holds data necessary for the web server to function
type Server struct {
	ctx               context.Context
	matrixWriteCloser matrix.WriteCloser
	appSettings       cfg.AppSettings
	metrics           serverMetrics
}

// BuildServer builds a Server instance based on the provided context.Context, a matrix.WriteCloser, and the cfg.AppSettings
func BuildServer(ctx context.Context, matrixWriteCloser matrix.WriteCloser, appSettings cfg.AppSettings) Server {
	return Server{
		ctx:               ctx,
		matrixWriteCloser: matrixWriteCloser,
		appSettings:       appSettings,
		metrics:           serverMetrics{},
	}
}

// Start the web server and listen for incoming requests
func (server Server) Start() (err error) {
	log.Print("starting webserver ...")
	mux := http.NewServeMux()
	mux.Handle("/api/v0/forward", http.HandlerFunc(
		func(response http.ResponseWriter, request *http.Request) {
			err = server.handleGrafanaAlert(response, request)
			server.metrics.totalForwardCount++
			if err != nil {
				server.metrics.failForwardCount++
				log.Print(err)
				response.WriteHeader(500)
			} else {
				server.metrics.successForwardCount++
			}
		},
	))
	mux.Handle("/metrics", http.HandlerFunc(
		func(response http.ResponseWriter, request *http.Request) {
			err = server.handleMetricsRequest(response)
			if err != nil {
				log.Print(err)
				response.WriteHeader(500)
			}
		},
	))

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
