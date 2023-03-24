package teacher_keys_2

import (
	"development/go/recipes/lib/building-blocks"
	"development/go/recipes/lib/building-blocks/helpers"
	"development/go/recipes/lib/building-blocks/structs/common"
	"development/go/recipes/lib/building-blocks/structs/response"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetPodGradesForStudents(
	jwtToken string,
	podParam common.ResourceIdParam,
) (response.StudentGradeForBlockAndPod, error) {
	resStudentGradesForPod := response.StudentGradeForBlockAndPod{}
	route, err := helpers.GetRoute(
		building_blocks.RouteTeacherKeysGetPodGradesForStudents,
		podParam.PodId,
		podParam.KeyId,
		podParam.BlockId,
	)
	if err != nil {
		fmt.Println(err)
		return resStudentGradesForPod, err
	}

	req, err := http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		fmt.Println(err)
		return resStudentGradesForPod, err
	}

	helpers.AddUserHeaders(jwtToken, req)

	res, err := helpers.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resStudentGradesForPod, err
	}

	defer helpers.CloseBody(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resStudentGradesForPod, err
	}

	err = json.Unmarshal(body, &resStudentGradesForPod)
	if err != nil {
		fmt.Println(err)
		return resStudentGradesForPod, err
	}
	return resStudentGradesForPod, nil
}
