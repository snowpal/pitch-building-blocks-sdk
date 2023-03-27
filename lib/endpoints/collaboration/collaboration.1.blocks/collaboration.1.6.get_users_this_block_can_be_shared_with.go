package collaboration

import (
	"development/go/recipes/lib"
	helpers2 "development/go/recipes/lib/helpers"
	"development/go/recipes/lib/structs/common"
	"development/go/recipes/lib/structs/response"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetUsersThisBlockCanBeSharedWith(
	jwtToken string,
	blockAclParam common.SearchUsersParam,
) ([]response.SearchUser, error) {
	resUsers := response.SearchUsers{}
	route, err := helpers2.GetRoute(
		lib.RouteCollaborationGetUsersThisBlockCanBeSharedWith,
		blockAclParam.ResourceIds.BlockId,
		blockAclParam.ResourceIds.KeyId,
		blockAclParam.SearchToken,
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

	helpers2.AddUserHeaders(jwtToken, req)

	res, err := helpers2.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resUsers.SearchUsers, err
	}

	defer helpers2.CloseBody(res.Body)

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
