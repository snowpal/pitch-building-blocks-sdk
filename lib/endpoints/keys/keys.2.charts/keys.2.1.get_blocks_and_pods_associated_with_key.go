package keys

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/snowpal/pitch-building-blocks-sdk/lib"
	helpers2 "github.com/snowpal/pitch-building-blocks-sdk/lib/helpers"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/response"
)

func GetBlocksAndPodsAssociatedWithKey(jwtToken string, keyId string) (response.UserKey, error) {
	resUserKey := response.UserKey{}
	route, err := helpers2.GetRoute(lib.RouteKeysGetBlocksAndPodsAssociatedWithKey, keyId)
	if err != nil {
		fmt.Println(err)
		return resUserKey, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		fmt.Println(err)
		return resUserKey, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers2.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resUserKey, err
	}

	defer helpers2.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resUserKey, err
	}

	err = json.Unmarshal(body, &resUserKey)
	if err != nil {
		fmt.Println(err)
		return resUserKey, err
	}
	return resUserKey, nil
}
