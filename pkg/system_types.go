package whmcs

type GetEmailTemplatesRequest struct {
	// The type of email template to retrieve
	// 	Optional
	Type string `json:"type"`

	// The language of the email template to retrieve, if none provided will return default language templates.
	// 	Optional
	Language string `json:"language"`
}

type EmailTemplate struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Subject string `json:"subject"`
	Custom  bool   `json:"custom"`
}

type GetEmailTemplatesResponse struct {
	ApiResponse

	// The total number of results returned
	Total int `json:"totalresults"`

	// The email templates entries returned
	EmailTemplates []EmailTemplate `json:"emailtemplates"`
}

type EmailCustomType string

const (
	GENERAL   EmailCustomType = "general"
	PRODUCT   EmailCustomType = "product"
	DOMAIN    EmailCustomType = "domain"
	INVOICE   EmailCustomType = "invoice"
	SUPPORT   EmailCustomType = "support"
	AFFILIATE EmailCustomType = "affiliate"
)

type SendEmailRequest struct {
	// The name of the client email template to send
	// 	Optional
	MessageName string `json:"messagename"`

	// The related id for the type of email template. Eg this should be the client id for a general type email
	// What you must provide for the Related ID depends upon the type of email being sent. The available options are:
	// 		- General Email Type = Client ID (tblclients.id)
	// 		- Product Email Type = Service ID (tblhosting.id)
	// 		- Domain Email Type = Domain ID (tbldomains.id)
	// 		- Invoice Email Type = Invoice ID (tblinvoices.id)
	// 		- Support Email Type = Ticket ID (tbltickets.id)
	// 		- Affiliate Email Type = Affiliate ID (tblaffiliates.id)
	//
	// 	Optional
	Id int `json:"id"`

	// The type of custom email template to send ('general', 'product', 'domain', 'invoice', 'support', 'affiliate')
	// 	Optional
	CustomType EmailCustomType `json:"customtype"`

	// The HTML message body to send for a custom email
	// 	Optional
	CustomMessage string `json:"custommessage"`

	// The subject to send for a custom email
	// 	Optional
	CustomSubject string `json:"customsubject"`

	// The custom variables to provide to the email template. Can be used for existing and custom emails.
	// 	Optional
	CustomVars map[string]any `json:"customvars"`
}

type SendEmailResponse struct {
	ApiResponse
}

type AdminEmailType string

const (
	ADMIN_SYSTEM  AdminEmailType = "system"
	ADMIN_ACCOUNT AdminEmailType = "account"
	ADMIN_SUPPORT AdminEmailType = "support"
)

type SendAdminEmailRequest struct {
	// The name of the admin email template to send
	// 	Optional
	MessageName string `json:"messagename"`

	// The HTML message body to send for a custom email
	// 	Optional
	CustomMessage string `json:"custommessage"`

	// The subject to send for a custom email
	// 	Optional
	CustomSubject string `json:"customsubject"`

	// Which type of admin notification will be send ('system', 'account', 'support')
	// 	Optional
	Type AdminEmailType `json:"type"`

	// The Id of the department the notification is for if 'support' Type
	// 	Optional
	DeptId int `json:"deptid"`

	// The merge fields to be used in the email template
	// 	Optional
	MergeFields map[string]any `json:"mergefields"`
}

type SendAdminEmailResponse struct {
	ApiResponse
}

type NotificationStatusStyle string

const (
	SUCCESS NotificationStatusStyle = "success"
	DANGER  NotificationStatusStyle = "danger"
	INFO    NotificationStatusStyle = "info"
)

type Attribute struct {
	// 	Required
	Label string `json:"label"`

	// 	Required
	Value string `json:"value"`

	// 	Optional
	AdditionalAttributes map[string]any `json:"-"`
}

type TriggerNotificationEventRequest struct {
	// A unique identifier string, used as a condition when making a notification rule.
	// 	Optional
	NotificationIdentifier string `json:"notification_identifier"`

	// The title for the notification
	// 	Optional
	Title string `json:"title"`

	// The message body for the notification
	// 	Optional
	Message string `json:"message"`

	// The follow up URL for the notification
	// 	Optional
	URL string `json:"url"`

	// A status description for the notification
	// 	Optional
	Status string `json:"status"`

	// A formatting style for the status of the notification, currently supports 'success', 'danger', and 'info'
	// 	Optional
	StatusStyle NotificationStatusStyle `json:"statusStyle"`

	// An array of Attributes to include in the notification. Requires at least label and value parameters. Other parameters are optional. See WHMCS\Notification\NotificationAttribute.
	// 	Optional
	Attributes []Attribute `json:"-"`
}

type TriggerNotificationEventResponse struct {
	ApiResponse
}
