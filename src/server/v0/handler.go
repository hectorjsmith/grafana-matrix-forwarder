package v0

import (
	"encoding/json"
	"fmt"
	"grafana-matrix-forwarder/model"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type Handler struct {
}

func (h Handler) ParseRequest(request *http.Request, logPayload bool) (roomIDs []string, alert model.Data, err error) {
	bodyBytes, err := getRequestBodyAsBytes(request)
	if err != nil {
		return
	}
	if logPayload {
		logRequestPayload(request, bodyBytes)
	}

	roomIDs, err = getRoomIDsFromURL(request.URL)
	if err != nil {
		return
	}

	alertPayload, err := getAlertPayloadFromRequestBody(bodyBytes)
	if err != nil {
		return
	}

	alert = alertPayload.ToForwarderData()
	return
}

func getRoomIDsFromURL(url *url.URL) ([]string, error) {
	roomIDs, ok := url.Query()["roomId"]
	if !ok || len(roomIDs) < 1 {
		return nil, fmt.Errorf("url param 'roomId' is missing")
	}
	return roomIDs, nil
}

func getAlertPayloadFromRequestBody(bodyBytes []byte) (alertPayload AlertPayload, err error) {
	err = json.Unmarshal(bodyBytes, &alertPayload)
	return
}

func getRequestBodyAsBytes(request *http.Request) ([]byte, error) {
	return ioutil.ReadAll(request.Body)
}

func logRequestPayload(request *http.Request, bodyBytes []byte) {
	log.Printf("%s request received at URL: %s", request.Method, request.URL.String())
	body := string(bodyBytes)
	fmt.Println(body)
}
