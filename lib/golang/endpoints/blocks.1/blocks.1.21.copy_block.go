package blocks_1

import (
	"development/go/recipes/lib/golang"
	"development/go/recipes/lib/golang/helpers"
	"development/go/recipes/lib/golang/structs/request"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func CopyBlock(jwtToken string, blockParam request.CopyMoveBlockParam) error {
	route, err := helpers.GetRoute(
		golang.RouteBlocksCopyBlock,
		blockParam.BlockId,
		blockParam.KeyId,
		blockParam.TargetKeyId,
		strings.Join(*blockParam.PodIds, ","),
		strconv.FormatBool(*blockParam.AllPods),
		strconv.FormatBool(*blockParam.AllTasks),
		strconv.FormatBool(*blockParam.AllChecklists),
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

	_, err = helpers.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
