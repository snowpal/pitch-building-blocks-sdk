package teacherKeys

import (
	"development/go/recipes/lib"
	helpers2 "development/go/recipes/lib/helpers"
	request2 "development/go/recipes/lib/structs/request"
	"development/go/recipes/lib/structs/response"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func AssignBlockGradeToStudent(
	jwtToken string,
	reqBody request2.UpdateScaleValueReqBody,
	classroomParam request2.ClassroomIdParam,
) (response.UpdateBlockScaleValue, error) {
	resBlockScaleValue := response.UpdateBlockScaleValue{}
	requestBody, err := helpers2.GetRequestBody(reqBody)
	if err != nil {
		fmt.Println(err)
		return resBlockScaleValue, err
	}
	payload := strings.NewReader(requestBody)
	route, err := helpers2.GetRoute(
		lib.RouteTeacherKeysAssignGradeToStudent,
		classroomParam.ResourceIds.BlockId,
		classroomParam.StudentId,
		classroomParam.ResourceIds.KeyId,
	)
	if err != nil {
		fmt.Println(err)
		return resBlockScaleValue, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodPatch, route, payload)
	if err != nil {
		fmt.Println(err)
		return resBlockScaleValue, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers2.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resBlockScaleValue, err
	}

	defer helpers2.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resBlockScaleValue, err
	}

	err = json.Unmarshal(body, &resBlockScaleValue)
	if err != nil {
		fmt.Println(err)
		return resBlockScaleValue, err
	}
	return resBlockScaleValue, nil
}
