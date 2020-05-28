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
	c := &client.Client{}
	c.Init(endpoints.Guangzhou).WithSecretId(accessID, secretKey)

	req := &models.AndroidPushRequest{
		BaseRequest:  &client.BaseRequest{Path: "/v3/push/app"},
		AudienceType: models.AudienceToken,
		TokenList:    []string{token},
		MessageType:  models.MessageTypeNotify,
		Message: models.AndroidMessage{
			Title:   "test",
			Content: "This is Content",
		},
	}
	log.Printf("Request: %s\n", req.ToJsonString())

	resp := &client.TPNSBaseResponse{}

	if err := c.Send(req, resp); err != nil {
		log.Printf("Error: %s\n", err.Error())
	} else {
		log.Printf("Response: %s\n", resp.ToJsonString())
	}
}
