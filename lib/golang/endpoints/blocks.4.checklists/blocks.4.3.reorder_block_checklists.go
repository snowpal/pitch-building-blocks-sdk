package blocks_4

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {

	url := "https://gateway.snowpal.com/blocks/%7B%7Bblocks.blockId%7D%7D/checklists/reorder?keyId=%7B%7Bkeys.custom.keyId1%7D%7D"
	method := "PATCH"

	payload := strings.NewReader(`{"checklistIds":"63db073e9d84870014bdfb58,63d136d4fea2a60015442478"}`)

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
