package project_keys_1

import (
	"development/go/recipes/lib/golang/helpers"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main(jwtToken string) error {

	url := "blocks/%s/project-pods/reorder?keyId=%s"
	method := "PATCH"

	payload := strings.NewReader(`{"sourceProjectListId":"source_project_list[project_list_id]","sourceProjectListPodIds":"source_project_list[pod_ids]","targetProjectListId":"target_project_list[project_list_id]","targetProjectListPodIds":"target_project_list[pod_ids]"}`)

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
