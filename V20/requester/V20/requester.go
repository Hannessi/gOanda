package V20

import (
	"github.com/hannessi/gOanda/V20/definitions"
)

type Requester interface {
	// account
	GetAccounts(GetAccountsRequest) (*GetAccountsResponse, error)
	GetAccount(GetAccountRequest) (*GetAccountResponse, error)
}

type GetAccountsRequest struct {}

type GetAccountsResponse struct {
	Accounts []definitions.Account `json:"accounts"`
}

type GetAccountRequest struct {}

type GetAccountResponse struct {
	Account definitions.Account `json:"account"`
}