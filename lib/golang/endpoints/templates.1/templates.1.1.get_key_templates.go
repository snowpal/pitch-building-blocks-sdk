package templates

import (
	"development/go/recipes/lib/golang"
	"development/go/recipes/lib/golang/helpers"
	"development/go/recipes/lib/golang/structs/response"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetKeyTemplates(jwtToken string) ([]response.KeyTemplatesWithType, error) {
	resKeyTemplates := response.KeyTemplates{}
	client := &http.Client{}
	route, err := helpers.GetRoute(golang.RouteTemplatesGetKeyTemplates)
	if err != nil {
		fmt.Println(err)
		return resKeyTemplates.Templates, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		fmt.Println(err)
		return resKeyTemplates.Templates, err
	}

	helpers.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = client.Do(req)
	if err != nil {
		fmt.Println(err)
		return resKeyTemplates.Templates, err
	}

	defer helpers.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resKeyTemplates.Templates, err
	}

	err = json.Unmarshal(body, &resKeyTemplates)
	if err != nil {
		fmt.Println(err)
		return resKeyTemplates.Templates, err
	}
	return resKeyTemplates.Templates, nil
}
