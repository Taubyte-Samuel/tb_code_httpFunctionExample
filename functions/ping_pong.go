package lib

import (
	"bitbucket.org/taubyte/go-sdk/event"
)

//export ping
func ping(e event.Event) uint32 {
	h, err := e.HTTP()
	if err != nil {
		return 1
	}

	h.Write([]byte("Whatever I want"))

	return 0
}
