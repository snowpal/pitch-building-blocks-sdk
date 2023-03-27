package user

import (
	"development/go/recipes/lib"
	helpers2 "development/go/recipes/lib/helpers"
	"development/go/recipes/lib/structs/response"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetUserByEmail(jwtToken string, email string) (response.User, error) {
	resUser := response.User{}
	route, err := helpers2.GetRoute(lib.RouteUserGetUserByEmail, email)
	if err != nil {
		fmt.Println(err)
		return resUser, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		fmt.Println(err)
		return resUser, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers2.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resUser, err
	}

	defer helpers2.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resUser, err
	}

	err = json.Unmarshal(body, &resUser)
	if err != nil {
		fmt.Println(err)
		return resUser, err
	}
	return resUser, nil
}
