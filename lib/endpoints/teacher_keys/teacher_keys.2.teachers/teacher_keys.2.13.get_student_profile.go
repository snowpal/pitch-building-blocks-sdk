package teacherKeys

import (
	"development/go/recipes/lib"
	helpers2 "development/go/recipes/lib/helpers"
	"development/go/recipes/lib/structs/request"
	"development/go/recipes/lib/structs/response"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetStudentProfile(jwtToken string, classroomParam request.ClassroomIdParam) (response.Student, error) {
	resStudent := response.Student{}
	route, err := helpers2.GetRoute(
		lib.RouteTeacherKeysGetStudentProfile,
		classroomParam.StudentId,
		classroomParam.ResourceIds.BlockId,
		classroomParam.ResourceIds.KeyId,
	)
	if err != nil {
		fmt.Println(err)
		return resStudent, err
	}

	req, err := http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		fmt.Println(err)
		return resStudent, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	res, err := helpers2.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resStudent, err
	}

	defer helpers2.CloseBody(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resStudent, err
	}

	err = json.Unmarshal(body, &resStudent)
	if err != nil {
		fmt.Println(err)
		return resStudent, err
	}
	return resStudent, nil
}
