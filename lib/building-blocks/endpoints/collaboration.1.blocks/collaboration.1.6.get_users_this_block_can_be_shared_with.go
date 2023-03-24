package collaboration_1

import (
	"development/go/recipes/lib/building-blocks"
	"development/go/recipes/lib/building-blocks/helpers"
	"development/go/recipes/lib/building-blocks/structs/common"
	"development/go/recipes/lib/building-blocks/structs/response"
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
	route, err := helpers.GetRoute(
		building_blocks.RouteCollaborationGetUsersThisBlockCanBeSharedWith,
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
