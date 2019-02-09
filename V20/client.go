package V20

import (
	"github.com/hannessi/gOanda"
	"github.com/hannessi/gOanda/V20/definitions"
	"github.com/hannessi/gOanda/V20/requester/V20"
	"github.com/hannessi/gOanda/V20/requester/V20/api"
	"github.com/hannessi/gOanda/V20/requester/V20/url"
)

func New(accountNumber string, apiKey string) *Client {

	//instantiate URL Manager
	urlManager := url.Manager{
		BaseUrl:   gOanda.V20_BASE_PRACTICE_URL,
		AccountId: accountNumber,
	}

	//instantiate requester
	requester := api.New(accountNumber, apiKey, urlManager)

	return &Client{
		AccountNumber: accountNumber,
		ApiKey:        apiKey,
		Requester:     requester,
	}
}

type Client struct {
	AccountNumber string
	ApiKey        string
	Requester     V20.Requester
}

// account endpoints
type GetAccountsResponse struct {
	Account []definitions.Account
}

func (c *Client) GetAccounts() (*GetAccountsResponse, error) {
	getAccountResponse, err := c.Requester.GetAccounts(V20.GetAccountsRequest{})
	if err != nil {
		return nil, err
	}
	return &GetAccountsResponse{
		Account: getAccountResponse.Accounts,
	}, nil
}

type GetAccountResponse struct {
	Account definitions.Account
}

func (c *Client) GetAccount() (*GetAccountResponse, error) {
	getAccountResponse, err := c.Requester.GetAccount(V20.GetAccountRequest{})
	if err != nil {
		return nil, err
	}
	return &GetAccountResponse{
		Account: getAccountResponse.Account,
	}, nil
}