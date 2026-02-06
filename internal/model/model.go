package model

type ShortenRequest struct {
	URL string `json:"url"`
}

type ShortenResponse struct {
	Result string `json:"result"`
}

var AcceptedContentTypes = []string{
	"application/json",
	"text/html",
}
