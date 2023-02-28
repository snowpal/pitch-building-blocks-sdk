package keys_1

import (
	constants "development/go/recipes/endpoints"
	"fmt"
	"io/ioutil"
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
	req.Header.Add("x-api-key", constants.ApiKey)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
