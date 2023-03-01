package project_keys_1

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {

	url := "https://gateway.snowpal.com/project-pods/63eaac0c36aab4001233aee0/assign?blockId=63e18fa801336b0035dc10f8&keyId=63ceb977db035900128d975d"
	method := "PATCH"

	payload := strings.NewReader(`{
    "userIds": "63cabea69e9c500014d4f9b7"
}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("User-Authorization", "eyJhbGciOiJIUzI1NiJ9.Iml2PSUyNCVFMSVFQkQlMTUlQUZBcCVEN0QlQzclOTUsdGFnPSUxRVUlNjAlOEIlMTVHJTA4JUE3Q0tnJTlDJUNEJUJDOCVGMSxrZXk9JTlDJUE5JThFJTE3JUZDWSVBOSVGM1glQzAlMUUlNUQlOEIlOEElQTElQkRxRiU4MkElREIlRkQlMjIlNUUlRUMlRDRkJURDJUI1JTBCJUE5Uyx2YWx1ZT0tQiVENSUxRiVBOSU4QyVCNSU5NyVDNXElOUMlRjQlMDclMjclMDclRkElRkUlQTclOUYlODYlODUlRDAlRDUlMDklMENyJTgxLSU3RlclQzBwYiU4NiVGMCUyNSVCRCVBMCU4QSVDOCU4RiVGMiU4RSUwMyVENiU5QlF1UDglQkQlNUUlRjUlREIlQkQlQ0QlNjAlOTMi.hGrD96JtTRMIMnwg6ncHlIGEd5g9QbX9NyxGAWzdz_I")
	req.Header.Add("x-api-key", "xANSs9CMtP24cMh2eNUeG2qh5PMlh46u6hceHEDW")
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