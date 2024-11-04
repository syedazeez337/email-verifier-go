package emailverifier

import "net"

func (ev *EmailVerifier) ValidateDomain() bool {
	_, err := net.LookupMX(ev.domain)
	if err != nil {
		ev.errors = append(ev.errors, "Domain does not have valid MX records")
		return false
	}
	return true
}