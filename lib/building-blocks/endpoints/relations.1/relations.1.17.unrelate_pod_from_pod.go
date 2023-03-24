package relations

import (
	"development/go/recipes/lib/building-blocks"
	"development/go/recipes/lib/building-blocks/helpers"
	"development/go/recipes/lib/building-blocks/structs/request"
	"fmt"
	"net/http"
)

func unrelatePodFromPod(jwtToken string, route string) error {
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

func UnrelateKeyPodFromKeyPod(jwtToken string, relationParam request.PodToPodRelationParam) error {
	route, err := helpers.GetRoute(
		building_blocks.RouteRelationsUnrelatePodFromPod,
		relationParam.PodId,
		relationParam.SourceKeyId,
		relationParam.TargetPodId,
		relationParam.TargetKeyId,
	)
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = unrelatePodFromPod(jwtToken, route)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func UnrelateKeyPodFromBlockPod(jwtToken string, relationParam request.PodToBlockPodRelationParam) error {
	route, err := helpers.GetRoute(
		building_blocks.RouteRelationsUnrelatePodFromBlockPod,
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
	err = unrelatePodFromPod(jwtToken, route)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func UnrelateBlockPodFromBlockPod(jwtToken string, relationParam request.BlockPodToBlockPodRelationParam) error {
	route, err := helpers.GetRoute(
		building_blocks.RouteRelationsUnrelateBlockPodFromBlockPod,
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
	err = unrelatePodFromPod(jwtToken, route)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}