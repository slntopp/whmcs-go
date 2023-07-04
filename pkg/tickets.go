package whmcs

import (
	"encoding/base64"
)

type TicketsService struct {
	wc *Client
}

func (c *Client) initTicketsService() {
	c.Tickets = TicketsService{wc: c}
}

func (s *TicketsService) GetSupportDepartments(req *GetSupportDepartmentsRequest) (*GetSupportDepartmentsResponse, error) {
	raw := make(map[string]any)
	if err := s.wc.Call("GetSupportDepartments", req, &raw); err != nil {
		return nil, err
	}

	t, ok := raw["departments"].(map[string]any)
	if ok {
		raw["departments"] = t["department"]
	}

	res, err := toStruct[GetSupportDepartmentsResponse](raw)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *TicketsService) GetSupportStatuses(req *GetSupportStatusesRequest) (*GetSupportStatusesResponse, error) {
	raw := make(map[string]any)
	if err := s.wc.Call("GetSupportStatuses", req, &raw); err != nil {
		return nil, err
	}

	t, ok := raw["statuses"].(map[string]any)
	if ok {
		raw["statuses"] = t["status"]
	}

	res, err := toStruct[GetSupportStatusesResponse](raw)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *TicketsService) GetTicket(req *GetTicketRequest) (*GetTicketResponse, error) {
	raw := make(map[string]any)
	if err := s.wc.Call("GetTicket", req, &raw); err != nil {
		return nil, err
	}

	r, ok := raw["replies"].(map[string]any)
	if ok {
		raw["replies"] = r["reply"]
		if rr, ok := raw["replies"].([]any); ok {
			for _, v := range rr {
				v := v.(map[string]any)
				if v["attachment"] == "" {
					v["attachments"] = nil
				}
				if v["replyid"] == "0" {
					v["replyid"] = 0
				}
			}
		}
	}

	n, ok := raw["notes"].(map[string]any)
	if ok {
		raw["notes"] = n["note"]
		if nn, ok := raw["notes"].([]any); ok {
			for _, v := range nn {
				v := v.(map[string]any)
				if v["attachment"] == "" {
					v["attachments"] = nil
				}
				if v["noteid"] == "0" {
					v["noteid"] = 0
				}
			}
		}
	}

	res, err := toStruct[GetTicketResponse](raw)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *TicketsService) GetTicketAttachment(req *GetTicketAttachmentRequest) (*GetTicketAttachmentResponse, error) {
	raw := make(map[string]any)
	if err := s.wc.Call("GetTicketAttachment", req, &raw); err != nil {
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

func (s *TicketsService) GetTicketCounts(req *GetTicketCountsRequest) (*GetTicketCountsResponse, error) {
	res := &GetTicketCountsResponse{}
	if err := s.wc.Call("GetTicketCounts", req, &res); err != nil {
		return nil, err
	}

	return res, nil
}

func (s *TicketsService) GetTicketNotes(req *GetTicketNotesRequest) (*GetTicketNotesResponse, error) {
	raw := make(map[string]any)
	if err := s.wc.Call("GetTicketNotes", req, &raw); err != nil {
		return nil, err
	}

	t, ok := raw["notes"].(map[string]any)
	if ok {
		raw["notes"] = t["note"]
		if nn, ok := raw["notes"].([]any); ok {
			for i, v := range nn {
				v, ok := v.(map[string]any)
				if !ok {
					continue
				}
				if v["attachments"] == "" {
					v["attachments"] = nil
				}
				if v["id"] == "0" {
					v["id"] = 0
				}
				nn[i] = v
			}
			raw["notes"] = nn
		}
	}

	res, err := toStruct[GetTicketNotesResponse](raw)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *TicketsService) GetTicketPredefinedCats(req *GetTicketPredefinedCatsRequest) (*GetTicketPredefinedCatsResponse, error) {
	raw := make(map[string]any)
	if err := s.wc.Call("GetTicketPredefinedCats", req, &raw); err != nil {
		return nil, err
	}

	t, ok := raw["categories"].(map[string]any)
	if ok {
		raw["categories"] = t["category"]
	}

	res, err := toStruct[GetTicketPredefinedCatsResponse](raw)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *TicketsService) GetTicketPredefinedReplies(req *GetTicketPredefinedRepliesRequest) (*GetTicketPredefinedRepliesResponse, error) {
	raw := make(map[string]any)
	if err := s.wc.Call("GetTicketPredefinedReplies", req, &raw); err != nil {
		return nil, err
	}

	t, ok := raw["predefinedreplies"].(map[string]any)
	if ok {
		raw["predefinedreplies"] = t["predefinedreply"]
	}

	res, err := toStruct[GetTicketPredefinedRepliesResponse](raw)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *TicketsService) GetTickets(req *GetTicketsRequest) (*GetTicketsResponse, error) {
	raw := make(map[string]any)
	if err := s.wc.Call("GetTickets", req, &raw); err != nil {
		return nil, err
	}

	t, ok := raw["tickets"].(map[string]any)
	if ok {
		raw["tickets"] = t["ticket"]
	}

	res, err := toStruct[GetTicketsResponse](raw)
	if err != nil {
		return nil, err
	}

	return res, nil
}
