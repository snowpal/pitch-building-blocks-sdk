package block_pods

import (
	"development/go/recipes/lib/golang"
	"development/go/recipes/lib/golang/helpers"
	"development/go/recipes/lib/golang/structs/common"
	"development/go/recipes/lib/golang/structs/response"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetPodsAvailableToBeLinked(jwtToken string, slimBlock common.SlimBlock) ([]response.Pod, error) {
	podsResponse := response.Pods{}

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, helpers.GetRoute(golang.RouteBlockPodsGetPodsAvailableToBeLinked, slimBlock.ID, slimBlock.Key.ID), nil)
	if err != nil {
		fmt.Println(err)
		return podsResponse.Pods, err
	}
	helpers.AddUserHeaders(jwtToken, req)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return podsResponse.Pods, err
	}
	defer helpers.CloseBody(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return podsResponse.Pods, err
	}

	err = json.Unmarshal(body, &podsResponse)
	if err != nil {
		return podsResponse.Pods, err
	}
	return podsResponse.Pods, nil
}
