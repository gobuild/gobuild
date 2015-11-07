# gorelease
[![Build Status](https://travis-ci.org/gorelease/gorelease.svg?branch=master)](https://travis-ci.org/gorelease/gorelease)
[![gorelease](https://dn-gorelease.qbox.me/gorelease-download-blue.svg)](http://gorelease.herokuapp.com/gorelease/gorelease)

For easily build go cross platform online and share your binary.

With this project help you will have a download page. ex: <http://gorelease.herokuapp.com/gorelease/gorelease>

## How to use gorelease
It is very simple to use gorelease. Before doing everything, make sure this repo have integrate with <https://travis-ci.org>

Add two lines to your `travis.yml`.

	after_success:
	  - bash -c "$(curl -fsSL http://bitly.com/gorelease)" gorelease

After doing that, your `travis.yml` is like

	language: go
	go:
	  - 1.5
	script:
	  - go test -v ./...
	after_success:
	  - bash -c "$(curl -fsSL http://bitly.com/gorelease)" gorelease

Go to this link <http://gorelease.herokuapp.com/token> to get your personal token.

Add the token to travis env setting page.

	GRTOKEN=grlkalsjdfads....

If you want to pack some other stuff. You need to prepare a file `gopack.yml` in your root of you project.

	$ go get github.com/gorelease/gopack
	$ gopack init # generate gopack.yml

Then modify the `gopack.yml` file. It's better to do some check use gopack tool.

	$ gopack pack # package test

After doing all the things. Every time you push to the github, the new builed binaries will get from <http://gorelease.herokuapp.com>

If you repo in **github.com/gorelease/gorelease**, the download page is <http://gorelease.herokuapp.com/gorelease/gorelease>

# 其他说明
这个项目是用来帮助发布在线编译go，以及发布其二进制文件。

这种方案是我想出来最稳定的一种方法了。完全依赖于各种开源服务,目前已经证明了这种方法是完全可行的。

* 使用[travis-ci平台](https://travis-ci.org) 进行go代码的跨平台编译
* 使用[七牛CDN](http://qiniu.com)来发布编译好的文件
* 另外附加上我写的一些脚本[scripts](scripts). 完成编译的工作
* 一个简单的[发布界面](http://gorelease.herokuapp.com/)。托管在heroku平台上

BTW, Qiniu CDN cache is 15mins, So you new released app will be refreshed after 15mins.

## Badge
[![gorelease](https://dn-gorelease.qbox.me/gorelease-download-blue.svg)](http://gorelease.herokuapp.com/dn-gobuild5.qbox.me/gorelease/master)

Just change the download link.

	[![gorelease](https://dn-gorelease.qbox.me/gorelease-download-blue.svg)](http://gorelease.herokuapp.com/your-repo-download-page)

## How to run this project
To run this project you need a redis-server. Addr and Password are read from ENV

	REDIS_ADDR=localhost:6379
	REDIS_PASSWORD=""
	GITHUB_CLIENT_ID=12...
	GITHUB_CLIENT_SECRET=l213.....

Use redis db:0

	$ go build
	$ ./gorelease -debug	

## Design
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
All pull request and suggestions are welcomed. Just make sure the you have tested the code.

另外目前的发布界面有点丑，非常期待欢迎前端高手的参与。

Have a good day.

## Thanks
* <https://travis-ci.org>
* <http://qiniu.com>
* <http://shields.io>
* <https://github.com/mitchellh/gox>
* <https://www.redislabs.com>

## LICENSE
This repository is under [MIT](LICENSE).
