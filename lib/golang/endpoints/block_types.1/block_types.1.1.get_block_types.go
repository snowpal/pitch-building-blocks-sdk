package block_types

import (
	"development/go/recipes/lib/golang"
	"development/go/recipes/lib/golang/helpers"
	"development/go/recipes/lib/golang/structs/response"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func GetBlockTypes(jwtToken string, includeCounts bool) ([]response.BlockType, error) {
	resBlockTypes := response.BlockTypes{}
	client := &http.Client{}
	route, err := helpers.GetRoute(golang.RouteBlockTypesGetBlockTypes, strconv.FormatBool(includeCounts))
	req, err := http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		fmt.Println(err)
		return resBlockTypes.BlockTypes, err
	}

	helpers.AddUserHeaders(jwtToken, req)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return resBlockTypes.BlockTypes, err
	}

	defer helpers.CloseBody(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resBlockTypes.BlockTypes, err
	}

	err = json.Unmarshal(body, &resBlockTypes)
	if err != nil {
		fmt.Println(err)
		return resBlockTypes.BlockTypes, err
	}
	return resBlockTypes.BlockTypes, nil
}
