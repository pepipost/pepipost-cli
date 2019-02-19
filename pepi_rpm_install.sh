#!/bin/bash
printf "GO Version:: `go version`\n\n"

printf "Pepipost as Linux Agent for Sending Email & Fetching Logs\n\n"

`cd /tmp/ && wget https://pepipost.com/Pepipost-1.0.0-1.x86_64.rpm`

printf "Installing RPM \n"

`rpm -Uvh Pepipost-1.0.0-1.x86_64.rpm`

printf "Installation Completed\n"

printf "Checking executing.. \npepipost --help OR pepipost -h"



