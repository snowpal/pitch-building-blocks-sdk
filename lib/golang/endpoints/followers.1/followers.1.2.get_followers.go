package followers

import (
	"development/go/recipes/lib/golang"
	"development/go/recipes/lib/golang/helpers"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Followers struct {
	Followers []Follower `json:"followers"`
}

type Follower struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

func GetFollowers(jwtToken string) ([]Follower, error) {
	resFollowers := Followers{}
	client := &http.Client{}
	route, err := helpers.GetRoute(golang.RouteFollowersGetFollowers)
	if err != nil {
		fmt.Println(err)
		return resFollowers.Followers, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		fmt.Println(err)
		return resFollowers.Followers, err
	}

	helpers.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = client.Do(req)
	if err != nil {
		fmt.Println(err)
		return resFollowers.Followers, err
	}

	defer helpers.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resFollowers.Followers, err
	}

	err = json.Unmarshal(body, &resFollowers)
	if err != nil {
		fmt.Println(err)
		return resFollowers.Followers, err
	}
	return resFollowers.Followers, nil
}
