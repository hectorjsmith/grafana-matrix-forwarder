package server

import (
	"grafana-matrix-forwarder/model"
	"log"
	"net/http"
)

type RequestHandler interface {
	ParseRequest(request *http.Request, logPayload bool) (roomIDs []string, alerts []model.AlertData, err error)
}

func (server *Server) HandleGrafanaAlert(handler RequestHandler, response http.ResponseWriter, request *http.Request) {
	err := server.handleGrafanaAlertInner(handler, response, request)
	if err != nil {
		server.metricsCollector.IncrementFailure()
		log.Print(err)
		response.WriteHeader(http.StatusInternalServerError)
	} else {
		server.metricsCollector.IncrementSuccess()
	}
}

func (server *Server) handleGrafanaAlertInner(handler RequestHandler, response http.ResponseWriter, request *http.Request) error {
	roomIDs, alerts, err := handler.ParseRequest(request, server.appSettings.LogPayload)

	server.metricsCollector.RecordAlerts(alerts)

	err = server.alertForwarder.ForwardEvents(roomIDs, alerts)
	if err != nil {
		return err
	}

	response.WriteHeader(http.StatusOK)
	_, err = response.Write([]byte("OK"))
	return err
}
