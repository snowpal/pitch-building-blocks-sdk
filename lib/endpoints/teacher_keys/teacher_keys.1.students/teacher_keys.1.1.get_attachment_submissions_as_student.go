package teacherKeys

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/snowpal/pitch-building-blocks-sdk/lib"
	helpers2 "github.com/snowpal/pitch-building-blocks-sdk/lib/helpers"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/common"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/response"
)

func GetAttachmentSubmissionsAsStudent(
	jwtToken string,
	submissionParam common.ResourceIdParam,
) ([]response.Attachment, error) {
	resAttachments := response.Attachments{}
	route, err := helpers2.GetRoute(
		lib.RouteTeacherKeysGetAttachmentSubmissionsAsStudent,
		submissionParam.PodId,
		submissionParam.KeyId,
		submissionParam.BlockId,
	)
	if err != nil {
		fmt.Println(err)
		return resAttachments.Attachments, err
	}

	req, err := http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		fmt.Println(err)
		return resAttachments.Attachments, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	res, err := helpers2.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resAttachments.Attachments, err
	}

	defer helpers2.CloseBody(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resAttachments.Attachments, err
	}

	err = json.Unmarshal(body, &resAttachments)
	if err != nil {
		fmt.Println(err)
		return resAttachments.Attachments, err
	}
	return resAttachments.Attachments, nil
}
