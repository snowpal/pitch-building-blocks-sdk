package blockPods

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/snowpal/pitch-building-blocks-sdk/lib"
	helpers2 "github.com/snowpal/pitch-building-blocks-sdk/lib/helpers"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/request"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/response"
)

func GetBlockPodChecklists(jwtToken string, checklistParam request.ChecklistIdParam) ([]response.Checklist, error) {
	resChecklists := response.Checklists{}
	route, err := helpers2.GetRoute(
		lib.RouteBlockPodsGetBlockPodChecklists,
		*checklistParam.PodId,
		checklistParam.KeyId,
		*checklistParam.BlockId,
	)
	if err != nil {
		return resChecklists.Checklists, err
	}

	req, err := http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		return resChecklists.Checklists, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	res, err := helpers2.MakeRequest(req)
	if err != nil {
		return resChecklists.Checklists, err
	}

	defer helpers2.CloseBody(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return resChecklists.Checklists, err
	}

	err = json.Unmarshal(body, &resChecklists)
	if err != nil {
		return resChecklists.Checklists, err
	}
	return resChecklists.Checklists, nil
}
