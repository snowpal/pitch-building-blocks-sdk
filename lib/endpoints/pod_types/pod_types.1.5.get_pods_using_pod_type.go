package podTypes

import (
	"development/go/recipes/lib"
	helpers2 "development/go/recipes/lib/helpers"
	"development/go/recipes/lib/structs/response"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetPodsUsingPodType(jwtToken string, podTypeId string) ([]response.Pod, error) {
	resPods := response.Pods{}
	route, err := helpers2.GetRoute(lib.RoutePodTypesGetPodsUsingPodType, podTypeId)
	if err != nil {
		fmt.Println(err)
		return resPods.Pods, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		fmt.Println(err)
		return resPods.Pods, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers2.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resPods.Pods, err
	}

	defer helpers2.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resPods.Pods, err
	}

	err = json.Unmarshal(body, &resPods)
	if err != nil {
		fmt.Println(err)
		return resPods.Pods, err
	}
	return resPods.Pods, nil
}
