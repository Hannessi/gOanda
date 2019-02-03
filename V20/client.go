package V20

import (
	"github.com/hannessi/gOanda/V20/requester/V20"
	"github.com/hannessi/gOanda/V20/requester/V20/api"
	"github.com/sirupsen/logrus"
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
	Requester     V20.Requester
}

// account endpoints
func (c *Client) GetAccounts() error {
	getAccountResponse, err := c.Requester.GetAccounts(V20.GetAccountsRequest{})
	if err != nil {
		return err
	}
	logrus.Info("GetAccountResponse: ", getAccountResponse)
	return nil
}
