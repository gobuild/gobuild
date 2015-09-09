# gorelease
[![Build Status](https://travis-ci.org/codeskyblue/gorelease.svg?branch=master)](https://travis-ci.org/codeskyblue/gorelease)
[![gorelease](https://dn-gorelease.qbox.me/gorelease-download-brightgreen.png)](http://gorelease.herokuapp.com/dn-gobuild5.qbox.me/gorelease/master)

gorelease - for easily public released go binary.

这个项目是用来帮助发布go的二进制文件。

这种方案是我想出来最稳定的一种方法了。完全依赖于各种开源服务,目前已经证明了这种方法是完全可行的。

* 使用[travis-ci平台](https://travis-ci.org) 进行go代码的跨平台编译
* 使用[七牛CDN](http://qiniu.com)来发布编译好的文件
* 另外附加上我写的一些脚本[scripts](scripts). 完全编译的工作
* 一个简单的发布界面。现在被我托管到了heroku平台,比如[这个项目自身的发布界面](http://gorelease.herokuapp.com/dn-gobuild5.qbox.me/gorelease/master)

## Step1
Save the following content into `.travis.yml`, and put it into your repository.

	language: go
	go:
	  - 1.4
	env:
	  - "PATH=/home/travis/gopath/bin:$PATH"
	before_install:
	  - go get github.com/mitchellh/gox
	  - gox -os="linux darwin windows" -build-toolchain
	script:
	  - go test -v ./...
	after_success:
	  - gox -os="linux darwin windows" -output "gorelease-temp/dist/{{.OS}}-{{.Arch}}/{{.Dir}}"
      - bash -c "$(curl -fsSL https://raw.githubusercontent.com/codeskyblue/gorelease/master/scripts/upload-qiniu.sh)"

当前的编译脚本是

	gox -os="linux darwin windows" -output "gorelease-temp/dist/{{.OS}}-{{.Arch}}/{{.Dir}}"

你也可以改成别的，文件最后都是要放到 `gorelease-temp/dist/<os>-<arch>/`下的

## Step2
You need a account in [QiniuCDN](http://www.qiniu.com)

In `travis-ci.org` setting page. Set three environment variables. ex: 例如

	ACCESS_KEY=LKJFLSkdjfkj23lkjrl23kjflkzsjdfljwerf2w3
	SECRET_KEY=kljdlFLSDKFJo9iwejflkjLkjsdfoijw4elfkjsd
	BUCKET=gorelease

`BUCKET`也就是空间地址, 没有空间的话，选择创建一个新的空间就可以了. 不妨把这3个变量找个地方存起来,以后其他项目还能用。不要拷贝我上面写的, 从七牛上面拷贝

BTW, Qiniu CDN cache is 15mins, So you new released app will be refreshed after 15mins.

## Step3
Get download address page.

在七牛的 **空间设置/域名设置** 里面把域名拷贝出来. ex: `dn-gobuild5.qbox.me`

如你的项目名是 `gorelease`, 地址 <http://gorelease.herokuapp.com/dn-gobuild5.qbox.me/gorelease/master> 即为下载地址页面

## Step4
The badge

[![gorelease](https://dn-gorelease.qbox.me/gorelease-download-brightgreen.png)](http://gorelease.herokuapp.com/dn-gobuild5.qbox.me/gorelease/master)

Just change the download link.

	[![gorelease](https://dn-gorelease.qbox.me/gorelease-download-brightgreen.png)](http://gorelease.herokuapp.com/dn-gobuild5.qbox.me/gorelease/master)

## Contribute
All pull request and suggestions are welcomed. Just make sure the you have tested the code.

另外目前的发布界面有点丑，非常期待欢迎前端高手的参与。

Have a good day.
## Thanks
* <https://travis-ci.org>
* <http://qiniu.com>
* <https://github.com/mitchellh/gox>
* <http://buckler.repl.ca/>

## LICENSE
This repository is under [MIT](LICENSE).
