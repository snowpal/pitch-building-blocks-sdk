package collaboration_1

import (
	"development/go/recipes/lib/golang"
	"development/go/recipes/lib/golang/helpers"
	"development/go/recipes/lib/golang/structs/common"
	"development/go/recipes/lib/golang/structs/response"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type ShareBlockWithPodsReqBody struct {
	Acl    string   `json:"blockAcl"`
	PodIds []string `json:"podIds"`
}

func ShareBlockWithUserWithPods(
	jwtToken string,
	reqBody ShareBlockWithPodsReqBody,
	blockAclParam common.AclParam,
) (response.Block, error) {
	resBlock := response.Block{}
	requestBody, err := helpers.GetRequestBody(reqBody)
	if err != nil {
		fmt.Println(err)
		return resBlock, err
	}
	payload := strings.NewReader(requestBody)
	route, err := helpers.GetRoute(
		golang.RouteCollaborationShareBlockWithCollaboratorAlongWithPods,
		*blockAclParam.BlockId,
		blockAclParam.UserId,
		blockAclParam.KeyId,
	)
	req, err := http.NewRequest(http.MethodPatch, route, payload)
	if err != nil {
		fmt.Println(err)
		return resBlock, err
	}

	helpers.AddUserHeaders(jwtToken, req)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return resBlock, err
	}

	defer helpers.CloseBody(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resBlock, err
	}

	err = json.Unmarshal(body, &resBlock)
	if err != nil {
		fmt.Println(err)
		return resBlock, err
	}
	return resBlock, nil
}
