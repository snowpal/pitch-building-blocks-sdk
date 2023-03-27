package blockPods

import (
	"development/go/recipes/lib"
	helpers2 "development/go/recipes/lib/helpers"
	common2 "development/go/recipes/lib/structs/common"
	"development/go/recipes/lib/structs/response"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type BlockPodGradesForStudents struct {
	ID       string             `json:"id"`
	Name     string             `json:"blockName"`
	Key      common2.SlimKey    `json:"key"`
	Pod      common2.SlimPod    `json:"pod"`
	Students []response.Student `json:"students"`
}

func GetBlockPodGradesForStudents(jwtToken string, podParam common2.ResourceIdParam) (BlockPodGradesForStudents, error) {
	resBlockPodGrades := BlockPodGradesForStudents{}
	route, err := helpers2.GetRoute(
		lib.RouteBlockPodsGetBlockPodGradesForStudents,
		podParam.PodId,
		podParam.KeyId,
		podParam.BlockId,
	)
	if err != nil {
		fmt.Println(err)
		return resBlockPodGrades, err
	}

	req, err := http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		fmt.Println(err)
		return resBlockPodGrades, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	res, err := helpers2.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resBlockPodGrades, err
	}

	defer helpers2.CloseBody(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resBlockPodGrades, err
	}

	err = json.Unmarshal(body, &resBlockPodGrades)
	if err != nil {
		fmt.Println(err)
		return resBlockPodGrades, err
	}
	return resBlockPodGrades, nil
}
