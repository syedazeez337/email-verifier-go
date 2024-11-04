package main

import (
	"bufio"
	emailverifier "emailverifier/app"
	"fmt"
	"os"
	"strings"
)

func main() {
	/*
	email := "syed@gmail.com"
	verifier := emailverifier.NewEmailVerifier(email)

	format := verifier.ValidateFormat()
	fmt.Println(format)
	*/

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("\nEnter an email to verify (or 'quit' to exit): ")
		scanner.Scan()
		email := scanner.Text()

		if strings.ToLower(email) == "quit" {
			break
		}

		verifier := emailverifier.NewEmailVerifier(email)

		formatValid := verifier.ValidateFormat()
		var domainValid bool
		if formatValid {
			domainValid = verifier.ValidateDomain()
		}

		fmt.Println("\nVerification Results:")
		fmt.Printf("Email: %s\n", email)
		fmt.Printf("Format Valid: %v\n", formatValid)
		if formatValid {
			fmt.Printf("Domain Valid: %v\n", domainValid)
		}

		if len(verifier.GetErrors()) > 0 {
			fmt.Println("\nValidation Errors:")
			for _, err := range verifier.GetErrors() {
				fmt.Printf("- %s\n", err)
			}
		}
	}
}
