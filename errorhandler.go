package errorHandler

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/mailgun/mailgun-go"
)

func IsError(err error, errClass string, mgConf MailgunConf) bool {
	if err != nil {
		switch errClass {
		case "fatal":
			Sendmail(mgConf, "PANIC!!!", "A panicy error occured: "+err.Error())
			log.Fatal("Panic: ", err, " - ", mgConf.Receiver, " has been notified by mail")
		case "panic":
			Sendmail(mgConf, "PANIC!!!", "A panicy error occured: "+err.Error())
			log.Panic("Panic: ", err, " - ", mgConf.Receiver, " has been notified by mail")
		case "log&mail":
			Sendmail(mgConf, "A simple error occured", "A simple error occured: "+err.Error()+" - "+mgConf.Receiver)
			log.Print("A simple error occured: ", err, " - ", mgConf.Receiver, " has been notified by mail")
		case "log":
			log.Print("A simple error occured: ", err)
		default:
			Sendmail(mgConf, "Poorly set up logger", `Your logger was poorly set up.\n
			Make sure to choose from the official list of error classes in all functions\n
			See the github page for a list`)
			log.Print(`Your logger was poorly set up.\n
			Make sure to choose from the official list of error classes in all functions\n
			See the github page for a list`)
		}
		return true
	} else {
		return false
	}
}

// LocalDescription is for you to set in order to describe what file or what server is actually sending the error log
type MailgunConf struct {
	MailgunDomain    string
	PrivateAPIKey    string
	Sender           string
	Receiver         string
	LocalDescription string
}

func Sendmail(mgConf MailgunConf, subject string, body string) {
	// Create an instance of the Mailgun Client
	mg := mailgun.NewMailgun(mgConf.MailgunDomain, mgConf.PrivateAPIKey)

	// The message object allows you to add attachments and Bcc recipients
	message := mg.NewMessage(mgConf.Sender, subject, body, mgConf.Receiver)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// Send the message	with a 10 second timeout
	resp, id, err := mg.Send(ctx, message)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("ID: %s Resp: %s\n", id, resp)
}
