package helpers

import (
	"development/go/recipes/lib/golang"
	"net/http"
)

func AddUserHeaders(jwtToken string, req *http.Request) {
	req.Header.Add("User-Authorization", jwtToken)
	addHeaders(req)
}

func AddAppHeaders(req *http.Request) {
	addHeaders(req)
}

func addHeaders(req *http.Request) {
	req.Header.Add("x-api-key", golang.XApiKey)
	req.Header.Add("Content-Type", "application/json")
}

func GetRoute(route string) string {
	return golang.HostDev + route
}
