package teacher_keys_2

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	url := "https://gateway.snowpal.com/classroom-pods/63ec12a8f1a0b80010ec6889/submissions/attachments/as-teacher?studentId=63cabea69e9c500014d4f9b7&blockId=63cfd9321bff200012b0ec40&keyId=63ceb862db035900148d975d"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("User-Authorization", "eyJhbGciOiJIUzI1NiJ9.Iml2PSUyNCVFMSVFQkQlMTUlQUZBcCVEN0QlQzclOTUsdGFnPSUxRVUlNjAlOEIlMTVHJTA4JUE3Q0tnJTlDJUNEJUJDOCVGMSxrZXk9JTlDJUE5JThFJTE3JUZDWSVBOSVGM1glQzAlMUUlNUQlOEIlOEElQTElQkRxRiU4MkElREIlRkQlMjIlNUUlRUMlRDRkJURDJUI1JTBCJUE5Uyx2YWx1ZT0tQiVENSUxRiVBOSU4QyVCNSU5NyVDNXElOUMlRjQlMDclMjclMDclRkElRkUlQTclOUYlODYlODUlRDAlRDUlMDklMENyJTgxLSU3RlclQzBwYiU4NiVGMCUyNSVCRCVBMCU4QSVDOCU4RiVGMiU4RSUwMyVENiU5QlF1UDglQkQlNUUlRjUlREIlQkQlQ0QlNjAlOTMi.hGrD96JtTRMIMnwg6ncHlIGEd5g9QbX9NyxGAWzdz_I")
	req.Header.Add("x-api-key", "xANSs9CMtP24cMh2eNUeG2qh5PMlh46u6hceHEDW")

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
