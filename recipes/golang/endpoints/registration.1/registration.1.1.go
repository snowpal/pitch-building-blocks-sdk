package registration_1

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

const Url = "https://gateway-dev.snowpal.com/app/users/sign-up"
const ApiKey = "wf8sHELzWp9MGizZME5Zsjk4IntZS0e8mdYMYjjg"
const method = "POST"

func Signup(email string) {
	payload := strings.NewReader(fmt.Sprintf(`{
		"email": %s,
		"password": "Welcome1!",
		"confirmPassword": "Welcome1!"
	}`, email))

	client := &http.Client{}
	req, err := http.NewRequest(method, Url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("x-api-key", ApiKey)
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(res.Body)

	body, _ := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
