package block_types

import (
	"development/go/recipes/lib/building-blocks"
	"development/go/recipes/lib/building-blocks/helpers"
	"development/go/recipes/lib/building-blocks/structs/response"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func GetBlockTypes(jwtToken string, includeCounts bool) ([]response.BlockType, error) {
	resBlockTypes := response.BlockTypes{}
	route, err := helpers.GetRoute(building_blocks.RouteBlockTypesGetBlockTypes, strconv.FormatBool(includeCounts))
	if err != nil {
		fmt.Println(err)
		return resBlockTypes.BlockTypes, err
	}
	req, err := http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		fmt.Println(err)
		return resBlockTypes.BlockTypes, err
	}

	helpers.AddUserHeaders(jwtToken, req)

	res, err := helpers.MakeRequest(req)
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