package main

import (
	"encoding/json"
	"fmt"
	"github.com/sendgrid/sendgrid-go"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const (
	CONTENTTYPE     string = "Content-Type"
	APPLICATIONJSON string = "application/json"
)

func MiddlewareSendMail(w http.ResponseWriter, r *http.Request) {
	var response Response

	addHeaders(w, r)
	handleOptions(w, r)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response = Response{Code: 500, StatusCode: "500", Status: "ERROR", Message: fmt.Sprintf(" Could not read body data (MiddlewareSendMail) %v\n", err), Payload: ""}
	} else {

		apiKey := os.Getenv("API_KEY")
		url := os.Getenv("MAIL_URL")
		host := os.Getenv("MAIL_HOST")
		request := sendgrid.GetRequest(apiKey, url, host)
		request.Method = "POST"
		request.Body = body

		res, err := sendgrid.API(request)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			response = Response{Code: 500, StatusCode: "500", Status: "ERROR", Message: err.Error(), Payload: ""}
			logger.Error(fmt.Sprintf("sending email : %v\n", err))
		} else {
			w.WriteHeader(http.StatusOK)
			response = Response{Code: 200, StatusCode: "200", Status: "OK", Message: "Mail sent succesfully", Payload: res.Body}
			logger.Info(fmt.Sprintf("Mail sent succesfully : %v\n", response))
		}
	}
	b, _ := json.MarshalIndent(response, "", "	")
	fmt.Fprintf(w, string(b))
}

func MiddlewareCreateBatchId(w http.ResponseWriter, r *http.Request) {
	var response Response

	apiKey := os.Getenv("SENDGRID_API_KEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/mail/send", host)
	request.Method = "POST"
	res, err := sendgrid.API(request)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		response = Response{Code: 500, StatusCode: "500", Status: "ERROR", Message: err.Error(), Payload: ""}
	} else {
		w.WriteHeader(http.StatusOK)
		response = Response{Code: 200, StatusCode: "200", Status: "OK", Message: "Batch ID succesfully created", Payload: res.Body}
	}
	b, _ := json.MarshalIndent(response, "", "	")
	fmt.Fprintf(w, string(b))
}

func IsAlive(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "{ \"version\" : \""+os.Getenv("VERSION")+"\" , \"name\": \"Auth\" }")
}

// headers (with cors) utility
func addHeaders(w http.ResponseWriter, r *http.Request) {
	var request []string
	for name, headers := range r.Header {
		name = strings.ToLower(name)
		for _, h := range headers {
			request = append(request, fmt.Sprintf("%v: %v", name, h))
		}
	}

	logger.Trace(fmt.Sprintf("Headers : %s", request))

	w.Header().Set(CONTENTTYPE, APPLICATIONJSON)
	// use this for cors
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

}

// simple options handler
func handleOptions(w http.ResponseWriter, r *http.Request) bool {
	if r.Method == "OPTIONS" {
		return true
	}
	return false
}
