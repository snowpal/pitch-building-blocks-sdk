package recipes

import "time"

func sleepWindow(sleepTime time.Duration) {
	time.Sleep(time.Second * sleepTime)
}

func SleepBefore() {
	sleepWindow(1)
}

func SleepAfter() {
	sleepWindow(2)
}
