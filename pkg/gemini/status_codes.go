package gemini

const (
	INPUT                       int = 10
	SENSITIVE_INPUT             int = 11
	SUCCESS                     int = 20
	TEMPORARY_REDIRECT          int = 30
	PERMANENT_REDIRECT          int = 31
	TEMPORARY_FAILURE           int = 40
	SERVER_UNAVAILABLE          int = 41
	CGI_ERROR                   int = 42
	PROXY_ERROR                 int = 43
	SLOW_DOWN                   int = 44
	PERMANENT_FALIURE           int = 50
	NOT_FOUND                   int = 51
	GONE                        int = 52
	PROXY_REQUEST_REFUSED       int = 53
	BAD_REQUEST                 int = 59
	CLIENT_CERTIFICATE_REQUIRED int = 60
	CERTIFICATE_NOT_AUTHORIZED  int = 61
	CERTIFICATE_NOT_VALID       int = 62
)
