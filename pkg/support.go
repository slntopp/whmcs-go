package whmcs

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
)

type SupportService struct {
	wc *Client
}

func (c *Client) initSupportService() {
	c.Support = SupportService{wc: c}
}

func (s *SupportService) AddAnnouncement(req *AddAnnouncementRequest) (*AddAnnouncementResponse, error) {
	res := &AddAnnouncementResponse{}
	if err := s.wc.call("AddAnnouncement", req, &res); err != nil {
		return nil, err
	}

	return res, nil
}

func (s *SupportService) AddCancelRequest(req *AddCancelRequest_Request) (*AddCancelRequest_Response, error) {
	res := &AddCancelRequest_Response{}
	if err := s.wc.call("AddCancelRequest", req, &res); err != nil {
		return nil, err
	}

	return res, nil
}

func (s *SupportService) AddClientNote(req *AddClientNoteRequest) (*AddClientNoteResponse, error) {
	res := &AddClientNoteResponse{}
	if err := s.wc.call("AddClientNote", req, &res); err != nil {
		return nil, err
	}

	return res, nil
}

func (s *SupportService) AddTicketNote(req *AddTicketNoteRequest) (*AddTicketNoteResponse, error) {
	attachs := make([]interface{}, len(req.Attachments))
	for i, v := range req.Attachments {
		attachs[i] = map[string]interface{}{
			"name": v.Name,
			"data": base64.StdEncoding.EncodeToString(v.Data),
		}
	}

	_json, err := json.Marshal(attachs)
	if err != nil {
		return nil, err
	}

	// req.Attachments = nil
	params, err := toMap(req)
	if err != nil {
		return nil, err
	}
	params["attachments"] = base64.StdEncoding.EncodeToString(_json)

	res := &AddTicketNoteResponse{}
	if err := s.wc.call("AddTicketNote", params, &res); err != nil {
		return nil, err
	}

	return res, nil
}

func (s *SupportService) AddTicketReply(req *AddTicketReplyRequest) (*AddTicketReplyResponse, error) {
	attachs := make([]interface{}, len(req.Attachments))
	for i, v := range req.Attachments {
		attachs[i] = map[string]interface{}{
			"name": v.Name,
			"data": base64.StdEncoding.EncodeToString(v.Data),
		}
	}

	// req.Attachments = nil
	params, err := toMap(req)
	if err != nil {
		return nil, err
	}

	_json, err := json.Marshal(attachs)
	if err != nil {
		return nil, err
	}
	params["attachments"] = base64.StdEncoding.EncodeToString(_json)

	res := &AddTicketReplyResponse{}
	if err := s.wc.call("AddTicketReply", params, &res); err != nil {
		return nil, err
	}

	return res, nil
}

func (s *SupportService) BlockTicketSender(req *BlockTicketSenderRequest) (*BlockTicketSenderResponse, error) {
	res := &BlockTicketSenderResponse{}
	if err := s.wc.call("BlockTicketSender", req, &res); err != nil {
		return nil, err
	}

	return res, nil
}

func (s *SupportService) DeleteAnnouncement(req *DeleteAnnouncementRequest) (*DeleteAnnouncementResponse, error) {
	res := &DeleteAnnouncementResponse{}
	if err := s.wc.call("DeleteAnnouncement", req, &res); err != nil {
		return nil, err
	}

	return res, nil
}

func (s *SupportService) DeleteTicket(req *DeleteTicketRequest) (*DeleteTicketResponse, error) {
	res := &DeleteTicketResponse{}
	if err := s.wc.call("DeleteTicket", req, &res); err != nil {
		return nil, err
	}

	return res, nil
}

func (s *SupportService) DeleteTicketNote(req *DeleteTicketNoteRequest) (*DeleteTicketNoteResponse, error) {
	res := &DeleteTicketNoteResponse{}
	if err := s.wc.call("DeleteTicketNote", req, &res); err != nil {
		return nil, err
	}

	return res, nil
}

func (s *SupportService) GetAnnouncements(req *GetAnnouncementsRequest) (*GetAnnouncementsResponse, error) {
	raw := make(map[string]interface{})
	if err := s.wc.call("GetAnnouncements", req, &raw); err != nil {
		return nil, err
	}

	raw["announcements"] = raw["announcements"].(map[string]interface{})["announcement"]
	// for _, v := range raw["announcements"].([]interface{}) {
	// 	v := v.(map[string]interface{})
	// 	if v["published"] == 1.0 {
	// 		v["published"] = true
	// 	} else {
	// 		v["published"] = false
	// 	}
	// }

	res, err := toStruct[GetAnnouncementsResponse](raw)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *SupportService) MergeTicket(req *MergeTicketRequest) (*MergeTicketResponse, error) {
	var ids []string
	for _, v := range req.MergeTicketIds {
		ids = append(ids, fmt.Sprint(v))
	}

	params, err := toMap(req)
	if err != nil {
		return nil, err
	}
	params["mergeticketids"] = strings.Join(ids, ",")

	res := &MergeTicketResponse{}
	if err := s.wc.call("MergeTicket", req, &res); err != nil {
		return nil, err
	}

	return res, nil
}

func (s *SupportService) OpenTicket(req *OpenTicketRequest) (*OpenTicketResponse, error) {
	attachs := make([]interface{}, len(req.Attachments))
	for i, v := range req.Attachments {
		attachs[i] = map[string]interface{}{
			"name": v.Name,
			"data": base64.StdEncoding.EncodeToString(v.Data),
		}
	}

	// req.Attachments = nil
	params, err := toMap(req)
	if err != nil {
		return nil, err
	}

	_json, err := json.Marshal(attachs)
	if err != nil {
		return nil, err
	}
	params["attachments"] = base64.StdEncoding.EncodeToString(_json)

	res := &OpenTicketResponse{}
	if err := s.wc.call("OpenTicket", params, &res); err != nil {
		return nil, err
	}

	return res, nil
}

func (s *SupportService) UpdateTicket(req *UpdateTicketRequest) (*UpdateTicketResponse, error) {
	res := &UpdateTicketResponse{}
	if err := s.wc.call("UpdateTicket", req, &res); err != nil {
		return nil, err
	}

	return res, nil
}
