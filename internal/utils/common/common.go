package common

import "os"

func Exit(msg string) {
	println(msg)
	os.Exit(1)
}
