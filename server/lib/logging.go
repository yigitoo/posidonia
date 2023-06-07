package lib

import "log"

func LogError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
