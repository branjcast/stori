package utils

import (
	"encoding/base64"
	"fmt"
	"github.com/go-mail/mail"
	"io/ioutil"
	"log"
	"os"
	"stori/src/types"
	"strconv"
	"strings"
)

var template string

func SendMail(client types.TClient) string {
	template := formatBody(client)

	pwd, _ := base64.StdEncoding.DecodeString(os.Getenv("EMAIL_PASSWORD"))

	password := string(pwd)
	smtpHost := os.Getenv("SMTP_SERVER")
	smtpPort := os.Getenv("SMTP_SERVER_PORT")
	smtpPortI, _ := strconv.Atoi(smtpPort)
	mailFrom := os.Getenv("EMAIL")

	log.Println("SMTP ENVS")
	log.Println("HOST: " + smtpHost + "\nPORT: " + smtpPort + "\nFROM: " + mailFrom + "\nPWD: " + password)

	m := mail.NewMessage()
	m.SetHeader("From", mailFrom)
	m.SetHeader("To", client.Email)
	m.SetHeader("Subject", "Stori - Transactional report")
	m.SetBody("text/html", template)

	d := mail.NewDialer(smtpHost, smtpPortI, os.Getenv("EMAIL"), password)

	if err := d.DialAndSend(m); err != nil {

		panic(err)

	}

	return template
}

func formatBody(client types.TClient) string {
	fileContent, err := ioutil.ReadFile("./assets/email/summary.html")
	if err != nil {
		log.Fatal("[ERROR] Reading HTML Template ", err)
	}

	template = string(fileContent)
	template = strings.Replace(template, "${CLIENT_NAME}", client.Name, -1)
	template = strings.Replace(template, "${CLIENT_TOTAL_BALANCE}",
		fmt.Sprintf("%.2f", client.TotalBalance), -1)

	template = strings.Replace(template, "${CLIENT_AMOUNT_BY_ACCOUNT}", "<tr><td>Average debit amount: "+fmt.Sprintf("-%.2f", client.AvgDebit)+" <br/>"+
		"Average credit amount: "+fmt.Sprintf("%.2f", client.AvgCredit)+"</td></tr>", -1)

	template = strings.Replace(template, "${CLIENT_MENSUAL_BALANCE_TR}", addTr(client), -1)
	return template
}

func addTr(client types.TClient) string {

	newTr := ""
	for _, trx := range client.GroupedBy {
		newTr +=
			"<tr><td>" + trx.Month + "</td> " + "<td>" + fmt.Sprintf("%d", len(trx.Trxs)) + "</td></tr>"
	}

	return newTr

}
