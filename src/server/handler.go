package server

import (
	"grafana-matrix-forwarder/model"
	"log"
	"net/http"
)

type RequestHandler interface {
	ParseRequest(request *http.Request, logPayload bool) (roomIDs []string, alert model.Data, err error)
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
	roomIDs, alert, err := handler.ParseRequest(request, server.appSettings.LogPayload)

	log.Printf("alert received (%s) - forwarding to rooms: %v", alert.Id, roomIDs)
	server.metricsCollector.RecordAlert(alert)

	err = server.alertForwarder.ForwardEvent(roomIDs, alert)
	if err != nil {
		return err
	}

	response.WriteHeader(http.StatusOK)
	_, err = response.Write([]byte("OK"))
	return err
}
