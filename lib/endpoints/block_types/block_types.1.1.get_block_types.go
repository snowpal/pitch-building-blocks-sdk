package block_types

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/snowpal/pitch-building-blocks-sdk/lib"
	helpers2 "github.com/snowpal/pitch-building-blocks-sdk/lib/helpers"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/response"
)

func GetBlockTypes(jwtToken string, includeCounts bool) ([]response.BlockType, error) {
	resBlockTypes := response.BlockTypes{}
	route, err := helpers2.GetRoute(lib.RouteBlockTypesGetBlockTypes, strconv.FormatBool(includeCounts))
	if err != nil {
		fmt.Println(err)
		return resBlockTypes.BlockTypes, err
	}
	req, err := http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		fmt.Println(err)
		return resBlockTypes.BlockTypes, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	res, err := helpers2.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resBlockTypes.BlockTypes, err
	}

	defer helpers2.CloseBody(res.Body)

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
