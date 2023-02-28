package keys_1

import (
	constants "development/go/recipes/endpoints"
	"fmt"
	"io"
	"net/http"
)

func GetAllKeys(jwtToken string) {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, constants.UrlGeyKeys, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("User-Authorization", jwtToken)
	req.Header.Add("x-api-key", constants.XApiKey)

	res, _ := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(res.Body)

	body, _ := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
