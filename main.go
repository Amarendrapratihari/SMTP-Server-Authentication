package main

import (
	email_auth "cachatto/emailserver/emailmfa"
	"fmt"
)

func main() {
	// MFA EMAIL AUTHENTICATION
	err := email_auth.SendEmail("reciver@gmail.com", "OTP", "Dear customer, the one time password (OTP) to reset your password", "45777")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email sent successfully.")

}
