package relations

import (
	"development/go/recipes/lib/golang"
	"development/go/recipes/lib/golang/helpers"
	"development/go/recipes/lib/golang/structs/request"
	"fmt"
	"net/http"
)

func relateBlockToPod(jwtToken string, route string) error {
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

func RelateBlockToKeyPod(jwtToken string, relationParam request.BlockToPodRelationParam) error {
	route, err := helpers.GetRoute(
		golang.RouteRelationsRelatePodToBlock,
		relationParam.BlockId,
		relationParam.TargetPodId,
		relationParam.TargetKeyId,
	)
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = relateBlockToPod(jwtToken, route)
	if err != nil {
		return err
	}
	return nil
}

func RelateBlockToBlockPod(jwtToken string, relationParam request.BlockToBlockPodRelationParam) error {
	route, err := helpers.GetRoute(
		golang.RouteRelationsRelateBlockPodToBlock,
		relationParam.BlockId,
		relationParam.TargetPodId,
		relationParam.TargetKeyId,
		relationParam.TargetBlockId,
	)
	if err != nil {
		fmt.Println(err)
		return err
	}

	if err != nil {
		fmt.Println(err)
		return err
	}
	err = relateBlockToPod(jwtToken, route)
	if err != nil {
		return err
	}
	return nil
}
