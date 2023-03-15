package dashboard_2

import (
	"development/go/recipes/lib/golang"
	"development/go/recipes/lib/golang/helpers"
	"development/go/recipes/lib/golang/structs/response"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetPodsBasedOnPodTypes(jwtToken string) ([]response.PodTypesKey, error) {
	resPodTypesKeys := response.PodTypesKeys{}
	client := &http.Client{}
	route, err := helpers.GetRoute(golang.RouteDashboardGetPodsBasedOnPodTypes)
	if err != nil {
		fmt.Println(err)
		return *resPodTypesKeys.Keys, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		fmt.Println(err)
		return *resPodTypesKeys.Keys, err
	}

	helpers.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = client.Do(req)
	if err != nil {
		fmt.Println(err)
		return *resPodTypesKeys.Keys, err
	}

	defer helpers.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return *resPodTypesKeys.Keys, err
	}

	err = json.Unmarshal(body, &resPodTypesKeys)
	if err != nil {
		fmt.Println(err)
		return *resPodTypesKeys.Keys, err
	}
	return *resPodTypesKeys.Keys, nil
}
