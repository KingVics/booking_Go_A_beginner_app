// package main
package helper

import "strings"

func ValidateUserInput(userName string, lastName string, email string, userTicker uint, remainTicket uint) (bool, bool, bool) {
	isValidNames := len(userName) >= 2 && len(lastName) >= 2
	isEmail := strings.Contains(email, "@")
	isValidTicket := userTicker > 0 && userTicker <= remainTicket

	return isValidNames, isEmail, isValidTicket
}