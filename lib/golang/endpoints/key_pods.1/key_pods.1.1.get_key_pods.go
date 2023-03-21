package key_pods_1

import (
	"development/go/recipes/lib/golang"
	"development/go/recipes/lib/golang/helpers"
	"development/go/recipes/lib/golang/structs/request"
	"development/go/recipes/lib/golang/structs/response"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func GetKeyPods(jwtToken string, podsParam request.GetPodsParam) ([]response.Pod, error) {
	resPods := response.Pods{}
	route, err := helpers.GetRoute(
		golang.RouteKeyPodsGetKeyPods,
		podsParam.KeyId,
		strconv.Itoa(podsParam.BatchIndex),
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
	res, err = helpers.MakeRequest(req)
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
