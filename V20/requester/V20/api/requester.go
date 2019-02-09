package api

import (
	"encoding/json"
	"errors"
	"github.com/hannessi/gOanda/V20/requester/V20"
	"github.com/hannessi/gOanda/V20/requester/V20/url"
	"io/ioutil"
	"net/http"
)

func New(
	AccountId string,
	Token string,
	UrlManager url.Manager,
) *Requester {
	return &Requester{
		AccountId:  AccountId,
		Token:      Token,
		UrlManager: UrlManager,
	}
}

type Requester struct {
	AccountId  string
	Token      string
	UrlManager url.Manager
}

func (r *Requester) GetAccounts(request V20.GetAccountsRequest) (*V20.GetAccountsResponse, error) {
	httpRequest, err := http.NewRequest("GET", r.UrlManager.GetAccounts(), nil)
	if err != nil {
		return nil, errors.New("could not create new HTTP request: " + err.Error())
	}

	httpRequest.Header.Add("Content-type", "application/json")
	httpRequest.Header.Add("Authorization", "Bearer "+r.Token)

	httpResponse, err := http.DefaultClient.Do(httpRequest)
	if err != nil {
		return nil, errors.New("could not send HTTP request: " + err.Error())
	}

	defer httpResponse.Body.Close()
	body, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		return nil, errors.New("could not read body of response: " + err.Error())
	}

	response := &V20.GetAccountsResponse{}

	err = json.Unmarshal(body, response)
	if err != nil {
		return nil, errors.New("could not unmarshal the body of the response: " + err.Error())
	}

	return response, nil
}

func (r *Requester) GetAccount(request V20.GetAccountRequest) (*V20.GetAccountResponse, error) {
	httpRequest, err := http.NewRequest("GET", r.UrlManager.GetAccount(), nil)
	if err != nil {
		return nil, errors.New("could not create new HTTP request: " + err.Error())
	}

	httpRequest.Header.Add("Content-type", "application/json")
	httpRequest.Header.Add("Authorization", "Bearer "+r.Token)

	httpResponse, err := http.DefaultClient.Do(httpRequest)
	if err != nil {
		return nil, errors.New("could not send HTTP request: " + err.Error())
	}

	defer httpResponse.Body.Close()
	body, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		return nil, errors.New("could not read body of response: " + err.Error())
	}

	response := &V20.GetAccountResponse{}

	err = json.Unmarshal(body, response)
	if err != nil {
		return nil, errors.New("could not unmarshal the body of the response: " + err.Error())
	}

	return response, nil
}
