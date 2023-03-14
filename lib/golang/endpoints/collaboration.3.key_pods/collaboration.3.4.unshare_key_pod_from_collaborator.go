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

func UnshareKeyPodWithCollaborator(jwtToken string, podAclParam common.AclParam) (response.Pod, error) {
	resPod := response.Pod{}
	client := &http.Client{}
	route, err := helpers.GetRoute(
		golang.RouteCollaborationUnshareKeyPodFromCollaborator,
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

	res, err := client.Do(req)
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
