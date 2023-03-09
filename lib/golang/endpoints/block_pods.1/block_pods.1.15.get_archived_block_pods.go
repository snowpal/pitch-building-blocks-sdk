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
	"strconv"
)

func GetArchivedBlockPods(jwtToken string, slimBlock common.SlimBlock, batchIndex int) ([]response.Pod, error) {
	podsResponse := response.Pods{}

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, helpers.GetRoute(golang.RouteBlockPodsGetArchivedBlockPods, strconv.Itoa(batchIndex), slimBlock.Key.ID, slimBlock.ID), nil)
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
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return podsResponse.Pods, err
	}
	fmt.Println(string(body))

	err = json.Unmarshal(body, &podsResponse)
	if err != nil {
		return podsResponse.Pods, err
	}
	return podsResponse.Pods, nil
}
