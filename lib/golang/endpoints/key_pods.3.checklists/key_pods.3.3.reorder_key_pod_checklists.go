package key_pods_3

import (
	"development/go/recipes/lib/golang/helpers"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main(jwtToken string) {

	url := "%7B%7BbaseUrl%7D%7D/pods/%7B%7Bpods.podId%7D%7D/checklists/reorder?keyId=%7B%7Bkeys.custom.keyId1%7D%7D"
	method := "PATCH"

	payload := strings.NewReader(`{"checklistIds":"63d26cfea82d3900135ffb8a,63d26a77a82d3900145ffb8a"}`)

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
