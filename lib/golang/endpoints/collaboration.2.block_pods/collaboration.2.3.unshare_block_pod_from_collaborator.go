package collaboration_2

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

func UnshareBlockPodWithCollaborator(jwtToken string, podAclParam common.AclParam) (response.Pod, error) {
	resPod := response.Pod{}
	route, err := helpers.GetRoute(
		golang.RouteCollaborationUnshareBlockPodFromCollaborator,
		*podAclParam.PodId,
		podAclParam.UserId,
		podAclParam.KeyId,
		*podAclParam.BlockId,
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
