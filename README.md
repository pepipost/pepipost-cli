# Pepipost as Linux Agent :penguin:

### Introduction

  Developers always :heart: the freedom of working on open source and code on terminal. So why not make this combo easy for all our developer friends who love to perform task by just using single command on cli. 
  
  Taking into consideration Pepipost is trying to make community which will help users to easily integrate this open source project on their Linux servers. So that Sending mail and keeping analytics becomes easier. 

#### What is the Challenge ?

  You just need to contribute a linux Agent which will help clients with any 4 functionality from the following
  
  * Command line interface for
  
    * Sending mail
    * Extracting Reports
    * Adding Domain
    * Whitelisting and Blacklisting Emailids
    * Template addition / Deletion 
    * Subadmin 
  
#### How can i start my challenge ?

  Using the below information you can easily get started with your challenge 
  
  * You can just have read through this great [article](https://akrabat.com/the-beginners-guide-to-contributing-to-a-github-project/) which will help you in contribution to this project, followed by our [CONTRIBUTING.md]()  
  
  * [Sign-up](https://app.pepipost.com/index.php/signup/registeruser) to pepipost for apikey which will be required for your [Utilities](#utilities), activate your account and send your first test mail using our sandbox domain.
  
  * Pepipost APIs from our [developers documentation](https://developers.pepipost.com) (This includes all the open APIs which is describe above in challenge). 
  
  * SDK from [Github reposistory](htt.ps://github.com/pepipost)(SDKs has only Email sending functionality)
  
<a name="#utilities"></a>
#### Referenes/Utilities

  * For Sending email cli interface
  
  ```bash
   pepipost -sendemail -key 'Secret apikey here' -f 'from@mypepi.com' -t 'recipient@gmail.com' -s 'Pepipost Linux Agent' -b 'Hi, this is pepipost Agent for sending email' -a "myattachmentfile.pdf" 
  
   options are
    -key apikey 
    -s Subject
    -f from
    -t to
    -B bcc
    -b Body (should accept both html and text)
    -a file to attach 
    -h X-apiheader
    
  ```
  
  * For Extracting Report cli interface 
  
  ```bash
    pepipost -fetchreports -key 'secret apikey here' -events 'open,click,unsub,dropped' -sdate 'YYYY-MM-DD' -edate 'YYYY-MM-DD' -t 'recipient@gmail.com'
  
    options are
    -key apikey
    -events comma separated or any delimiter for appending flags in api
    -sdate startdate
    -edate enddate
    -t toemailid
  ```
  
   Similary more cli can be made that totally depends on you.

#### What are requirements to get my pull request(PRs) accepted ?

  * Brief introduction of your challenge and making.
  * Prerequisites.
  * Installation Guide.
  * Example
  * Usage cases for each functionality.
  * Blog representing you hard work.
  
  You can publish your challenge by creating a pull request here in this reposistory.
  
### cheers :beer:

### All the Best :thumbsup:

## Happy Coding...! :tada:
  
  
  
