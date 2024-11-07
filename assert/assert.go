package assert

import (
	"fmt"
	"log"
)

/*
given an error, if it is not nil, logs a fatal message and exits program.

params:

	err error
	msg string
*/
func Error(err error, msg string) {
	if err != nil {
		log.Fatalln(fmt.Sprint(msg, ":", err.Error()))
	}
}

/*
given an error, if it is not nil, logs a message.

params:

	err error
	msg string
*/
func Info(err error, msg string) {
	if err != nil {
		log.Println(msg)
	}
}
