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

func GetBlockPods(jwtToken string, block common.SlimBlock, batchIndex int) ([]response.Pod, error) {
	podsResponse := response.Pods{}
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, helpers.GetRoute(golang.RouteBlockPodsGetBlockPods, block.ID,
		strconv.Itoa(batchIndex), block.Key.ID), nil)
	if err != nil {
		fmt.Println(err)
		return podsResponse.Pods, err
	}

	helpers.AddUserHeaders(jwtToken, req)

	res, _ := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return podsResponse.Pods, err
	}

	defer helpers.CloseBody(res.Body)

	body, _ := io.ReadAll(res.Body)
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
