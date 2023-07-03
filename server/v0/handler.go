package v0

import (
	"encoding/json"
	"grafana-matrix-forwarder/model"
	"grafana-matrix-forwarder/server/util"
	"net/http"
)

type Handler struct {
}

func (h Handler) ParseRequest(request *http.Request, logPayload bool) (roomIDs []string, alerts []model.AlertData, err error) {
	bodyBytes, err := util.GetRequestBodyAsBytes(request)
	if err != nil {
		return
	}
	if logPayload {
		util.LogRequestPayload(request, bodyBytes)
	}

	roomIDs, err = util.GetRoomIDsFromURL(request.URL)
	if err != nil {
		return
	}

	alertPayload, err := getAlertPayloadFromRequestBody(bodyBytes)
	if err != nil {
		return
	}

	alerts = make([]model.AlertData, 1)
	alerts[0] = alertPayload.ToForwarderData()
	return
}

func getAlertPayloadFromRequestBody(bodyBytes []byte) (alertPayload alertPayload, err error) {
	err = json.Unmarshal(bodyBytes, &alertPayload)
	return
}
