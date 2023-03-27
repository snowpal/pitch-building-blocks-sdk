package blockPods

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/snowpal/pitch-building-blocks-sdk/lib"
	helpers2 "github.com/snowpal/pitch-building-blocks-sdk/lib/helpers"
	common2 "github.com/snowpal/pitch-building-blocks-sdk/lib/structs/common"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/response"
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
