package teacherKeys

import (
	"development/go/recipes/lib"
	"development/go/recipes/lib/helpers"
	"development/go/recipes/lib/structs/common"
	"development/go/recipes/lib/structs/request"
	"fmt"
	"net/http"
	"strings"
)

func PublishStudentGradesForABlock(
	jwtToken string,
	reqBody request.PublishGradesReqBody,
	podParam common.ResourceIdParam,
) error {
	requestBody, err := helpers.GetRequestBody(reqBody)
	if err != nil {
		fmt.Println(err)
		return err
	}
	payload := strings.NewReader(requestBody)
	route, err := helpers.GetRoute(
		lib.RouteTeacherKeysPublishStudentGradesForABlock,
		podParam.BlockId,
		podParam.KeyId,
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
