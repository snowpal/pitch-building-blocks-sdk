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

func GetPodsBasedOnPodTypesInKey(jwtToken string, keyId string) (response.PodTypesKey, error) {
	resPodTypesKey := response.PodTypesKey{}
	route, err := helpers2.GetRoute(lib.RouteKeysGetPodsBasedOnPodTypesInKey, keyId)
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

	helpers2.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers2.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resPodTypesKey, err
	}

	defer helpers2.CloseBody(res.Body)

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
