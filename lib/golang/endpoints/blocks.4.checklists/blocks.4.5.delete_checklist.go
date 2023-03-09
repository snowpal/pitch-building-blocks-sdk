package blocks_4

import (
	"development/go/recipes/lib/golang/helpers"
	"fmt"
	"io"
	"net/http"
)

func main(jwtToken string) error {

	url := "blocks/%s/checklists/%s?keyId=%s"
	method := "DELETE"

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
