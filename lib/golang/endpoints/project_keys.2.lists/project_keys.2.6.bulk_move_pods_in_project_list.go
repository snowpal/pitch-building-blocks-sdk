package project_keys_2

import (
	"development/go/recipes/lib/golang/helpers"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main(jwtToken string) {

	url := "project-block-lists/%s/pods/move?blockId=%s&keyId=%s&targetKeyId=20tal56wic25724389v358p&targetBlockId=03bwn23nuw726786247d981k&targetProjectListId=17niw72bvd2481501237a125k&podIds=67mnw82huw218472984b398h,31pqs82niz262832984b754b"
	method := "PATCH"

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
