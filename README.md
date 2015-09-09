# gorelease
[![Build Status](https://travis-ci.org/codeskyblue/gorelease.svg?branch=master)](https://travis-ci.org/codeskyblue/gorelease)
[![gorelease](https://dn-gorelease.qbox.me/gorelease-download-brightgreen.png)](http://gorelease.herokuapp.com/7xln6q.dl1.z0.glb.clouddn.com/gorelease/master)

gorelease - for easily public released go binary.

这个项目是用来帮助发布go的二进制文件。

这种方案是我想出来最稳定的一种方法了。完全依赖于各种开源服务,目前已经证明了这种方法是完全可行的。

* 需要依赖 [travis-ci平台](https://travis-ci.org) 这个搞不好需要翻墙才行。
* 依赖[七牛](http://qiniu.com) 还有我写的qiniu上传工具 <https://github.com/codeskyblue/qsync>
* 另外附加上我写的一些脚本. 就在这个项目的[scripts](scripts)目录下
* 一个简单的发布界面。我托管到了heroku平台,比如[这个项目自身的发布界面](http://gorelease.herokuapp.com/7xln6q.dl1.z0.glb.clouddn.com/gorelease/master)

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

## Step2
You need a account in [QiniuCDN](http://www.qiniu.com)

In `travis-ci.org` setting page. set three env vars (Copied from qiniu). for example

`BUCKET`也就是空间地址, 没有空间的话，选择创建一个新的空间就可以了.

	ACCESS_KEY=LKJFLSkdjfkj23lkjrl23kjflkzsjdfljwerf2w3
	SECRET_KEY=kljdlFLSDKFJo9iwejflkjLkjsdfoijw4elfkjsd
	BUCKET=gorelease

## Step3
Get download address page.

在七牛的 **空间设置/域名设置** 里面把域名拷贝出来. ex: `7xln6q.dl1.z0.glb.clouddn.com`

如你的项目名是 `gorelease`, 地址 <http://10.240.187.174:4000/7xln6q.dl1.z0.glb.clouddn.com/gorelease/master> 即为下载地址页面

Good luck.

## Step4
The badge

[![gorelease](https://dn-gorelease.qbox.me/gorelease-download-brightgreen.png)](http://gorelease.herokuapp.com/7xln6q.dl1.z0.glb.clouddn.com/gorelease/master)

Just change the link.

	[![gorelease](https://dn-gorelease.qbox.me/gorelease-download-brightgreen.png)](http://gorelease.herokuapp.com/7xln6q.dl1.z0.glb.clouddn.com/gorelease/master)

## Contribute
All pull request and suggestions are welcomed. Just make sure the you have tested the code.

另外目前的发布界面有点丑，非常期待欢迎前端高手的参与。

## LICENSE
[MIT](LICENSE)
