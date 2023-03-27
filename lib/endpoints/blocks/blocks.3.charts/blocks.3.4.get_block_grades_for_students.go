package blocks

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

type BlockGradesForStudents struct {
	Block response.BlockGrade `json:"block"`
	Pods  []response.PodGrade `json:"pods"`
}

func GetBlockGradesForStudents(jwtToken string, gradeParam common.ResourceIdParam) (BlockGradesForStudents, error) {
	resBlockGrades := BlockGradesForStudents{}
	route, err := helpers2.GetRoute(
		lib.RouteBlocksGetBlockGradesForStudents,
		gradeParam.BlockId,
		gradeParam.KeyId,
	)
	if err != nil {
		fmt.Println(err)
		return resBlockGrades, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		fmt.Println(err)
		return resBlockGrades, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers2.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resBlockGrades, err
	}

	defer helpers2.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resBlockGrades, err
	}

	err = json.Unmarshal(body, &resBlockGrades)
	if err != nil {
		fmt.Println(err)
		return resBlockGrades, err
	}
	return resBlockGrades, nil
}
