package functions

import "log"

func CheckFatal(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
