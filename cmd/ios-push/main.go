package main

import (
	"log"
	"os"
	"time"

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
	now := time.Now()

	req := &models.IOSPushRequest{
		BaseRequest:  &client.BaseRequest{Path: "/v3/push/app"},
		AudienceType: models.AudienceToken,
		TokenList:    []string{token},
		MessageType:  models.MessageTypeNotify,
		Message: models.IOSMessage{
			Title:   "test",
			Content: "This is Content @" + now.Format(time.RFC3339),
		},
		Environment: "dev",
	}
	log.Printf("AccessID: %s\n", accessID)
	log.Printf("IOS Request: %s\n", req.ToJsonString())

	resp := &client.TPNSBaseResponse{}

	if err := c.Send(req, resp); err != nil {
		log.Printf("Error: %s\n", err.Error())
	} else {
		log.Printf("IOS Response: %s\n", resp.ToJsonString())
	}
}
