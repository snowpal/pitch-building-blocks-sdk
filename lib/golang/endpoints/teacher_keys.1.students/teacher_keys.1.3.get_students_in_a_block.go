package teacher_keys_1

import (
	"development/go/recipes/lib/golang"
	"development/go/recipes/lib/golang/helpers"
	"development/go/recipes/lib/golang/structs/common"
	"development/go/recipes/lib/golang/structs/response"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetStudentsInABlock(jwtToken string, blockParam common.ResourceIdParam) ([]response.Student, error) {
	resStudents := response.Students{}
	route, err := helpers.GetRoute(golang.RouteTeacherKeysGetStudentsInABlock, blockParam.BlockId, blockParam.KeyId)
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

	res, err := client.Do(req)
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
