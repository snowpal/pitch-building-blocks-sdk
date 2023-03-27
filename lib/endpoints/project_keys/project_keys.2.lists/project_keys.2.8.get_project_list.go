package projectKeys

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/snowpal/pitch-building-blocks-sdk/lib"
	helpers2 "github.com/snowpal/pitch-building-blocks-sdk/lib/helpers"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/request"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/response"
)

func GetProjectList(
	jwtToken string,
	projectListParam request.ProjectListIdParam,
) (response.ProjectList, error) {
	resProjectList := response.ProjectList{}
	route, err := helpers2.GetRoute(
		lib.RouteProjectKeysGetProjectList,
		projectListParam.BlockId,
		projectListParam.ProjectListId,
		projectListParam.KeyId,
	)
	if err != nil {
		fmt.Println(err)
		return resProjectList, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodGet, route, nil)
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
