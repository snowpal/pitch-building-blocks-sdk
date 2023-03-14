package key_pods_1

import (
	"development/go/recipes/lib/golang"
	"development/go/recipes/lib/golang/helpers"
	"development/go/recipes/lib/golang/structs/request"
	"fmt"
	"net/http"
	"strconv"
)

func CopyKeyPod(jwtToken string, podParam request.CopyMovePodParam) error {
	client := &http.Client{}
	route, err := helpers.GetRoute(
		golang.RouteKeyPodsCopyKeyPod,
		podParam.PodId,
		podParam.KeyId,
		strconv.FormatBool(*podParam.AllTasks),
		strconv.FormatBool(*podParam.AllChecklists),
		podParam.TargetKeyId,
		*podParam.TargetBlockId,
	)
	if err != nil {
		fmt.Println(err)
		return err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodPost, route, nil)
	if err != nil {
		fmt.Println(err)
		return err
	}

	helpers.AddUserHeaders(jwtToken, req)

	_, err = client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
