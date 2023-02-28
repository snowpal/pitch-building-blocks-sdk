package attributes_1

import (
	"development/go/recipes/lib/golang"
	"fmt"
	"io"
	"net/http"
)

func GetResourceAttributes(jwtToken string) {

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, golang.UrlGetAttributes, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("User-Authorization", jwtToken)
	req.Header.Add("x-api-key", golang.XApiKey)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
