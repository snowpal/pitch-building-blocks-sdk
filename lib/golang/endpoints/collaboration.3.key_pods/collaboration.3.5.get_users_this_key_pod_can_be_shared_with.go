package collaboration_3

import (
	"development/go/recipes/lib/golang"
	"development/go/recipes/lib/golang/helpers"
	"development/go/recipes/lib/golang/structs/common"
	"development/go/recipes/lib/golang/structs/response"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetUsersThisKeyPodCanBeSharedWith(
	jwtToken string,
	podAclParam common.SearchUsersParam,
) ([]response.SearchUser, error) {
	resUsers := response.SearchUsers{}
	route, err := helpers.GetRoute(
		golang.RouteCollaborationGetUsersThisKeyPodCanBeSharedWith,
		podAclParam.ResourceIds.PodId,
		podAclParam.ResourceIds.KeyId,
		podAclParam.SearchToken,
	)
	if err != nil {
		fmt.Println(err)
		return resUsers.SearchUsers, err
	}

	req, err := http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		fmt.Println(err)
		return resUsers.SearchUsers, err
	}

	helpers.AddUserHeaders(jwtToken, req)

	res, err := helpers.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resUsers.SearchUsers, err
	}

	defer helpers.CloseBody(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resUsers.SearchUsers, err
	}

	err = json.Unmarshal(body, &resUsers)
	if err != nil {
		fmt.Println(err)
		return resUsers.SearchUsers, err
	}
	return resUsers.SearchUsers, nil
}
