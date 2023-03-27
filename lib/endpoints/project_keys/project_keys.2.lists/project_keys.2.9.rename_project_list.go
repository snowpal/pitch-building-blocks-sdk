package projectKeys

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/snowpal/pitch-building-blocks-sdk/lib"
	helpers2 "github.com/snowpal/pitch-building-blocks-sdk/lib/helpers"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/request"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/response"
)

func RenameProjectList(
	jwtToken string,
	reqBody request.AddProjectListReqBody,
	projectListParam request.ProjectListIdParam,
) (response.ProjectList, error) {
	resProjectList := response.ProjectList{}
	requestBody, err := helpers2.GetRequestBody(reqBody)
	if err != nil {
		fmt.Println(err)
		return resProjectList, err
	}
	payload := strings.NewReader(requestBody)
	route, err := helpers2.GetRoute(
		lib.RouteProjectKeysRenameProjectList,
		projectListParam.BlockId,
		projectListParam.ProjectListId,
		projectListParam.KeyId,
	)
	if err != nil {
		fmt.Println(err)
		return resProjectList, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodPatch, route, payload)
	if err != nil {
		fmt.Println(err)
		return resProjectList, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers2.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resProjectList, err
	}

	defer helpers2.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resProjectList, err
	}

	err = json.Unmarshal(body, &resProjectList)
	if err != nil {
		fmt.Println(err)
		return resProjectList, err
	}
	return resProjectList, nil
}
