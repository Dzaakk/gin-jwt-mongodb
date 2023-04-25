package helpers

import "log"

func FatalError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
