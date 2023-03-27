package podTypes

import (
	"development/go/recipes/lib"
	helpers2 "development/go/recipes/lib/helpers"
	"development/go/recipes/lib/structs/response"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
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
