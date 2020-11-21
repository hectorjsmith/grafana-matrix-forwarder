package server

import (
	"context"
	"encoding/json"
	"fmt"
	"grafana-matrix-forwarder/grafana"
	"grafana-matrix-forwarder/matrix"
	"io/ioutil"
	"log"
	"maunium.net/go/mautrix"
	"net/http"
	"time"
)

func Start(ctx context.Context, matrixClient *mautrix.Client, serverHost string, serverPort int) (err error) {
	log.Print("starting webserver ...")
	mux := http.NewServeMux()
	mux.Handle("/api/v0/forward", http.HandlerFunc(
		func(response http.ResponseWriter, request *http.Request) {
			err = handleGrafanaAlert(response, request, matrixClient)
			if err != nil {
				log.Print(err)
				response.WriteHeader(500)
			}
		},
	))

	serverAddr := fmt.Sprintf("%s:%d", serverHost, serverPort)
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

func handleGrafanaAlert(response http.ResponseWriter, request *http.Request, matrixClient *mautrix.Client) error {
	roomId, err := getRoomIdFromUrl(request)
	if err != nil {
		return err
	}
	log.Printf("alert received - forwarding to room: %s", roomId)

	alert, err := getAlertPayloadFromRequestBody(request)
	if err != nil {
		return err
	}

	err = matrix.SendAlert(matrixClient, alert, roomId)
	if err != nil {
		return err
	}
	response.WriteHeader(200)
	_, err = response.Write([]byte("OK"))
	return nil
}

func getRoomIdFromUrl(request *http.Request) (string, error) {
	roomIds, ok := request.URL.Query()["roomId"]
	if !ok || len(roomIds[0]) < 1 {
		return "", fmt.Errorf("url param 'roomId' is missing")
	}
	return roomIds[0], nil
}

func getAlertPayloadFromRequestBody(request *http.Request) (alertPayload grafana.AlertPayload, err error) {
	bodyBytes, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(bodyBytes, &alertPayload)
	return
}
