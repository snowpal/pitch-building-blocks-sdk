package podTypes

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/snowpal/pitch-building-blocks-sdk/lib"
	helpers2 "github.com/snowpal/pitch-building-blocks-sdk/lib/helpers"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/response"
)

func GetPodTypes(jwtToken string, includeCounts bool) ([]response.PodType, error) {
	resPodTypes := response.PodTypes{}
	route, err := helpers2.GetRoute(lib.RoutePodTypesGetPodTypes, strconv.FormatBool(includeCounts))
	if err != nil {
		fmt.Println(err)
		return resPodTypes.PodTypes, err
	}
	req, err := http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		fmt.Println(err)
		return resPodTypes.PodTypes, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	res, err := helpers2.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resPodTypes.PodTypes, err
	}

	defer helpers2.CloseBody(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resPodTypes.PodTypes, err
	}

	err = json.Unmarshal(body, &resPodTypes)
	if err != nil {
		fmt.Println(err)
		return resPodTypes.PodTypes, err
	}
	return resPodTypes.PodTypes, nil
}
