package main

import (
	"time"
	"token/Tbuilder"
)

func main() {
	for {
		Tbuilder.PrintToken()
		time.Sleep(5 * time.Second)
	}
}
