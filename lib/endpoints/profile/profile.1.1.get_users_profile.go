package profiles

import (
	"development/go/recipes/lib"
	helpers2 "development/go/recipes/lib/helpers"
	"development/go/recipes/lib/structs/response"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetUsersProfile(jwtToken string) (response.Profile, error) {
	resProfile := response.Profile{}
	route, err := helpers2.GetRoute(lib.RouteProfileGetUsersProfile)
	if err != nil {
		fmt.Println(err)
		return resProfile, err
	}
	req, err := http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		fmt.Println(err)
		return resProfile, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	res, err := helpers2.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resProfile, err
	}

	defer helpers2.CloseBody(res.Body)

	body, _ := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resProfile, err
	}

	err = json.Unmarshal(body, &resProfile)
	if err != nil {
		fmt.Println(err)
		return resProfile, err
	}
	return resProfile, nil
}
