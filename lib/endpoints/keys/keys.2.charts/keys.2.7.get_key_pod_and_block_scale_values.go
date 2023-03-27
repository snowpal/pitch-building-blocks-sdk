package keys

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

func GetKeyPodAndBlockScaleValues(jwtToken string, scaleParam request.ScaleIdParam) (response.ScaleValues, error) {
	resScaleValues := response.ScaleValues{}
	route, err := helpers2.GetRoute(
		lib.RouteKeysGetKeyPodAndBlockScaleValues,
		scaleParam.KeyId,
		scaleParam.ScaleId,
	)
	if err != nil {
		fmt.Println(err)
		return resScaleValues, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		fmt.Println(err)
		return resScaleValues, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers2.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resScaleValues, err
	}

	defer helpers2.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resScaleValues, err
	}

	err = json.Unmarshal(body, &resScaleValues)
	if err != nil {
		fmt.Println(err)
		return resScaleValues, err
	}
	return resScaleValues, nil
}
