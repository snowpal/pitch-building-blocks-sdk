package project_keys_2

import (
	"development/go/recipes/lib/golang/helpers"
	"fmt"
	"io"
	"net/http"
)

func main(jwtToken string) error {

	url := "blocks/%s/project-block-lists/%s?keyId=%s"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPatch, helpers.GetRoute(golang), nil)

	if err != nil {
		fmt.Println(err)
		return err
	}
	helpers.AddUserHeaders(jwtToken, req)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err
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
		return err
	}
	fmt.Println(string(body))
}
