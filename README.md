# gorelease
[![Build Status](https://travis-ci.org/gorelease/gorelease.svg?branch=master)](https://travis-ci.org/gorelease/gorelease)
[![gorelease](https://dn-gorelease.qbox.me/gorelease-download-blue.svg)](http://gorelease.herokuapp.com/dn-gobuild5.qbox.me/gorelease/master)

For easily build go cross platform online and share your binary.

With this project help you will have a download page. ex: <http://gorelease.herokuapp.com/gorelease/gorelease>

## Not stable
This project is not stable for now.

## How to use gorelease
It is very simple to use gorelease. Before doing everything, make sure this repo have integrate with <https://travis-ci.org>

Just add two lines to your `travis.yml`.

	after_success:
	  - bash -c "$(curl -fsSL http://bitly.com/gorelease)" gorelease

The simplest `travis.yml` is like

	language: go
	go:
	  - 1.4
	script:
	  - go test -v ./...
	after_success:
	  - bash -c "$(curl -fsSL http://bitly.com/gorelease)" gorelease

Go to this link <http://gorelease.herokuapp.com/token> to get your personal token.

Add the token to travis env setting page.

	GORELEASE_TOKEN=grlkalsjdfads....

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
