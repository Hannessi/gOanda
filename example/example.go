package main

import "gitlab.com/personal/go/gOanda/V20"

func main() {

	// create new client
	oandaClient := V20.New(V20.NewRequest{
		AccountNumber: "123021334",
	})

	oandaClient.GetAccounts()
}
