package projectKeys

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/snowpal/pitch-building-blocks-sdk/lib"
	helpers2 "github.com/snowpal/pitch-building-blocks-sdk/lib/helpers"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/request"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/response"
)

type ProjectPodByTemplateParam struct {
	ProjectListId string
	PodParam      request.PodByTemplateParam
}

func AddBlockPodBasedOnTemplate(
	jwtToken string,
	reqBody request.AddPodReqBody,
	projectPodParam ProjectPodByTemplateParam,
) (response.Pod, error) {
	resPod := response.Pod{}
	requestBody, err := helpers2.GetRequestBody(reqBody)
	if err != nil {
		fmt.Println(err)
		return resPod, err
	}
	payload := strings.NewReader(requestBody)

	var route string
	route, err = helpers2.GetRoute(
		lib.RouteProjectKeysAddProjectPodBasedOnTemplate,
		*projectPodParam.PodParam.BlockId,
		projectPodParam.PodParam.KeyId,
		projectPodParam.ProjectListId,
		projectPodParam.PodParam.TemplateId,
		strconv.FormatBool(projectPodParam.PodParam.ExcludeTasks),
	)
	if err != nil {
		fmt.Println(err)
		return resPod, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodPost, route, payload)
	if err != nil {
		fmt.Println(err)
		return resPod, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers2.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resPod, err
	}

	defer helpers2.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resPod, err
	}

	err = json.Unmarshal(body, &resPod)
	if err != nil {
		fmt.Println(err)
		return resPod, err
	}
	return resPod, err
}
