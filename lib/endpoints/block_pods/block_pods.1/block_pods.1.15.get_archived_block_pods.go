package blockPods

import (
	"development/go/recipes/lib"
	helpers2 "development/go/recipes/lib/helpers"
	"development/go/recipes/lib/structs/request"
	"development/go/recipes/lib/structs/response"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func GetArchivedBlockPods(jwtToken string, podsParam request.GetPodsParam) ([]response.Pod, error) {
	resPods := response.Pods{}
	route, err := helpers2.GetRoute(
		lib.RouteBlockPodsGetArchivedBlockPods,
		strconv.Itoa(podsParam.BatchIndex),
		podsParam.KeyId,
		*podsParam.BlockId,
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

	helpers2.AddUserHeaders(jwtToken, req)

	var res *http.Response
	_, err = helpers2.MakeRequest(req)
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
