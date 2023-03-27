package blocks

import (
	"development/go/recipes/lib"
	helpers2 "development/go/recipes/lib/helpers"
	"development/go/recipes/lib/structs/request"
	"development/go/recipes/lib/structs/response"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetBlockComments(jwtToken string, commentParam request.CommentIdParam) ([]response.Comment, error) {
	resComments := response.Comments{}
	route, err := helpers2.GetRoute(
		lib.RouteBlocksGetBlockComments,
		*commentParam.BlockId,
		commentParam.KeyId,
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

	helpers2.AddUserHeaders(jwtToken, req)

	res, err := helpers2.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resComments.Comments, err
	}

	defer helpers2.CloseBody(res.Body)

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
