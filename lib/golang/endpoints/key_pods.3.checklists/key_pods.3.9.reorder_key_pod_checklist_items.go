package key_pods_3

import (
	"development/go/recipes/lib/golang/helpers"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main(jwtToken string) {

	url := "https://gateway.snowpal.com/pods/%7B%7Bpods.podId%7D%7D/checklists/63d26a77a82d3900145ffb8a/checklist-items/reorder?keyId=%7B%7Bkeys.custom.keyId1%7D%7D"
	method := "PATCH"

	payload := strings.NewReader(`{"checklistItemIds":"63e510314c721500376f9f2f,63d95d371610fe0014b7c4db"}`)

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
