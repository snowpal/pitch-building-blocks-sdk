package followers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/snowpal/pitch-building-blocks-sdk/lib"
	helpers2 "github.com/snowpal/pitch-building-blocks-sdk/lib/helpers"
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
	route, err := helpers2.GetRoute(lib.RouteFollowersGetFollowers)
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

	helpers2.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers2.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resFollowers.Followers, err
	}

	defer helpers2.CloseBody(res.Body)

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
