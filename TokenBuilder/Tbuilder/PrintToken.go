package Tbuilder

import (
	"fmt"
	"math/rand"
	"time"
)

func PrintToken() { //Can be changable parametricaly as wished style token code
	rand.Seed(time.Now().Unix())
	minSpecialChar := 0
	minNum := 0
	minUpperCase := 0
	passwordLength := 6
	minlowercase := 6
	password := generateToken(passwordLength, minSpecialChar, minNum, minUpperCase, minlowercase)
	fmt.Println(password)
}
