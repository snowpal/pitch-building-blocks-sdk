package attributes_1

import (
	"development/go/recipes/lib/golang"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func UpdateKeyAttrs(jwtToken string, keyId string) {
	fmt.Println("TODO: Replace with struct")
	payload := strings.NewReader(`{"showAttribute":"show_attribute","attributeNames":"view[attribute_names]"}`)
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf(golang.UrlUpdateKeyAttributes, keyId), payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("User-Authorization", jwtToken)
	req.Header.Add("x-api-key", golang.XApiKey)
	req.Header.Add("Content-Type", "application/json")

	res, _ := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(res.Body)

	body, _ := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
