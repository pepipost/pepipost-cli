#!/bin/bash
printf "GO Version:: `go version`\n\n"

printf "Installing Pepipost as Linux Agent for Sending Email\n\n"

`go install cmd/pepipost/pepipost.go`

printf "Compiling....... \n\n"

printf "Compilation Completed\n\n"

`sudo cp ../../bin/pepipost /usr/bin/`

printf "Setting enviroment for Pepipost..... Done.\n\n"

printf "All Set to Rock\n\n"

printf "Usage :: pepipost -h\n\n"

