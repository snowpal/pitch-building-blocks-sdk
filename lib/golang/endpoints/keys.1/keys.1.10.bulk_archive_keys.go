package keys_1

import (
	"development/go/recipes/lib/golang"
	"development/go/recipes/lib/golang/helpers"
	"fmt"
	"net/http"
	"strings"
)

func BulkArchiveKeys(jwtToken string, keyIds []string) error {
	client := &http.Client{}
	route, err := helpers.GetRoute(golang.RouteKeysBulkArchiveKeys, strings.Join(keyIds, ","))
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

	_, err = client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}