package whmcs

type AddAnnouncementRequest struct {
	// Date in the format YYYY-MM-DD HH:MM:SS
	// 	Required
	Date formattedTime `json:"date,omitempty"`

	// 	Required
	Title *string `json:"title,omitempty"`

	// Announcement text.
	// 	Required
	AnnouncementText *string `json:"announcement,omitempty"`

	// Pass as true to publish.
	// 	Optional
	Published *boolean `json:"published,omitempty"`
}

type AddAnnouncementResponse struct {
	ApiResponse

	// The id of the new announcement
	AnnouncementId *int `json:"announcementid,omitempty"`
}

type AddCancelRequestType string

const (
	IMMEDIATE             AddCancelRequestType = "Immediate"
	END_OF_BILLING_PEROID AddCancelRequestType = "End of Billing Period"
)

type AddCancelRequest_Request struct {
	// The Service ID to cancel
	// 	Required
	ServieId *int `json:"serviceid,omitempty"`

	// The type of cancellation. 'Immediate' or 'End of Billing Period'
	// 	Optional
	Type AddCancelRequestType `json:"type,omitempty"`

	// The customer reason for cancellation
	// 	Optional
	Reason *string `json:"reason,omitempty"`
}

type AddCancelRequest_Response struct {
	ApiResponse

	// The id of the service the request was for
	ServiceId *int `json:"serviceid,omitempty"`

	// The id of the user the service belongs to
	UserId *int `json:"userid,omitempty"`
}

type AddClientNoteRequest struct {
	// The Client ID to apply the note to
	// 	Required
	UserId *int `json:"userid,omitempty"`

	// The note to add
	// 	Required
	NoteText *string `json:"notes,omitempty"`

	// Should the note be made sticky. Makes the note 'sticky' and displays the note throughout the client’s
	// account and on any tickets they submit in the admin area
	// 	Optional
	Sticky *boolean `json:"sticky,omitempty"`
}

type AddClientNoteResponse struct {
	ApiResponse

	// 	The id of the newly created note
	NoteId *int `json:"noteid,omitempty"`
}

type Attachment struct {
	// File name
	// 	Required
	Name *string `json:"name,omitempty"`

	// File data
	// 	Required
	Data []byte `json:"data,omitempty"`
}

type AddTicketNoteRequest struct {
	// The content of the ticket note
	// 	Required
	MessageText *string `json:"message,omitempty"`

	// The Client Ticket Number ID to apply the note to
	// 	Optional
	TicketNum *string `json:"ticketnum,omitempty"`

	// The id of the ticket in the database. Either TicketNum or TicketId is required
	// 	Optional
	TicketId *int `json:"ticketid,omitempty"`

	// Should markdown be used on the ticket note output
	// 	Optional
	Markdown *boolean `json:"markdown,omitempty"`

	// Optional array of file attachments.
	// 	Optional
	Attachments []Attachment `json:"-"`

	// The date and time the ticket note will show as created. Format: ISO8601 or YYYY-MM-DD HH:mm:ss.
	// Ticket creation date cannot be in the future
	// 	Optional
	CreatedDate formattedTime `json:"created,omitempty"`
}

type AddTicketNoteResponse struct {
	ApiResponse
}

type AddTicketReplyRequest struct {
	// The id of the ticket.
	// 	Required
	TicketId *int `json:"ticketid,omitempty"`

	// The content of the ticket reply
	// 	Required
	MessageText *string `json:"message,omitempty"`

	// Should markdown be used on the ticket note output
	// 	Optional
	Markdown *boolean `json:"markdown,omitempty"`

	// Pass a clientid to associate the ticket reply with a specific client
	// 	Optional
	ClientId *int `json:"clientid,omitempty"`

	// Pass a contactid to associate the ticket reply with a specific contact belonging to ClientId
	// 	Optional
	ContactId *int `json:"contactid,omitempty"`

	// The admin username to associate the ticket reply with
	// 	Optional
	AdminUsername *string `json:"adminusername,omitempty"`

	// The name to associate with the ticket reply if not an admin or client response
	// 	Optional
	Name *string `json:"name,omitempty"`

	// The email to associate with the ticket reply if not an admin or client response
	// 	Optional
	Email *string `json:"email,omitempty"`

	// The status to set on the ticket after the reply is made if the default status on admin/client response is not required. See GetSupportStatuses API command
	// 	Optional
	Status *string `json:"status,omitempty"`

	// Set to true to stop the ticket reply email being sent
	// 	Optional
	NoEmail *boolean `json:"noemail,omitempty"`

	// A base64 encoded array of the custom fields to update
	// 	Optional
	CustomFields map[string]any `json:"-"`

	// Optional array of file attachments.
	// 	Optional
	Attachments []Attachment `json:"-"`

	// The date and time the ticket note will show as created. Format: ISO8601 or YYYY-MM-DD HH:mm:ss.
	// Ticket creation date cannot be in the future
	// 	Optional
	CreatedDate formattedTime `json:"created,omitempty"`
}

type AddTicketReplyResponse struct {
	ApiResponse
}

type BlockTicketSenderRequest struct {
	// The ticket the sender opened
	// 	Required
	TicketId *int `json:"ticketid,omitempty"`

	// Should the ticket also be deleted
	// 	Optional
	Delete *boolean `json:"delete,omitempty"`
}

type BlockTicketSenderResponse struct {
	ApiResponse
}

type DeleteAnnouncementRequest struct {
	// The id of the announcement to be deleted
	// 	Required
	AnnouncementId *int `json:"announcementid,omitempty"`
}

type DeleteAnnouncementResponse struct {
	ApiResponse
}

type DeleteTicketRequest struct {
	// The id of the Ticket to be deleted
	// 	Required
	TicketId *int `json:"ticketid,omitempty"`
}

type DeleteTicketResponse struct {
	ApiResponse
}

type DeleteTicketNoteRequest struct {
	// The id of the Ticket note to be deleted
	// 	Required
	NoteId *int `json:"noteid,omitempty"`
}

type DeleteTicketNoteResponse struct {
	ApiResponse
}

type DeleteTicketReplyRequest struct {
	// The id of the Ticket the reply belongs to
	// 	Required
	TicketId *int `json:"ticketid,omitempty"`

	// The id of the Ticket reply to be deleted
	// 	Required
	ReplyId *int `json:"replyid,omitempty"`
}

type DeleteTicketReplyResponse struct {
	ApiResponse
}

type GetAnnouncementsRequest struct {
	// The offset for the returned announcement data (default: 0)
	// 	Optional
	LimitStart *int `json:"limitstart,omitempty"`

	// The number of records to return (default: 25)
	// 	Optional
	LimitNum *int `json:"limitnum,omitempty"`
}

type Announcement struct {
	Id *int `json:"id"`

	ParentId *int `json:"parentid"`

	Title *string `json:"title"`

	AnnouncementText *string `json:"announcement"`

	Date *formattedTime `json:"date"`

	Published *boolean `json:"published"`

	Language *string `json:"language"`

	CreatedAt *formattedTime `json:"created_at"`

	UpdatedAt *formattedTime `json:"updated_at"`
}

type GetAnnouncementsResponse struct {
	ApiResponse

	// The total number of results available
	Total *int `json:"totalresults"`

	// The starting number for the returned results
	StartNum *int `json:"startnumber"`

	// The number of results returned
	NumReturned *int `json:"numreturned"`

	// The announcement entries returned
	Announcements []Announcement `json:"announcements"`
}

type MergeTicketRequest struct {
	// The unique ticket id that MergeTicketIds will be merged into
	// 	Required
	TicketId *int `json:"ticketid,omitempty"`

	// A list of ticket ids to merge into TicketId
	// 	Required
	MergeTicketIds []int `json:"mergeticketids,omitempty"`

	// An optional subject to be set on the TicketId
	// 	Optional
	NewSubject *string `json:"newsubject,omitempty"`
}

type MergeTicketResponse struct {
	ApiResponse

	TicketId *int `json:"ticketid"`
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
	DeptId *int `json:"deptid,omitempty"`

	// The subject of the ticket
	// 	Required
	Subject *string `json:"subject,omitempty"`

	// The message of the ticket
	// 	Required
	MessageText *string `json:"message,omitempty"`

	// If applicable, the Client ID to create the ticket for.
	// 	Optional
	ClientId *int `json:"clientid,omitempty"`

	// If applicable, the user ID to create the ticket for (if ClientId is also passed).
	// 	Optional
	UserId *int `json:"userid,omitempty"`

	// If applicable, the Contact ID to create the ticket for (if ClientId and no UserId is also passed).
	// 	Optional
	ContactId *int `json:"contactid,omitempty"`

	// The name of the person opening the ticket (if not a client)
	// 	Optional
	Name *string `json:"name,omitempty"`

	// The email address of the person opening the ticket (if not a client)
	// 	Optional
	Email *string `json:"email,omitempty"`

	// The priority of the ticket ('Low', 'Medium', 'High')
	// 	Optional
	Priority *TicketPriority `json:"priority,omitempty"`

	// The date and time the ticket message will show as sent. Format: ISO8601 or YYYY-MM-DD HH:mm:ss
	// 	Optional
	CreatedDate formattedTime `json:"created,omitempty"`

	// The service to associate the ticket with (only one of ServiceId or DomainId)
	// 	Optional
	ServiceId *int `json:"serviceid,omitempty"`

	// The domain to associate the ticket with (only one of ServiceId or DomainId)
	// 	Optional
	DomainId *int `json:"domainid,omitempty"`

	// Is an Admin opening the ticket
	// 	Optional
	Admin *boolean `json:"admin,omitempty"`

	// Pass 'true' for this value to prevent the ticket email from being sent.
	// 	Optional
	NoEmail *boolean `json:"noemail,omitempty"`

	// Should markdown be used on the ticket output
	// 	Optional
	Markdown *boolean `json:"markdown,omitempty"`

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
	TicketId *int `json:"id"`

	// The unique ticket id displayed to the client, and to load the ticket in the client area
	ClientTicketId *string `json:"tid"`

	// The code to access the ticket in the client area
	Code *string `json:"c"`
}

type UpdateTicketRequest struct {
	// The ticket Id to update
	// 	Required
	TicketId *int `json:"ticketid,omitempty"`

	// The department id of the ticket
	// 	Optional
	DeptId *int `json:"deptid,omitempty"`

	// The status of the ticket
	// 	Optional
	Status *string `json:"status,omitempty"`

	// The subject of the ticket
	// 	Optional
	Subject *string `json:"subject,omitempty"`

	// If applicable, the Client ID to update the ticket for.
	// 	Optional
	UserId *int `json:"userid,omitempty"`

	// The name of the person opening the ticket (if not a client)
	// 	Optional
	Name *string `json:"name,omitempty"`

	// The email address of the person opening the ticket (if not a client)
	// 	Optional
	Email *string `json:"email,omitempty"`

	// The cc email addresses for the ticket
	// 	Optional
	CCEmail *string `json:"cc,omitempty"`

	// The priority of the ticket ('Low', 'Medium', 'High')
	// 	Optional
	Priority TicketPriority `json:"priority,omitempty"`

	// The date and time the initial message will show as created. Format: ISO8601 or YYYY-MM-DD HH:mm:ss
	// 	Optional
	CreatedDate formattedTime `json:"created,omitempty"`

	// The id of the admin to flag the ticket to
	// 	Optional
	Flag *int `json:"flag,omitempty"`

	// Remove the flag from the ticket
	// 	Optional
	RemoveFlag *boolean `json:"removeFlag,omitempty"`

	// Update the ticket message
	// 	Optional
	MessageText *string `json:"message,omitempty"`

	// Should markdown be used on the ticket output.
	// 	Optional
	Markdown *boolean `json:"markdown,omitempty"`

	// Base64 encoded serialized array of custom field values
	// 	Optional
	CustomFields map[string]any `json:"-"`
}

type UpdateTicketResponse struct {
	ApiResponse

	// The ticket id that has been updated
	TicketId *int `json:"ticketid"`
}

type UpdateTicketReplyRequest struct {
	// The reply id to update.
	// 	Required
	ReplyId *int `json:"replyid,omitempty"`

	// The message to be updated
	// 	Required
	MessageText *string `json:"message,omitempty"`

	// Should markdown be used on the ticket message. Existing value is used if not supplied.
	// 	Optional
	Markdown *boolean `json:"markdown,omitempty"`

	// The date and time the ticket reply will show as sent. Format: ISO8601 or YYYY-MM-DD HH:mm:ss
	// 	Optional
	CreatedDate formattedTime `json:"created,omitempty"`
}

type UpdateTicketReplyResponse struct {
	ApiResponse

	// The reply id that has been updated
	ReplyId *int `json:"replyid"`
}
