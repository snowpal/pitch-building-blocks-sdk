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

func GetKey(jwtToken string) (response.Key, error) {
	resKey := response.Key{}
	route, err := helpers2.GetRoute(lib.RouteKeysGetKey)
	if err != nil {
		fmt.Println(err)
		return resKey, err
	}
	req, err := http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		fmt.Println(err)
		return resKey, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	res, err := helpers2.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resKey, err
	}

	defer helpers2.CloseBody(res.Body)

	body, _ := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resKey, err
	}

	err = json.Unmarshal(body, &resKey)
	if err != nil {
		fmt.Println(err)
		return resKey, err
	}
	return resKey, nil
}
