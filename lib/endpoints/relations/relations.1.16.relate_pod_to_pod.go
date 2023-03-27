package relations

import (
	"fmt"
	"net/http"

	"github.com/snowpal/pitch-building-blocks-sdk/lib"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/helpers"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/request"
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
		lib.RouteRelationsRelatePodToPod,
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
		lib.RouteRelationsRelatePodToBlockPod,
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
		lib.RouteRelationsRelateBlockPodToBlockPod,
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
