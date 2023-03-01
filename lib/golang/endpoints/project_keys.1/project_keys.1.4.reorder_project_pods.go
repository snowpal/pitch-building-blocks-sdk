package project_keys_1

import (
	"development/go/recipes/lib/golang/helpers"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main(jwtToken string) {

	url := "https://gateway.snowpal.com/blocks/:id/project-pods/reorder?keyId=63ceb29edb035900138d975d"
	method := "PATCH"

	payload := strings.NewReader(`{"sourceProjectListId":"source_project_list[project_list_id]","sourceProjectListPodIds":"source_project_list[pod_ids]","targetProjectListId":"target_project_list[project_list_id]","targetProjectListPodIds":"target_project_list[pod_ids]"}`)

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
