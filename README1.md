# Linux-Agent-Pepipost :penguin:

## Introduction

Linux agent for Pepipost is developed to send email using command line interface(CLI). This CLI not only sends email but also keeps track of emails. you can send emails using both SMTP and API.

This CLI is developed in generic way. So that, even if you are not using Pepipost you can send email through SMTP using this CLI.

For apiEmail and tracking of the email you need to have [Pepipost](http://www.pepipost.com/?utm_campaign=GitHublinuxagent&utm_medium=Github&utm_source=Githublinuxagent) account.

## Installation

`git clone github.com/pepipost/linux-agent-pepipost $GOPATH/src/`

## Setting Enviroment variable.

We recommend to set SMTP username, Password and Pepipost **APIKEY** has enviroment variable, which can be available after login to [Pepipost](https://app.pepipost.com/) app -> settings -> Integration 

replace the credentials with your account and copy paste below in your ~/.bashrc or ~/.bash_profile or checkout [setting enviroment variable](https://www.digitalocean.com/community/tutorials/how-to-read-and-set-environmental-and-shell-variables-on-a-linux-vps) / [info on setting env](https://codeburst.io/linux-environment-variables-53cea0245dc9)


```shell 
export PEPISMTPPASS="myaccount@345"
export PEPISMTPUSER="smtpusernameispepi"
export PEPIKEY="e8a874fbd6f35fffdsafdsfa230879205"

```

## Usage

* Sending SMTP email using command line interface.

`pepipost smtpEmail -h`

   **OR**

```shell
  
 pepipost smtpEmail 
 
	-d	smtp.pepipost.com 
	-P	25/465/587
	-u	smtpusername                                            #this will override your enviroment variable(not recommended)
	-p	smtpuserpassword                                        #this will override your enviroment variable(not recommended)
	-fn	fromname
	-f	info@yourregisterdomainwithpepipost.com 
	-t	myrecipient1@gmail.com,myrecipient2@gmail.com           #can pass multiple toaddress comma separated 
	-c	myrecipientcc1@gmail.com, myrecipientcc2@gmail.com      #can pass multiple cc address comma separated
	-B	mybcc1@gmail.com,mybcc2@gmail.com                       #can pass multiple bcc address comma separated
	-r 	replytoid@something.com                                 #replytoid for your email
	-s 	"Subject Line :: SMTPMAIL v1.0.0"                       #subject of your email
	-b 	"This is text email body from pepipost."                #body of your email
	-z	"<html><body></body></html>"                            #can pass html as text command line
	-Z 	"/tmp/htmldemo.html"                                    #can pass html file from path  
	-h 	"X-PepipostHeader,DEMO-786-XXX-435"                     #customized header you want to pass
	-a 	"/tmp/mypdfforpepi.pdf"                                 #attachment if any.

```

* Sending API email using command line interface

`pepipost apiEmail -h`

   **OR**

```shell

 pepipost apiEmail

	-K  "kfdjkanf dasjkfi38yirekjbfidsnk"               #this will override your enviroment variable(not recommended)
	-fn	fromname
	-f	info@yourregisterdomainwithpepipost.com
	-t	recipient1@gmail.com 
	-c	rcc1@gmail.com,rcc2@gmail.com                   #can pass multiple cc address comma separated
	-B	bcc@gmail.com
	-r 	replytoid@something.com
	-s 	"Subject Line :: APIMAIL v1.0.0"                #subject of your email
	-b 	"This is text email from pepipost."             #body of your email
	-z	"<html><body>mybody</body></html>"              #can pass html as text command line 
	-Z 	"/tmp/htmldemo.html"                            #can pass html file from path 
	-h 	"X-PepipostHeader,DEMO-786-XXX-435"             #customized header you want to pass
	-a 	"/tmp/mypdfforpepi.pdf"                         #attachment if any can pass multiple

```
* Fetching logs for email activity


