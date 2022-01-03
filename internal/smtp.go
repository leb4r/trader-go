package internal

import "net/smtp"

func SendMail(to, from string, message []byte, host, port, password string) error {

	auth := smtp.PlainAuth("", from, password, host)

	err := smtp.SendMail(host+":"+port, auth, from, []string{to}, message)
	if err != nil {
		return err
	}
	return nil
}
