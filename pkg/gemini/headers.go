package gemini

import "strings"

type Request struct {
	Host   string
	Port   string
	Path   string
	Schema string
}

func NewRequest(url string) Request {
	r := Request{}
	var parts []string
	if strings.Contains(url, "://") {
		parts = strings.Split(url, "://")
		r.Schema = parts[0]
	} else {
		r.Schema = "gemini"
	}
	return r
}
