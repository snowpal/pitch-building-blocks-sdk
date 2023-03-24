package teacher_keys_2

import (
	"development/go/recipes/lib/building-blocks"
	"development/go/recipes/lib/building-blocks/helpers"
	"development/go/recipes/lib/building-blocks/structs/request"
	"development/go/recipes/lib/building-blocks/structs/response"
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
	route, err := helpers.GetRoute(
		building_blocks.RouteTeacherKeysGetBlockAndPodsGradesForAStudentAsTeacher,
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

	helpers.AddUserHeaders(jwtToken, req)

	res, err := helpers.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resStudentGrades, err
	}

	defer helpers.CloseBody(res.Body)

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
