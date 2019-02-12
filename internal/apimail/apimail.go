package apimail

import (
	"os"
	//"fmt"
	"strings"
	"time"
	"io/ioutil"
	"path/filepath"
	b64 "encoding/base64"
	"github.com/pepipost/pepipost-sdk-go/pepipost_lib/pepipost_pkg"
	"github.com/pepipost/pepipost-sdk-go/pepipost_lib/models_pkg"
	"github.com/briandowns/spinner"
	"github.com/urfave/cli"
)

var(
	apimailflags []cli.Flag
)

func init(){
	apimailflags = []cli.Flag{
		cli.StringFlag{ //apikey
			Name:  "secretkey, K",
			Usage: "Pass api-Key here. NOTE :: we highly recommend to set your apikey as enviroment variable for security reason",
		},
		cli.StringFlag{ //fromid
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
			Usage: "html data to be passed in file and give path as argument to the command ie. -Z '/tmp/mydata.html'",
		}, cli.StringFlag{ //reply-to-id
			Name:  "replytoid, r",
			Usage: "Set your reply-to-id here",
		}, cli.StringFlag{ //passing custom header
			Name:  "xheader, x",
			Value: "X-PepipostHeader,DEMO-786-XXX-999",
			Usage: "Pass custom Header in an email for recipients",
		},cli.StringFlag{ //passing custom header
			Name:  "xccheader, xcc",
			Value: "X-myCCheader,DEMO-786-XXX-lll",
			Usage: "Pass custom Header in an email for cc recipients here",
		},cli.StringFlag{ //adding attachement in an email
			Name:  "attach, a",
			Usage: "attachment to be sent in an email",
		},cli.IntFlag{ //adding template in an email
			Name:  "templateid, tid",
			Usage: "template to be sent in an email ",
		},
	}
}

func Getflags() []cli.Flag{
	return apimailflags
}

func Sendmail(a *cli.Context) (string,int){

	if a.Args().Present() {

		spin := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
		spin.Start()

		client := PepipostClient.NewPEPIPOST()
		email := client.Email()
		ApiKey := os.Getenv("PEPIKEY")

		if len(a.String("K")) != 0 {
			ApiKey = a.String("K")
		}

		username := a.String("t")
		fromEmail := a.String("f")
		fromName := a.String("fn")
		subject := a.String("s")

		//plain body text to be passed
		body :=  a.String("b")

		//EmailBody passing in html format while taking from path
		if len(a.String("Z")) != 0 {
			body = getHtmlContent(a.String("Z"))
		} else if len(a.String("z")) != 0 {
			body = a.String("z")
		}

		Body := &models_pkg.EmailBody{}
		Body.Personalizations = make([]*models_pkg.Personalizations,3)
		Body.Personalizations[0] = &models_pkg.Personalizations{}
		Body.Personalizations[0].Recipient = &username

		if len(a.String("c")) != 0 {
			rcpt_cc := strings.Split(a.String("c"), ",")
			Body.Personalizations[0].RecipientCc = &rcpt_cc
		}

		Body.From = models_pkg.From{}
		Body.From.FromEmail = &fromEmail
		Body.From.FromName = &fromName
		Body.Subject = &subject
		Body.Content = &body

		if len(a.String("a")) !=0 {
			Body.Attachments = make([]*models_pkg.EmailBodyAttachments,3)
			attach := strings.Split(a.String("a"), ",")
			for index,element := range attach{
				mydata := getHtmlContent(element) //Getting buff
				if mydata == "Error while fetching attachment"{
					return mydata,1
				}
				fdata := b64.StdEncoding.EncodeToString([]byte(mydata))
				fname := filepath.Base(element)
				Body.Attachments[index] = &models_pkg.EmailBodyAttachments{}
				Body.Attachments[index].FileContent = &fdata
				Body.Attachments[index].FileName = &fname
			}
		}

		//	email settings

		Body.Settings = models_pkg.Settings{}
		defaultval := int64(1)
		Body.Settings.Footer = &defaultval
		Body.Settings.Clicktrack = &defaultval
		Body.Settings.Opentrack = &defaultval
		Body.Settings.Unsubscribe = &defaultval

		if len(a.String("B")) !=0{
			bccid := a.String("B")
			Body.Settings.Bcc = &bccid
		}

		if len(a.String("r")) !=0{
			rtoid := a.String("r")
			Body.ReplyToId = &rtoid
		}
		if a.Int("tid") !=0{
			templateid := int64(a.Int("tid"))
			Body.TemplateId = &templateid
		}

		var err error
		var result *models_pkg.SendEmailResponse
		result, err = email.CreateSendEmail(&ApiKey, Body)
		errstatus := 0
		if err != nil{
			//TODO: Use err variable here
			spin.Stop()
			return result.Message,errstatus
		}else{
			//TODO: Use result variable here
			spin.Stop()
			errstatus := 1
			result1 := result.Message + "\nError Message :: " + result.ErrorInfo.ErrorMessage 
			return  result1,errstatus
		}
	}
	cli.ShowSubcommandHelp(a)
	return "NO Arguments Passed for apiEmail", 3
}

func getHtmlContent(path string) string {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		str := "Error while fetching attachment"
		return str
	}
	return string(data)
}

