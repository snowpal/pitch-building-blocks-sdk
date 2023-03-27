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

func GetPodGradesForStudents(
	jwtToken string,
	podParam common.ResourceIdParam,
) (response.StudentGradeForBlockAndPod, error) {
	resStudentGradesForPod := response.StudentGradeForBlockAndPod{}
	route, err := helpers2.GetRoute(
		lib.RouteTeacherKeysGetPodGradesForStudents,
		podParam.PodId,
		podParam.KeyId,
		podParam.BlockId,
	)
	if err != nil {
		fmt.Println(err)
		return resStudentGradesForPod, err
	}

	req, err := http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		fmt.Println(err)
		return resStudentGradesForPod, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	res, err := helpers2.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resStudentGradesForPod, err
	}

	defer helpers2.CloseBody(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resStudentGradesForPod, err
	}

	err = json.Unmarshal(body, &resStudentGradesForPod)
	if err != nil {
		fmt.Println(err)
		return resStudentGradesForPod, err
	}
	return resStudentGradesForPod, nil
}
