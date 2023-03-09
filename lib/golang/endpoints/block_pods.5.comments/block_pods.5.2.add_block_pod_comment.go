package block_pods_5

import (
	"development/go/recipes/lib/golang/helpers"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main(jwtToken string) error {

	url := "block-pods/%s/comments?keyId=%s&blockId=%s"
	method := "POST"

	payload := strings.NewReader(`{"commentText":"comment[comment_text]","taggedUserIds":"comment[tagged_user_ids]"}`)

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
