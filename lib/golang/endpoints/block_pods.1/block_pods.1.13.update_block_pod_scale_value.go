package block_pods

import (
	"development/go/recipes/lib/golang/helpers"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Scale struct {
	ScaleValue string `json:"scaleValue"`
}

func UpdateBlockPodScaleValue(jwtToken string, scale Scale) error {
	url := "block-pods/%s/scale-value?keyId=%s&blockId=%s"
	requestBody, err := helpers.GetRequestBody(scale)
	if err != nil {
		return err
	}

	payload := strings.NewReader(requestBody)
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPatch, url, payload)
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

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println(string(body))
	return nil
}
