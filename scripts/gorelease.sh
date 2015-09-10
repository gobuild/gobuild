#!/bin/bash -
#
# gorelease.sh: build and publish
#
# shorted url: http://bitly.com/gorelease
# need login to view stats: https://bitly.com/a/stats

set -eu
set -o pipefail

bash -c "$(curl -fsSL https://raw.githubusercontent.com/gorelease/gorelease/master/scripts/build-standalone.sh)" args0 "windows linux darwin"
bash -c "$(curl -fsSL https://raw.githubusercontent.com/gorelease/gorelease/master/scripts/upload-qiniu.sh)"
