#!/bin/bash
printf "GO Version:: `go version`\n\n"

printf "Initializing Linux Agent for Sending Email\n\n"

printf "Compiling....... \n\n"
`go get github.com/urfave/cli`
printf "Compiled file 1..\n"
`go get github.com/briandowns/spinner`
printf "Compiled file 2..\n"
`go get github.com/olekukonko/tablewriter`
printf "Compiled file 3..\n\n"

printf "Compilation Completed\n\n"

`go install cmd/pepipost/pepipost.go`
printf "Installation Completed..\n"


`sudo cp ../../bin/pepipost /usr/bin/`

printf "Setting enviroment for Pepipost..... Done.\n\n"

printf "All Set to Rock\n\n"

printf "Usage :: pepipost -h\n\n"

