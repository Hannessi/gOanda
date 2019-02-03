package api

import "github.com/hannessi/gOanda/V20/requester/V20"

func New() *Requester {
	return &Requester{}
}

type Requester struct {

}

func (r *Requester) GetAccounts (request V20.GetAccountsRequest) (*V20.GetAccountsResponse, error) {
	return nil, nil
}