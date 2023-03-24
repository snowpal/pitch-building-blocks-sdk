package collaboration_2

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

func UnshareBlockPodWithCollaborator(jwtToken string, podAclParam common.AclParam) (response.Pod, error) {
	resPod := response.Pod{}
	route, err := helpers.GetRoute(
		building_blocks.RouteCollaborationUnshareBlockPodFromCollaborator,
		podAclParam.ResourceIds.PodId,
		podAclParam.UserId,
		podAclParam.ResourceIds.KeyId,
		podAclParam.ResourceIds.BlockId,
	)
	req, err := http.NewRequest(http.MethodPatch, route, nil)
	if err != nil {
		fmt.Println(err)
		return resPod, err
	}

	helpers.AddUserHeaders(jwtToken, req)

	res, err := helpers.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resPod, err
	}

	defer helpers.CloseBody(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resPod, err
	}

	err = json.Unmarshal(body, &resPod)
	if err != nil {
		fmt.Println(err)
		return resPod, err
	}
	return resPod, nil
}
