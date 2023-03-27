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

func GetBlockAndPodsGradesForAStudentAsTeacher(
	jwtToken string,
	classroomParam request.ClassroomIdParam,
) (response.StudentGradeForBlockAndPod, error) {
	resStudentGrades := response.StudentGradeForBlockAndPod{}
	route, err := helpers2.GetRoute(
		lib.RouteTeacherKeysGetBlockAndPodsGradesForAStudentAsTeacher,
		classroomParam.ResourceIds.BlockId,
		classroomParam.StudentId,
		classroomParam.ResourceIds.KeyId,
	)
	if err != nil {
		fmt.Println(err)
		return resStudentGrades, err
	}

	req, err := http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		fmt.Println(err)
		return resStudentGrades, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	res, err := helpers2.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resStudentGrades, err
	}

	defer helpers2.CloseBody(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resStudentGrades, err
	}

	err = json.Unmarshal(body, &resStudentGrades)
	if err != nil {
		fmt.Println(err)
		return resStudentGrades, err
	}
	return resStudentGrades, nil
}
