package project_keys_1

import (
	"development/go/recipes/lib/golang"
	"development/go/recipes/lib/golang/helpers"
	"development/go/recipes/lib/golang/structs/request"
	"fmt"
	"net/http"
	"strconv"
)

func CopyProjectBlock(jwtToken string, blockParam request.CopyMoveBlockParam) error {
	route, err := helpers.GetRoute(
		golang.RouteProjectKeysCopyProjectBlock,
		blockParam.BlockId,
		blockParam.KeyId,
		blockParam.TargetKeyId,
		strconv.FormatBool(*blockParam.AllPods),
		strconv.FormatBool(*blockParam.AllTasks),
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

	res, err := helpers.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
