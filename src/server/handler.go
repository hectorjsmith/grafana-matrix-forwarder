package server

import (
	"fmt"
	"grafana-matrix-forwarder/model"
	"log"
	"net/http"
	"strings"
)

type RequestHandler interface {
	ParseRequest(request *http.Request, logPayload bool) (roomIDs []string, alerts []model.AlertData, err error)
}

func (server *Server) HandleGrafanaAlert(handler RequestHandler, response http.ResponseWriter, request *http.Request) {
	if !server.isAuthorised(request) {
		log.Print("unauthorised request (credentials do not match)")
		response.WriteHeader(http.StatusUnauthorized)
		return
	}
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
	if err != nil {
		return err
	}

	server.metricsCollector.RecordAlerts(alerts)

	err = server.alertForwarder.ForwardEvents(roomIDs, alerts)
	if err != nil {
		return err
	}

	response.WriteHeader(http.StatusOK)
	_, err = response.Write([]byte("OK"))
	return err
}

func (server *Server) isAuthorised(request *http.Request) bool {
	if strings.ToLower(server.appSettings.AuthScheme) == "bearer" {
		authHeader := request.Header.Get("Authorization")
		requiredToken := fmt.Sprintf("Bearer %s", server.appSettings.AuthCredentials)
		return authHeader == requiredToken
	}
	return true
}
