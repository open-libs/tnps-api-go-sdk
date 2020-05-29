package models

import (
	"encoding/json"

	"github.com/open-libs/tpns-api-go-sdk/pkg/client"
)

type TimeHourMin struct {
	Hour string `json:"hour"`
	Min  string `json:"min"`
}

type TimeRange struct {
	Start TimeHourMin `json:"start"`
	End   TimeHourMin `json:"end"`
}

type AndroidMessage struct {
	Title                 string                 `json:"title,omitempty"`
	Content               string                 `json:"content,omitempty"`
	AcceptTime            []TimeRange            `json:"accept_time,omitempty"`
	XgMediaResources      string                 `json:"xg_media_resources,omitempty"`
	XgMediaAudioResources string                 `json:"xg_media_audio_resources,omitempty"`
	Android               map[string]interface{} `json:"android,omitempty"`
}

type APSMessage struct {
	Alert          map[string]interface{} `json:"alert,omitempty"`
	BadgeType      int64                  `json:"badge_type,omitempty"`
	Category       string                 `json:"category,omitempty"`
	Sound          string                 `json:"sound,omitempty"`
	MutableContent int64                  `json:"mutable-content,omitempty"`
}

type IOSContent struct {
	APS           APSMessage `json:"aps,omitempty"`
	CustomContent string     `json:"custom_content,omitempty"`
}

type IOSMessage struct {
	Title            string     `json:"title,omitempty"`
	Content          string     `json:"content,omitempty"`
	IOS              IOSContent `json:"ios"`
	XgMediaResources string     `json:"xg_media_resources,omitempty"`
	Xg               string     `json:"xg,omitempty"`
}

const (
	AudienceAll                = "all"
	AudienceTag                = "tag"
	AudienceToken              = "token"
	AudienceTokenList          = "token_list"
	AudienceAccount            = "account"
	AudienceAccountList        = "account_list"
	AudiencePackageAccountPush = "package_account_push"

	MessageTypeNotify  = "notify"
	MessageTypeMessage = "message"
)

type AndroidPushRequest struct {
	*client.BaseRequest

	/*	AudienceType 推送目标：
		all：全量推送
		tag：标签推送
		token：单设备推送
		token_list：设备列表推送
		account：单账号推送
		account_list：账号列表推送
		package_account_push：号码包推送
	*/
	AudienceType string `json:"audience_type,omitempty"`
	/* MessageType 	消息类型：
	notify：通知
	message：透传消息/静默消息
	*/
	MessageType string `json:"message_type,omitempty"`

	// Message 消息体
	Message AndroidMessage `json:"message,omitempty"`

	Environment string `json:"environment,omitempty"`
	UploadID    int64  `json:"upload_id,omitempty"`

	// tag_rules
	TokenList []string `json:"token_list,omitempty"`

	AccountList []string `json:"account_list,omitempty"`
}

func (r *AndroidPushRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AndroidPushRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type IOSPushRequest struct {
	*client.BaseRequest

	/*	AudienceType 推送目标：
		all：全量推送
		tag：标签推送
		token：单设备推送
		token_list：设备列表推送
		account：单账号推送
		account_list：账号列表推送
		package_account_push：号码包推送
	*/
	AudienceType string `json:"audience_type,omitempty"`
	/* MessageType 	消息类型：
	notify：通知
	message：透传消息/静默消息
	*/
	MessageType string `json:"message_type,omitempty"`

	// Message 消息体
	Message IOSMessage `json:"message,omitempty"`

	Environment string `json:"environment,omitempty"`
	UploadID    int64  `json:"upload_id,omitempty"`

	// tag_rules
	TokenList []string `json:"token_list,omitempty"`

	AccountList []string `json:"account_list,omitempty"`
}

func (r *IOSPushRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *IOSPushRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
