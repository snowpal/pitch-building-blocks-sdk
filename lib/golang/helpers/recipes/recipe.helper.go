package recipes

import (
	"development/go/recipes/lib/golang"
	"development/go/recipes/lib/golang/structs/response"
	"fmt"
	"time"
)

func sleepWindow(sleepTime time.Duration) {
	time.Sleep(time.Second * sleepTime)
}

func SleepBefore() {
	sleepWindow(1)
}

func SleepAfter() {
	sleepWindow(2)
}

// ValidateDependencies We require that the first recipe be run before anything else as it registers a bunch of users.
// To verify if it was actually run, we do this "random" check.
func ValidateDependencies() (response.User, error) {
	user, err := SignIn("a3@gmail.com", golang.Password)
	fmt.Println(user)
	fmt.Println(err)
	if err != nil {
		return user, err
	}

	return user, nil
}
