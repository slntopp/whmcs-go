package whmcs

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
)

type Client struct {
	api      string
	username string
	password string

	client *http.Client

	Support SupportService
	Tickets TicketsService
}

func NewClient(api, username, password string, dangerMode bool) (*Client, error) {
	if api == "" || username == "" || password == "" {
		return nil, fmt.Errorf("empty credentials")
	}

	client := &http.Client{}
	if dangerMode {
		client.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		}
	}

	c := &Client{
		api:      api,
		username: username,
		password: password,
		client:   client,
	}

	c.initSupportService()
	c.initTicketsService()

	return c, nil
}

type ApiResponse struct {
	// The result of the operation: success or error
	Result string `json:"result"`

	// Error message. Nil if success
	Message *string `json:"message"`
}

func toMap(in any) (map[string]interface{}, error) {
	_json, err := json.Marshal(in)
	if err != nil {
		return nil, err
	}
	out := make(map[string]interface{})
	if err := json.Unmarshal(_json, &out); err != nil {
		return nil, err
	}
	return out, nil
}

func toStruct[T any](in map[string]interface{}) (*T, error) {
	_json, err := json.Marshal(in)
	if err != nil {
		return nil, err
	}
	out := new(T)
	if err := json.Unmarshal(_json, &out); err != nil {
		return nil, err
	}
	return out, nil
}

func (c *Client) call(action string, apiReq, apiRes any) error {
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	writer.WriteField("username", c.username)
	writer.WriteField("password", c.password)
	writer.WriteField("action", action)
	writer.WriteField("responsetype", "json")

	params, ok := apiReq.(map[string]interface{})
	var err error
	if !ok && apiReq != nil {
		params, err = toMap(apiReq)
		if err != nil {
			return err
		}
	}
	for k, v := range params {
		writer.WriteField(k, fmt.Sprint(v))
	}

	if err := writer.Close(); err != nil {
		return nil
	}

	req, err := http.NewRequest("POST", c.api, payload)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	errRes := ApiResponse{}
	if err := json.Unmarshal(body, &errRes); err != nil {
		return err
	}

	if errRes.Message != nil {
		return fmt.Errorf("%s error: %s", action, *errRes.Message)
	}

	if apiRes != nil {
		return json.Unmarshal(body, apiRes)
	}
	return nil
}
