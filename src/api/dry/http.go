package dry

import (
	"net/http"
)

func GetApiBase(req *http.Request) string {
	proto := req.Header.Get("X-forwarded-proto")
	host := req.Header.Get("x-forwarded-host")
	if proto == "" {
		proto = "http"
	}
	if host == "" {
		host = req.Host
	}
	Assert(host != "", req)
	Assert(proto != "", req)
	return proto + "://" + host
}
