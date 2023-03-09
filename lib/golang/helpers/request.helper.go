package helpers

import (
	"development/go/recipes/lib/golang"
	"encoding/json"
	"net/http"
	"strings"
)

func AddUserHeaders(jwtToken string, req *http.Request) {
	req.Header.Add("User-Authorization", jwtToken)
	addHeaders(req)
}

func AddAppHeaders(req *http.Request) {
	addHeaders(req)
}

func GetRoute(route string, args ...string) string {
	requestRoute := ""
	tokens := strings.Split(route, "%s")
	for index, token := range tokens {
		requestRoute += token + args[index]
	}
	return golang.HostDev + requestRoute
}

func GetRequestBody(obj interface{}) (string, error) {
	var res string
	marshaled, err := json.Marshal(obj)
	if err != nil {
		return res, err
	}

	res = string(marshaled)
	return res, nil
}

// Private Methods

func addHeaders(req *http.Request) {
	req.Header.Add("x-api-key", golang.XApiKey)
	req.Header.Add("Content-Type", "application/json")
}
