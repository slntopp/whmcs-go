package whmcs

type AuthenticationService struct {
	wc *Client
}

func (c *Client) initAuthenticationService() {
	c.Authentication = AuthenticationService{wc: c}
}

func (s *AuthenticationService) ValidateLogin(req *ValidateLoginRequest) (*ValidateLoginResponse, error) {
	res := &ValidateLoginResponse{}
	if err := s.wc.Call("ValidateLogin", req, &res); err != nil {
		return nil, err
	}

	return res, nil
}
