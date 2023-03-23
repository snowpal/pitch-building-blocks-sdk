package collaboration_1

import (
	"development/go/recipes/lib/golang"
	"development/go/recipes/lib/golang/helpers"
	"development/go/recipes/lib/golang/structs/common"
	"fmt"
	"net/http"
	"strings"
)

type BlockBulkShareReqBody struct {
	Acl      string `json:"blockAcl"`
	BlockIds string `json:"blockIds"`
}

func ShareBlocksWithCollaborators(
	jwtToken string,
	reqBody BlockBulkShareReqBody,
	blockAclParam common.AclParam,
) error {
	requestBody, err := helpers.GetRequestBody(reqBody)
	if err != nil {
		fmt.Println(err)
		return err
	}
	payload := strings.NewReader(requestBody)
	route, err := helpers.GetRoute(
		golang.RouteCollaborationBulkShareBlocksWithCollaborators,
		blockAclParam.UserId,
		blockAclParam.ResourceIds.KeyId,
	)
	if err != nil {
		fmt.Println(err)
		return err
	}

	req, err := http.NewRequest(http.MethodPatch, route, payload)
	if err != nil {
		fmt.Println(err)
		return err
	}

	helpers.AddUserHeaders(jwtToken, req)

	_, err = helpers.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
