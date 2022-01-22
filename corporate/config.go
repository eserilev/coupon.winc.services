package corporate

func CwApiBaseUrl() string {
	return "http://localhost:3000"
}

func CorporateOrderRelativePath(userGuid string) string {
	return "/__s/v2/tmp/todo/orders/corporate?userGuid=" + userGuid
}
