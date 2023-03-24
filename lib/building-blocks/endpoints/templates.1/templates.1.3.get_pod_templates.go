package templates

import (
	"development/go/recipes/lib/building-blocks"
	"development/go/recipes/lib/building-blocks/helpers"
	"development/go/recipes/lib/building-blocks/structs/response"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetPodTemplates(jwtToken string) ([]response.PodTemplate, error) {
	resPodTemplates := response.PodTemplates{}
	route, err := helpers.GetRoute(building_blocks.RouteTemplatesGetPodTemplates)
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

	helpers.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resPodTemplates.Templates, err
	}

	defer helpers.CloseBody(res.Body)

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
