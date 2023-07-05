package whmcs

import (
	"bytes"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	assert := assert.New(t)

	c, err := NewClient("test", "test", "test", false)
	assert.NoError(err)
	assert.NotNil(c)
	assert.NotNil(c.Authentication)
	assert.NotNil(c.Support)
	assert.NotNil(c.Tickets)
	assert.NotNil(c.System)

	c, err = NewClient("", "", "", false)
	assert.Error(err)
	assert.Nil(c)
}

type MockTransport struct {
}

func (t *MockTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.ParseMultipartForm(req.ContentLength)

	if req.MultipartForm.Value["action"][0] == "testSuccess" {
		body := []byte(`{"result": "success", "message": null}`)
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewBuffer(body)),
		}, nil
	} else {
		body := []byte(`{"result": "error", "message": "some error"}`)
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewBuffer(body)),
		}, nil
	}
}

func TestCall(t *testing.T) {
	assert := assert.New(t)

	client := &Client{
		api:      "https://api.whmcs.com",
		username: "username",
		password: "password",
		client:   &http.Client{Transport: &MockTransport{}},
	}

	apiReq := map[string]interface{}{"param1": "value1"}
	apiRes := make(map[string]interface{})

	err := client.Call("testSuccess", apiReq, &apiRes)
	assert.NoError(err)

	expectedRes := map[string]interface{}{
		"result":  "success",
		"message": nil,
	}
	assert.Equal(expectedRes, apiRes)

	apiRes = make(map[string]interface{})

	err = client.Call("testError", apiReq, &apiRes)
	assert.Error(err)
	assert.Equal(map[string]interface{}{}, apiRes)
}
