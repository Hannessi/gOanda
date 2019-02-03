package V20

type Requester interface {
	GetAccounts(GetAccountsRequest) (*GetAccountsResponse, error)
}

type GetAccountsRequest struct {

}

type GetAccountsResponse struct {

}
