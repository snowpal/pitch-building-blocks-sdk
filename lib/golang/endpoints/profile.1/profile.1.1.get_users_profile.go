package profiles

import (
	"development/go/recipes/lib/golang"
	"development/go/recipes/lib/golang/helpers"
	"development/go/recipes/lib/golang/structs/response"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetUsersProfile(jwtToken string) (response.Profile, error) {
	resProfile := response.Profile{}
	client := &http.Client{}
	route, err := helpers.GetRoute(golang.RouteProfileGetUsersProfile)
	if err != nil {
		fmt.Println(err)
		return resProfile, err
	}
	req, err := http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		fmt.Println(err)
		return resProfile, err
	}

	helpers.AddUserHeaders(jwtToken, req)

	res, _ := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return resProfile, err
	}

	defer helpers.CloseBody(res.Body)

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
