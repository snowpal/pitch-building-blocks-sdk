package block_pods

import (
	"development/go/recipes/lib/golang"
	"development/go/recipes/lib/golang/helpers"
	"fmt"
	"io"
	"net/http"
)

func GetBlockPods(jwtToken string) {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf(golang.RouteGetBlockPods, "", "", ""), nil)

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
