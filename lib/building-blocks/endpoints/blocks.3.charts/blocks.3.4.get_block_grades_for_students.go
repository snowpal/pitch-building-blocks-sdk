package blocks_3

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

type BlockGradesForStudents struct {
	Block response.BlockGrade `json:"block"`
	Pods  []response.PodGrade `json:"pods"`
}

func GetBlockGradesForStudents(jwtToken string, gradeParam common.ResourceIdParam) (BlockGradesForStudents, error) {
	resBlockGrades := BlockGradesForStudents{}
	route, err := helpers.GetRoute(
		building_blocks.RouteBlocksGetBlockGradesForStudents,
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

	helpers.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resBlockGrades, err
	}

	defer helpers.CloseBody(res.Body)

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
