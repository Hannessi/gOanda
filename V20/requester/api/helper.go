package api

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
)

type method string

func(m method) String() string {
	return string(m)
}

const GET method = "GET"
const POST method = "POST"
const PATCH method = "PATCH"

func HttpRequestWrapper(method method, url string, requestBody io.Reader, response interface{}, token string) error {
	httpRequest, err := http.NewRequest(method.String(), url, requestBody)
	if err != nil {
		return errors.New("could not create new HTTP request: " + err.Error())
	}

	httpRequest.Header.Add("Content-type","application/json")

	if token != "" {
		httpRequest.Header.Add("Authorization", "Bearer "+token)
	}

	httpResponse, err := http.DefaultClient.Do(httpRequest)
	if err != nil {
		return errors.New("could not send HTTP request: " + err.Error())
	}

	defer httpResponse.Body.Close()
	responseBody, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		return  errors.New("could not read body of response: " + err.Error())
	}

	err = json.Unmarshal(responseBody, response)
	if err != nil {
		return errors.New("could not unmarshal the body of the response: " + err.Error())
	}
	return nil
}