package keys

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/snowpal/pitch-building-blocks-sdk/lib"
	helpers2 "github.com/snowpal/pitch-building-blocks-sdk/lib/helpers"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/common"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/response"
)

func GetKeysLinkedToPod(jwtToken string, keyParam common.ResourceIdParam) ([]response.Key, error) {
	resKeys := response.Keys{}
	route, err := helpers2.GetRoute(lib.RouteKeysGetKeysLinkedToPod, keyParam.PodId, keyParam.KeyId)
	if err != nil {
		fmt.Println(err)
		return resKeys.Keys, err
	}
	req, err := http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		fmt.Println(err)
		return resKeys.Keys, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	res, err := helpers2.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resKeys.Keys, err
	}

	defer helpers2.CloseBody(res.Body)

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
