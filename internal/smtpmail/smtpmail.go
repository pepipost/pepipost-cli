package smtpmail

import(
	"os"
	"io"
	"fmt"
	"bytes"
	"time"
	"strings"
	"net/smtp"
	"io/ioutil"
	"path/filepath"
	"linux-agent-pepipost/vendor/github.com/urfave/cli"
	"linux-agent-pepipost/vendor/github.com/briandowns/spinner"
	"linux-agent-pepipost/vendor/github.com/domodwyer/mailyak"
)

var(
	smtpmailflags []cli.Flag
)

func init(){
	smtpmailflags = []cli.Flag{
		cli.StringFlag{ //host
			Name:  "Host, d",
			Value: "smtp.pepipost.com",
			Usage: "server to sendemail (smtp.example.com)",
		}, cli.StringFlag{ //Port
			Name:  "Port, P",
			Value: "25",
			Usage: "SMTP PeerPort",
		}, cli.StringFlag{ //smtpuser
			Name:  "user, u",
			Usage: "smtpauth user",
		}, cli.StringFlag{ //smtppass
			Name:  "pass, p",
			Usage: "smtpauth passwd",
		}, cli.StringFlag{ //fromid
			Name:  "From, f",
			Value: "info@nullvalue.net",
			Usage: "fromid address",
		}, cli.StringFlag{ //fromname
			Name:  "Fromname, fn",
			Value: "",
			Usage: "From name for your email",
		}, cli.StringFlag{ //Subject
			Name:  "Subject, s",
			Value: "Test Email",
			Usage: "Email subject",
		}, cli.StringFlag{ //body
			Name:  "body, b",
			Value: "This is Simple Text email body",
			Usage: "Email body string",
		}, cli.StringFlag{ //recipient
			Name:  "To, t",
			Value: "info@nullvalue.net",
			Usage: "recipient email address",
		}, cli.StringFlag{ //recipients_cc
			Name:  "cc, c",
			Usage: "CC recipients",
		}, cli.StringFlag{ //recipients_bcc
			Name:  "bcc, B",
			Usage: "bcc recipients for your email",
		}, cli.StringFlag{ //html data
			Name:  "htmltext, z",
			Value: "",
			Usage: "html data in text format",
		}, cli.StringFlag{ //html data from particular path in filename.html
			Name:  "htmlpath, Z",
			Usage: "html data in text format",
		}, cli.StringFlag{ //reply-to-id
			Name:  "replytoid, r",
			Usage: "Emailid you want to keep send in reply to id",
		}, cli.StringFlag{ //passing custom header
			Name:  "header, x",
			Usage: "Pass custom Header in an email",
		}, cli.StringFlag{ //adding attachement in an email
			Name:  "attach, a",
			Usage: "attachment to be sent in an email",
		},
	}
}

func Getflags() []cli.Flag{
	return smtpmailflags
}

func Sendmail(a *cli.Context) (string , error) {

	fmt.Print("\nSending Email through SMTP\n")

	spin := spinner.New(spinner.CharSets[43], 100*time.Millisecond)
	spin.Start()


	smtpuser := os.Getenv("PEPISMTPUSER")
	if len(a.String("u")) != 0 {
		smtpuser = a.String("u")
	}
	smtppass := os.Getenv("PEPISMTPPASS")

	if len(a.String("p")) != 0{
		smtppass = a.String("p")
	}

	//Setting Authentication
	domain := a.String("d") + ":" + a.String("P")
	mail := mailyak.New(domain, smtp.PlainAuth("", smtpuser, smtppass, a.String("d"))) //auth

	//Setting To email address
	toemailids := strings.Split(a.String("t"), ",")
	mail.To(toemailids...)

	//Setting CC email address
	if len(a.String("c")) != 0 {
		ccemailids := strings.Split(a.String("c"), ",")
		mail.Cc(ccemailids...)
	}

	//Setting BCC email address
	if len(a.String("B")) != 0 {
		bccemails := strings.Split(a.String("B"), ",")
		mail.Bcc(bccemails...)
	}

	mail.From(a.String("f"))                //fromaddress
	mail.FromName(a.String("fn"))           //fromname
	mail.Subject(a.String("s"))             //ToSubject
	mail.ReplyTo(a.String("r"))             //reply to id in an email
	mail.Plain().Set(a.String("b"))         //EmailBody in simple plain text

	//passing header in SMTP mail
	if len (a.String("x")) !=0{
		emailHeaders := strings.Split(a.String("x"), ",")
		mail.AddHeader(emailHeaders[0],emailHeaders[1])
	}


	//Sending Attachemnet in email
	if len(a.String("a")) !=0{
		mydata := getHtmlContent(a.String("a")) //Getting buff
		buf := &bytes.Buffer{}
		io.WriteString(buf, mydata)
		mail.Attach(filepath.Base(a.String("a")), buf) //Attachment in an email
	}

	//EmailBody passing in html format while taking from path
	if len(a.String("Z")) != 0 {
		html := getHtmlContent(a.String("Z"))
		mail.HTML().Set(html)
	} else if len(a.String("z")) != 0 {
		mail.HTML().Set(a.String("z"))
	}

	//sending email
	if err := mail.Send(); err != nil {
		spin.Stop()
		return "",err
	} else {
		spin.Stop()
		fmt.Println("Email Sent Successfully \n")
		return "",err
	}

}

func getHtmlContent(path string) string {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("File reading error", err)
		str := "Error while File reading"
		return str
	}
	return string(data)
}

