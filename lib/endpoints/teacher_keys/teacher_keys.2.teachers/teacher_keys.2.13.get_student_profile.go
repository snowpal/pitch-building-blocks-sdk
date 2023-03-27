package teacherKeys

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/snowpal/pitch-building-blocks-sdk/lib"
	helpers2 "github.com/snowpal/pitch-building-blocks-sdk/lib/helpers"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/request"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/response"
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
