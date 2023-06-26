package recipes

import (
	"time"

	"github.com/snowpal/pitch-building-blocks-sdk/lib"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/response"
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
	user, err := SignIn(lib.DefaultEmail, lib.Password)
	if err != nil {
		return user, err
	}

	return user, nil
}
