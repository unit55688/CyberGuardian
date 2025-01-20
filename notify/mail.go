package notify

import (
	"fmt"
	"net/smtp"

	"CyberGuardian/logger"
)

func MailNotify(szServer string, szPort string, szSender string, szPassword string, szReceiver string, szSubject string, szBody string, szResult string, szResultFail string) {
	message := fmt.Sprintf("From: %s\nTo: %s\nSubject: %s\n\n%s", szSender, szReceiver, szSubject, szBody)

	auth := smtp.PlainAuth("", szSender, szPassword, szServer)
	err := smtp.SendMail(szServer+":"+szPort, auth, szSender, []string{szReceiver}, []byte(message))
	if err != nil {
		logger.WARN(szResultFail)
	} else {
		logger.INFO(szResult)
	}
}
