package keys_2

import (
	"development/go/recipes/lib/building-blocks"
	"development/go/recipes/lib/building-blocks/helpers"
	"development/go/recipes/lib/building-blocks/structs/response"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetPodsBasedOnPodTypesInKey(jwtToken string, keyId string) (response.PodTypesKey, error) {
	resPodTypesKey := response.PodTypesKey{}
	route, err := helpers.GetRoute(building_blocks.RouteKeysGetPodsBasedOnPodTypesInKey, keyId)
	if err != nil {
		fmt.Println(err)
		return resPodTypesKey, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		fmt.Println(err)
		return resPodTypesKey, err
	}

	helpers.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resPodTypesKey, err
	}

	defer helpers.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resPodTypesKey, err
	}

	err = json.Unmarshal(body, &resPodTypesKey)
	if err != nil {
		fmt.Println(err)
		return resPodTypesKey, err
	}
	return resPodTypesKey, nil
}