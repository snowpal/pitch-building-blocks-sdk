package collaboration

import (
	"net/http"
	"strings"

	"github.com/snowpal/pitch-building-blocks-sdk/lib"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/helpers"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/common"
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
		return err
	}
	payload := strings.NewReader(requestBody)
	route, err := helpers.GetRoute(
		lib.RouteCollaborationBulkShareBlocksWithCollaborators,
		blockAclParam.UserId,
		blockAclParam.ResourceIds.KeyId,
	)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPatch, route, payload)
	if err != nil {
		return err
	}

	helpers.AddUserHeaders(jwtToken, req)

	_, err = helpers.MakeRequest(req)
	if err != nil {
		return err
	}
	return nil
}
