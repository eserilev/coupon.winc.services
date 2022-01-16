package config

func CwApiBaseUrl() string {
	return "http://cwapi-staging.cloudapp.net/"
}

func CorporateOrderRelativePath(userGuid string) string {
	return "winc/users/" + userGuid + "gift-checkout"
}

func CwApiUserName() string {
	return "api@clubw.com"
}

func CwApiPassword() string {
	return "randomPassword"
}
