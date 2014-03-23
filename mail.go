package mail

import (
	"bytes"
	"os/exec"
	"strings"
	"time"
)

const CRLF = "\r\n"

type Mail struct {
	From    string
	To      []string
	Subject string
	Body    string
}

func NewMail(from string, to []string, subject string, body string) *Mail {
	return &Mail{
		From:    from,
		To:      to,
		Subject: subject,
		Body:    body,
	}
}

func (m *Mail) Send() error {
	var recipients = ""
	//TODO: Attachments
	//var boundary string = "f00cc1c5b16b00b5cafebabe741213"
	for _, i := range m.To {
		recipients += i + ","
	}
	recipients = strings.TrimRight(recipients, ",")
	b := bytes.NewBuffer(nil)
	b.WriteString("Date: " + time.Now().Format("Mon, 2 Jan 2006 15:04:05 -0700") + CRLF)
	b.WriteString("From: " + m.From + CRLF)
	b.WriteString("To: " + recipients + CRLF)
	b.WriteString("Subject: " + m.Subject + CRLF)
	b.WriteString("MIME-Version: 1.0" + CRLF)
	b.WriteString("Content-Type: text/plain; charset=utf-8\n")
	b.WriteString(m.Body)
	cmd := exec.Command("sendmail", "-t", "-i")
	cmd.Stdin = b
	return cmd.Run()
}
