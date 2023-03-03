package blocks_1

import (
	"development/go/recipes/lib/golang"
	"development/go/recipes/lib/golang/helpers"
	"development/go/recipes/lib/golang/structs/response"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func AddBlock(jwtToken string) {
	block := response.SlimBlock{Name: "SlimBlock A"}

	blockBody, err := json.Marshal(block)
	if err != nil {
		return
	}
	payload := strings.NewReader(string(blockBody))
	client := &http.Client{}
	fmt.Println("TODO: Replace with Key ID")
	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf(helpers.GetRoute(golang.RouteBlocksAddBlock), ""), payload)

	if err != nil {
		fmt.Println(err)
		return
	}

	helpers.AddUserHeaders(jwtToken, req)
	res, _ := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(res.Body)

	body, _ := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
