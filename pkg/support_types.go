package whmcs

type AddAnnouncementRequest struct {
	// Date in the format YYYY-MM-DD HH:MM:SS
	// 	Required
	Date FormattedTime `json:"date"`

	// 	Required
	Title string `json:"title"`

	// Announcement text.
	// 	Required
	AnnouncementText string `json:"announcement"`

	// Pass as true to publish.
	// 	Optional
	Published Boolean `json:"published"`
}

type AddAnnouncementResponse struct {
	ApiResponse

	// The id of the new announcement
	AnnouncementId int `json:"announcementid"`
}

type AddCancelRequestType string

const (
	IMMEDIATE             AddCancelRequestType = "Immediate"
	END_OF_BILLING_PEROID AddCancelRequestType = "End of Billing Period"
)

type AddCancelRequest_Request struct {
	// The Service ID to cancel
	// 	Required
	ServieId int `json:"serviceid"`

	// The type of cancellation. 'Immediate' or 'End of Billing Period'
	// 	Optional
	Type AddCancelRequestType `json:"type"`

	// The customer reason for cancellation
	// 	Optional
	Reason string `json:"reason"`
}

type AddCancelRequest_Response struct {
	ApiResponse

	// The id of the service the request was for
	ServiceId int `json:"serviceid"`

	// The id of the user the service belongs to
	UserId int `json:"userid"`
}

type AddClientNoteRequest struct {
	// The Client ID to apply the note to
	// 	Required
	UserId int `json:"userid"`

	// The note to add
	// 	Required
	NoteText string `json:"notes"`

	// Should the note be made sticky. Makes the note 'sticky' and displays the note throughout the client’s
	// account and on any tickets they submit in the admin area
	// 	Optional
	Sticky Boolean `json:"sticky"`
}

type AddClientNoteResponse struct {
	ApiResponse

	// 	The id of the newly created note
	NoteId int `json:"noteid"`
}

type Attachment struct {
	// File name
	// 	Required
	Name string `json:"name"`

	// File data
	// 	Required
	Data []byte `json:"data"`
}

type AddTicketNoteRequest struct {
	// The content of the ticket note
	// 	Required
	MessageText string `json:"message"`

	// The Client Ticket Number ID to apply the note to
	// 	Optional
	TicketNum string `json:"ticketnum"`

	// The id of the ticket in the database. Either TicketNum or TicketId is required
	// 	Optional
	TicketId int `json:"ticketid"`

	// Should markdown be used on the ticket note output
	// 	Optional
	Markdown Boolean `json:"markdown"`

	// Optional array of file attachments.
	// 	Optional
	Attachments []Attachment `json:"-"`

	// The date and time the ticket note will show as created. Format: ISO8601 or YYYY-MM-DD HH:mm:ss.
	// Ticket creation date cannot be in the future
	// 	Optional
	CreatedDate FormattedTime `json:"created"`
}

type AddTicketNoteResponse struct {
	ApiResponse
}

type AddTicketReplyRequest struct {
	// The id of the ticket.
	// 	Required
	TicketId int `json:"ticketid"`

	// The content of the ticket reply
	// 	Required
	MessageText string `json:"message"`

	// Should markdown be used on the ticket note output
	// 	Optional
	Markdown Boolean `json:"markdown"`

	// Pass a clientid to associate the ticket reply with a specific client
	// 	Optional
	ClientId int `json:"clientid"`

	// Pass a contactid to associate the ticket reply with a specific contact belonging to ClientId
	// 	Optional
	ContactId int `json:"contactid"`

	// The admin username to associate the ticket reply with
	// 	Optional
	AdminUsername string `json:"adminusername"`

	// The name to associate with the ticket reply if not an admin or client response
	// 	Optional
	Name string `json:"name"`

	// The email to associate with the ticket reply if not an admin or client response
	// 	Optional
	Email string `json:"email"`

	// The status to set on the ticket after the reply is made if the default status on admin/client response is not required. See GetSupportStatuses API command
	// 	Optional
	Status string `json:"status"`

	// Set to true to stop the ticket reply email being sent
	// 	Optional
	NoEmail Boolean `json:"noemail"`

	// A base64 encoded array of the custom fields to update
	// 	Optional
	CustomFields map[string]any `json:"-"`

	// Optional array of file attachments.
	// 	Optional
	Attachments []Attachment `json:"-"`

	// The date and time the ticket note will show as created. Format: ISO8601 or YYYY-MM-DD HH:mm:ss.
	// Ticket creation date cannot be in the future
	// 	Optional
	CreatedDate FormattedTime `json:"created"`
}

type AddTicketReplyResponse struct {
	ApiResponse
}

type BlockTicketSenderRequest struct {
	// The ticket the sender opened
	// 	Required
	TicketId int `json:"ticketid"`

	// Should the ticket also be deleted
	// 	Optional
	Delete Boolean `json:"delete"`
}

type BlockTicketSenderResponse struct {
	ApiResponse
}

type DeleteAnnouncementRequest struct {
	// The id of the announcement to be deleted
	// 	Required
	AnnouncementId int `json:"announcementid"`
}

type DeleteAnnouncementResponse struct {
	ApiResponse
}

type DeleteTicketRequest struct {
	// The id of the Ticket to be deleted
	// 	Required
	TicketId int `json:"ticketid"`
}

type DeleteTicketResponse struct {
	ApiResponse
}

type DeleteTicketNoteRequest struct {
	// The id of the Ticket note to be deleted
	// 	Required
	NoteId int `json:"noteid"`
}

type DeleteTicketNoteResponse struct {
	ApiResponse
}

type DeleteTicketReplyRequest struct {
	// The id of the Ticket the reply belongs to
	// 	Required
	TicketId int `json:"ticketid"`

	// The id of the Ticket reply to be deleted
	// 	Required
	ReplyId int `json:"replyid"`
}

type DeleteTicketReplyResponse struct {
	ApiResponse
}

type GetAnnouncementsRequest struct {
	// The offset for the returned announcement data (default: 0)
	// 	Optional
	LimitStart int `json:"limitstart"`

	// The number of records to return (default: 25)
	// 	Optional
	LimitNum int `json:"limitnum"`
}

type Announcement struct {
	Id int `json:"id"`

	ParentId int `json:"parentid"`

	Title string `json:"title"`

	AnnouncementText string `json:"announcement"`

	Date FormattedTime `json:"date"`

	Published Boolean `json:"published"`

	Language string `json:"language"`

	CreatedAt FormattedTime `json:"created_at"`

	UpdatedAt FormattedTime `json:"updated_at"`
}

type GetAnnouncementsResponse struct {
	// The total number of results available
	Total int `json:"totalresults"`

	// The starting number for the returned results
	StartNum int `json:"startnumber"`

	// The number of results returned
	NumReturned int `json:"numreturned"`

	// The announcement entries returned
	Announcements []Announcement `json:"announcements"`
}

type MergeTicketRequest struct {
	// The unique ticket id that mergeticketids will be merged into
	// 	Required
	TicketId int `json:"ticketid"`

	// A list of ticket ids to merge into ticketid
	// 	Required
	MergeTicketIds []int `json:"mergeticketids"`

	// An optional subject to be set on the ticketid
	// 	Optional
	NewSubject string `json:"newsubject"`
}

type MergeTicketResponse struct {
	ApiResponse

	TicketId int `json:"ticketid"`
}

type TicketPriority string

const (
	LOW    TicketPriority = "Low"
	MEDIUM TicketPriority = "Medium"
	HIGH   TicketPriority = "High"
)

type OpenTicketRequest struct {
	// The department to open the ticket in
	// 	Required
	DeptId int `json:"deptid"`

	// The subject of the ticket
	// 	Required
	Subject string `json:"subject"`

	// The message of the ticket
	// 	Required
	MessageText string `json:"message"`

	// If applicable, the Client ID to create the ticket for.
	// 	Optional
	ClientId int `json:"clientid"`

	// If applicable, the user ID to create the ticket for (if ClientId is also passed).
	// 	Optional
	UserId int `json:"userid"`

	// If applicable, the Contact ID to create the ticket for (if ClientId and no UserId is also passed).
	// 	Optional
	ContactId int `json:"contactid"`

	// The name of the person opening the ticket (if not a client)
	// 	Optional
	Name string `json:"name"`

	// The email address of the person opening the ticket (if not a client)
	// 	Optional
	Email string `json:"email"`

	// The priority of the ticket ('Low', 'Medium', 'High')
	// 	Optional
	Priority TicketPriority `json:"priority"`

	// The date and time the ticket message will show as sent. Format: ISO8601 or YYYY-MM-DD HH:mm:ss
	// 	Optional
	CreatedDate FormattedTime `json:"created"`

	// The service to associate the ticket with (only one of ServiceId or DomainId)
	// 	Optional
	ServiceId int `json:"serviceid"`

	// The domain to associate the ticket with (only one of ServiceId or DomainId)
	// 	Optional
	DomainId int `json:"domainid"`

	// Is an Admin opening the ticket
	// 	Optional
	Admin Boolean `json:"admin"`

	// Pass 'true' for this value to prevent the ticket email from being sent.
	// 	Optional
	NoEmail Boolean `json:"noemail"`

	// Should markdown be used on the ticket output
	// 	Optional
	Markdown Boolean `json:"markdown"`

	// Base64 encoded serialized array of custom field values
	// 	Optional
	CustomFields map[string]any `json:"-"`

	// Optional array of file attachments.
	// 	Optional
	Attachments []Attachment `json:"-"`
}

type OpenTicketResponse struct {
	ApiResponse

	// The unique id of the newly created ticket
	TicketId int `json:"id"`

	// The unique ticket id displayed to the client, and to load the ticket in the client area
	ClientTicketId string `json:"tid"`

	// The code to access the ticket in the client area
	Code string `json:"c"`
}

type UpdateTicketRequest struct {
	// The ticket Id to update
	// 	Required
	TicketId int `json:"ticketid"`

	// The department id of the ticket
	// 	Optional
	DeptId int `json:"deptid"`

	// The status of the ticket
	// 	Optional
	Status string `json:"status"`

	// The subject of the ticket
	// 	Optional
	Subject string `json:"subject"`

	// If applicable, the Client ID to update the ticket for.
	// 	Optional
	UserId int `json:"userid"`

	// The name of the person opening the ticket (if not a client)
	// 	Optional
	Name string `json:"name"`

	// The email address of the person opening the ticket (if not a client)
	// 	Optional
	Email string `json:"email"`

	// The cc email addresses for the ticket
	// 	Optional
	CCEmail string `json:"cc"`

	// The priority of the ticket ('Low', 'Medium', 'High')
	// 	Optional
	Priority TicketPriority `json:"priority"`

	// The date and time the initial message will show as created. Format: ISO8601 or YYYY-MM-DD HH:mm:ss
	// 	Optional
	CreatedDate FormattedTime `json:"created"`

	// The id of the admin to flag the ticket to
	// 	Optional
	Flag int `json:"flag"`

	// Remove the flag from the ticket
	// 	Optional
	RemoveFlag Boolean `json:"removeFlag"`

	// Update the ticket message
	// 	Optional
	MessageText string `json:"message"`

	// Should markdown be used on the ticket output.
	// 	Optional
	Markdown Boolean `json:"markdown"`

	// Base64 encoded serialized array of custom field values
	// 	Optional
	CustomFields map[string]any `json:"-"`
}

type UpdateTicketResponse struct {
	ApiResponse

	// The ticket id that has been updated
	TicketId int `json:"ticketid"`
}

type UpdateTicketReplyRequest struct {
	// The reply id to update.
	// 	Required
	ReplyId int `json:"replyid"`

	// The message to be updated
	// 	Required
	MessageText string `json:"message"`

	// Should markdown be used on the ticket message. Existing value is used if not supplied.
	// 	Optional
	Markdown Boolean `json:"markdown"`

	// The date and time the ticket reply will show as sent. Format: ISO8601 or YYYY-MM-DD HH:mm:ss
	// 	Optional
	CreatedDate FormattedTime `json:"created"`
}

type UpdateTicketReplyResponse struct {
	ApiResponse

	// The reply id that has been updated
	ReplyId int `json:"replyid"`
}
