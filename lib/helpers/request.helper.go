package helpers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/snowpal/pitch-building-blocks-sdk/lib"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/config"
	"golang.org/x/exp/slices"
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
	res = lib.GatewayHost + route
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

func MakeRequest(req *http.Request) (*http.Response, error) {
	client := &http.Client{}
	res, err := client.Do(req)
	successCodes := []int{200, 201, 202, 204}
	if err != nil || !slices.Contains(successCodes, res.StatusCode) {
		return res, errors.New("API Request Failed")
	}
	return res, nil
}

func addHeaders(req *http.Request) {
	req.Header.Add("x-api-key", config.GetValue("XApiKey"))
	req.Header.Add("x-snowpal-product-code", config.GetValue("XProductCode"))
	req.Header.Add("Content-Type", "application/json")
}
