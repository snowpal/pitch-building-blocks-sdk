package blocks_1

import (
	"development/go/recipes/lib/golang"
	"development/go/recipes/lib/golang/structs"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetBlocks(jwtToken string) (structs.ResourceAttributes, error) {
	fmt.Println("TODO: Replace structs.ResourceAttributes with Blocks struct")
	var resourceAttrs structs.ResourceAttributes
	client := &http.Client{}
	
	fmt.Println("TODO: Fix arguments")
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf(golang.UrlGetBlocks, "", "", "", ""), nil)

	if err != nil {
		fmt.Println(err)
		return resourceAttrs, err
	}
	req.Header.Add("User-Authorization", jwtToken)
	req.Header.Add("x-api-key", golang.XApiKey)

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
