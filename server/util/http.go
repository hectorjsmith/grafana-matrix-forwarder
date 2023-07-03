package util

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func GetRoomIDsFromURL(url *url.URL) ([]string, error) {
	roomIDs, ok := url.Query()["roomId"]
	if !ok || len(roomIDs) < 1 {
		return nil, fmt.Errorf("url param 'roomId' is missing")
	}
	return roomIDs, nil
}

func GetRequestBodyAsBytes(request *http.Request) ([]byte, error) {
	return ioutil.ReadAll(request.Body)
}

func LogRequestPayload(request *http.Request, bodyBytes []byte) {
	log.Printf("%s request received at URL: %s", request.Method, request.URL.String())
	body := string(bodyBytes)
	fmt.Println(body)
}
