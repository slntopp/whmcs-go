package whmcs

import (
	"encoding/base64"
	"fmt"

	"github.com/elliotchance/phpserialize"
)

type SystemService struct {
	wc *Client
}

func (c *Client) initSystemService() {
	c.System = SystemService{wc: c}
}

func (s *SystemService) GetEmailTemplates(req *GetEmailTemplatesRequest) (*GetEmailTemplatesResponse, error) {
	raw := make(map[string]any)
	if err := s.wc.Call("GetEmailTemplates", req, &raw); err != nil {
		return nil, err
	}

	t, ok := raw["emailtemplates"].(map[string]any)
	if ok {
		raw["emailtemplates"] = t["emailtemplate"]
	}

	res, err := toStruct[GetEmailTemplatesResponse](raw)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *SystemService) SendEmail(req *SendEmailRequest) (*SendEmailResponse, error) {
	params, err := toMap(req)
	if err != nil {
		return nil, err
	}

	bytes, err := phpserialize.Marshal(req.CustomVars, nil)
	if err != nil {
		return nil, err
	}
	params["customvars"] = base64.StdEncoding.EncodeToString(bytes)

	res := &SendEmailResponse{}
	if err := s.wc.Call("SendEmail", params, &res); err != nil {
		return nil, err
	}

	return res, nil
}

func (s *SystemService) SendAdminEmail(req *SendAdminEmailRequest) (*SendAdminEmailResponse, error) {
	params, err := toMap(req)
	if err != nil {
		return nil, err
	}

	bytes, err := phpserialize.Marshal(req.MergeFields, nil)
	if err != nil {
		return nil, err
	}
	params["mergefields"] = base64.StdEncoding.EncodeToString(bytes)

	res := &SendAdminEmailResponse{}
	if err := s.wc.Call("SendAdminEmail", params, &res); err != nil {
		return nil, err
	}

	return res, nil
}

func (s *SystemService) TriggerNotificationEvent(req *TriggerNotificationEventRequest) (*TriggerNotificationEventResponse, error) {
	params, err := toMap(req)
	if err != nil {
		return nil, err
	}

	attrs := make([]any, len(req.Attributes))
	for i, a := range req.Attributes {
		attr, err := toMap(a)
		if err != nil {
			return nil, err
		}
		for k, v := range a.AdditionalAttributes {
			if k == "label" || k == "value" {
				return nil, fmt.Errorf("additional attribute with key 'label' or 'value' not allowed")
			}
			attr[k] = v
		}
		attrs[i] = attr
	}
	params["attributes"] = attrs

	res := &TriggerNotificationEventResponse{}
	if err := s.wc.Call("TriggerNotificationEvent", params, &res); err != nil {
		return nil, err
	}

	return res, nil
}
