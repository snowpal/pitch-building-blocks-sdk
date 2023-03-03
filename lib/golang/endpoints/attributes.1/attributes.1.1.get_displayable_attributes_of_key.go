package attributes_1

import (
	"development/go/recipes/lib/golang"
	"development/go/recipes/lib/golang/helpers"
	"development/go/recipes/lib/golang/structs/response"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetResourceAttrs(jwtToken string) (response.ResourceAttributes, error) {
	var resourceAttrs response.ResourceAttributes

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, golang.RouteAttributesGetDisplayableAttributesOfKey, nil)

	if err != nil {
		fmt.Println(err)
		return resourceAttrs, err
	}
	helpers.AddUserHeaders(jwtToken, req)

	res, _ := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return resourceAttrs, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(res.Body)

	body, _ := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resourceAttrs, err
	}
	fmt.Println(string(body))

	err = json.Unmarshal(body, &resourceAttrs)
	if err != nil {
		return resourceAttrs, err
	}

	return resourceAttrs, nil
}
