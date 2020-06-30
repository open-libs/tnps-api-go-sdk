package main

import (
	"log"
	"os"

	"github.com/open-libs/tpns-api-go-sdk/pkg/client"
	"github.com/open-libs/tpns-api-go-sdk/pkg/client/endpoints"
	"github.com/open-libs/tpns-api-go-sdk/pkg/models"
)

func main() {

	accessID := os.Args[1]
	secretKey := os.Args[2]
	token := os.Args[3]

	account := ""
	if len(os.Args) > 4 {
		account = os.Args[4]
	}

	c := &client.Client{}
	c.Init(endpoints.Guangzhou).WithSecretId(accessID, secretKey)

	req := &models.BindAccountRequest{
		BaseRequest:   &client.BaseRequest{Path: "/v3/device/account/batchoperate"},
		OperatorType:  models.OverWriteAccount,
		TokenAccounts: []models.TokenAccountInfo{{Token: token, AccountList: []models.AccountInfo{{Account: account}}}},
	}
	log.Printf("AccessID: %s\n", accessID)
	log.Printf("Request: %s\n", req.ToJsonString())

	resp := &models.TPNSBindAccountResponse{}

	if err := c.Send(req, resp); err != nil {
		log.Printf("Error: %s\n", err.Error())
	} else {
		log.Printf("Response: %s\n", resp.ToJsonString())
	}

	bindReq := &models.QueryBindRequest{
		BaseRequest:  &client.BaseRequest{Path: "/v3/device/account/query"},
		OperatorType: models.QueryAccountByToken,
		TokenList:    []string{token},
	}
	log.Printf("Request: %s\n", bindReq.ToJsonString())

	bindResp := &models.TPNSQueryBindResponse{}

	if err := c.Send(bindReq, bindResp); err != nil {
		log.Printf("Error: %s\n", err.Error())
	} else {
		log.Printf("Response: %s\n", bindResp.ToJsonString())
	}
}
