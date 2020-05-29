package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"

	"net/http"
)

type Response interface {
	ToJsonString() string
	FromJsonString(s string) error
}

type TPNSBaseResponse struct {
	Seq         int64  `json:"seq"`
	PushID      string `json:"push_id"`
	RetCode     int64  `json:"ret_code"`
	Environment string `json:"environment"`
	ErrMsg      string `json:"err_msg"`
	Result      string `json:"result"`
}

func (r *TPNSBaseResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *TPNSBaseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

func ParseFromHttpResponse(hr *http.Response, response Response) (err error) {
	defer hr.Body.Close()
	body, err := ioutil.ReadAll(hr.Body)
	if err != nil {
		msg := fmt.Sprintf("Fail to read response body because %s", err)
		return errors.New("ClientError.IOError: " + msg)
	}
	if hr.StatusCode != 200 {
		msg := fmt.Sprintf("Request fail with http status code: %s, with body: %s", hr.Status, body)
		return errors.New("ClientError.HttpStatusCodeError: " + msg)
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		msg := fmt.Sprintf("Fail to parse json content: %s, because: %s", body, err)
		return errors.New("ClientError.ParseJsonError: " + msg)
	}
	return
}
