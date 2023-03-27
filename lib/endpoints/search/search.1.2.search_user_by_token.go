package search

import (
	"development/go/recipes/lib"
	helpers2 "development/go/recipes/lib/helpers"
	"development/go/recipes/lib/structs/response"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func SrchUserByToken(jwtToken string, searchToken string) ([]response.SearchUser, error) {
	resSearchUsers := response.SearchUsers{}
	route, err := helpers2.GetRoute(lib.RouteSearchSearchUserByToken, searchToken)
	if err != nil {
		fmt.Println(err)
		return resSearchUsers.SearchUsers, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		fmt.Println(err)
		return resSearchUsers.SearchUsers, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers2.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resSearchUsers.SearchUsers, err
	}

	defer helpers2.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resSearchUsers.SearchUsers, err
	}

	err = json.Unmarshal(body, &resSearchUsers)
	if err != nil {
		fmt.Println(err)
		return resSearchUsers.SearchUsers, err
	}
	return resSearchUsers.SearchUsers, nil
}
