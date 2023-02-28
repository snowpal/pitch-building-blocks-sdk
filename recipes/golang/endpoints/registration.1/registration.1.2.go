package registration_1

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func SignIn() {

	url := "https://gateway-dev.snowpal.com/app/users/sign-in"
	method := "POST"

	payload := strings.NewReader(`{
		"email": "apiuser3@gmail.com",
		"password": "Welcome1!"
	}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("x-api-key", "wf8sHELzWp9MGizZME5Zsjk4IntZS0e8mdYMYjjg")
	req.Header.Add("Content-Type", "application/json")

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
