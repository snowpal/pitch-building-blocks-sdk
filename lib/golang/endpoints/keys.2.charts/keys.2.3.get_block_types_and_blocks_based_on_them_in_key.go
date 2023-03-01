package keys_2

import (
	"development/go/recipes/lib/golang/helpers"
	"fmt"
	"io"
	"net/http"
)

func main(jwtToken string) {

	url := "https://gateway.snowpal.com/charts/keys/%7B%7Bkeys.custom.keyId1%7D%7D/block-types"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	helpers.AddUserHeaders(jwtToken, req)

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
