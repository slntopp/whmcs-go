package whmcs

type ValidateLoginRequest struct {
	// User Email Address
	// 	Required
	Email *string `json:"email,omitempty"`

	// Password to validate
	// 	Required
	Password *string `json:"password2,omitempty"`
}

type ValidateLoginResponse struct {
	ApiResponse

	// User ID
	UserId *int `json:"userid,omitempty"`

	// Login session token - returned if Two-Factor Authentication is not required for the account
	PasswordHash *string `json:"passwordhash,omitempty"`

	// True if Two-Factor Authentication is enabled for the given account
	TwoFactorEnabled *bool `json:"twoFactorEnabled,omitempty"`
}
