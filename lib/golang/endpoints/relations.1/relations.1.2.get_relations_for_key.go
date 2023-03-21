package relations

import (
	"development/go/recipes/lib/golang"
	"development/go/recipes/lib/golang/helpers"
	"development/go/recipes/lib/golang/structs/response"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetRelationsForKey(jwtToken string, keyId string) (response.Relationships, error) {
	resRelations := response.Relations{}
	route, err := helpers.GetRoute(golang.RouteRelationsGetRelationsForKey, keyId)
	if err != nil {
		fmt.Println(err)
		return resRelations.Relationships, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		fmt.Println(err)
		return resRelations.Relationships, err
	}

	helpers.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resRelations.Relationships, err
	}

	defer helpers.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resRelations.Relationships, err
	}

	err = json.Unmarshal(body, &resRelations)
	if err != nil {
		fmt.Println(err)
		return resRelations.Relationships, err
	}
	return resRelations.Relationships, nil
}
