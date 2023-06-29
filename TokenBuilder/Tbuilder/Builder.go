package Tbuilder

import (
	"math/rand"
	"strings"
)

var (
	lowerCharSet   = "abcdedfghijklmnopqrst"
	upperCharSet   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	specialCharSet = "!@#$%&*"
	numberSet      = "0123456789"
	allCharSet     = lowerCharSet + upperCharSet + specialCharSet + numberSet
)

func generateToken(TokenLength, TokenminSpecialChar, TokenminNum, TokenminUpperCase, TokenminlowerCase int) string {
	var token strings.Builder

	for i := 0; i < TokenminSpecialChar; i++ { //For set min character
		random := rand.Intn(len(specialCharSet))
		token.WriteString(string(specialCharSet[random]))
	}

	for i := 0; i < TokenminNum; i++ { //For set min Number
		random := rand.Intn(len(numberSet))
		token.WriteString(string(numberSet[random]))
	}

	for i := 0; i < TokenminUpperCase; i++ { //FOr set min UpperCase
		random := rand.Intn(len(upperCharSet))
		token.WriteString(string(upperCharSet[random]))
	}

	for i := 0; i < TokenminlowerCase; i++ { //For set min lowercase
		random := rand.Intn(len(lowerCharSet))
		token.WriteString(string(lowerCharSet[random]))
	}

	remainingLength := TokenLength - TokenminSpecialChar - TokenminNum - TokenminUpperCase - TokenminlowerCase // After the calculation, the remainder is filled randomly.

	for i := 0; i < remainingLength; i++ {
		random := rand.Intn(len(allCharSet))
		token.WriteString(string(allCharSet[random]))
	}

	_token := []rune(token.String()) //To return value
	return string(_token)

}
