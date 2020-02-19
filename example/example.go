package main

import (
	"flag"
	"github.com/hannessi/gOanda"
	"github.com/sirupsen/logrus"
)

func main() {

	token := flag.String("oanda-api-key", "", "OANDA server API key")
	account := flag.String("oanda-account-number", "", "OANDA account number")

	flag.Parse()

	// create new client
	oandaClient := gOanda.New(*account, *token)

	//getAccountsResponse, err := oandaClient.GetAccounts()
	//if err != nil {
	//	logrus.Error(err.Error())
	//}
	//
	//for _, account := range getAccountsResponse.Accounts {
	//	logrus.Info(account.String())
	//}
	//
	//getAccountResponse, err := oandaClient.GetAccount()
	//if err != nil {
	//	logrus.Error(err.Error())
	//}
	//
	//logrus.Info(getAccountResponse.Account.String())

	getOrdersResponse, err := oandaClient.GetOrders(gOanda.GetOrdersRequest{
		InstrumentName: gOanda.INSTRUMENT_NAME_USDCHF,
	})
	if err != nil {
		logrus.Error(err.Error())
	}

	logrus.Info(getOrdersResponse)

	getPendingOrderResponse, err := oandaClient.GetPendingOrders()
	if err != nil {
		logrus.Error(err.Error())
	}

	logrus.Info(getPendingOrderResponse)


}
