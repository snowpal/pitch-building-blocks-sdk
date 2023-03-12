package blocks_5

import (
	"development/go/recipes/lib/golang"
	"development/go/recipes/lib/golang/helpers"
	"development/go/recipes/lib/golang/structs/request"
	"fmt"
	"net/http"
)

func DeleteBlockComment(jwtToken string, commentParam request.CommentIdParam) error {
	client := &http.Client{}
	route, err := helpers.GetRoute(
		golang.RouteBlocksDeleteBlockComment,
		*commentParam.CommentId,
		commentParam.KeyId,
		*commentParam.BlockId,
	)
	req, err := http.NewRequest(http.MethodDelete, route, nil)
	if err != nil {
		fmt.Println(err)
		return err
	}

	helpers.AddUserHeaders(jwtToken, req)

	_, err = client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
