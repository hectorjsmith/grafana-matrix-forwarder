package server

import (
	"context"
	"encoding/json"
	"fmt"
	"grafana-matrix-forwarder/cfg"
	"grafana-matrix-forwarder/grafana"
	"grafana-matrix-forwarder/matrix"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// Server data structure that holds data necessary for the web server to function
type Server struct {
	ctx               context.Context
	matrixWriteCloser matrix.WriteCloser
	appSettings       cfg.AppSettings
}

// BuildServer builds a Server instance based on the provided context.Context, a matrix.WriteCloser, and the cfg.AppSettings
func BuildServer(ctx context.Context, matrixWriteCloser matrix.WriteCloser, appSettings cfg.AppSettings) Server {
	return Server{
		ctx:               ctx,
		matrixWriteCloser: matrixWriteCloser,
		appSettings:       appSettings,
	}
}

// Start the web server and listen for incoming requests
func (server Server) Start() (err error) {
	log.Print("starting webserver ...")
	mux := http.NewServeMux()
	mux.Handle("/api/v0/forward", http.HandlerFunc(
		func(response http.ResponseWriter, request *http.Request) {
			err = server.handleGrafanaAlert(response, request)
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

func (server Server) handleGrafanaAlert(response http.ResponseWriter, request *http.Request) error {
	bodyBytes, err := getRequestBodyAsBytes(request)
	if err != nil {
		return err
	}
	if server.appSettings.LogPayload {
		logPayload(request, bodyBytes)
	}

	roomID, err := getRoomIDFromURL(request)
	if err != nil {
		return err
	}

	alert, err := getAlertPayloadFromRequestBody(bodyBytes)
	if err != nil {
		return err
	}

	log.Printf("alert received (%s) - forwarding to room: %s", alert.FullRuleID(), roomID)

	err = grafana.ForwardAlert(server.matrixWriteCloser.GetWriter(), roomID, alert)
	if err != nil {
		return err
	}

	response.WriteHeader(200)
	_, err = response.Write([]byte("OK"))
	return err
}

func logPayload(request *http.Request, bodyBytes []byte) {
	log.Printf("%s request received at URL: %s", request.Method, request.URL.String())
	body := string(bodyBytes)
	fmt.Println(body)
}

func getRoomIDFromURL(request *http.Request) (string, error) {
	roomIds, ok := request.URL.Query()["roomId"]
	if !ok || len(roomIds[0]) < 1 {
		return "", fmt.Errorf("url param 'roomId' is missing")
	}
	return roomIds[0], nil
}

func getAlertPayloadFromRequestBody(bodyBytes []byte) (alertPayload grafana.AlertPayload, err error) {
	err = json.Unmarshal(bodyBytes, &alertPayload)
	return
}

func getRequestBodyAsBytes(request *http.Request) ([]byte, error) {
	return ioutil.ReadAll(request.Body)
}
