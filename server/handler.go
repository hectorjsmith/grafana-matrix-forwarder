package server

import (
	"encoding/json"
	"fmt"
	"grafana-matrix-forwarder/grafana"
	"io/ioutil"
	"log"
	"net/http"
)

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

	err = grafana.ForwardAlert(server.matrixWriteCloser.GetWriter(), roomID, alert, server.appSettings.ResolveMode)
	if err != nil {
		return err
	}

	response.WriteHeader(200)
	_, err = response.Write([]byte("OK"))
	return err
}

func (server Server) handleMetricsRequest(response http.ResponseWriter) (err error) {
	metric, err := buildMetricString("up", "gauge", 1.0)
	_, err = response.Write([]byte(metric))
	return
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
