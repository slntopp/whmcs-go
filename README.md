# WHMCS-GO
WHMCS API bindings for GoLang

## Usage
```
client, err := whmcs.NewClient(urlToWhmcs, username, password, dangerMode)
if err != nil {
	return err
}

login, err := client.Authentication.ValidateLogin(&whmcs.ValidateLoginRequest{
	Email:    whmcs.String("example@mail.com"),
	Password: whmcs.String("password"),
})
if err != nil {
	return err
}
```

## Implemented methods
Methods are grouped in the same way as in the [official documentation](https://developers.whmcs.com/api/api-index/).
- Authentication
  - ValidateLogin
- Tickets (all)
- Support (all)
- System
  - GetEmailTemplates
  - SendAdminEmail
  - SendEmail
  - TriggerNotificationEvent
