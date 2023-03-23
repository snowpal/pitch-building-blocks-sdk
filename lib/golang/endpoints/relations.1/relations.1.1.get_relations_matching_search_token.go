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

type SearchKeyRelationParam struct {
	Token        string
	CurrentKeyId string
}

type SearchBlockRelationParam struct {
	Token          string
	CurrentBlockId string
	KeyId          string
}

type SearchPodRelationParam struct {
	Token        string
	CurrentPodId string
	KeyId        string
}

type SearchBlockPodRelationParam struct {
	Token        string
	CurrentPodId string
	KeyId        string
	BlockId      string
}

func searchRelationsMatchingSearchToken(jwtToken string, route string) ([]response.SearchResource, error) {
	resSearchResources := response.SearchResources{}
	req, err := http.NewRequest(http.MethodGet, route, nil)
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

func SearchRelationsForKeyMatchingSearchToken(
	jwtToken string,
	relationParam SearchKeyRelationParam,
) ([]response.SearchResource, error) {
	var searchResults []response.SearchResource
	route, err := helpers.GetRoute(
		golang.RouteRelationsGetRelationsForKeyMatchingSearchToken,
		relationParam.Token,
		relationParam.CurrentKeyId,
	)
	if err != nil {
		fmt.Println(err)
		return searchResults, err
	}
	searchResults, err = searchRelationsMatchingSearchToken(jwtToken, route)
	if err != nil {
		fmt.Println(err)
		return searchResults, err
	}
	return searchResults, nil
}

func SearchRelationsForBlockMatchingSearchToken(
	jwtToken string,
	relationParam SearchBlockRelationParam,
) ([]response.SearchResource, error) {
	var searchResults []response.SearchResource
	route, err := helpers.GetRoute(
		golang.RouteRelationsGetRelationsForBlockMatchingSearchToken,
		relationParam.Token,
		relationParam.CurrentBlockId,
		relationParam.KeyId,
	)
	if err != nil {
		fmt.Println(err)
		return searchResults, err
	}
	searchResults, err = searchRelationsMatchingSearchToken(jwtToken, route)
	if err != nil {
		fmt.Println(err)
		return searchResults, err
	}
	return searchResults, nil
}

func SearchRelationsForPodMatchingSearchToken(
	jwtToken string,
	relationParam SearchPodRelationParam,
) ([]response.SearchResource, error) {
	var searchResults []response.SearchResource
	route, err := helpers.GetRoute(
		golang.RouteRelationsGetRelationsForKeyMatchingSearchToken,
		relationParam.Token,
		relationParam.CurrentPodId,
		relationParam.KeyId,
	)
	if err != nil {
		fmt.Println(err)
		return searchResults, err
	}
	searchResults, err = searchRelationsMatchingSearchToken(jwtToken, route)
	if err != nil {
		fmt.Println(err)
		return searchResults, err
	}
	return searchResults, nil
}

func SearchRelationsForBlockPodMatchingSearchToken(
	jwtToken string,
	relationParam SearchBlockPodRelationParam,
) ([]response.SearchResource, error) {
	var searchResults []response.SearchResource
	route, err := helpers.GetRoute(
		golang.RouteRelationsGetRelationsForBlockPodMatchingSearchToken,
		relationParam.Token,
		relationParam.CurrentPodId,
		relationParam.KeyId,
		relationParam.BlockId,
	)
	if err != nil {
		fmt.Println(err)
		return searchResults, err
	}
	searchResults, err = searchRelationsMatchingSearchToken(jwtToken, route)
	if err != nil {
		fmt.Println(err)
		return searchResults, err
	}
	return searchResults, nil
}
