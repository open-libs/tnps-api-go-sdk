package main

import (
	"log"
	"os"
	"time"

	"github.com/open-libs/tpns-api-go-sdk/pkg/client"
	"github.com/open-libs/tpns-api-go-sdk/pkg/models"
)

func main() {

	accessID := os.Args[1]
	secretKey := os.Args[2]
	token := os.Args[3]
	c := &client.Client{}
	// c.Init(endpoints.Guangzhou).WithSecretId(accessID, secretKey).WithAuthMethod(client.Signature)
	// c.Init("http://9.134.4.90:8080").WithSecretId(accessID, secretKey).WithAuthMethod(client.Signature)
	c.Init("http://api.tpns.tencent.com").WithSecretId(accessID, secretKey).WithAuthMethod(client.Signature)

	now := time.Now()

	req := &models.AndroidPushRequest{
		BaseRequest:  &client.BaseRequest{Path: "/v3/push/app"},
		AudienceType: models.AudienceToken,
		TokenList:    []string{token},
		MessageType:  models.MessageTypeNotify,
		Message: models.AndroidMessage{
			Title:   "test",
			Content: "This is Content @" + now.Format(time.RFC3339),
		},
	}
	log.Printf("AccessID: %s\n", accessID)
	log.Printf("Android Request: %s\n", req.ToJsonString())

	resp := &client.TPNSBaseResponse{}

	if err := c.Send(req, resp); err != nil {
		log.Printf("Error: %s\n", err.Error())
	} else {
		log.Printf("Response: %s\n", resp.ToJsonString())
	}
}
