package models

import (
	"encoding/json"

	"github.com/open-libs/tpns-api-go-sdk/pkg/client"
)

const (
	/*
		操作类型参数，取值范围为[1,5]；值对应功能说明如下：
			1：Token 追加设置 account。
			2：Token 覆盖设置 account。
			3：Token 删除绑定的多个 account。
			4：Token 删除绑定的所有 account。
			5： account 删除绑定的所有 Token。
	*/
	AppendAccount      = 1
	OverWriteAccount   = 2
	DeleteMultiAccount = 3
	DeleteAllAccount   = 4
	DeleteAllToken     = 5

	/*
			操作类型：
		1：根据 Account 批量查询对应 Token
		2：根据 Token 查询 Account
	*/
	QueryTokenByAccount = 1
	QueryAccountByToken = 2
)

type AccountInfo struct {
	Account string `json:"account,omitempty"`
}

type TokenAccountInfo struct {
	Token       string        `json:"token,omitempty"`
	AccountList []AccountInfo `json:"account_list,omitempty"`
}

type BindAccountRequest struct {
	*client.BaseRequest

	OperatorType int32 `json:"operator_type,omitempty"`

	Platform      string             `json:"platform,omitempty"`
	TokenAccounts []TokenAccountInfo `json:"token_accounts,omitempty"`
	TokenList     []string           `json:"token_list,omitempty"`
	AccountList   []AccountInfo      `json:"account_list,omitempty"`
}

func (r *BindAccountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BindAccountRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TPNSBindAccountResponse struct {
	RetCode int64    `json:"ret_code"`
	ErrMsg  string   `json:"err_msg"`
	Result  []string `json:"result"`
}

func (r *TPNSBindAccountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *TPNSBindAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryBindRequest struct {
	*client.BaseRequest

	OperatorType int32 `json:"operator_type,omitempty"`

	TokenList   []string      `json:"token_list,omitempty"`
	AccountList []AccountInfo `json:"account_list,omitempty"`
}

func (r *QueryBindRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryBindRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TPNSQueryBindResponse struct {
	RetCode       int64              `json:"ret_code"`
	ErrMsg        string             `json:"err_msg"`
	Result        []string           `json:"result"`
	TokenAccounts []TokenAccountInfo `json:"token_accounts,omitempty"`
}

func (r *TPNSQueryBindResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *TPNSQueryBindResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
