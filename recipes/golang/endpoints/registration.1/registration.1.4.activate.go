package registration_1

import (
	"development/go/recipes/endpoints"
	"fmt"
	"io"
	"net/http"
)

func Activate() {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPatch, endpoints.Url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("x-api-key", endpoints.ApiKey)

	res, _ := client.Do(req)
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

	body, _ := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
