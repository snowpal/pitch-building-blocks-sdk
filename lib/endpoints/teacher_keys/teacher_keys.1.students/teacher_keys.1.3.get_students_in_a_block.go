package teacherKeys

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/snowpal/pitch-building-blocks-sdk/lib"
	helpers2 "github.com/snowpal/pitch-building-blocks-sdk/lib/helpers"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/common"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/response"
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
