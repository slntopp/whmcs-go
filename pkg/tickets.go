package whmcs

import "encoding/base64"

type TicketsService struct {
	wc *Client
}

func (c *Client) initTicketsService() {
	c.Tickets = TicketsService{wc: c}
}

func (s *TicketsService) GetSupportDepartments(req *GetSupportDepartmentsRequest) (*GetSupportDepartmentsResponse, error) {
	raw := make(map[string]interface{})
	if err := s.wc.call("GetSupportDepartments", req, &raw); err != nil {
		return nil, err
	}

	raw["departments"] = raw["departments"].(map[string]interface{})["department"]

	res, err := toStruct[GetSupportDepartmentsResponse](raw)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *TicketsService) GetSupportStatuses(req *GetSupportStatusesRequest) (*GetSupportStatusesResponse, error) {
	raw := make(map[string]interface{})
	if err := s.wc.call("GetSupportStatuses", req, &raw); err != nil {
		return nil, err
	}

	raw["statuses"] = raw["statuses"].(map[string]interface{})["status"]

	res, err := toStruct[GetSupportStatusesResponse](raw)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *TicketsService) GetTicket(req *GetTicketRequest) (*GetTicketResponse, error) {
	raw := make(map[string]interface{})
	if err := s.wc.call("GetTicket", req, &raw); err != nil {
		return nil, err
	}

	raw["replies"] = raw["replies"].(map[string]interface{})["reply"]
	for _, v := range raw["replies"].([]any) {
		v := v.(map[string]any)
		if v["attachment"] == "" {
			v["attachments"] = nil
		}
		// if _, ok := v["attachments"].([]any)[0].([]any); ok {
		// 	v["attachments"] = nil
		// }
	}
	raw["notes"] = raw["notes"].(map[string]interface{})["note"]
	for _, v := range raw["replies"].([]any) {
		v := v.(map[string]any)
		if v["attachment"] == "" {
			v["attachments"] = nil
		}
	}

	res, err := toStruct[GetTicketResponse](raw)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *TicketsService) GetTicketAttachment(req *GetTicketAttachmentRequest) (*GetTicketAttachmentResponse, error) {
	raw := make(map[string]interface{})
	if err := s.wc.call("GetTicketAttachment", req, &raw); err != nil {
		return nil, err
	}

	data, err := base64.StdEncoding.DecodeString(raw["data"].(string))
	if err != nil {
		return nil, err
	}
	raw["data"] = data

	res, err := toStruct[GetTicketAttachmentResponse](raw)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *SupportService) GetTicketCounts(req *GetTicketCountsRequest) (*GetTicketCountsResponse, error) {
	res := &GetTicketCountsResponse{}
	if err := s.wc.call("GetTicketCounts", req, &res); err != nil {
		return nil, err
	}

	return res, nil
}

func (s *TicketsService) GetTicketPredefinedCats(req *GetTicketPredefinedCatsRequest) (*GetTicketPredefinedCatsResponse, error) {
	raw := make(map[string]interface{})
	if err := s.wc.call("GetTicketPredefinedCats", req, &raw); err != nil {
		return nil, err
	}

	raw["categories"] = raw["categories"].(map[string]interface{})["category"]

	res, err := toStruct[GetTicketPredefinedCatsResponse](raw)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *TicketsService) GetTicketPredefinedReplies(req *GetTicketPredefinedRepliesRequest) (*GetTicketPredefinedRepliesResponse, error) {
	raw := make(map[string]interface{})
	if err := s.wc.call("GetTicketPredefinedReplies", req, &raw); err != nil {
		return nil, err
	}

	raw["predefinedreplies"] = raw["predefinedreplies"].(map[string]interface{})["predefinedreply"]

	res, err := toStruct[GetTicketPredefinedRepliesResponse](raw)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *TicketsService) GetTickets(req *GetTicketsRequest) (*GetTicketsResponse, error) {
	raw := make(map[string]interface{})
	if err := s.wc.call("GetTickets", req, &raw); err != nil {
		return nil, err
	}

	raw["tickets"] = raw["tickets"].(map[string]interface{})["ticket"]

	res, err := toStruct[GetTicketsResponse](raw)
	if err != nil {
		return nil, err
	}

	return res, nil
}
