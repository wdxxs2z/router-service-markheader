package headers

import (
	"fmt"
	"net/http"
	"net/url"
)

const (
	RouteServiceSignature    = "X-CF-Proxy-Signature"
	RouteServiceForwardedUrl = "X-CF-Forwarded-Url"
	RouteServiceMetadata     = "X-CF-Proxy-Metadata"
        RouteServiceMark         = "X-CF-Mark"
)

type RouteServiceHeaders struct {
	Signature string
	Metadata  string
	UrlString string
        Mark      string
	ParsedUrl *url.URL
}

func NewRouteServiceHeaders() *RouteServiceHeaders {

	return &RouteServiceHeaders{
		Signature: "",
		Metadata:  "",
		UrlString: "",
                Mark:      "",
	}
}

func (r *RouteServiceHeaders) ParseHeadersAndClean(headers *http.Header) error {
	var err error
	r.Signature = headers.Get(RouteServiceSignature)
	r.UrlString = headers.Get(RouteServiceForwardedUrl)
	r.Metadata = headers.Get(RouteServiceMetadata)
        r.Mark = headers.Get(RouteServiceMark)

	r.ParsedUrl, err = url.Parse(r.UrlString)
	headers.Del(RouteServiceForwardedUrl)
	return err

}

func (r *RouteServiceHeaders) IsValidRequest() bool {
	if r.Signature == "" || r.Metadata == "" || r.UrlString == "" {
		return false

	}
	return true

}

func (r *RouteServiceHeaders) IsMarkRequest(mark string) bool {
	if r.Mark != mark {
	   	return false
	}
	return true
}

func (r *RouteServiceHeaders) String() string {
	return fmt.Sprintf("X-CF-Proxy-Signature: %v \n X-CF-Forwarded-Url: %v \n X-CF-Proxy-Metadata: %v", r.Signature, r.UrlString, r.Metadata)

}
