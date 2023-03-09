package attributes_1

import (
	"development/go/recipes/lib/golang"
	"development/go/recipes/lib/golang/helpers"
	"development/go/recipes/lib/golang/structs/common"
	"development/go/recipes/lib/golang/structs/request"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func UpdateBlockPodAttrs(jwtToken string, pod common.SlimPod, attribute request.ResourceAttribute) error {
	requestBody, err := helpers.GetRequestBody(attribute)
	if err != nil {
		return err
	}

	payload := strings.NewReader(requestBody)
	client := &http.Client{}
	req, err := http.NewRequest(
		http.MethodGet,
		helpers.GetRoute(golang.RouteAttributesUpdateBlockPodDisplayAttributes, pod.ID, pod.Key.ID, pod.Block.ID),
		payload)

	if err != nil {
		fmt.Println(err)
		return err
	}
	helpers.AddUserHeaders(jwtToken, req)

	res, _ := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(res.Body)

	body, _ := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(string(body))
	return nil
}
