package keys_1

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetAllKeys() {

	url := "https://gateway-dev.snowpal.com/keys"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("User-Authorization", "eyJhbGciOiJIUzI1NiJ9.Iml2PSVDQ3hxbyU4RCVGRGdMJUFBJTA4JThEXyx0YWc9JTk4JTVDJTkzJUVGJTVDfiVCOCU4QTMlRDclQjBzJUFEJUU5JUMzJUQzLGtleT1TJTE0JTA5JTI3JTA0JUQ3JTEwJUFDJUEwSiVDMyVBQSU3RiVFNyVFQyVBOXglMDB4JTkwJUU1JTA1Uk8lRTklOEElN0YlQkFvJTkySiUxQyx2YWx1ZT1mJUY0JUFEJTJGJUYxJUU1JTI2JUM5aCVEQiU4RiU5QyUwMUslMjklQ0UlQkJLeiUxMSVFNCU1QiU4NXdUJUYzJTkxJTI0JUYyJTEwJUUyNyUyQ1NiJTI3JTgzJTAyTCUxREglMjclQTclMjElMTIlQ0FHMCUzRiVGNSVDOGs0byUyRiU4NTglQTQi.sX02yyGbIgD7gK9u_sP6SYJVt6-FQE0iSrvyY-KAe3M")
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
