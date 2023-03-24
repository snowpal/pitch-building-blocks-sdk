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

func GetBlockGradesForStudents(
	jwtToken string,
	blockParam common.ResourceIdParam,
) (response.StudentGradeForBlockAndPod, error) {
	resStudentGradesForBlock := response.StudentGradeForBlockAndPod{}
	route, err := helpers.GetRoute(
		building_blocks.RouteTeacherKeysGetBlockGradesForStudents,
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

	helpers.AddUserHeaders(jwtToken, req)

	res, err := helpers.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resStudentGradesForBlock, err
	}

	defer helpers.CloseBody(res.Body)

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
