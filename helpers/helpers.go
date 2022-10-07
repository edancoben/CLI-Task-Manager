package helpers

import (
	"fmt"
	"os"
)

func Must(err error, msg string) {
	if err != nil {
		fmt.Println(msg)
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
