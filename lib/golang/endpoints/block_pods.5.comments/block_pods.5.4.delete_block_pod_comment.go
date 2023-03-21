package block_pods_5

import (
	"development/go/recipes/lib/golang"
	"development/go/recipes/lib/golang/helpers"
	"development/go/recipes/lib/golang/structs/request"
	"fmt"
	"net/http"
)

func DeleteBlockPodComment(jwtToken string, commentParam request.CommentIdParam) error {
	route, err := helpers.GetRoute(
		golang.RouteBlockPodsDeleteBlockPodComment,
		*commentParam.CommentId,
		commentParam.KeyId,
		*commentParam.BlockId,
		*commentParam.PodId,
	)
	req, err := http.NewRequest(http.MethodDelete, route, nil)
	if err != nil {
		fmt.Println(err)
		return err
	}

	helpers.AddUserHeaders(jwtToken, req)

	res, err := helpers.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
