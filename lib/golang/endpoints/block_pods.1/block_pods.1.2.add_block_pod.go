package block_pods

import (
	"development/go/recipes/lib/golang"
	"development/go/recipes/lib/golang/helpers"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func AddBlockPod(jwtToken string) {
	fmt.Println("TODO: Replace with struct")
	payload := strings.NewReader(`{"podName":"new block pod name 1 assessment[assessment_name]"}`)
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPost, helpers.GetRoute(fmt.Sprintf(golang.RouteBlockPodsAddBlockPod, "", "")), payload)

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
