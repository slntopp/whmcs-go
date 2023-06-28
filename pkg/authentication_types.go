package whmcs

type ValidateLoginRequest struct {
	// User Email Address
	// 	Required
	Email string `json:"email"`

	// Password to validate
	// 	Required
	Password string `json:"password2"`
}

type ValidateLoginResponse struct {
	ApiResponse

	// User ID
	UserId int `json:"userid"`

	// Login session token - returned if Two-Factor Authentication is not required for the account
	PasswordHash string `json:"passwordhash"`

	// True if Two-Factor Authentication is enabled for the given account
	TwoFactorEnabled bool `json:"twoFactorEnabled"`
}
