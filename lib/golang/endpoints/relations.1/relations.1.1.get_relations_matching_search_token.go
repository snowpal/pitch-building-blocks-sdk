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

type SearchRelationParam struct {
	Token          string
	KeyId          *string
	BlockId        *string
	CurrentKeyId   *string
	CurrentBlockId *string
	CurrentPodId   *string
}

func SearchRelationsMatchingSearchToken(
	jwtToken string,
	relationParam SearchRelationParam,
) ([]response.SearchResource, error) {
	resSearchResources := response.SearchResources{}
	route, err := helpers.GetRoute(
		golang.RouteRelationsGetRelationsMatchingSearchToken,
		relationParam.Token,
		*relationParam.CurrentKeyId,
		*relationParam.CurrentBlockId,
		*relationParam.CurrentPodId,
		*relationParam.KeyId,
		*relationParam.BlockId,
	)
	if err != nil {
		fmt.Println(err)
		return resSearchResources.Results, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		fmt.Println(err)
		return resSearchResources.Results, err
	}

	helpers.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resSearchResources.Results, err
	}

	defer helpers.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resSearchResources.Results, err
	}

	err = json.Unmarshal(body, &resSearchResources)
	if err != nil {
		fmt.Println(err)
		return resSearchResources.Results, err
	}
	return resSearchResources.Results, nil
}
