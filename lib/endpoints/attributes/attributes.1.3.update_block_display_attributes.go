package attributes

import (
	"net/http"
	"strings"

	"github.com/snowpal/pitch-building-blocks-sdk/lib"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/helpers"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/common"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/request"
)

func UpdateBlockAttrs(jwtToken string, blockParam common.ResourceIdParam, attribute request.ResourceAttributeReqBody) error {
	requestBody, err := helpers.GetRequestBody(attribute)
	if err != nil {
		return err
	}
	payload := strings.NewReader(requestBody)

	var route string
	route, err = helpers.GetRoute(
		lib.RouteAttributesUpdateBlockDisplayAttributes,
		blockParam.BlockId,
		blockParam.KeyId,
	)
	if err != nil {
		return err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodPatch, route, payload)
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
