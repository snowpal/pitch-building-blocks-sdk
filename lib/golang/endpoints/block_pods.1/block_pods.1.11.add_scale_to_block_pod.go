package block_pods

import (
	"development/go/recipes/lib/golang"
	"development/go/recipes/lib/golang/helpers"
	"fmt"
	"net/http"
)

type AddScaleIdParam struct {
	KeyId   string
	BlockId string
	PodId   string
	ScaleId string
}

func AddScaleToBlockPod(jwtToken string, podParam AddScaleIdParam) error {
	client := &http.Client{}
	route, err := helpers.GetRoute(
		golang.RouteBlockPodsAddScaleToBlockPod,
		podParam.PodId,
		podParam.ScaleId,
		podParam.KeyId,
		podParam.BlockId,
	)
	if err != nil {
		fmt.Println(err)
		return err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodPatch, route, nil)
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
