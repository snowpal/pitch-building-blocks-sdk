package blocks_1

import (
	"development/go/recipes/lib/golang"
	"development/go/recipes/lib/golang/helpers"
	"development/go/recipes/lib/golang/structs/response"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetBlocks(jwtToken string) (response.ResourceAttributes, error) {
	fmt.Println("TODO: Replace structs.ResourceAttributes with Blocks struct")
	var resourceAttrs response.ResourceAttributes
	client := &http.Client{}

	fmt.Println("TODO: Fix arguments")
	req, err := http.NewRequest(http.MethodGet,
		fmt.Sprintf(helpers.GetRoute(golang.RouteBlocksGetBlocksInAKey), "", "", "", ""), nil)

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
	defer helpers.CloseBody(res.Body)

	body, _ := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resourceAttrs, err
	}

	err = json.Unmarshal(body, &resourceAttrs)
	if err != nil {
		return resourceAttrs, err
	}

	return resourceAttrs, nil
}
