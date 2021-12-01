package src

import (
	"net/mail"
)

func validateEmail(email string) error {
	_, err := mail.ParseAddress(email)
	return err
}

func (r *Recipients) validate() error {
	receivers := r.To
	receivers = append(receivers, r.Cc...)
	receivers = append(receivers, r.Bcc...)

	return validateEmails(receivers)
}

func validateEmails(emails []string) error {
	for _, email := range emails {
		err := validateEmail(email)
		if err != nil {
			return err
		}

	}

	return nil
}
