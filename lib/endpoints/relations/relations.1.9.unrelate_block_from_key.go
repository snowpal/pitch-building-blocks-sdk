package relations

import (
	"net/http"

	"github.com/snowpal/pitch-building-blocks-sdk/lib"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/helpers"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/request"
)

func UnrelateKeyFromBlock(jwtToken string, relationParam request.KeyToBlockRelationParam) error {
	route, err := helpers.GetRoute(
		lib.RouteRelationsUnrelateBlockFromKey,
		relationParam.KeyId,
		relationParam.TargetBlockId,
	)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPatch, route, nil)
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
