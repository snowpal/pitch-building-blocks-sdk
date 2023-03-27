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
	"strings"
)

func AddBlockNote(
	jwtToken string,
	reqBody request.NoteReqBody,
	commentParam request.NoteIdParam,
) (response.Note, error) {
	resNote := response.Note{}
	requestBody, err := helpers2.GetRequestBody(reqBody)
	if err != nil {
		fmt.Println(err)
		return resNote, err
	}
	payload := strings.NewReader(requestBody)
	route, err := helpers2.GetRoute(
		lib.RouteBlocksAddBlockNote,
		*commentParam.BlockId,
		commentParam.KeyId,
	)
	req, err := http.NewRequest(http.MethodPost, route, payload)
	if err != nil {
		fmt.Println(err)
		return resNote, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	res, err := helpers2.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resNote, err
	}

	defer helpers2.CloseBody(res.Body)

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
