package helpers

import "io"

func CloseBody(Body io.ReadCloser) {
	err := Body.Close()
	if err != nil {
		return
	}
}
