package emailverifier

type EmailVerifier struct {
	email string
	username string
	domain string
	errors []string
}