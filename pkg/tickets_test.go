package whmcs

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestClient_GetSupportDepartments(t *testing.T) {
	request := &GetSupportDepartmentsRequest{}

	response := &GetSupportDepartmentsResponse{
		ApiResponse: ApiResponse{
			Result:  "success",
			Message: nil,
		},
		Total: Int(1),
		Departments: []Department{
			{
				Id:            Int(1),
				Name:          String("Support"),
				AwaitingReply: Int(5),
				OpenTickets:   Int(10),
			},
		},
	}

	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseMultipartForm(r.ContentLength)
		if err != nil {
			t.Fatal(err)
		}

		if r.MultipartForm.Value["action"][0] != "GetSupportDepartments" {
			t.Errorf("Expected action to be 'GetSupportDepartments', got '%s'", r.FormValue("action"))
		}

		_json := []byte(`{
			"result": "success",
			"totalresults": 1,
			"departments": {
				"department": [
					{
						"id": 1,
						"name": "Support",
						"awaitingreply": 5,
						"opentickets": 10
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

	result, err := client.Tickets.GetSupportDepartments(request)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, response, result)
}

func TestClient_GetSupportStatuses(t *testing.T) {
	request := &GetSupportStatusesRequest{}

	response := &GetSupportStatusesResponse{
		ApiResponse: ApiResponse{
			Result:  "success",
			Message: nil,
		},
		Total: Int(1),
		Statuses: []Status{
			{
				Title: String("Open"),
				Count: Int(12),
				Color: String("#779500"),
			},
			{
				Title: String("Answered"),
				Count: Int(43),
				Color: String("#000000"),
			},
			{
				Title: String("Customer-Reply"),
				Count: Int(6),
				Color: String("#ff6600"),
			},
			{
				Title: String("On Hold"),
				Count: Int(0),
				Color: String("#224488"),
			},
			{
				Title: String("In Progress"),
				Count: Int(3),
				Color: String("#cc0000"),
			},
			{
				Title: String("Closed"),
				Count: Int(3562),
				Color: String("#888888"),
			},
		},
	}

	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseMultipartForm(r.ContentLength)
		if err != nil {
			t.Fatal(err)
		}

		if r.MultipartForm.Value["action"][0] != "GetSupportStatuses" {
			t.Errorf("Expected action to be 'GetSupportStatuses', got '%s'", r.FormValue("action"))
		}

		_json := []byte(`{
    "result": "success",
    "totalresults": 1,
    "statuses": {
        "status": [
            {
                "title": "Open",
                "count": 12,
                "color": "#779500"
            },
            {
                "title": "Answered",
                "count": 43,
                "color": "#000000"
            },
            {
                "title": "Customer-Reply",
                "count": 6,
                "color": "#ff6600"
            },
            {
                "title": "On Hold",
                "count": 0,
                "color": "#224488"
            },
            {
                "title": "In Progress",
                "count": 3,
                "color": "#cc0000"
            },
            {
                "title": "Closed",
                "count": 3562,
                "color": "#888888"
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

	result, err := client.Tickets.GetSupportStatuses(request)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, response, result)
}

func TestTicketsService_GetTicket(t *testing.T) {
	request := &GetTicketRequest{}

	response := &GetTicketResponse{
		ApiResponse: ApiResponse{
			Result:  "success",
			Message: nil,
		},
		TicketId: Int(1),
		Ticket: Ticket{
			DeptName:      String("Sample Support Department"),
			OpenedDate:    Time(time.Date(2016, time.January, 1, 6, 26, 29, 0, time.UTC)),
			LastReplyDate: Time(time.Date(2016, time.January, 1, 6, 30, 16, 0, time.UTC)),
			Replies: []Reply{
				{
					ReplyId: Int(0),
					Name:    String("Cynthia Reilly"),
					Date:    Time(time.Date(2016, time.January, 1, 6, 26, 29, 0, time.UTC)),
					Attachments: []IndexedAttachment{
						{
							Filename: String("attachment_name.png"),
							Index:    Int(0),
						},
					},
					Admin:              String(""),
					AttachmentsRemoved: true,
				},
				{
					ReplyId:            Int(1),
					Name:               String(""),
					Date:               Time(time.Date(2016, time.January, 1, 6, 27, 1, 0, time.UTC)),
					Attachment:         String(""),
					Attachments:        nil,
					AttachmentsRemoved: false,
					Admin:              String("Demo Admin"),
					Rating:             Int(0),
				},
			},
			Notes: []Note{
				{
					NoteId:             Int(1),
					Date:               Time(time.Date(2016, time.January, 1, 6, 26, 42, 0, time.UTC)),
					Message:            String("This is a ticket note"),
					Attachment:         String(""),
					Attachments:        nil,
					AttachmentsRemoved: false,
					Admin:              String("Demo Admin"),
				},
			},
		},
	}

	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseMultipartForm(r.ContentLength)
		if err != nil {
			t.Fatal(err)
		}

		if r.MultipartForm.Value["action"][0] != "GetTicket" {
			t.Errorf("Expected action to be 'GetTicket', got '%s'", r.FormValue("action"))
		}

		_json := []byte(`{
    "result": "success",
    "ticketid": 1,
    "deptname": "Sample Support Department",
    "date": "2016-01-01 06:26:29",
    "lastreply": "2016-01-01 06:30:16",
    "replies": {
        "reply": [
            {
                "replyid": "0",
                "name": "Cynthia Reilly",
                "date": "2016-01-01 06:26:29",
                "attachments": [
                    {
                        "filename": "attachment_name.png",
                        "index": 0
                    }
                ],
                "attachments_removed": true,
                "admin": ""
            },
            {
                "replyid": 1,
                "name": "",
                "date": "2016-01-01 06:27:01",
                "attachment": "",
                "attachments": [
                    []
                ],
                "attachments_removed": false,
                "admin": "Demo Admin",
                "rating": 0
            }
        ]
    },
    "notes": {
        "note": [
            {
                "noteid": 1,
                "date": "2016-01-01 06:26:42",
                "message": "This is a ticket note",
                "attachment": "",
                "attachments": [
                    []
                ],
                "attachments_removed": false,
                "admin": "Demo Admin"
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

	result, err := client.Tickets.GetTicket(request)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, response, result)
}
