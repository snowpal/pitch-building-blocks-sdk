package collaboration

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/snowpal/pitch-building-blocks-sdk/lib"
	helpers2 "github.com/snowpal/pitch-building-blocks-sdk/lib/helpers"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/common"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/response"
)

func UnshareBlockPodWithCollaborator(jwtToken string, podAclParam common.AclParam) (response.Pod, error) {
	resPod := response.Pod{}
	route, err := helpers2.GetRoute(
		lib.RouteCollaborationUnshareBlockPodFromCollaborator,
		podAclParam.ResourceIds.PodId,
		podAclParam.UserId,
		podAclParam.ResourceIds.KeyId,
		podAclParam.ResourceIds.BlockId,
	)
	req, err := http.NewRequest(http.MethodPatch, route, nil)
	if err != nil {
		return resPod, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	res, err := helpers2.MakeRequest(req)
	if err != nil {
		return resPod, err
	}

	defer helpers2.CloseBody(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return resPod, err
	}

	err = json.Unmarshal(body, &resPod)
	if err != nil {
		return resPod, err
	}
	return resPod, nil
}
