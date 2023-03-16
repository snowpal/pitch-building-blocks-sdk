package relations

import (
	"development/go/recipes/lib/golang"
	"development/go/recipes/lib/golang/helpers"
	"development/go/recipes/lib/golang/structs/request"
	"fmt"
	"net/http"
)

func RelatePodToPod(jwtToken string, relationParam request.PodToPodRelationParam) error {
	client := &http.Client{}
	route, err := helpers.GetRoute(
		golang.RouteRelationsRelatePodToPod,
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

	_, err = client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
