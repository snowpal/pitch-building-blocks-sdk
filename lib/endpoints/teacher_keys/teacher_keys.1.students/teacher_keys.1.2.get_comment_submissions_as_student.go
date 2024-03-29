package teacherKeys

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/snowpal/pitch-building-blocks-sdk/lib"
	helpers2 "github.com/snowpal/pitch-building-blocks-sdk/lib/helpers"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/common"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/response"
)

func GetCommentSubmissionsAsStudent(
	jwtToken string,
	submissionParam common.ResourceIdParam,
) ([]response.Comment, error) {
	resComments := response.Comments{}
	route, err := helpers2.GetRoute(
		lib.RouteTeacherKeysGetCommentSubmissionsAsStudent,
		submissionParam.PodId,
		submissionParam.KeyId,
		submissionParam.BlockId,
	)
	if err != nil {
		return resComments.Comments, err
	}

	req, err := http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		return resComments.Comments, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	res, err := helpers2.MakeRequest(req)
	if err != nil {
		return resComments.Comments, err
	}

	defer helpers2.CloseBody(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return resComments.Comments, err
	}

	err = json.Unmarshal(body, &resComments)
	if err != nil {
		return resComments.Comments, err
	}
	return resComments.Comments, nil
}
