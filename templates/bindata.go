// Code generated by go-bindata.
// sources:
// analysis.tmpl
// forkme.tmpl
// homepage.tmpl
// release.tmpl
// DO NOT EDIT!

package templates

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
	"os"
	"time"
	"io/ioutil"
	"path/filepath"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name string
	size int64
	mode os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _analysisTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x44\x91\x41\x8f\xdb\x20\x10\x85\xef\x91\xf2\x1f\x7c\x03\x2a\x82\x93\x28\x39\x54\x8e\x2b\x25\x6a\xd5\x5b\x7a\x69\x4f\x96\x0f\x53\x3c\x61\x89\x6c\xf0\x02\x5e\x2b\x8a\xf3\xdf\xd7\xd8\xd9\xdd\x0b\x30\x6f\x1e\x9f\xe6\xc1\xc1\x4b\xa7\xdb\xf0\x63\xb9\x48\x12\x7a\xe9\x8c\x0c\xda\x1a\xaa\xb9\xe7\x96\x2b\xee\x38\xf0\x86\xdd\x75\x41\x7e\x5b\xab\x6a\x3c\x1a\xa8\x6f\x41\x4b\xff\xe7\xff\x15\x65\x20\x65\xee\x32\x5d\xb8\x32\x8f\xcb\x30\x7c\xde\x67\xf7\x89\x17\x55\xf1\x9a\xcf\xdb\x30\x14\x25\x13\x6d\xe7\x5f\x28\x38\xd5\x35\x68\x82\x67\x0f\x3e\x35\xeb\x7c\xf3\xcd\x60\x9f\xfc\x84\x80\x94\x65\x90\x7b\x21\x1d\x8e\xc5\xaf\x1a\xa3\x91\x5a\xc6\x23\xb1\x19\x1b\x0a\xc3\x53\xf5\xa7\xdb\x5f\x50\x67\x68\x70\xec\x17\xeb\x32\x03\x01\xfe\x66\x64\xbe\x19\x4f\xde\xc9\x5c\x65\x8d\x68\xc1\x8d\xd6\xb3\xad\x50\x68\xe3\xd1\x85\x13\x5e\xac\x43\x1a\x83\x45\xe4\x83\xd1\x5e\x9b\xca\xf6\xbc\xb2\x72\x9a\x8a\x93\xf9\x4d\x08\x27\x69\xda\xf7\xbd\x50\x53\xf6\x15\x7c\x84\x17\xd2\x36\xe9\x57\x75\xf5\xa3\x53\x01\x61\xd9\x72\x11\x89\x0a\x28\x99\xa7\x27\x3c\x21\xff\x8e\xab\xfd\x7e\xb3\xdd\x6d\xd7\xdf\x57\xbb\x28\x40\x17\xec\xe4\x9d\x9d\x1e\x4d\x15\xe5\x16\x14\xbe\x69\xec\x9f\x98\x43\xfa\xfc\x98\xf7\x00\x00\x00\xff\xff\x6c\x97\x19\x15\xa1\x01\x00\x00")

func analysisTmplBytes() ([]byte, error) {
	return bindataRead(
		_analysisTmpl,
		"analysis.tmpl",
	)
}

func analysisTmpl() (*asset, error) {
	bytes, err := analysisTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "analysis.tmpl", size: 417, mode: os.FileMode(384), modTime: time.Unix(1441776279, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _forkmeTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x54\x51\xd1\x8e\xdb\x20\x10\xfc\x15\xe4\xf7\x0b\x0e\x36\x4b\x9c\xbb\xdc\x63\xdb\xbf\x88\x16\xbc\x38\xe8\x6c\xb0\x00\xab\x4a\xbf\xbe\x6b\x47\x55\x75\x5a\x59\x9e\x5d\x66\x18\x18\x3e\x50\x3c\x32\xf9\x5b\xf3\xa8\x75\x2d\x57\x29\xa7\x50\x1f\x9b\x3d\xb9\xb4\x48\x97\x46\x2a\x5f\x4f\x3b\x6f\x24\xa7\x94\x69\x26\x2c\xd4\x7c\x7e\x84\x65\x12\xa5\x3e\x67\xba\x35\x6b\x2a\xa1\x86\x14\xaf\x02\x6d\x49\xf3\x56\xe9\x5d\xd4\xb4\x5e\x45\xfb\x2e\x72\x98\x1e\xf5\x40\x36\xe5\x91\xf2\x0e\x1b\x51\xb2\xfb\x6f\xe6\x70\x49\xa7\x97\xe3\x56\x28\xbb\x14\x2b\xc5\x7a\x98\x23\x80\x31\xb6\xbd\xb8\x41\x6b\xf4\x97\xbe\x6d\x7d\xdf\x3b\x50\xc3\xc5\xf7\x2d\x99\xf1\x3c\x38\xa7\xad\x1a\x25\x5c\x4c\xcf\xd5\x9a\xae\x43\xe5\x95\xe7\x7f\xa7\x08\xce\x30\xc2\xd9\x20\x78\x60\x6c\x8c\xd9\x67\x1d\x77\xa3\xf2\x60\x60\x30\x3d\xeb\x34\x28\xe6\x2b\x18\x40\x71\x31\x93\x59\x1e\x00\xf6\x99\x65\xbd\xd6\xaf\x55\xb3\x7b\x68\xd6\x71\x77\x36\x83\xf6\x1d\x40\xff\xef\x53\x64\x5a\xf6\x30\x8d\xc0\xb9\xde\x9a\x1f\x29\x7f\x89\x85\x44\x8a\xe2\x67\xa8\xbf\x36\xdb\x88\x11\x2b\xbe\x39\x8c\x29\x06\x87\xf3\xdb\xb7\x04\x4a\x77\xc2\x05\xff\xa4\x88\xbf\xcb\x71\xef\x57\x1a\x32\x07\x6b\x53\x2c\xd2\xf3\x76\x0b\xdd\x8f\x2c\xef\x53\xc6\xe7\x9d\xcf\xc5\x75\x5a\xe3\xc4\x4f\x21\xf1\xf3\x6f\x00\x00\x00\xff\xff\x80\xe2\xbd\x27\xc2\x01\x00\x00")

func forkmeTmplBytes() ([]byte, error) {
	return bindataRead(
		_forkmeTmpl,
		"forkme.tmpl",
	)
}

func forkmeTmpl() (*asset, error) {
	bytes, err := forkmeTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "forkme.tmpl", size: 450, mode: os.FileMode(384), modTime: time.Unix(1441776271, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _homepageTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xac\x53\x4d\x6f\xd5\x30\x10\x3c\xbf\xfe\x8a\xc5\x67\x1a\x8b\x4a\x48\xa8\x72\x7c\xa1\x70\x85\xc3\x43\x88\xe3\x26\xde\xc6\xd6\xf3\x57\x6d\x07\x1a\x55\xfd\xef\xd8\x49\xa0\x8f\x16\x6e\x9c\x9c\xcc\x7a\x76\x66\x77\x12\xf1\xea\xe6\xd3\xfb\xe3\xb7\xcf\x1f\x40\x17\x67\xe5\x85\x68\x07\x58\xf4\x53\xcf\xc8\xb3\x06\x10\x2a\x79\x71\x10\x8e\x0a\xc2\xa8\x31\x65\x2a\x3d\xfb\x72\xfc\x78\xf9\x8e\x35\xbc\x98\x62\x49\x4e\x21\x91\x25\xcc\x24\xf8\x06\xd4\x8a\x35\xfe\x04\x15\xee\x59\x2e\x8b\xa5\xac\x89\x0a\x03\x9d\xe8\xb6\x67\x7c\x0c\xce\x05\xdf\x8d\x39\x37\x11\xbe\xa9\x88\x21\xa8\xa5\x51\x95\xf9\x0e\x46\xf5\x2c\xe2\x44\x4d\xe5\x20\xf4\x9b\x73\x8d\xfa\xd6\xc0\x28\xbf\x92\xad\x9d\xe8\x35\x64\x22\x10\xb8\x77\xd7\xa5\xc4\x7c\xcd\xf9\x64\x8a\x9e\x87\xae\xde\xa8\x7a\x8a\xf2\x69\x19\xec\x4c\xfc\x77\x23\x26\x55\x18\x05\x47\x09\xb7\x21\x81\xab\x30\x18\x5f\x1f\x1d\x16\x53\xcd\x09\x1e\x57\x99\xea\xa6\x9d\xd5\xc4\x95\xbc\x09\x3f\xbc\x0d\xa8\x00\x95\x4a\x94\x73\xf5\x72\xb5\x15\xa3\xbc\x33\xde\xcc\x90\xe7\x41\x05\x87\xc6\x83\xc7\xe6\x8c\xee\xaf\x41\x34\x79\xa9\xfc\xe5\x14\x86\xd9\x58\xf5\xb6\xbb\x1b\xc2\x7d\xe7\xea\x28\x6b\x65\x57\x6a\x4d\x96\x30\xa7\xba\xb5\x18\x5e\xd0\xcf\xe6\xff\x37\x69\x48\xe8\x47\x7d\x4e\x73\x98\x0b\xa5\x97\x9c\xa3\x26\x0f\x4f\x44\x65\x72\x01\xf5\x6c\x3a\x30\x19\xda\xf5\xc3\x1f\xab\x6d\x9b\xfd\xe5\xa5\xd3\x94\xc2\x69\xc6\x18\xd7\x35\xff\x65\xc4\xa7\xbb\x7c\xb3\xc2\xe4\x7f\x68\xd2\x52\x5b\x07\xd9\x33\xe2\x7b\x48\x0f\x0f\x85\x5c\xb4\x58\x08\x58\x4d\xf2\xe4\x88\x41\xf7\xf8\xf8\xac\x82\x1e\xed\x92\x4d\xde\x6b\x3b\x5b\xf0\xed\xfb\xab\x99\xae\x3f\xc3\xcf\x00\x00\x00\xff\xff\x2a\x14\x0f\x14\x1d\x03\x00\x00")

func homepageTmplBytes() ([]byte, error) {
	return bindataRead(
		_homepageTmpl,
		"homepage.tmpl",
	)
}

func homepageTmpl() (*asset, error) {
	bytes, err := homepageTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "homepage.tmpl", size: 797, mode: os.FileMode(420), modTime: time.Unix(1441776291, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _releaseTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x7c\x54\x4d\x6f\xdb\x38\x10\x3d\x3b\xbf\x82\xcb\x35\x16\xf6\xc1\x14\x72\x58\xa0\x68\x64\x01\x69\xd2\xa6\x28\x8a\x38\xc8\xc7\xa1\xa7\x80\x96\xc6\x12\x13\x8a\x54\x48\xca\x89\x21\xf8\xbf\x77\x48\xca\xb2\x9d\x06\xbd\x88\xc4\x9b\xe1\xe3\x9b\x99\x27\xa6\xff\x5c\x2e\x2e\xee\x7f\xdd\x7c\x25\x95\xab\x65\x76\x92\xfa\x85\x48\xae\xca\x39\x05\x45\x3d\x00\xbc\xc8\x4e\x46\x69\x0d\x8e\x93\xbc\xe2\xc6\x82\x9b\xd3\x87\xfb\x6f\xb3\x4f\xd4\xe3\x52\xa8\x67\x62\x40\xce\xa9\x75\x1b\x09\xb6\x02\x70\x94\x54\x06\x56\x73\x9a\xe4\xba\xae\xb5\x62\xb9\xb5\x21\xd7\xe6\x46\x34\x8e\x58\x93\x63\xec\xe9\xa5\x05\xb3\x61\xb5\x50\xec\x09\xc3\x69\x12\xa3\x3e\xcf\x09\x27\x21\x2b\x35\xd2\x02\xb7\x40\x66\xa4\xeb\xd8\x35\xaf\x61\xbb\x4d\x93\x18\x3c\x49\x93\xa8\x2c\x5d\xea\x62\xe3\x0f\x15\x62\x4d\x44\x31\xa7\x0d\x2f\xc1\xdf\x36\x4a\xab\xd3\x6c\x38\x47\x52\x5b\x73\x29\x3d\xf0\xc5\x70\x95\x57\x9e\x2a\x42\xc8\x74\x1a\xf2\x3d\x43\x2e\xb9\xb5\x73\xba\x6c\x85\x2c\x66\x42\xad\x74\xa0\x1a\xa5\x4d\x58\xb0\x04\x67\xb4\x2a\xb3\x2b\x8d\xa7\xe3\x16\x99\x1b\xae\xc2\xd5\xa5\x5e\x83\xb1\x42\x63\xe3\x5a\xf5\xac\xf4\xab\xc2\x2c\x0c\x46\x8a\xa4\xf9\x90\xea\xa1\x29\xb8\x83\xe2\x23\xbe\x36\x84\x66\xdc\xfd\x85\x2f\x4d\x50\x76\xd8\x38\xbe\xf4\x8d\xf1\x21\x67\xfa\x3b\x5c\x95\x5d\x2d\x16\x77\xd8\xb5\xea\x10\x39\xbf\xbd\xf8\x7e\x8c\xfd\xc4\x31\x1e\x23\x8b\x86\x0d\x00\x6e\x22\x63\xd7\x61\xf3\x4a\x20\xec\x36\x8e\xc6\x6e\xb7\xef\x2e\x2c\x7c\x8b\x17\x77\x61\x52\xc5\x11\x78\x6e\x62\xd7\xdf\xc1\xfe\xe6\xf7\x70\xd8\x8c\x52\xde\xfb\x68\xc8\xa2\xd9\x25\x36\x41\x6a\x8e\xed\xe2\x7d\xfa\xee\xe0\x81\x46\x50\x45\xd0\x85\xd0\xae\x27\x83\x3d\x56\x5a\x3b\x30\xfd\x54\x77\x17\x24\x34\xab\x74\x0d\xde\x3a\x3b\xe2\xff\x72\xdd\x6c\xce\x6e\xf4\x2b\x18\x28\xc8\x72\x43\x86\xe4\xca\xb9\xc6\x7e\x4e\x92\x52\xb8\xaa\x5d\x32\xf4\x38\xfa\xbc\x00\xfb\xbc\x59\xca\x16\x92\xc1\xb6\x74\xef\xe0\x9e\x74\x37\xab\x61\x3d\xfa\x21\x90\xa8\xd1\x0a\x94\xb3\x49\x8d\x62\x94\xeb\x97\x3f\x7e\x8f\x61\x37\x1a\x4f\x56\xad\xca\x1d\x3a\x6e\x32\xed\xbc\xea\x35\x37\x64\x1c\x6d\x73\xee\xc8\x9c\x8c\x27\xf4\xdf\xbd\x8b\xa6\x67\x3e\x67\xcc\x4a\x70\x3f\xee\x16\xd7\x93\xa1\x94\x42\xcd\x4a\x1d\x3c\xff\x3f\x7b\x59\xea\x37\x56\x1f\xd4\x81\x3b\xdb\xae\x93\x10\xb6\x28\x06\xed\x3d\x0d\xad\x67\xb6\xcd\x73\xb0\x76\x2f\xc2\x80\x8d\x3a\xa2\x92\xb6\xb9\x17\x35\xa0\x8c\xbe\x90\x56\x89\x37\x9f\xc3\xa2\xa4\x47\x87\xd1\xa8\x29\xe6\xaf\x8c\xae\xaf\xf5\x2b\x1e\x88\x27\x59\x0f\x4c\x76\x49\x43\x69\xcc\xc1\x9b\x9b\xf4\x61\x8c\xf6\x61\xac\x76\xff\x0f\x4e\x63\x92\xbf\xaf\xd4\x8f\x3d\xda\x33\x6d\xfb\x0a\xc0\x18\x6d\x3e\xd4\x8f\x5c\x38\x24\x76\xf0\x10\x4c\x59\x25\x0a\x98\x0c\x0c\x7e\x0d\xdf\x83\xd9\x74\x9d\x83\xba\x91\xa8\x91\x50\xae\xb8\xdc\x58\x61\x29\x61\xe8\xc6\x34\x89\x0f\x15\xbe\x36\xe1\xa5\xfd\x1d\x00\x00\xff\xff\xf3\xa1\x98\xd4\x7a\x05\x00\x00")

func releaseTmplBytes() ([]byte, error) {
	return bindataRead(
		_releaseTmpl,
		"release.tmpl",
	)
}

func releaseTmpl() (*asset, error) {
	bytes, err := releaseTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "release.tmpl", size: 1402, mode: os.FileMode(420), modTime: time.Unix(1441869767, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if (err != nil) {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"analysis.tmpl": analysisTmpl,
	"forkme.tmpl": forkmeTmpl,
	"homepage.tmpl": homepageTmpl,
	"release.tmpl": releaseTmpl,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"analysis.tmpl": &bintree{analysisTmpl, map[string]*bintree{
	}},
	"forkme.tmpl": &bintree{forkmeTmpl, map[string]*bintree{
	}},
	"homepage.tmpl": &bintree{homepageTmpl, map[string]*bintree{
	}},
	"release.tmpl": &bintree{releaseTmpl, map[string]*bintree{
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
        data, err := Asset(name)
        if err != nil {
                return err
        }
        info, err := AssetInfo(name)
        if err != nil {
                return err
        }
        err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
        if err != nil {
                return err
        }
        err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
        if err != nil {
                return err
        }
        err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
        if err != nil {
                return err
        }
        return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
        children, err := AssetDir(name)
        // File
        if err != nil {
                return RestoreAsset(dir, name)
        }
        // Dir
        for _, child := range children {
                err = RestoreAssets(dir, filepath.Join(name, child))
                if err != nil {
                        return err
                }
        }
        return nil
}

func _filePath(dir, name string) string {
        cannonicalName := strings.Replace(name, "\\", "/", -1)
        return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

