package client

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"strings"

	"encoding/base64"
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
	debug      bool
}

func (c *Client) Send(request Request, response Response) (err error) {
	if c.authMethod == "Signature" {
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
	if c.debug {
		outbytes, err := httputil.DumpRequest(httpRequest, true)
		if err != nil {
			log.Printf("[ERROR] dump request failed because %s", err)
			return err
		}
		log.Printf("[DEBUG] http request = %s", outbytes)
	}
	httpResponse, err := c.httpClient.Do(httpRequest)
	if err != nil {
		msg := fmt.Sprintf("Fail to get response because %s", err)
		log.Println(msg)
		return errors.New("ClientError.NetworkError")
	}
	err = ParseFromHttpResponse(httpResponse, response)
	return err
}

func (c *Client) sendWithSignature(request Request, response Response) (err error) {
	headers := map[string]string{
		// "Host":               request.GetDomain(),
		"AccessId": c.credential.AccessID,
		// "TimeStamp": request.GetParams()["Timestamp"],
	}

	headers["Content-Type"] = "application/json"

	// headers["Authorization"] = authorization
	url := c.GetEndPoint() + request.GetPath()

	reqBody := request.ToJsonString()

	httpRequest, err := http.NewRequest("POST", url, strings.NewReader(reqBody))
	if err != nil {
		return err
	}
	for k, v := range headers {
		httpRequest.Header[k] = []string{v}
	}
	if c.debug {
		outbytes, err := httputil.DumpRequest(httpRequest, true)
		if err != nil {
			log.Printf("[ERROR] dump request failed because %s", err)
			return err
		}
		log.Printf("[DEBUG] http request = %s", outbytes)
	}
	httpResponse, err := c.httpClient.Do(httpRequest)
	if err != nil {
		msg := fmt.Sprintf("Fail to get response because %s", err)
		log.Println(msg)
		return errors.New("ClientError.NetworkError")
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
	c.authMethod = "Basic"
	c.debug = false
	log.SetFlags(log.LstdFlags | log.Lshortfile)
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

func (c *Client) WithSignatureMethod(method string) *Client {
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
