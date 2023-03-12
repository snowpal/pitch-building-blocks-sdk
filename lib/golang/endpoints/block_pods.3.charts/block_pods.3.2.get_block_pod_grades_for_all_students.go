package block_pods

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

type BlockPodGrade struct {
	ID       string             `json:"id"`
	Name     string             `json:"blockName"`
	Key      common.SlimKey     `json:"key"`
	Pod      common.SlimPod     `json:"pod"`
	Students []response.Student `json:"students"`
}

func GetBlockPodGradesForAllStudents(jwtToken string, podParam common.ResourceIdParam) (BlockPodGrade, error) {
	resBlockPodGrades := BlockPodGrade{}
	client := &http.Client{}
	route, err := helpers.GetRoute(
		golang.RouteBlockPodsGetBlockPodGradesForAllStudents,
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

	helpers.AddUserHeaders(jwtToken, req)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return resBlockPodGrades, err
	}

	defer helpers.CloseBody(res.Body)

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
