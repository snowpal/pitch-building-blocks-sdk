package relations

import (
	"development/go/recipes/lib"
	helpers2 "development/go/recipes/lib/helpers"
	"development/go/recipes/lib/structs/common"
	"development/go/recipes/lib/structs/response"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetRelationsForBlock(jwtToken string, relationParam common.ResourceIdParam) (response.Relationships, error) {
	resRelations := response.Relations{}
	route, err := helpers2.GetRoute(
		lib.RouteRelationsGetRelationsForBlock,
		relationParam.BlockId,
		relationParam.KeyId,
	)
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

	helpers2.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers2.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resRelations.Relationships, err
	}

	defer helpers2.CloseBody(res.Body)

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
