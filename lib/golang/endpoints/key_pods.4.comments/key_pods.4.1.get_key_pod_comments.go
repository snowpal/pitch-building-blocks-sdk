package key_pods_4

import (
	"development/go/recipes/lib/golang"
	"development/go/recipes/lib/golang/helpers"
	"development/go/recipes/lib/golang/structs/request"
	"development/go/recipes/lib/golang/structs/response"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetKeyPodComments(jwtToken string, commentParam request.CommentIdParam) ([]response.Comment, error) {
	resComments := response.Comments{}
	route, err := helpers.GetRoute(
		golang.RouteKeyPodsGetKeyPodComments,
		*commentParam.PodId,
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
