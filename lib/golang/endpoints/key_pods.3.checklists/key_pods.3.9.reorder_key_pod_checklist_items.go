package key_pods_3

import (
	"development/go/recipes/lib/golang/helpers"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main(jwtToken string) error {

	url := "pods/%s/checklists/%s/checklist-items/reorder?keyId=%s"
	method := "PATCH"

	payload := strings.NewReader(`{"checklistItemIds":"63e510314c721500376f9f2f,63d95d371610fe0014b7c4db"}`)

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
