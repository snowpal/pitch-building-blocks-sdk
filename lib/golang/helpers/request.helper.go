package helpers

import (
	"development/go/recipes/lib/golang"
	"encoding/json"
	"errors"
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

func GetRoute(route string, args ...string) (string, error) {
	var res string
	if len(strings.Split(route, "%s"))-1 != len(args) {
		return res, errors.New("substitution does not match args count")
	}
	for _, arg := range args {
		route = strings.Replace(route, "%s", arg, 1)
	}
	res = golang.HostDev + route
	return res, nil
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
