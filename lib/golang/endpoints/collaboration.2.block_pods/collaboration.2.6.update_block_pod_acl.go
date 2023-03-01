package collaboration_2

import (
	"development/go/recipes/lib/golang/helpers"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main(jwtToken string) {

	url := "https://gateway.snowpal.com/block-pods/%7B%7BblockPods.podId%7D%7D/users/63cabea69e9c500014d4f9b7/acl?keyId=%7B%7Bkeys.custom.keyId1%7D%7D&blockId=%7B%7Bblocks.blockId%7D%7D"
	method := "PATCH"

	payload := strings.NewReader(`{"podAcl":"write"}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

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
