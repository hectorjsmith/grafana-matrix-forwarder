package service

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type sentMatrixEvent struct {
	EventID           string
	SentFormattedBody string
}

const (
	alertMapFileName = "grafanaToMatrixMap.json"
)

func (f *Forwarder) storeMatrixEvent(alertID string, msgID, body string) {
	f.alertToSentEventMap[alertID] = sentMatrixEvent{
		EventID:           msgID,
		SentFormattedBody: body,
	}
}

func (f *Forwarder) deleteMatrixEvent(alertID string) {
	delete(f.alertToSentEventMap, alertID)
}

func (f *Forwarder) prePopulateAlertMap() {
	fileData, err := ioutil.ReadFile(alertMapFileName)
	if err == nil {
		err = json.Unmarshal(fileData, &f.alertToSentEventMap)
	}

	if err != nil {
		log.Printf("failed to load alert map - falling back on an empty map (%v)", err)
	}
}

func (f *Forwarder) persistAlertMap() {
	if !f.alertMapPersistenceEnabled {
		return
	}

	jsonData, err := json.Marshal(f.alertToSentEventMap)
	if err == nil {
		err = ioutil.WriteFile(alertMapFileName, jsonData, os.ModePerm)
	}

	if err != nil {
		log.Printf("failed to persist alert map - functionality disabled (%v)", err)
		f.alertMapPersistenceEnabled = false
	}
}
