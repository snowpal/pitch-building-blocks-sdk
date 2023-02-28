package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type HeaderKeys struct {
	ApiKey   string
	UserAuth string
}

var headerKeys = HeaderKeys{
	ApiKey:   "x-api-key",
	UserAuth: "User-Authorization",
}

const AwsXApiKey = "xANSs9CMtP24cMh2eNUeG2qh5PMlh46u6hceHEDW"
const BaseUrl = "https://gateway.snowpal.com"

type MethodType int

const (
	GET MethodType = iota
	POST
	PATCH
	DELETE
)

var Method = map[MethodType]string{
	GET:    "GET",
	POST:   "POST",
	PATCH:  "PATCH",
	DELETE: "DELETE",
}

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	JwtToken string `json:"jwtToken"`
}

type Key struct {
	ID   string `json:"id"`
	Name string `json:"keyName"`
}

func newRequest(route string, method MethodType, body io.Reader) (*http.Request, error) {
	url := fmt.Sprintf("%s/%s", BaseUrl, route)
	req, err := http.NewRequest(Method[method], url, body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	req.Header.Add(headerKeys.ApiKey, AwsXApiKey)
	return req, err
}

func newRequestWithAuth(route string, method MethodType, userAuth string, body io.Reader) (*http.Request, error) {
	req, err := newRequest(route, method, body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	req.Header.Add(headerKeys.UserAuth, userAuth)
	return req, err
}

func makeRequest(req *http.Request) (json.RawMessage, error) {
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	if res.StatusCode == 200 || res.StatusCode == 201 {
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				fmt.Println(err)
			}
		}(res.Body)
		body, err := io.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		return body, err
	} else if res.StatusCode == 204 {
		return nil, nil
	}
	return nil, errors.New(res.Status)
}

func registerUser() User {
	var user User
	body := strings.NewReader(`{
		"email": "apiuser4@gmail.com",
		"password": "Welcome1!",
		"confirmPassword": "Welcome1!"
	}`)
	req, err := newRequest("/app/users/sign-up", POST, body)
	if err != nil {
		fmt.Println(err)
		return user
	}
	res, err := makeRequest(req)
	if err != nil {
		fmt.Println(err)
		return user
	}
	err = json.Unmarshal(res, &user)
	if err != nil {
		fmt.Println(err)
		return user
	}
	fmt.Println(fmt.Sprintf("User is registered, and id is %s", user.ID))
	return user
}

func activateUser(userId string) bool {
	req, err := newRequest(fmt.Sprintf("/app/user-verified/%s", userId), PATCH, nil)
	if err != nil {
		fmt.Println(err)
		return false
	}
	_, err = makeRequest(req)
	if err != nil {
		fmt.Println(err)
		return false
	}
	fmt.Println("User is activated")
	return true
}

func loginUser() User {
	body := strings.NewReader(`{
		"email": "apiuser4@gmail.com",
		"password": "Welcome1!"
	}`)
	var user User
	req, err := newRequest("/app/users/sign-in", POST, body)
	if err != nil {
		fmt.Println(err)
		return user
	}
	res, err := makeRequest(req)
	if err != nil {
		fmt.Println(err)
		return user
	}
	err = json.Unmarshal(res, &user)
	if err != nil {
		fmt.Println(err)
		return user
	}
	fmt.Println("User is logged in")
	return user
}

func getAllKeys(userAuth string) []Key {
	req, err := newRequestWithAuth("/keys", GET, userAuth, nil)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	res, err := makeRequest(req)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	keysResponse := map[string][]Key{}
	err = json.Unmarshal(res, &keysResponse)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return keysResponse["keys"]
}

func printKeys(keys []Key) {
	for index, key := range keys {
		fmt.Println(fmt.Sprintf("Key%d: %s", index, key.Name))
	}
}

func addNewKey(userAuth string) Key {
	var key Key
	req, err := newRequestWithAuth("/keys", POST, userAuth, nil)
	if err != nil {
		fmt.Println(err)
		return key
	}
	res, err := makeRequest(req)
	if err != nil {
		fmt.Println(err)
		return key
	}
	err = json.Unmarshal(res, &key)
	if err != nil {
		fmt.Println(err)
		return key
	}
	fmt.Println(fmt.Sprintf("Key is created with %s id.", key.ID))
	return key
}

func getProfile(userAuth string) User {
	var user User
	req, err := newRequestWithAuth("/profiles", GET, userAuth, nil)
	if err != nil {
		fmt.Println(err)
		return user
	}
	res, err := makeRequest(req)
	if err != nil {
		fmt.Println(err)
		return user
	}
	err = json.Unmarshal(res, &user)
	if err != nil {
		fmt.Println(err)
		return user
	}
	fmt.Println("User is fetched.")
	return user
}

func updateUsername(userAuth string, username string) bool {
	req, err := newRequestWithAuth(fmt.Sprintf("/profiles/username/%s", username), PATCH, userAuth, nil)
	if err != nil {
		fmt.Println(err)
		return false
	}
	_, err = makeRequest(req)
	if err != nil {
		fmt.Println(err)
		return false
	}
	fmt.Println("Username is updated")
	return true
}

func createRecipe1() {
	fmt.Println("Recipe1")
	user := registerUser()
	if user.ID == "" {
		return
	}
	succeed := activateUser(user.ID)
	if !succeed {
		return
	}
	user = loginUser()
	if user.JwtToken == "" {
		return
	}
	keys := getAllKeys(user.JwtToken)
	if keys != nil {
		printKeys(keys)
	}
}

func createRecipe2() {
	fmt.Println("Recipe2")
	user := loginUser()
	if user.JwtToken == "" {
		return
	}
	key := addNewKey(user.JwtToken)
	if key.ID == "" {
		return
	}
	keys := getAllKeys(user.JwtToken)
	if keys != nil {
		printKeys(keys)
	}
}

func createRecipe3() {
	fmt.Println("Recipe3")
	user := loginUser()
	if user.JwtToken == "" {
		return
	}
	user = getProfile(user.JwtToken)
	if user.Username == "" {
		return
	}
	fmt.Println(fmt.Sprintf("Username: %s", user.Username))
	succeed := updateUsername(user.JwtToken, "unique_username")
	if !succeed {
		return
	}
	user = getProfile(user.JwtToken)
	fmt.Println(fmt.Sprintf("After update Username: %s", user.Username))
}

func main() {
	createRecipe1()
	createRecipe2()
	createRecipe3()
}
