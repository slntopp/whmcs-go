package whmcs

import "encoding/json"

type GetSupportDepartmentsRequest struct {
	// Pass as true to not adhere to the departments the API user is a member of.
	// 	Optional
	IgnoreDeptAssignment *bool `json:"ignore_dept_assignments,omitempty"`
}

type Department struct {
	Id            *int    `json:"id"`
	Name          *string `json:"name"`
	AwaitingReply *int    `json:"awaitingreply"`
	OpenTickets   *int    `json:"opentickets"`
}

type GetSupportDepartmentsResponse struct {
	ApiResponse

	// The total number of results available
	Total *int `json:"totalresults"`

	// An array of department information
	Departments []Department `json:"departments"`
}

type GetSupportStatusesRequest struct {
	// Obtain counts for a specific department id
	//	Optional
	DeptId *int `json:"deptid,omitempty"`
}

type Status struct {
	Title *string `json:"title"`
	Count *int    `json:"count"`
	Color *string `json:"color,omitempty"`
}

type GetSupportStatusesResponse struct {
	ApiResponse

	// The total number of results available
	Total *int `json:"totalresults"`

	// An array of status information
	Statuses []Status `json:"statuses"`
}

type GetTicketRequest struct {
	// A specific Client Ticket Number to find tickets for.
	// 	Optional
	TicketNum *string `json:"ticketnum,omitempty"`

	// A specific ticket ID to find tickets for (either TicketNum or TicketId is required).
	// 	Optional
	TicketId *int `json:"ticketid,omitempty"`

	// ASC or DESC. The order to use to organise the ticket replies.
	// 	Optional
	RepliesSort *string `json:"repliessort,omitempty"`
}

type IndexedAttachment struct {
	Filename *string `json:"filename"`
	Index    *int    `json:"index"`
}

type Note struct {
	NoteId             *int                `json:"noteid"`
	Date               *formattedTime      `json:"date"`
	Message            *string             `json:"message"`
	Attachment         *string             `json:"attachment"`
	Attachments        []IndexedAttachment `json:"attachments"`
	AttachmentsRemoved any                 `json:"attachments_removed,omitempty"` // number(autocast to float64) or bool
	Admin              *string             `json:"admin,omitempty"`
}

type Reply struct {
	ReplyId            *int                `json:"replyid"`
	UserId             *int                `json:"userid"`
	ContactId          *int                `json:"contactid"`
	Name               *string             `json:"name"`
	Email              *string             `json:"email"`
	RequestorName      *string             `json:"requestor_name"`
	RequestorEmail     *string             `json:"requestor_email"`
	RequestorType      *string             `json:"requestor_type"`
	Date               *formattedTime      `json:"date"`
	Message            *string             `json:"message"`
	Attachment         *string             `json:"attachment"`
	Attachments        []IndexedAttachment `json:"attachments"`
	AttachmentsRemoved any                 `json:"attachments_removed,omitempty"` // number(autocast to float64) or bool
	Admin              *string             `json:"admin"`
	Rating             *int                `json:"rating,omitempty"`
}

type Ticket struct {
	// The unique ID of the ticket. If null check TicketId
	Id *int `json:"id"`

	// The unique ticket number string displayed to end users.
	ClientTicketId *string `json:"tid"`

	// The client unique access of the ticket.
	ClientAccess *string `json:"c"`

	// The ID of the department the ticket belongs to.
	DeptId *int `json:"deptid"`

	// The name of the department the ticket belongs to.
	DeptName *string `json:"deptname"`

	// The user ID the ticket belongs to.
	UserId *int `json:"userid"`

	// The contact ID the ticket was opened by.
	ContactId *int `json:"contactid"`

	// The ticket submitter’s name.
	Name *string `json:"name"`

	// The ticket submitter’s email.
	Email *string `json:"email"`

	// The ticket submitter’s name.
	RequestorName *string `json:"requestor_name"`

	// The ticket submitter’s type.
	RequestorType *string `json:"requestor_type"`

	// The ticket submitter’s email.
	RequestorEmail *string `json:"requestor_email"`

	// The CC email addresses for the ticket.
	CCEmail *string `json:"cc"`

	// The date the ticket was opened on. Format: Y-m-d H:i:s
	OpenedDate *formattedTime `json:"date"`

	// The subject of the ticket.
	Subject *string `json:"subject"`

	// The status of the ticket.
	Status *string `json:"status"`

	// The priority of the ticket.
	Priority *TicketPriority `json:"priority"`

	// The name of the admin user who opened the ticket.
	Admin *string `json:"admin"`

	// The date the ticket was last replied to. Format: Y-m-d H:i:s
	LastReplyDate *formattedTime `json:"lastreply"`

	// The ID of the admin user a ticket is flagged to.
	Flag *int `json:"flag"`

	// The ID of the service associated with the ticket (Sx for services and Dx for domains).
	Service *string `json:"service"`

	// An array of replies on the ticket.
	Replies []Reply `json:"replies,omitempty"`

	// An array of notes on the ticket.
	Notes []Note `json:"notes,omitempty"`

	Attachment         *string             `json:"attachment"`
	Attachments        []IndexedAttachment `json:"attachments"`
	AttachmentsRemoved *bool               `json:"attachments_removed"`
}

type GetTicketResponse struct {
	ApiResponse

	// The unique ID of the ticket.
	TicketId *int `json:"ticketid"`

	Ticket
}

type RelatedEntityType string

const (
	TICKET RelatedEntityType = "ticket"
	REPLY  RelatedEntityType = "reply"
	NOTE   RelatedEntityType = "note"
)

type GetTicketAttachmentRequest struct {
	// One of ticket, reply, note
	// 	Required
	Type RelatedEntityType `json:"type,omitempty"`

	// The unique id for the type
	// 	Required
	RelatedId *int `json:"relatedid,omitempty"`

	// The numerical index of the attachment to get
	// 	Required
	Index *int `json:"index,omitempty"`
}

type GetTicketAttachmentResponse struct {
	ApiResponse

	Filename *string `json:"filename"`
	Data     []byte  `json:"data"`
}

type GetTicketCountsRequest struct {
	// Pass as true to not adhere to the departments the API user is a member of.
	// 	Optional
	IgnoreDepartmentAssignments *bool `json:"ignoreDepartmentAssignments,omitempty"`

	// Pass as true to not adhere to the departments the API user is a member of.
	// 	Optional
	IncludeCountsByStatus *bool `json:"includeCountsByStatus,omitempty"`
}

type Open struct {
	Status
}

type Answered struct {
	Status
}

type CustomerReply struct {
	Status
}

type Closed struct {
	Status
}

type OnHold struct {
	Status
}

type InProgress struct {
	Status
}

type Statuses struct {
	Open          Open          `json:"open"`
	Answered      Answered      `json:"answered"`
	CustomerReply CustomerReply `json:"customerreply"`
	Closed        Closed        `json:"closed"`
	OnHold        OnHold        `json:"onhold"`
	InProgress    InProgress    `json:"inprogress"`
}

func (s *Statuses) Array() []Status {
	_json, _ := json.Marshal(s)
	sMap := make(map[string]any)
	json.Unmarshal(_json, &sMap)

	arr := make([]any, 0)
	for _, v := range sMap {
		arr = append(arr, v)
	}

	_json, _ = json.Marshal(arr)
	res := make([]Status, 0)
	json.Unmarshal(_json, &res)

	return res
}

type GetTicketCountsResponse struct {
	ApiResponse

	FilteredDepartments []int     `json:"filteredDepartments"`
	AllActive           *int      `json:"allActive"`
	AwaitingReply       *int      `json:"awaitingReply"`
	FlaggedTickets      *int      `json:"flaggedTickets"`
	Status              *Statuses `json:"status"`
}

type GetTicketNotesRequest struct {
	// Obtain the ticket for the specific ticket id
	// 	Required
	TicketId *int `json:"ticketid,omitempty"`
}

type GetTicketNotesResponse struct {
	ApiResponse

	// The total number of results being returned
	Total *int `json:"totalresults"`

	// An array of notes information
	Notes []Note `json:"notes"`
}

type GetTicketPredefinedCatsRequest struct{}

type Category struct {
	Id         *int    `json:"id"`
	ParentId   *int    `json:"parentid"`
	Name       *string `json:"name"`
	ReplyCount *int    `json:"replycount"`
}

type GetTicketPredefinedCatsResponse struct {
	ApiResponse

	// The total number of results being returned
	Total *int `json:"totalresults"`

	// An array of the reply categories
	Categories []Category `json:"categories"`
}

type GetTicketPredefinedRepliesRequest struct {
	// Obtain predefined replies for a specific category id
	// 	Optional
	CategoryId *int `json:"catid,omitempty"`
}

type PredefReply struct {
	Name  *string `json:"name"`
	Reply *string `json:"reply"`
}

type GetTicketPredefinedRepliesResponse struct {
	ApiResponse

	// The total number of results returned
	Total *int `json:"totalresults"`

	// An array of predefined replies
	Replies []PredefReply `json:"predefinedreplies"`
}

type TicketStatus string

const (
	AWAITING_REPLY     TicketStatus = "Awaiting Reply"
	ALL_ACTIVE_TICKETS TicketStatus = "All Active Tickets"
	MY_FLAGGED_TICKETS TicketStatus = "My Flagged Tickets"
)

type GetTicketsRequest struct {
	// The offset for the returned quote data (default: 0)
	// 	Optional
	LimitStart *int `json:"limitstart,omitempty"`

	// The number of records to return (default: 25)
	// 	Optional
	LimitNum *int `json:"limitnum,omitempty"`

	// Obtain tickets in a specific department
	// 	Optional
	DeptId *int `json:"deptid,omitempty"`

	// Find tickets for a specific client id
	// 	Optional
	ClientId *int `json:"clientid,omitempty"`

	// Find tickets for a specific non-client email address
	// 	Optional
	Email *string `json:"email,omitempty"`

	// Find tickets matching a specific status. Any configured status plus: Awaiting Reply, All Active Tickets, My Flagged Tickets
	// 	Optional
	Status TicketStatus `json:"status,omitempty"`

	// Find tickets containing a specific subject - uses approximate string matching.
	// 	Optional
	Subject *string `json:"subject,omitempty"`

	// Pass as true to not adhere to the departments the API user is a member of.
	// 	Optional
	IgnoreDepartmentAssignments *bool `json:"ignore_dept_assignments,omitempty"`
}

type GetTicketsResponse struct {
	ApiResponse

	// The total number of results available
	Total *int `json:"totalresults"`

	// The starting number for the returned results
	StartNumber *int `json:"startnumber"`

	// The number of results returned
	NumrRturned *int `json:"numreturned"`

	// An array of tickets matching the criteria passed
	Tickets []Ticket `json:"tickets"`

	// The ticket submitter’s name.
	RequestorName *string `json:"requestor_name"`

	// The ticket submitter’s type.
	RequestorType *string `json:"requestor_type"`

	// The ticket submitter’s email.
	RequestorEmail *string `json:"requestor_email"`

	// The ticket submitter’s owner.
	OwnerName *string `json:"owner_name"`
}
