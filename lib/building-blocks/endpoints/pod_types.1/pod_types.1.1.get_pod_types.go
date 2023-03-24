package pod_types_1

import (
	"development/go/recipes/lib/building-blocks"
	"development/go/recipes/lib/building-blocks/helpers"
	"development/go/recipes/lib/building-blocks/structs/response"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func GetPodTypes(jwtToken string, includeCounts bool) ([]response.PodType, error) {
	resPodTypes := response.PodTypes{}
	route, err := helpers.GetRoute(building_blocks.RoutePodTypesGetPodTypes, strconv.FormatBool(includeCounts))
	if err != nil {
		fmt.Println(err)
		return resPodTypes.PodTypes, err
	}
	req, err := http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		fmt.Println(err)
		return resPodTypes.PodTypes, err
	}

	helpers.AddUserHeaders(jwtToken, req)

	res, err := helpers.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resPodTypes.PodTypes, err
	}

	defer helpers.CloseBody(res.Body)

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