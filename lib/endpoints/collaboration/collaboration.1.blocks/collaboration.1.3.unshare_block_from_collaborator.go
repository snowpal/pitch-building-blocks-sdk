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

func UnshareBlockWithCollaborator(jwtToken string, blockAclParam common.AclParam) (response.Block, error) {
	resBlock := response.Block{}
	route, err := helpers2.GetRoute(
		lib.RouteCollaborationUnshareBlockFromCollaborator,
		blockAclParam.ResourceIds.BlockId,
		blockAclParam.UserId,
		blockAclParam.ResourceIds.KeyId,
	)
	req, err := http.NewRequest(http.MethodPatch, route, nil)
	if err != nil {
		fmt.Println(err)
		return resBlock, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	res, err := helpers2.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resBlock, err
	}

	defer helpers2.CloseBody(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resBlock, err
	}

	err = json.Unmarshal(body, &resBlock)
	if err != nil {
		fmt.Println(err)
		return resBlock, err
	}
	return resBlock, nil
}
