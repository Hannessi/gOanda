package V20

import (
	"github.com/sirupsen/logrus"
	"gitlab.com/personal/go/gOanda/V20/requester/api"
)

func New(new NewRequest) *Client {

	//instantiate requester
	requester := api.New()

	return &Client{
		AccountNumber: new.AccountNumber,
		Requester:     requester,
	}
}

type NewRequest struct {
	AccountNumber string
}

type Client struct {
	AccountNumber string
	Requester     Requester
}

// account endpoints
func (c *Client) GetAccounts() error {
	getAccountResponse, err := c.Requester.GetAccounts(GetAccountsRequest{})
	if err != nil {
		return err
	}
	 logrus.Info("GetAccountResponse: ", getAccountResponse)
	return nil
}
