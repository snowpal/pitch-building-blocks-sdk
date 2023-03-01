package keys_3

import (
	"development/go/recipes/lib/golang/helpers"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main(jwtToken string) {

	url := "https://gateway.snowpal.com/keys/:id/checklists/:checklist-id/checklist-items/:checklist-item-id"
	method := "PATCH"

	payload := strings.NewReader(`{"checklistItemText":"checklist_item[checklist_item_title]","completed":"checklist_item[completed]"}`)

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
