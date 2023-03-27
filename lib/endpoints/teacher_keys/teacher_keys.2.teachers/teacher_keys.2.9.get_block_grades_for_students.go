package teacherKeys

import (
	"development/go/recipes/lib"
	helpers2 "development/go/recipes/lib/helpers"
	"development/go/recipes/lib/structs/common"
	"development/go/recipes/lib/structs/response"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetBlockGradesForStudents(
	jwtToken string,
	blockParam common.ResourceIdParam,
) (response.StudentGradeForBlockAndPod, error) {
	resStudentGradesForBlock := response.StudentGradeForBlockAndPod{}
	route, err := helpers2.GetRoute(
		lib.RouteTeacherKeysGetBlockGradesForStudents,
		blockParam.BlockId,
		blockParam.KeyId,
	)
	if err != nil {
		fmt.Println(err)
		return resStudentGradesForBlock, err
	}

	req, err := http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		fmt.Println(err)
		return resStudentGradesForBlock, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	res, err := helpers2.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resStudentGradesForBlock, err
	}

	defer helpers2.CloseBody(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resStudentGradesForBlock, err
	}

	err = json.Unmarshal(body, &resStudentGradesForBlock)
	if err != nil {
		fmt.Println(err)
		return resStudentGradesForBlock, err
	}
	return resStudentGradesForBlock, nil
}
