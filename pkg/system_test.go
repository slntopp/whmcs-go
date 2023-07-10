package whmcs

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/elliotchance/phpserialize"
	"github.com/stretchr/testify/assert"
)

func TestSystemService_GetEmailTemplates(t *testing.T) {
	request := &GetEmailTemplatesRequest{}

	response := &GetEmailTemplatesResponse{
		ApiResponse: ApiResponse{
			Result:  "success",
			Message: nil,
		},
		Total: Int(2),
		EmailTemplates: []EmailTemplate{
			{
				Id:      Int(38),
				Name:    String("Automated Password Reset"),
				Subject: String("Your new password for {$company_name}"),
				Custom:  Bool(false),
			},
			{
				Id:      Int(64),
				Name:    String("Client Email Address Verification"),
				Subject: String("Confirm Your Registration"),
				Custom:  Bool(false),
			},
		},
	}

	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseMultipartForm(r.ContentLength)
		if err != nil {
			t.Fatal(err)
		}

		if r.MultipartForm.Value["action"][0] != "GetEmailTemplates" {
			t.Errorf("Expected action to be 'GetEmailTemplates', got '%s'", r.FormValue("action"))
		}

		_json := []byte(`{
			"result": "success",
			"totalresults": 2,
			"emailtemplates": {
				"emailtemplate": [
					{
						"id": 38,
						"name": "Automated Password Reset",
						"subject": "Your new password for {$company_name}",
						"custom": false
					},
					{	
						"id": 64,
						"name": "Client Email Address Verification",
						"subject": "Confirm Your Registration",
						"custom": false
					}
				]
			}
}`)

		_, _ = w.Write(_json)
	}))
	defer mockServer.Close()

	client, err := NewClient(mockServer.URL, "username", "password", false)
	if err != nil {
		t.Fatal(err)
	}

	result, err := client.System.GetEmailTemplates(request)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, response, result)
}

func TestSystemService_SendEmail(t *testing.T) {
	request := &SendEmailRequest{
		CustomVars: map[string]any{
			"test":  int64(1),
			"test2": "2",
		},
	}

	response := &SendEmailResponse{
		ApiResponse: ApiResponse{
			Result:  "success",
			Message: nil,
		},
	}

	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseMultipartForm(r.ContentLength)
		if err != nil {
			t.Fatal(err)
		}

		if r.MultipartForm.Value["action"][0] != "SendEmail" {
			t.Errorf("Expected action to be 'SendEmail', got '%s'", r.FormValue("action"))
		}

		dec, err := base64.StdEncoding.DecodeString(r.MultipartForm.Value["customvars"][0])
		assert.NoError(t, err)
		vars := make(map[any]any)
		assert.NoError(t, phpserialize.Unmarshal(dec, &vars))
		res := make(map[string]any)
		for k, v := range vars {
			res[fmt.Sprint(k)] = v
		}
		assert.Equal(t, request.CustomVars, res)

		_json := []byte(`{
			"result": "success"
		}`)

		_, _ = w.Write(_json)
	}))
	defer mockServer.Close()

	client, err := NewClient(mockServer.URL, "username", "password", false)
	if err != nil {
		t.Fatal(err)
	}

	result, err := client.System.SendEmail(request)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, response, result)
}

func TestSystemService_TriggerNotificationEvent(t *testing.T) {
	cases := []struct {
		Request *TriggerNotificationEventRequest
		isErr   bool
	}{
		{
			Request: &TriggerNotificationEventRequest{
				Attributes: []Attribute{
					{
						Label:                String("key"),
						Value:                String("value"),
						AdditionalAttributes: map[string]any{"label": "test"},
					},
				},
			},
			isErr: true,
		},
		{
			Request: &TriggerNotificationEventRequest{
				Attributes: []Attribute{
					{
						Label:                String("key"),
						Value:                String("value"),
						AdditionalAttributes: map[string]any{"some": "test"},
					},
				},
			},
			isErr: false,
		},
	}

	response := &TriggerNotificationEventResponse{
		ApiResponse: ApiResponse{
			Result:  "success",
			Message: nil,
		},
	}

	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseMultipartForm(r.ContentLength)
		if err != nil {
			t.Fatal(err)
		}

		if r.MultipartForm.Value["action"][0] != "TriggerNotificationEvent" {
			t.Errorf("Expected action to be 'TriggerNotificationEvent', got '%s'", r.FormValue("action"))
		}

		_json := []byte(`{
			"result": "success"
		}`)

		_, _ = w.Write(_json)
	}))
	defer mockServer.Close()

	client, err := NewClient(mockServer.URL, "username", "password", false)
	if err != nil {
		t.Fatal(err)
	}

	for _, v := range cases {
		result, err := client.System.TriggerNotificationEvent(v.Request)
		if v.isErr {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
			assert.Equal(t, response, result)
		}
	}
}
