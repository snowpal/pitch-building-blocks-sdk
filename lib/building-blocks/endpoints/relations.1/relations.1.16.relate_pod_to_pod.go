package relations

import (
	"development/go/recipes/lib/building-blocks"
	"development/go/recipes/lib/building-blocks/helpers"
	"development/go/recipes/lib/building-blocks/structs/request"
	"fmt"
	"net/http"
)

func relatePodToPod(jwtToken string, route string) error {
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

func RelateKeyPodToKeyPod(jwtToken string, relationParam request.PodToPodRelationParam) error {
	route, err := helpers.GetRoute(
		building_blocks.RouteRelationsRelatePodToPod,
		relationParam.PodId,
		relationParam.SourceKeyId,
		relationParam.TargetPodId,
		relationParam.TargetKeyId,
	)
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = relatePodToPod(jwtToken, route)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func RelateKeyPodToBlockPod(jwtToken string, relationParam request.PodToBlockPodRelationParam) error {
	route, err := helpers.GetRoute(
		building_blocks.RouteRelationsRelatePodToBlockPod,
		relationParam.PodId,
		relationParam.SourceKeyId,
		relationParam.TargetPodId,
		relationParam.TargetKeyId,
		relationParam.TargetBlockId,
	)
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = relatePodToPod(jwtToken, route)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func RelateBlockPodToBlockPod(jwtToken string, relationParam request.BlockPodToBlockPodRelationParam) error {
	route, err := helpers.GetRoute(
		building_blocks.RouteRelationsRelateBlockPodToBlockPod,
		relationParam.PodId,
		relationParam.SourceKeyId,
		relationParam.SourceBlockId,
		relationParam.TargetPodId,
		relationParam.TargetKeyId,
		relationParam.TargetBlockId,
	)
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = relatePodToPod(jwtToken, route)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}