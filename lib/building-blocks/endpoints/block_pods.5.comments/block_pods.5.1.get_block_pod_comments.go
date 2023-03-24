package block_pods_5

import (
	"development/go/recipes/lib/building-blocks"
	"development/go/recipes/lib/building-blocks/helpers"
	"development/go/recipes/lib/building-blocks/structs/request"
	"development/go/recipes/lib/building-blocks/structs/response"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetBlockPodComments(jwtToken string, commentParam request.CommentIdParam) ([]response.Comment, error) {
	resComments := response.Comments{}
	route, err := helpers.GetRoute(
		building_blocks.RouteBlockPodsGetBlockPodComments,
		*commentParam.PodId,
		commentParam.KeyId,
		*commentParam.BlockId,
	)
	if err != nil {
		fmt.Println(err)
		return resComments.Comments, err
	}

	req, err := http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		fmt.Println(err)
		return resComments.Comments, err
	}

	helpers.AddUserHeaders(jwtToken, req)

	res, err := helpers.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resComments.Comments, err
	}

	defer helpers.CloseBody(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resComments.Comments, err
	}

	err = json.Unmarshal(body, &resComments)
	if err != nil {
		fmt.Println(err)
		return resComments.Comments, err
	}
	return resComments.Comments, nil
}