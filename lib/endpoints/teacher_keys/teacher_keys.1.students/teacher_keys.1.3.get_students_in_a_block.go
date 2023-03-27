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

func GetStudentsInABlock(jwtToken string, blockParam common.ResourceIdParam) ([]response.Student, error) {
	resStudents := response.Students{}
	route, err := helpers2.GetRoute(lib.RouteTeacherKeysGetStudentsInABlock, blockParam.BlockId, blockParam.KeyId)
	if err != nil {
		fmt.Println(err)
		return resStudents.Students, err
	}

	req, err := http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		fmt.Println(err)
		return resStudents.Students, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	res, err := helpers2.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resStudents.Students, err
	}

	defer helpers2.CloseBody(res.Body)

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
