package registration_1

import (
	"development/go/recipes/lib/golang"
	"development/go/recipes/lib/golang/helpers"
	"fmt"
	"io"
	"net/http"
)

func Activate(userId string) {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPatch, fmt.Sprintf(helpers.GetRoute(golang.RouteRegistrationActivateUser), userId), nil)

	if err != nil {
		fmt.Println(err)
		return err
	}

	helpers.AddAppHeaders(req)
	res, _ := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer helpers.CloseBody(res.Body)

	body, _ := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println("..activated user", string(body))
}
