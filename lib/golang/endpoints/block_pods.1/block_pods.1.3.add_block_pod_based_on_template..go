package block_pods

import (
	"development/go/recipes/lib/golang"
	"development/go/recipes/lib/golang/helpers"
	"development/go/recipes/lib/golang/structs/common"
	"development/go/recipes/lib/golang/structs/request"
	"development/go/recipes/lib/golang/structs/response"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

type PodByTemplate struct {
	KeyId        string `json:"keyId"`
	BlockId      string `json:"blockId"`
	TemplateId   string `json:"templateId"`
	ExcludeTasks *bool  `json:"excludeTasks"`
}

func AddBlockPodBasedOnTemplate(jwtToken string, block common.SlimBlock, pod request.Pod, template request.Template, excludeTasks bool) (response.Pod, error) {
	podResponse := response.Pod{}

	requestBody, err := helpers.GetRequestBody(pod)
	if err != nil {
		fmt.Println(err)
		return podResponse, err
	}
	payload := strings.NewReader(requestBody)
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPost, helpers.GetRoute(golang.RouteBlockPodsAddBlockPodBasedOnTemplate, block.ID, block.Key.ID, template.ID, strconv.FormatBool(excludeTasks)), payload)

	if err != nil {
		fmt.Println(err)
		return podResponse, err
	}
	helpers.AddUserHeaders(jwtToken, req)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return podResponse, err
	}
	defer helpers.CloseBody(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return podResponse, err
	}

	err = json.Unmarshal(body, &podResponse)
	if err != nil {
		fmt.Println(err)
		return podResponse, err
	}
	return podResponse, err
}
