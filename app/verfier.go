package emailverifier

import (
	"strings"
)

func NewEmailVerifier(email string) *EmailVerifier {
	return &EmailVerifier{
		email:  email,
		errors: []string{},
	}
}

func (ev *EmailVerifier) ValidateFormat() bool {
	parts := strings.Split(ev.email, "@")
	if len(parts) != 2 {
		ev.errors = append(ev.errors, "Email must contain exactly one @ symbol")
		return false
	}

	ev.username = parts[0]
	ev.domain = parts[1]

	// fmt.Println(ev)

	if len(ev.username) == 0 {
		ev.errors = append(ev.errors, "Username cannot be empty")
		return false
	}

	if len(ev.domain) < 3 || !strings.Contains(ev.domain, ".") {
		ev.errors = append(ev.errors, "Invalid domain format")
		return false
	}

	return true
}
