package templates

import (
	"development/go/recipes/lib"
	helpers2 "development/go/recipes/lib/helpers"
	"development/go/recipes/lib/structs/response"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetPodTemplates(jwtToken string) ([]response.PodTemplate, error) {
	resPodTemplates := response.PodTemplates{}
	route, err := helpers2.GetRoute(lib.RouteTemplatesGetPodTemplates)
	if err != nil {
		fmt.Println(err)
		return resPodTemplates.Templates, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		fmt.Println(err)
		return resPodTemplates.Templates, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers2.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resPodTemplates.Templates, err
	}

	defer helpers2.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resPodTemplates.Templates, err
	}

	err = json.Unmarshal(body, &resPodTemplates)
	if err != nil {
		fmt.Println(err)
		return resPodTemplates.Templates, err
	}
	return resPodTemplates.Templates, nil
}
