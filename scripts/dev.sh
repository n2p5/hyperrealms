#!/bin/bash

set -euxo pipefail

cd tmp
touch blah5.txt
ls -alh
sleep 10
ls -alh
sleep 5
# find . -name "*.go" -type f | entr -n -r go run .