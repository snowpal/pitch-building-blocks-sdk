package teacher_keys_1

import (
	"fmt"
	"io"
	"net/http"
)

func main(jwtToken string) {

	url := "classroom-pods/63ec12a8f1a0b80010ec6889/submissions/comments/as-student?blockId=63cfd9321bff200012b0ec40&keyId=63ebe3f3c003770011dbec0a"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("x-api-key", "xANSs9CMtP24cMh2eNUeG2qh5PMlh46u6hceHEDW")
	req.Header.Add("User-Authorization", "eyJhbGciOiJIUzI1NiJ9.Iml2PSVFNyU1RSUxOCVBOCUwQSVBNCU3QyVGNiUxMUElQzVTLHRhZz0lMjQlRUIlREFGJTlEU2glOTMlQzAlQjclRTYlRTUlOUQlOTZBJUEyLGtleT1fJUY0JTAyJTAzJTA5JTEzeSU4NCUxRnAlOTdEJUQyJUYxJUU1JUI0JTE2JUMzZCVCMTAlODRxJUY4JTlEJUQ2JTJGJUQzJUE0JTk2JTBCJUFFLHZhbHVlPVczJTYwJTA4JUYyJUQxJUQ4JTg5WiVGNiUxMjklQ0ElQTclQzklMUQlMTYlRDRYJTVCVSVGMCVGNCU1QiVCMCU5NWdPJTdGJUI5JTFFJUM2JUU0JTNENSVERiUwQyU4QyU0MFUlM0YweiVBQSVFQyUwMCU4QXglQzYlQjglNDBqJTJGJUQxJURFJTBEJUY0JTkxIg.45DoNYCo3Ty7ElMlvYjE-qKTTaePwShIJo7gwd29Xhw")

	res, err := client.Do(req)
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

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
