package blocks_pods_2

import (
	"development/go/recipes/lib/golang/helpers"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main(jwtToken string) error {

	url := "block-pods/%s/attachments?keyId=%s&blockId=%s"
	method := "PATCH"

	payload := strings.NewReader(`{"files":[{"fileName":"file_name","fileURL":"file_url"}]}`)

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPatch, helpers.GetRoute(golang), payload)

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
