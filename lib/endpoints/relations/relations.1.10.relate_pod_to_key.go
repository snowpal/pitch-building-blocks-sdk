package relations

import (
	"development/go/recipes/lib"
	"development/go/recipes/lib/helpers"
	"development/go/recipes/lib/structs/request"
	"fmt"
	"net/http"
)

func relateKeyToPod(jwtToken string, route string) error {
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

func RelateKeyToBlockPod(jwtToken string, relationParam request.KeyToPodRelationParam) error {
	route, err := helpers.GetRoute(
		lib.RouteRelationsRelatePodToKey,
		relationParam.KeyId,
		relationParam.TargetPodId,
		relationParam.TargetKeyId,
	)
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = relateKeyToPod(jwtToken, route)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func RelateKeyToPod(jwtToken string, relationParam request.KeyToBlockPodRelationParam) error {
	route, err := helpers.GetRoute(
		lib.RouteRelationsRelateBlockPodToKey,
		relationParam.KeyId,
		relationParam.TargetPodId,
		relationParam.TargetKeyId,
		relationParam.TargetBlockId,
	)
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = relateKeyToPod(jwtToken, route)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
