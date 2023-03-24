package keys_4

import (
	"development/go/recipes/lib/building-blocks"
	"development/go/recipes/lib/building-blocks/helpers"
	"development/go/recipes/lib/building-blocks/structs/request"
	"development/go/recipes/lib/building-blocks/structs/response"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func AddKeyNote(
	jwtToken string,
	reqBody request.NoteReqBody,
	commentParam request.NoteIdParam,
) (response.Note, error) {
	resNote := response.Note{}
	requestBody, err := helpers.GetRequestBody(reqBody)
	if err != nil {
		fmt.Println(err)
		return resNote, err
	}
	payload := strings.NewReader(requestBody)
	route, err := helpers.GetRoute(
		building_blocks.RouteKeysAddKeyNote,
		commentParam.KeyId,
	)
	req, err := http.NewRequest(http.MethodPost, route, payload)
	if err != nil {
		fmt.Println(err)
		return resNote, err
	}

	helpers.AddUserHeaders(jwtToken, req)

	res, err := helpers.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resNote, err
	}

	defer helpers.CloseBody(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resNote, err
	}

	err = json.Unmarshal(body, &resNote)
	if err != nil {
		fmt.Println(err)
		return resNote, err
	}
	return resNote, nil
}