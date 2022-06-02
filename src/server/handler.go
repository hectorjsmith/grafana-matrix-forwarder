package server

import (
	"encoding/json"
	"fmt"
	"grafana-matrix-forwarder/server/v0"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func (server *Server) handleGrafanaAlert(response http.ResponseWriter, request *http.Request) error {
	bodyBytes, err := getRequestBodyAsBytes(request)
	if err != nil {
		return err
	}
	if server.appSettings.LogPayload {
		logPayload(request, bodyBytes)
	}

	roomIDs, err := getRoomIDsFromURL(request.URL)
	if err != nil {
		return err
	}

	alert, err := getAlertPayloadFromRequestBody(bodyBytes)
	if err != nil {
		return err
	}

	log.Printf("alert received (%s) - forwarding to rooms: %v", alert.FullRuleID(), roomIDs)
	server.metricsCollector.RecordAlert(alert)

	for _, roomID := range roomIDs {
		err = server.alertForwarder.ForwardAlert(roomID, alert.ToForwarderData())
		if err != nil {
			return err
		}
	}

	response.WriteHeader(http.StatusOK)
	_, err = response.Write([]byte("OK"))
	return err
}

func logPayload(request *http.Request, bodyBytes []byte) {
	log.Printf("%s request received at URL: %s", request.Method, request.URL.String())
	body := string(bodyBytes)
	fmt.Println(body)
}

func getRoomIDsFromURL(url *url.URL) ([]string, error) {
	roomIDs, ok := url.Query()["roomId"]
	if !ok || len(roomIDs) < 1 {
		return nil, fmt.Errorf("url param 'roomId' is missing")
	}
	return roomIDs, nil
}

func getAlertPayloadFromRequestBody(bodyBytes []byte) (alertPayload v0.AlertPayload, err error) {
	err = json.Unmarshal(bodyBytes, &alertPayload)
	return
}

func getRequestBodyAsBytes(request *http.Request) ([]byte, error) {
	return ioutil.ReadAll(request.Body)
}
