package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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

func newRequest(route string, method MethodType) (*http.Request, error) {
	url := fmt.Sprintf("%s/%s", BaseUrl, route)
	req, err := http.NewRequest(Method[method], url, nil)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	req.Header.Add(headerKeys.ApiKey, AwsXApiKey)
	return req, err
}

func newRequestWithAuth(route string, method MethodType, userAuth string) (*http.Request, error) {
	req, err := newRequest(route, method)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	req.Header.Add(headerKeys.UserAuth, userAuth)
	return req, err
}

func makeRequest(req *http.Request) json.RawMessage {
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return body
}

func registerUser() User {
	var user User
	req, err := newRequest("/app/users/sign-up", POST)
	if err != nil {
		fmt.Println(err)
		return user
	}
	response := makeRequest(req)
	err = json.Unmarshal(response, &user)
	if err != nil {
		fmt.Println(err)
		return user
	}
	fmt.Println(fmt.Sprintf("User is registered, and id is %s", user.ID))
	return user
}

func activateUser(userId string) bool {
	req, err := newRequest(fmt.Sprintf("/app/user-verified/%s", userId), PATCH)
	if err != nil {
		fmt.Println(err)
		return false
	}
	makeRequest(req)
	fmt.Println("User is activated")
	return true
}

func loginUser() User {
	var user User
	req, err := newRequest("/app/users/sign-in", POST)
	if err != nil {
		fmt.Println(err)
		return user
	}
	response := makeRequest(req)
	err = json.Unmarshal(response, &user)
	if err != nil {
		fmt.Println(err)
		return user
	}
	fmt.Println("User is logged in")
	return user
}

func getAllKeys(userAuth string) []Key {
	req, err := newRequestWithAuth("/keys", GET, userAuth)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	response := makeRequest(req)
	keysResponse := map[string][]Key{}
	err = json.Unmarshal(response, &keysResponse)
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
	req, err := newRequestWithAuth("/keys", POST, userAuth)
	if err != nil {
		fmt.Println(err)
		return key
	}
	response := makeRequest(req)
	err = json.Unmarshal(response, &key)
	if err != nil {
		fmt.Println(err)
		return key
	}
	fmt.Println(fmt.Sprintf("Key is created with %s id.", key.ID))
	return key
}

func getProfile(userAuth string) User {
	var user User
	req, err := newRequestWithAuth("/profiles", GET, userAuth)
	if err != nil {
		fmt.Println(err)
		return user
	}
	response := makeRequest(req)
	err = json.Unmarshal(response, &user)
	if err != nil {
		fmt.Println(err)
		return user
	}
	fmt.Println("User is fetched.")
	return user
}

func updateUsername(userAuth string, username string) bool {
	req, err := newRequestWithAuth(fmt.Sprintf("/profiles/username/%s", username), PATCH, userAuth)
	if err != nil {
		fmt.Println(err)
		return false
	}
	makeRequest(req)
	fmt.Println("Username is updated")
	return true
}

func main() {
	fmt.Println("Recipe1")
	user := registerUser()
	activateUser(user.ID)
	user = loginUser()
	keys := getAllKeys(user.JwtToken)
	if keys != nil {
		printKeys(keys)
	}

	fmt.Println("Recipe2")
	user = loginUser()
	addNewKey(user.JwtToken)
	keys = getAllKeys(user.JwtToken)
	if keys != nil {
		printKeys(keys)
	}

	fmt.Println("Recipe3")
	user = loginUser()
	user = getProfile(user.JwtToken)
	fmt.Println(fmt.Sprintf("Username: %s", user.Username))
	updateUsername(user.JwtToken, "unique_username")
	user = getProfile(user.JwtToken)
	fmt.Println(fmt.Sprintf("Username: %s", user.Username))
}
