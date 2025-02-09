#!/bin/bash

set -euxo pipefail

find . -name "*.go" -type f | entr -r dagger --progress plain call dev --src=. up