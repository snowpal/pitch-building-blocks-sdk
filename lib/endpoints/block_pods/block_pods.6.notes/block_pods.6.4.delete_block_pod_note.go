package blockPods

import (
	"net/http"

	"github.com/snowpal/pitch-building-blocks-sdk/lib"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/helpers"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/request"
)

func DeleteBlockPodNote(jwtToken string, commentParam request.NoteIdParam) error {
	route, err := helpers.GetRoute(
		lib.RouteBlockPodsDeleteBlockPodNote,
		*commentParam.NoteId,
		commentParam.KeyId,
		*commentParam.BlockId,
		*commentParam.PodId,
	)
	req, err := http.NewRequest(http.MethodDelete, route, nil)
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
