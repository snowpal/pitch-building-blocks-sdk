package teacher_keys_1

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

func GetStudentsInABlock(jwtToken string, blockParam common.ResourceIdParam) ([]response.Student, error) {
	resStudents := response.Students{}
	route, err := helpers.GetRoute(building_blocks.RouteTeacherKeysGetStudentsInABlock, blockParam.BlockId, blockParam.KeyId)
	if err != nil {
		fmt.Println(err)
		return resStudents.Students, err
	}

	req, err := http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		fmt.Println(err)
		return resStudents.Students, err
	}

	helpers.AddUserHeaders(jwtToken, req)

	res, err := helpers.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resStudents.Students, err
	}

	defer helpers.CloseBody(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resStudents.Students, err
	}

	err = json.Unmarshal(body, &resStudents)
	if err != nil {
		fmt.Println(err)
		return resStudents.Students, err
	}
	return resStudents.Students, nil
}
