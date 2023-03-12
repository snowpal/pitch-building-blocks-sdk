package collaboration_2

import (
	"development/go/recipes/lib/golang/helpers"
	"fmt"
	"io"
	"net/http"
)

func main(jwtToken string) error {

	url := "search/block-pods/%s/shareable/users?keyId=%s&blockId=%s&token=%s"
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
	defer helpers.CloseBody(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return err
	}
}
