package blocks_pods_6

import (
	"development/go/recipes/lib/golang"
	"development/go/recipes/lib/golang/helpers"
	"development/go/recipes/lib/golang/structs/request"
	"development/go/recipes/lib/golang/structs/response"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetBlockPodNotes(jwtToken string, noteParam request.NoteIdParam) ([]response.Note, error) {
	resNotes := response.Notes{}
	client := &http.Client{}
	route, err := helpers.GetRoute(
		golang.RouteBlockPodsGetBlockPodNotes,
		*noteParam.PodId,
		noteParam.KeyId,
		*noteParam.BlockId,
	)
	if err != nil {
		fmt.Println(err)
		return resNotes.Notes, err
	}

	req, err := http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		fmt.Println(err)
		return resNotes.Notes, err
	}

	helpers.AddUserHeaders(jwtToken, req)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return resNotes.Notes, err
	}

	defer helpers.CloseBody(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resNotes.Notes, err
	}

	err = json.Unmarshal(body, &resNotes)
	if err != nil {
		fmt.Println(err)
		return resNotes.Notes, err
	}
	return resNotes.Notes, nil
}
