package pod_types_1

import (
	"development/go/recipes/lib/golang"
	"development/go/recipes/lib/golang/helpers"
	"development/go/recipes/lib/golang/structs/request"
	"development/go/recipes/lib/golang/structs/response"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func UpdatePodType(jwtToken string, reqBody request.PodTypeReqBody, podTypeId string) (response.PodType, error) {
	resPodType := response.PodType{}
	requestBody, err := helpers.GetRequestBody(reqBody)
	if err != nil {
		fmt.Println(err)
		return resPodType, err
	}
	payload := strings.NewReader(requestBody)
	route, err := helpers.GetRoute(golang.RoutePodTypesUpdatePodType, podTypeId)
	if err != nil {
		fmt.Println(err)
		return resPodType, err
	}

	req, err := http.NewRequest(http.MethodPatch, route, payload)
	if err != nil {
		fmt.Println(err)
		return resPodType, err
	}

	helpers.AddUserHeaders(jwtToken, req)

	res, err := helpers.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resPodType, err
	}

	defer helpers.CloseBody(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resPodType, err
	}

	err = json.Unmarshal(body, &resPodType)
	if err != nil {
		fmt.Println(err)
		return resPodType, err
	}
	return resPodType, nil
}
