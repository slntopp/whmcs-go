package whmcs

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSupportService_GetAnnouncements(t *testing.T) {
	request := &GetAnnouncementsRequest{}

	zeroTime := formattedTime("0000-00-00 00:00:00")
	response := &GetAnnouncementsResponse{
		ApiResponse: ApiResponse{
			Result:  "success",
			Message: nil,
		},
		Total:       Int(1),
		StartNum:    Int(0),
		NumReturned: Int(1),
		Announcements: []Announcement{
			{
				Id:               Int(1),
				Date:             Time(time.Date(2016, 2, 24, 21, 27, 4, 0, time.UTC)),
				Title:            String("Thank you for choosing WHMCS!"),
				AnnouncementText: String("<p>Welcome to <a title=\"WHMCS\" href=\"https://whmcs.com\" target=\"_blank\">WHMCS</a>! You have made a great choice and we want to help you get up and running as quickly as possible.</p><p>This is a sample announcement. Announcements are a great way to keep your customers informed about news and special offers. You can edit or delete this announcement by logging into the admin area and navigating to <em>Support &gt; Announcements</em>.</p><p>If at any point you get stuck, our support team is available 24x7 to assist you. Simply visit <a title=\"www.whmcs.com/support\" href=\"https://www.whmcs.com/support\" target=\"_blank\">www.whmcs.com/support</a> to request assistance.</p>"),
				Published:        BoolInt(true),
				ParentId:         Int(0),
				Language:         String(""),
				CreatedAt:        &zeroTime,
				UpdatedAt:        &zeroTime,
			},
		},
	}

	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseMultipartForm(r.ContentLength)
		assert.NoError(t, err)
		// if err != nil {
		// 	t.Fatal(err)
		// }

		assert.Equal(t, "GetAnnouncements", r.MultipartForm.Value["action"][0])
		// if r.MultipartForm.Value["action"][0] != "GetAnnouncements" {
		// 	t.Errorf("Expected action to be 'GetAnnouncements', got '%s'", r.FormValue("action"))
		// }

		_json := []byte(`{
			"result": "success",
			"totalresults": 1,
			"startnumber": 0,
			"numreturned": 1,
			"announcements": {
				"announcement": [
					{
						"id": 1,
						"date": "2016-02-24 21:27:04",
						"title": "Thank you for choosing WHMCS!",
						"announcement": "<p>Welcome to <a title=\"WHMCS\" href=\"https:\/\/whmcs.com\" target=\"_blank\">WHMCS<\/a>! You have made a great choice and we want to help you get up and running as quickly as possible.<\/p><p>This is a sample announcement. Announcements are a great way to keep your customers informed about news and special offers. You can edit or delete this announcement by logging into the admin area and navigating to <em>Support &gt; Announcements<\/em>.<\/p><p>If at any point you get stuck, our support team is available 24x7 to assist you. Simply visit <a title=\"www.whmcs.com\/support\" href=\"https:\/\/www.whmcs.com\/support\" target=\"_blank\">www.whmcs.com\/support<\/a> to request assistance.<\/p>",
						"published": 1,
						"parentid": 0,
						"language": "",
						"created_at": "0000-00-00 00:00:00",
						"updated_at": "0000-00-00 00:00:00"
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

	result, err := client.Support.GetAnnouncements(request)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotPanics(t, func() {
		result.Announcements[0].CreatedAt.Parse()
	})

	assert.Equal(t, response, result)
}
