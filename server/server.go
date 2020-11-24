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
	"maunium.net/go/mautrix"
	"net/http"
	"time"
)

func Start(ctx context.Context, matrixClient *mautrix.Client, settings cfg.AppSettings) (err error) {
	log.Print("starting webserver ...")
	mux := http.NewServeMux()
	mux.Handle("/api/v0/forward", http.HandlerFunc(
		func(response http.ResponseWriter, request *http.Request) {
			err = handleGrafanaAlert(response, request, matrixClient, settings)
			if err != nil {
				log.Print(err)
				response.WriteHeader(500)
			}
		},
	))

	serverAddr := fmt.Sprintf("%s:%d", settings.ServerHost, settings.ServerPort)
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

	<-ctx.Done()

	log.Printf("shutting down ...")

	ctxShutDown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()

	if _, err = matrixClient.Logout(); err != nil {
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

func handleGrafanaAlert(response http.ResponseWriter, request *http.Request, matrixClient *mautrix.Client, settings cfg.AppSettings) error {
	bodyBytes, err := getRequestBodyAsBytes(request)
	if err != nil {
		return err
	}
	if settings.LogPayload {
		logPayload(request, bodyBytes)
	}

	roomId, err := getRoomIdFromUrl(request)
	if err != nil {
		return err
	}
	log.Printf("alert received - forwarding to room: %s", roomId)

	alert, err := getAlertPayloadFromRequestBody(bodyBytes)
	if err != nil {
		return err
	}

	err = matrix.SendAlert(matrixClient, alert, roomId)
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

func getRoomIdFromUrl(request *http.Request) (string, error) {
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
