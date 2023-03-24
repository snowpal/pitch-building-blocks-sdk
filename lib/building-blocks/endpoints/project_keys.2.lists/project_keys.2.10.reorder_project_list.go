package project_keys_2

import (
	"development/go/recipes/lib/building-blocks"
	"development/go/recipes/lib/building-blocks/helpers"
	"development/go/recipes/lib/building-blocks/structs/common"
	"development/go/recipes/lib/building-blocks/structs/response"
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
	requestBody, err := helpers.GetRequestBody(reqBody)
	if err != nil {
		fmt.Println(err)
		return resProjectLists.ProjectLists, err
	}
	payload := strings.NewReader(requestBody)
	route, err := helpers.GetRoute(
		building_blocks.RouteProjectKeysReorderProjectList,
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

	helpers.AddUserHeaders(jwtToken, req)

	res, err := helpers.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resProjectLists.ProjectLists, err
	}

	defer helpers.CloseBody(res.Body)

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
