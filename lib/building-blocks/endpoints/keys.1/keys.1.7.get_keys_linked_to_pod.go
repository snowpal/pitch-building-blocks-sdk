package keys_1

import (
	"development/go/recipes/lib/building-blocks"
	"development/go/recipes/lib/building-blocks/helpers"
	"development/go/recipes/lib/building-blocks/structs/common"
	"development/go/recipes/lib/building-blocks/structs/response"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetKeysLinkedToPod(jwtToken string, keyParam common.ResourceIdParam) ([]response.Key, error) {
	resKeys := response.Keys{}
	route, err := helpers.GetRoute(building_blocks.RouteKeysGetKeysLinkedToPod, keyParam.PodId, keyParam.KeyId)
	if err != nil {
		fmt.Println(err)
		return resKeys.Keys, err
	}
	req, err := http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		fmt.Println(err)
		return resKeys.Keys, err
	}

	helpers.AddUserHeaders(jwtToken, req)

	res, err := helpers.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resKeys.Keys, err
	}

	defer helpers.CloseBody(res.Body)

	body, _ := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resKeys.Keys, err
	}

	err = json.Unmarshal(body, &resKeys)
	if err != nil {
		fmt.Println(err)
		return resKeys.Keys, err
	}
	return resKeys.Keys, nil
}
