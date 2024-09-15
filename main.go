package main

import (
	"fmt"
	"log"
	"math/rand/v2"
	"net/smtp"
	"strconv"
)

func SendVeriEmail(email string, code int) {
	from := "Enter SMTP Email"
	password := "Enter SMTP Password"
	msg := "From: " + from + "\n" +
		"To: " + email + "\n" +
		"Subject: Go Email Verification\n\n" +
		"The verification code is " + strconv.Itoa(code)
	err := smtp.SendMail("Enter SMTP provider",
		smtp.PlainAuth("", from, password, "Enter SMTP provider"),
		from, []string{email}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return
	}
	fmt.Println("Email has been sent. Enter the verification code: ")
	var clientCode int
	fmt.Scan(&clientCode)
	if code == clientCode {
		fmt.Println("Correct. Email is verified")
		return
	} else {
		fmt.Println("Incorrect.")
		return
	}
}

func main() {
	var email string
	fmt.Println("Type the email you want verified: ")
	fmt.Scan(&email)
	fmt.Printf("Is the email you entered correct: %s? (Y/N)\n", email)
	var verifyEmail string
	fmt.Scan(&verifyEmail)
	if verifyEmail == "Y" {
		verifyCode := rand.IntN(8999) + 1000
		SendVeriEmail(email, verifyCode)
	} else {
		fmt.Println("Restart the program to enter the email.")
	}

}
