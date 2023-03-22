package relations

import (
	"development/go/recipes/lib/golang"
	"development/go/recipes/lib/golang/helpers"
	"development/go/recipes/lib/golang/structs/request"
	"fmt"
	"net/http"
)

func UnrelatePodToPod(jwtToken string, relationParam request.PodToPodRelationParam) error {
	route, err := helpers.GetRoute(
		golang.RouteRelationsUnrelatePodFromPod,
		relationParam.PodId,
		relationParam.SourceKeyId,
		*relationParam.SourceBlockId,
		relationParam.TargetPodId,
		relationParam.TargetKeyId,
		*relationParam.TargetBlockId,
	)
	if err != nil {
		fmt.Println(err)
		return err
	}

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
