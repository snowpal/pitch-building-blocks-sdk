package relations

import (
	"fmt"
	"net/http"

	"github.com/snowpal/pitch-building-blocks-sdk/lib"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/helpers"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/request"
)

func unrelateBlockToPod(jwtToken string, route string) error {
	req, err := http.NewRequest(http.MethodPatch, route, nil)
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

func UnrelateBlockFromKeyPod(jwtToken string, relationParam request.BlockToPodRelationParam) error {
	route, err := helpers.GetRoute(
		lib.RouteRelationsUnrelatePodFromBlock,
		relationParam.BlockId,
		relationParam.TargetPodId,
		relationParam.TargetKeyId,
	)
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = unrelateBlockToPod(jwtToken, route)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func UnrelateBlockFromBlockPod(jwtToken string, relationParam request.BlockToBlockPodRelationParam) error {
	route, err := helpers.GetRoute(
		lib.RouteRelationsUnrelateBlockPodFromBlock,
		relationParam.BlockId,
		relationParam.TargetPodId,
		relationParam.TargetKeyId,
		relationParam.TargetBlockId,
	)
	if err != nil {
		fmt.Println(err)
		return err
	}

	err = unrelateBlockToPod(jwtToken, route)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
