#!/bin/bash

set -euxo pipefail

find . -name "*.go" -type f | entr -r dagger -s call dev --src=. up