package keys_1

import (
	"development/go/recipes/lib/building-blocks"
	"development/go/recipes/lib/building-blocks/helpers"
	"fmt"
	"net/http"
	"strings"
)

func BulkArchiveKeys(jwtToken string, keyIds []string) error {
	route, err := helpers.GetRoute(building_blocks.RouteKeysBulkArchiveKeys, strings.Join(keyIds, ","))
	if err != nil {
		fmt.Println(err)
		return err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodPatch, route, nil)
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
