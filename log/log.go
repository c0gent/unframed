package log

import (
	"fmt"
)

func Error(s ...interface{}) {
	fmt.Println("Error:", s)
}

func Message(s ...interface{}) {
	fmt.Println("Message:", s)
}