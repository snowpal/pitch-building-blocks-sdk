package keys_1

import (
	"development/go/recipes/lib/golang"
	"development/go/recipes/lib/golang/helpers"
	"development/go/recipes/lib/golang/structs/response"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetKeysFilteredByType(jwtToken string, keyType string) ([]response.Key, error) {
	resKeys := response.Keys{}
	client := &http.Client{}
	route, err := helpers.GetRoute(golang.RouteKeysGetKeysFilteredByType, keyType)
	if err != nil {
		fmt.Println(err)
		return resKeys.Keys, err
	}
	req, err := http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		fmt.Println(err)
		return resKeys.Keys, err
	}

	helpers.AddUserHeaders(jwtToken, req)

	res, _ := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return resKeys.Keys, err
	}

	defer helpers.CloseBody(res.Body)

	body, _ := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resKeys.Keys, err
	}

	err = json.Unmarshal(body, &resKeys)
	if err != nil {
		fmt.Println(err)
		return resKeys.Keys, err
	}
	return resKeys.Keys, nil
}
