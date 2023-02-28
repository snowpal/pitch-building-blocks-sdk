package registration_1

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Activate() {

	url := "https://gateway-dev.snowpal.com/app/user-verified/63f66881a6f32f0037e35b24"
	method := "PATCH"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("x-api-key", "wf8sHELzWp9MGizZME5Zsjk4IntZS0e8mdYMYjjg")

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
