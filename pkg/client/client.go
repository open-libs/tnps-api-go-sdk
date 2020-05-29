package client

import (
	"crypto/hmac"
	"crypto/sha256"
	"errors"
	"fmt"
	"strconv"
	"time"

	"net/http"
	"strings"

	"encoding/base64"
	"encoding/hex"
)

const (
	Signature = "Signature"
	Basic     = "Basic"
)

type Client struct {
	endpoint   string
	httpClient *http.Client
	credential *Credential
	authMethod string
}

func (c *Client) Send(request Request, response Response) (err error) {
	if c.authMethod == Signature {
		return c.sendWithSignature(request, response)
	} else {
		return c.sendWithBasicAuth(request, response)
	}
}

func (c *Client) sendWithBasicAuth(request Request, response Response) (err error) {
	headers := map[string]string{}

	authorization := base64.StdEncoding.EncodeToString([]byte(
		fmt.Sprintf("%s:%s", c.credential.AccessID, c.credential.SecretKey)))
	headers["Authorization"] = "Basic " + authorization
	url := c.GetEndPoint() + request.GetPath()

	reqBody := request.ToJsonString()

	httpRequest, err := http.NewRequest("POST", url, strings.NewReader(reqBody))
	if err != nil {
		return err
	}
	for k, v := range headers {
		httpRequest.Header[k] = []string{v}
	}

	httpResponse, err := c.httpClient.Do(httpRequest)
	if err != nil {
		msg := fmt.Sprintf("Fail to get response because %s", err)
		return errors.New("ClientError.NetworkError: " + msg)
	}
	err = ParseFromHttpResponse(httpResponse, response)
	return err
}

func (c *Client) sendWithSignature(request Request, response Response) (err error) {

	now := time.Now()
	ts := strconv.FormatInt(now.Unix(), 10)
	headers := map[string]string{
		"AccessId":  c.credential.AccessID,
		"TimeStamp": ts,
	}

	headers["Content-Type"] = "application/json"

	url := c.GetEndPoint() + request.GetPath()

	reqBody := request.ToJsonString()

	stringToSign := c.credential.AccessID + ts + reqBody

	h := hmac.New(sha256.New, []byte(c.credential.SecretKey))
	h.Write([]byte(stringToSign))
	sha := hex.EncodeToString(h.Sum(nil))
	sign := base64.StdEncoding.EncodeToString([]byte(sha))

	headers["Sign"] = sign

	httpRequest, err := http.NewRequest("POST", url, strings.NewReader(reqBody))
	if err != nil {
		return err
	}
	for k, v := range headers {
		httpRequest.Header[k] = []string{v}
	}

	httpResponse, err := c.httpClient.Do(httpRequest)
	if err != nil {
		msg := fmt.Sprintf("Fail to get response because %s", err)
		return errors.New("ClientError.NetworkError: " + msg)
	}
	err = ParseFromHttpResponse(httpResponse, response)
	return err
}

func (c *Client) GetEndPoint() string {
	return c.endpoint
}

func (c *Client) Init(endpoint string) *Client {
	c.httpClient = &http.Client{}
	c.endpoint = endpoint
	c.authMethod = Basic
	return c
}

func (c *Client) WithSecretId(accessID, secretKey string) *Client {
	c.credential = NewCredential(accessID, secretKey)
	return c
}

func (c *Client) WithCredential(cred *Credential) *Client {
	c.credential = cred
	return c
}

func (c *Client) WithAuthMethod(method string) *Client {
	c.authMethod = method
	return c
}

func (c *Client) WithHttpTransport(transport http.RoundTripper) *Client {
	c.httpClient.Transport = transport
	return c
}

func NewClientWithSecretId(accessID, secretKey, region string) (client *Client, err error) {
	client = &Client{}
	client.Init(region).WithSecretId(accessID, secretKey)
	return
}
