package collaboration

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/snowpal/pitch-building-blocks-sdk/lib"
	helpers2 "github.com/snowpal/pitch-building-blocks-sdk/lib/helpers"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/common"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/response"
)

func GetUsersThisBlockPodCanBeSharedWith(
	jwtToken string,
	podAclParam common.SearchUsersParam,
) ([]response.SearchUser, error) {
	resUsers := response.SearchUsers{}
	route, err := helpers2.GetRoute(
		lib.RouteCollaborationGetUsersThisBlockPodCanBeSharedWith,
		podAclParam.ResourceIds.PodId,
		podAclParam.ResourceIds.KeyId,
		podAclParam.ResourceIds.BlockId,
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
