package projectKeys

import (
	"development/go/recipes/lib"
	helpers2 "development/go/recipes/lib/helpers"
	"development/go/recipes/lib/structs/common"
	"development/go/recipes/lib/structs/response"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type ReorderProjectListsReqBody struct {
	ProjectListIds string `json:"projectListIds"`
}

func ReorderProjectLists(
	jwtToken string,
	reqBody ReorderProjectListsReqBody,
	projectListParam common.ResourceIdParam,
) ([]response.ProjectList, error) {
	resProjectLists := response.ProjectLists{}
	requestBody, err := helpers2.GetRequestBody(reqBody)
	if err != nil {
		fmt.Println(err)
		return resProjectLists.ProjectLists, err
	}
	payload := strings.NewReader(requestBody)
	route, err := helpers2.GetRoute(
		lib.RouteProjectKeysReorderProjectList,
		projectListParam.BlockId,
		projectListParam.KeyId,
	)
	if err != nil {
		fmt.Println(err)
		return resProjectLists.ProjectLists, err
	}

	req, err := http.NewRequest(http.MethodPatch, route, payload)
	if err != nil {
		fmt.Println(err)
		return resProjectLists.ProjectLists, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	res, err := helpers2.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resProjectLists.ProjectLists, err
	}

	defer helpers2.CloseBody(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resProjectLists.ProjectLists, err
	}

	err = json.Unmarshal(body, &resProjectLists)
	if err != nil {
		fmt.Println(err)
		return resProjectLists.ProjectLists, err
	}
	return resProjectLists.ProjectLists, nil
}
