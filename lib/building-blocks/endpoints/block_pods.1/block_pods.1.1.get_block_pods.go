package block_pods

import (
	"development/go/recipes/lib/building-blocks"
	"development/go/recipes/lib/building-blocks/helpers"
	"development/go/recipes/lib/building-blocks/structs/request"
	"development/go/recipes/lib/building-blocks/structs/response"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func GetBlockPods(jwtToken string, podsParam request.GetPodsParam) ([]response.Pod, error) {
	resPods := response.Pods{}
	route, err := helpers.GetRoute(
		building_blocks.RouteBlockPodsGetBlockPods,
		*podsParam.BlockId,
		strconv.Itoa(podsParam.BatchIndex),
		podsParam.KeyId,
	)
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

	helpers.AddUserHeaders(jwtToken, req)

	var res *http.Response
	_, err = helpers.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resPods.Pods, err
	}

	defer helpers.CloseBody(res.Body)

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