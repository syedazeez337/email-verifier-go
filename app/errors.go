package emailverifier

func (ev *EmailVerifier) GetErrors() []string {
	return ev.errors
}