package main

import (
	"github.com/crossphoton/email-microservice/src"
)

func main() {
	// Raw request
	rawEmailRequest := src.RawSendEmailRequest{
		Recipients: []string{
			"email@e.crossphoton.tech",
			"spam@e.crossphoton.tech",
			"support@e.crossphoton.tech",
		},
		Body: []byte(RawEmailBody),
	}
	sendRawEmail(&rawEmailRequest)

	// Standard email
	emailRequest := src.SendEmailRequest{
		Recipients: &src.Recipients{
			To: []string{
				"test",
				"email@e.crossphoton.tech",
			},
			Cc: []string{
				"spam@e.crossphoton.tech",
			},
			Bcc: []string{
				"support@e.crossphoton.tech",
			},
		},
		Subject:     "Hi there. I hope you're good",
		ContentType: "text/html",
		Body:        "<h1>This is heading</h1><p>This is text</p>",
	}
	sendEmailStd(&emailRequest)
}

const RawEmailBody = `From: "Sender Name" <sender@example.com>
To: recipient@example.com
Subject: Customer service contact info
Content-Type: multipart/mixed;
    boundary="a3f166a86b56ff6c37755292d690675717ea3cd9de81228ec2b76ed4a15d6d1a"

--a3f166a86b56ff6c37755292d690675717ea3cd9de81228ec2b76ed4a15d6d1a
Content-Type: multipart/alternative;
    boundary="sub_a3f166a86b56ff6c37755292d690675717ea3cd9de81228ec2b76ed4a15d6d1a"

--sub_a3f166a86b56ff6c37755292d690675717ea3cd9de81228ec2b76ed4a15d6d1a
Content-Type: text/plain; charset=iso-8859-1
Content-Transfer-Encoding: quoted-printable

Please see the attached file for a list of customers to contact.

--sub_a3f166a86b56ff6c37755292d690675717ea3cd9de81228ec2b76ed4a15d6d1a
Content-Type: text/html; charset=iso-8859-1
Content-Transfer-Encoding: quoted-printable

<html>
<head></head>
<body>
<h1>Hello!</h1>
<p>Please see the attached file for a list of customers to contact.</p>
</body>
</html>

--sub_a3f166a86b56ff6c37755292d690675717ea3cd9de81228ec2b76ed4a15d6d1a--

--a3f166a86b56ff6c37755292d690675717ea3cd9de81228ec2b76ed4a15d6d1a
Content-Type: text/plain; name="customers.txt"
Content-Description: customers.txt
Content-Disposition: attachment;filename="customers.txt";
    creation-date="Sat, 05 Aug 2017 19:35:36 GMT";
Content-Transfer-Encoding: base64

SUQsRmlyc3ROYW1lLExhc3ROYW1lLENvdW50cnkKMzQ4LEpvaG4sU3RpbGVzLENhbmFkYQo5MjM4
OSxKaWUsTGl1LENoaW5hCjczNCxTaGlybGV5LFJvZHJpZ3VleixVbml0ZWQgU3RhdGVzCjI4OTMs
QW5heWEsSXllbmdhcixJbmRpYQ==

--a3f166a86b56ff6c37755292d690675717ea3cd9de81228ec2b76ed4a15d6d1a--
`
