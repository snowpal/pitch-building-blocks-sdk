package project_keys_2

import (
	"development/go/recipes/lib/golang/helpers"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main(jwtToken string) {

	url := "https://gateway.snowpal.com/project-block-lists/%7B%7BprojectLists.projectListId%7D%7D/pods/copy-all?keyId=%7B%7Bkeys.project.keyId1%7D%7D&targetKeyId=%7B%7Bkeys.project.keyId1%7D%7D&targetBlockId=%7B%7Bblocks.projectBlockId%7D%7D&targetProjectListId=%7B%7BprojectLists.projectListId%7D%7D&allPods=true&allTasks=false&blockId=%7B%7Bblocks.projectBlockId%7D%7D"
	method := "POST"

	payload := strings.NewReader(``)

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
