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
	customContent := ""
	if len(os.Args) > 4 {
		customContent = os.Args[4]
	}

	c.Init(endpoints.Guangzhou).WithSecretId(accessID, secretKey).WithAuthMethod(client.Basic)

	now := time.Now()

	req := &models.AndroidPushRequest{
		BaseRequest:  &client.BaseRequest{Path: "/v3/push/app"},
		AudienceType: models.AudienceToken,
		TokenList:    []string{token},
		MessageType:  models.MessageTypeNotify,
		Message: models.AndroidMessage{
			Title:   "test",
			Content: "This is Content @" + now.Format(time.RFC3339),
			Android: map[string]interface{}{
				"custom_content": customContent,
				"action":         map[string]interface{}{"action_type": 1, "activity": "com.kitlink.activity.LoginActivity"},
			},
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
