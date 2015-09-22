#!/bin/bash -
#
# gorelease.sh: build and publish
#
# shorted url: <http://bitly.com/gorelease>
# need login to view stats: https://bitly.com/a/stats
#
# ref:
# - [travis-ci environment](http://docs.travis-ci.com/user/environment-variables/)

set -e
set -o pipefail

echo "Is Pull Request: $TRAVIS_PULL_REQUEST"
# Set environment variables
GORELEASE_GO_VERSION="1.5"
BUILD_OS=${1:-"windows linux darwin"}
TMPDIR=$PWD/gorelease-temp
BRANCH=

if test -z "$TRAVIS"
then
	# Here for my test
	GORELEASE_TOKEN=12345678
	BRANCH=master
else
	#ACCESS_KEY=${ACCESS_KEY:?}
	#SECRET_KEY=${SECRET_KEY:?}
	#BUCKET=${BUCKET:?}
	GORELEASE_TOKEN=${GORELEASE_TOKEN:-"$GRTOKEN"}
	GORELEASE_TOKEN=${GORELEASE_TOKEN:?}
	BRANCH=${TRAVIS_BRANCH:-$TRAVIS_TAG}
fi
KEY_PREFIX=/gorelease${PWD#*/src/github.com}/${BRANCH:?}/

echo "Branch: $BRANCH"
echo "KeyPrefix: $KEY_PREFIX"

if test -n "$TRAVIS" -a "X$TRAVIS_GO_VERSION" != "X$GORELEASE_GO_VERSION"; then
	echo "Expect go$GORELEASE_GO_VERSION, but travis got go$TRAVIS_GO_VERSION"
	exit 0
fi

# Set build environment
if test -n "$TRAVIS"
then
	go get github.com/mitchellh/gox
	if test $GORELEASE_GO_VERSION != "1.5"
	then
		gox -os="${BUILD_OS}" -build-toolchain
	fi
	go get github.com/gorelease/qsync
else
	BUILD_OS="darwin"
fi


/bin/mkdir -p $TMPDIR
DISTDIR=$TMPDIR/dist

# FIXME(ssx): need support build pack
# build standalone
if test -f .gopack.yml
then
	go get github.com/gorelease/gopack
	gopack all \
		--output "$DISTDIR/{{.OS}}-{{.Arch}}/{{.Dir}}.zip" \
		--json "$DISTDIR/builds.json"
else
	gox -os "$BUILD_OS" -output "$DISTDIR/{{.OS}}-{{.Arch}}/{{.Dir}}"
cat > $DISTDIR/builds.json <<EOF
{
	"update_time": $(date +%s),
	"go_version": "$GORELEASE_GO_VERSION",
	"commit": "$TRAVIS_COMMIT"
}
EOF
fi

cat > $TMPDIR/conf.ini <<EOF
[qiniu]
uphost = http://up.qiniug.com
bucket = ""
accesskey = ""
secretkey = ""
keyprefix = $KEY_PREFIX

[local]
syncdir = $DISTDIR

[gorelease]
token = "$GORELEASE_TOKEN"
host = "qntoken.herokuapp.com"
EOF

set -eu

# upload
qsync -c $TMPDIR/conf.ini
