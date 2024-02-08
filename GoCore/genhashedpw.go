package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	pw := "abcdefghij"
	hpw, _ := bcrypt.GenerateFromPassword([]byte(pw), 12)
	fmt.Println(string(hpw))

}