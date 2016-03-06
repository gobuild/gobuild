# gobuild web
[![Build Status](https://travis-ci.org/gobuild/gobuild.svg?branch=master)](https://travis-ci.org/gobuild/gobuild)
[![gobuild-web](https://dn-gorelease.qbox.me/gorelease-download-blue.svg)](https://gobuild.io/gobuild/gobuild)

For easily build go cross platform online and share your binary.

With this project help you will have a download page. ex: <https://gobuild.io/gobuild/gopack>

## How to use
1. Open browser <https://gobuild.io>
2. Enter your repository name in the web.
3. Wait for some minute, the binary will be packaged done.

## Adanvanced
1. Define how to the build process.

	You need install `gopack` before doing anything.

	```
	go get -u -v github.com/gobuild/gopack
	# generate .gopack.yml
	gopack init
	```

	Picke some text editor(ex: Vim), Change the script part to something else(ex: `go build -tags hello`)

2. Test package in local

	I'm going to use pack a repo, then you will know how it works.
	```
	$ go get github.com/gobuild/gopack
	$ cd $GOPATH/src/github.com/gobuild/gopack
	$ gopack pack -o dist.zip
	Packaging ...
	Done
	$ unzip -t dist.zip
	Archive:  dist.zip
		testing: gopack                   OK
		testing: README.md                OK
		testing: LICENSE                  OK
	No errors detected in compressed data of dist.zip.
	```
3. Add badge to your repo readme.

## Badge
[![gorelease](https://dn-gorelease.qbox.me/gorelease-download-blue.svg)](https://gobuild.io/gobuild/gobuild)


## For developer
First [generate github token](https://help.github.com/articles/creating-an-access-token-for-command-line-use/)

To run this project you need a redis-server. Addr and Password are read from ENV

	REDIS_ADDR=localhost:6379
	REDIS_PASSWORD=""
	GITHUB_CLIENT_ID=12...
	GITHUB_CLIENT_SECRET=l213.....
    GITHUB_TOKEN=...
    MYSQL_URI=...

Use redis db:0

    $ bower install
	$ go build
	$ ./gobuild

Redis storage.

	> GET user:codeskyblue:github_token
	# github token

	> GET user:codeskyblue:token
	# web token, need to set in travis, ex
	grABCDEFG

	> SMEMBERS token:grABCDEFG:orgs
	# list token orgnization, which org can upload use this token
	1) "codeskyblue"
	2) "gorelease"

	> HGETALL orgs:codeskybule:repos
	# list repos under org, and the store domain
	1) "gosuv"
	2) "dn-gobuild5.qbox.me"
	3) "syncgit"
	4) ""

	> GET downloads:codeskyblue/gosuv
	# total number of downloads

	> GET downloads:codeskyblue/gosuv:linux-amd64
	# total number of download linux-amd64 binary

	> GET pageview:codeskyblue/gosuv
	# download page PV
	
redis data backup can use: <https://github.com/p/redis-dump-load>, or see a [script](scripts/redisdl.py)

## Contribute
Fix typo is very welcome.

Have a good day.

## Thanks
* <https://travis-ci.org>
* <http://qiniu.com>
* <http://shields.io>
* <https://github.com/mitchellh/gox>
* <https://www.redislabs.com>

## LICENSE
This repository is under [MIT](LICENSE).
