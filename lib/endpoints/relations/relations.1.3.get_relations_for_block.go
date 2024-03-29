package relations

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/snowpal/pitch-building-blocks-sdk/lib"
	helpers2 "github.com/snowpal/pitch-building-blocks-sdk/lib/helpers"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/common"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/response"
)

func GetRelationsForBlock(jwtToken string, relationParam common.ResourceIdParam) (response.Relationships, error) {
	resRelations := response.Relations{}
	route, err := helpers2.GetRoute(
		lib.RouteRelationsGetRelationsForBlock,
		relationParam.BlockId,
		relationParam.KeyId,
	)
	if err != nil {
		return resRelations.Relationships, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		return resRelations.Relationships, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers2.MakeRequest(req)
	if err != nil {
		return resRelations.Relationships, err
	}

	defer helpers2.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		return resRelations.Relationships, err
	}

	err = json.Unmarshal(body, &resRelations)
	if err != nil {
		return resRelations.Relationships, err
	}
	return resRelations.Relationships, nil
}
