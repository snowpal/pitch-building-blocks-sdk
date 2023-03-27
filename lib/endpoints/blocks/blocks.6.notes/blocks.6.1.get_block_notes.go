package blocks

import (
	"development/go/recipes/lib"
	helpers2 "development/go/recipes/lib/helpers"
	"development/go/recipes/lib/structs/request"
	"development/go/recipes/lib/structs/response"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetBlockNotes(jwtToken string, noteParam request.NoteIdParam) ([]response.Note, error) {
	resNotes := response.Notes{}
	route, err := helpers2.GetRoute(
		lib.RouteBlocksGetBlockNotes,
		*noteParam.BlockId,
		noteParam.KeyId,
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

	helpers2.AddUserHeaders(jwtToken, req)

	res, err := helpers2.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resNotes.Notes, err
	}

	defer helpers2.CloseBody(res.Body)

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
