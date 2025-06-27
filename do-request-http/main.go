package main

import (
	"bytes"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"time"
)

type ApiBgTrackingRequest struct {
	RegistrationNo string `json:"registrationNo"`
}

type ApiBgTrackingResponse struct {
	ResponseCode    string             `json:"responseCode"`
	ResponseMessage *json.RawMessage   `json:"responseMessage"`
	Data            *ApiBgTrackingData `json:"responseData"`
}

type ApiBgTrackingData struct {
	RegistrationNo  string `json:"registrationNo"`
	ReferenceNo     string `json:"referenceNo"`
	WarkatUrl       string `json:"warkatUrl"`
	WarkatUrlPublic string `json:"warkatUrlPublic"`
	Status          string `json:"status"`
	ModifiedDate    string `json:"modifiedDate"`
}

func main() {
	req := &ApiBgTrackingRequest{RegistrationNo: "REG-00206-250526-001"}

	payloadBytes, _ := json.Marshal(req)
	fmt.Println(string(payloadBytes))
	bodyBytes := bytes.NewReader(payloadBytes)

	httpRequest, err := http.NewRequest("POST", "http://api.close.dev.bri.co.id:5557/gateway/apiPortalBG/1.0/tracking", bodyBytes)
	if err != nil {
		panic(err)
	}

	// Headers
	// httpRequest.SetBasicAuth("bricams", "Bricams4dd0ns")
	basicAuth := base64.StdEncoding.EncodeToString([]byte("bricams:Bricams4dd0ns"))
	fmt.Println("Basic Auth:", basicAuth)
	httpRequest.Header.Add("Authorization", "Basic "+basicAuth)
	httpRequest.Header.Add("X-DATAHUB-PERSONAL-NUMBER", "12345678")
	httpRequest.Header.Add("X-DATAHUB-CHANNEL", "BRIMOTESTxx")
	// httpRequest.Header.Add("Content-Type", "application/json")
	// httpRequest.Header.Add("User-Agent", "PostmanRuntime/7.43.3")
	// httpRequest.Header.Add("Accept", "*/*")
	// httpRequest.Header.Add("Accept-Encoding", "gzip, deflate, br")
	// httpRequest.Header.Add("Connection", "keep-alive")

	// Transport & client setup
	var netTransport = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		Dial: (&net.Dialer{
			Timeout: 60 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 60 * time.Second,
	}
	client := &http.Client{
		Transport: netTransport,
		Timeout:   60 * time.Second,
	}

	// Perform request
	res, err := client.Do(httpRequest)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	// Read response
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	response := &ApiBgTrackingResponse{}
	err = json.Unmarshal(resBody, response)
	if err != nil {
		panic(err)
	}

	fmt.Println("response", response)
	fmt.Println("response code", response.ResponseCode)
	fmt.Println("response message", string(*response.ResponseMessage))
}
