package block_types

import (
	"development/go/recipes/lib/building-blocks"
	"development/go/recipes/lib/building-blocks/helpers"
	"fmt"
	"net/http"
)

func DeleteBlockType(jwtToken string, blockTypeId string) error {
	route, err := helpers.GetRoute(building_blocks.RouteBlockTypesDeleteBlockType, blockTypeId)
	if err != nil {
		fmt.Println(err)
		return err
	}

	req, err := http.NewRequest(http.MethodDelete, route, nil)
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