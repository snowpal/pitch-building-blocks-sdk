package registration_1

import (
	"development/go/recipes/endpoints"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func Signup(email string) {
	payload := strings.NewReader(fmt.Sprintf(`{
		"email": %s,
		"password": "Welcome1!",
		"confirmPassword": "Welcome1!"
	}`, email))

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPost, endpoints.Url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("x-api-key", endpoints.ApiKey)
	req.Header.Add("Content-Type", "application/json")

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
