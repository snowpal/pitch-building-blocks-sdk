package teacher_keys_2

import (
	"development/go/recipes/lib/golang"
	"development/go/recipes/lib/golang/helpers"
	"development/go/recipes/lib/golang/structs/request"
	"development/go/recipes/lib/golang/structs/response"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetStudentProfile(jwtToken string, classroomParam request.ClassroomIdParam) (response.Student, error) {
	resStudent := response.Student{}
	route, err := helpers.GetRoute(
		golang.RouteTeacherKeysGetStudentProfile,
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

	helpers.AddUserHeaders(jwtToken, req)

	res, err := helpers.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resStudent, err
	}

	defer helpers.CloseBody(res.Body)

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
