package key_pods_4

import (
	"development/go/recipes/lib/building-blocks"
	"development/go/recipes/lib/building-blocks/helpers"
	"development/go/recipes/lib/building-blocks/structs/request"
	"fmt"
	"net/http"
)

func DeleteKeyPodComment(jwtToken string, commentParam request.CommentIdParam) error {
	route, err := helpers.GetRoute(
		building_blocks.RouteKeyPodsDeleteKeyPodComment,
		*commentParam.CommentId,
		commentParam.KeyId,
		*commentParam.PodId,
	)
	req, err := http.NewRequest(http.MethodDelete, route, nil)
	if err != nil {
		fmt.Println(err)
		return err
	}

	helpers.AddUserHeaders(jwtToken, req)

	_, err = helpers.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
