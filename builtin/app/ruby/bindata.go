// Code generated by go-bindata.
// sources:
// data/aws-simple/build/build-ruby.sh
// data/aws-simple/build/template.json.tpl
// data/aws-simple/deploy/main.tf.tpl
// data/aws-vpc-public-private/build/build-ruby.sh
// data/aws-vpc-public-private/build/template.json.tpl
// data/aws-vpc-public-private/deploy/main.tf.tpl
// data/common/dev/Vagrantfile.tpl
// DO NOT EDIT!

package rubyapp

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

func bindataRead(data, name string) ([]byte, error) {
	gz, err := gzip.NewReader(strings.NewReader(data))
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

var _dataAwsSimpleBuildBuildRubySh = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x56\x7f\x53\x1b\x39\x12\xfd\x7f\x3e\x45\xaf\xa1\xc2\x5d\x15\x9a\x01\xea\xf6\x6e\xd7\x5e\xb6\x0e\x58\x48\xa8\x4a\x99\x94\x49\xee\x47\x25\x39\x6a\x3c\xd3\xb6\x15\xc6\xd2\x44\xd2\x18\x0c\xf1\x77\xbf\xd7\x1a\xe3\x1f\x21\x9b\xca\x3f\x20\x4b\xea\xd6\xeb\xee\xd7\xaf\x67\xe7\xa7\x6c\xa8\x4d\x36\xcc\xfd\x24\x49\x3c\x07\x52\x96\x8c\x6d\xcc\x72\xc9\xce\xf1\xbd\x8e\xcb\x5a\xd7\x3c\xca\x75\xb5\xdc\x0e\x2e\x2f\x38\x49\xb0\xb2\xee\x2f\x7f\xa5\xc7\x84\x88\x2a\x5b\xe4\x15\x79\xdb\xb8\x82\x47\xba\xe2\xe3\xdd\xc3\xf5\x76\xa5\x0d\x1b\x7b\xbc\x7b\x24\x5b\x5c\x4c\x2c\x75\xce\x07\x83\xab\x01\xe5\x81\x76\x1f\xd7\x46\x8b\xee\xee\x63\x7b\x77\xd1\xa3\xd7\xb9\x0f\xb0\x1f\xfb\x6e\x47\xcc\xc6\x8e\x6b\xb2\x21\x58\xca\x66\xb9\xcb\x70\x90\xf9\xb9\xc7\x3f\xfa\x42\x21\x62\x33\x74\x74\x90\x2c\x12\xa0\xab\x69\x2f\x82\xa3\xce\xee\xe3\xe9\xc9\xf5\xab\x9b\xeb\xab\x77\x83\xb3\xf3\x45\x47\x36\x5e\x5f\xf6\xcf\xfb\x57\x8b\xce\x1e\x01\x43\x92\x58\x96\x10\x70\xf0\xcf\x0e\x1d\xfd\xfe\xe2\x10\xee\xe0\x74\xcc\x8e\x54\x68\xdf\xfb\x9d\xb2\x92\x67\x99\x69\xaa\xaa\x47\x8b\xc4\x56\xd1\xa0\x0d\xe3\xbd\xdc\xf8\x48\x30\x96\xa3\x64\x87\x8a\xca\x36\xa5\x2a\xac\x19\xe9\x31\x15\xb9\x21\x6d\x02\xbb\x11\x3b\xa6\x3b\x1d\x26\x94\xd7\x81\x0a\x3b\x9d\xe6\xa6\xf4\xa4\x47\xa4\xc3\x9e\x27\x1f\x74\x55\xe1\x26\xd5\xce\x22\x4e\xef\xf1\x08\x75\xfe\x9d\xeb\xa0\xcd\x98\x46\x08\x64\xcb\x2d\x30\xc1\x45\x5d\x71\xe0\x34\x4d\x3b\x49\x63\x60\x4f\xef\xdf\x93\x1a\x2d\x93\xa3\x87\x59\xb4\xc8\xb4\xf1\x21\x37\x05\x67\x43\x6b\x83\x1a\x69\xa3\xfd\x84\x4b\xfa\xf8\xb1\x47\xa5\x45\x5a\x7d\xc5\x48\xeb\x41\xfa\x73\x52\x5a\x83\x9a\xca\xbb\x27\x65\x29\xcf\x0a\x52\x60\xa9\xad\xd7\xc1\x3a\xcd\x9e\x80\x99\x9a\xba\xcc\x05\x55\x7c\xd8\x32\xf9\xa6\xb4\x72\x55\x8d\xc1\x9a\x78\xc8\xcf\xb6\x23\x08\x04\xa8\xe6\x54\xcf\xc3\xc4\x1a\xe5\xed\x28\xdc\xe5\x8e\x15\x02\xae\xd9\x05\xf1\xfe\x8d\x3d\x25\x99\xb2\x26\x3a\x42\x59\x0d\xc0\xb8\xa0\x26\x21\xd4\x7e\xfd\x48\x59\x2a\x39\x07\x3d\x5a\xa4\xf3\xf8\x4e\x9d\x77\x8b\x89\xd3\x5e\x55\x9c\x67\xc6\x96\x9c\x7e\xf2\x5b\xc0\xc4\xee\xb9\xcd\xd0\xe9\xf1\x24\x0c\xed\x7d\xe6\x9a\xe1\x5c\x99\xf1\x96\xcd\x2d\xcf\xf1\xde\x8c\x94\xac\x3c\xbb\x19\x58\x32\xb9\xad\xbb\x59\xb6\xfa\x9d\x36\x43\x94\xa3\x49\x81\xbc\xfb\xcb\x01\x6e\x3a\x2e\x66\xf1\x3a\xfd\xfc\xf7\xc3\x8b\x5f\x4f\x7f\x3d\x3b\x39\xfb\xdb\xc1\xe9\xd1\xc5\x3f\x92\x48\xa1\xbd\x92\x87\x14\x43\x82\x1b\xeb\xbd\x42\x4f\xe6\x92\xee\xb4\x9e\x34\x5e\x5b\x53\xe7\xde\xb3\x01\x21\xc5\x67\x06\x18\xd9\x6a\x87\x82\x6b\x7c\x98\xd3\x34\xd7\x66\x0f\xc4\x8d\x40\x03\x33\x65\x1c\x8a\x78\xb5\xed\x2d\x9f\x56\xda\x87\xb4\x5c\x5b\xc6\x8d\x4d\x66\xff\x59\x2d\x13\xbe\x97\xa4\xd3\xe0\xdd\xe9\x7f\x6f\xfe\x75\x3e\xb8\xbe\xbc\xea\x1f\x77\x1e\x1f\x49\xf2\x73\x83\x80\x05\x22\x2d\x16\x9d\x96\x3b\x97\x6d\xa9\x85\x3f\x03\x5c\xd8\xa7\x37\x4f\x2f\xee\x53\x7f\xac\xcd\xfd\x7e\x64\x91\x0d\x13\xa0\xaf\xf3\xe2\x36\x1f\x03\x9d\x70\x69\xf9\xce\x1f\xe7\xa7\x97\x27\xfd\x9b\x8b\xc1\x55\xff\xed\x79\xff\x8f\x63\x63\x4d\xec\xa0\xbc\x08\x7a\xf6\x5d\x6a\x0d\x1f\x1c\x8d\xa1\x57\x53\x76\x45\xe3\x34\x64\x67\xd8\xe8\xaa\x54\x2c\x00\x82\xfc\xfe\x00\xc2\xa3\x35\xea\xcf\x0a\x51\xd3\x03\x96\x87\xe3\xb8\xfc\x0e\xf5\xc4\x46\xe8\xf3\xc9\xc7\xa5\x44\xbd\xbb\x99\x8b\xe7\x3b\xd1\x63\x34\x93\x78\x15\xdf\x83\xb9\x9e\x56\x99\x7f\x96\xa7\xd3\xc6\x94\x15\x2a\xb2\xd9\x50\x63\x9e\xae\x42\x1b\xb6\xe7\xa0\x92\xb1\xca\xe9\xe5\xff\xd2\x16\xad\xa7\x73\xf1\x5f\x84\xb6\x63\xeb\xe8\x25\xba\x98\xde\x96\x1a\x46\x35\x65\xde\xcd\x32\x91\x29\xb4\x49\xdd\x9e\x85\xdc\xd1\xc3\x3d\xc4\x22\x4c\xeb\xd5\x51\x1a\xc6\x0f\xa4\xce\xbe\xba\xbf\xad\x08\x75\xa5\x0b\xb4\x3f\xf2\xd2\xf8\xaf\x20\xa3\x9f\x64\x0f\xf0\x4a\xed\xf3\x61\xc5\xa5\x92\x98\xef\xac\x2b\xb1\x37\xe6\xc2\x7a\xea\x74\x68\xdb\xf1\x35\x87\x88\x1c\x49\x9f\x6a\x2f\x54\xf2\x5b\x4e\xd1\x20\x77\x86\xd4\x60\x65\xd6\xfd\x16\xbc\xb3\x28\x8b\xa8\x39\x3c\xc5\xa4\x47\x1f\x10\xe3\xb7\x13\x0d\x91\xf5\x90\xb1\xcf\x8d\x76\xd0\x3d\x91\xd2\x8d\xee\x91\x44\x07\xca\x71\x9e\x7b\x6b\x04\x34\xb1\x99\x69\x67\xcd\x14\x94\xa1\xbb\x89\xc8\x36\x28\x05\x1d\x87\x37\xa8\x67\x49\x7c\xcf\x45\x13\xe4\xaa\x07\x97\x6e\xd1\x6a\x8d\x77\x71\x8e\xc2\x72\x7f\xfd\x0b\x14\xac\xf6\x09\x6d\x98\xd2\x25\x9e\xa8\xbc\x70\xb6\x06\xc7\x4c\xa8\xe6\x70\x66\x98\x31\x00\x80\xc0\x16\xb8\x4a\x13\xc8\x8e\x0c\x00\xf4\x05\xb5\x2a\x9f\xd2\x49\x5d\xb3\x89\x89\x07\x04\x09\xc4\xf8\x66\x34\xd2\x85\x86\x8f\x94\xba\xea\x4b\x5b\x4c\x8f\xb8\x94\xa6\xbd\x43\x9f\xfd\x4f\x40\xd0\x9b\x93\xb7\xaf\x7a\x1f\x4c\xb6\xd7\xca\x40\xcc\x48\xfb\x37\x15\xd7\xdf\xb0\xda\xa1\x2b\x24\xb4\x4b\x32\xf9\xc5\x1a\x0d\xb1\x91\x26\x99\x62\x1e\xa2\xf2\x24\x4c\xdf\x71\x8d\xc0\xfa\x08\x4c\xe2\x72\x3c\xb5\x33\x46\x40\x5a\xa4\xbd\xad\x8b\x24\x1a\x51\x43\x93\x08\xb2\xcb\x2d\x12\x37\xdd\x74\x26\xfb\x5e\x71\x2c\x46\x09\x75\x1a\xe5\x4d\x15\xda\x5a\xb2\xe7\xf8\x25\x81\x41\x84\xb2\xd4\x98\x4d\x52\x24\xf4\x96\xb4\x2a\x96\xfe\x29\x81\x2b\xe8\x6a\xa9\x30\x25\xad\x31\xee\xa3\xa1\x02\xfc\xc5\xc1\x8b\xba\xeb\x96\x08\x25\x3a\x1f\x4c\xf0\x8c\x1a\x41\xf7\xe8\x69\xd4\x4e\x10\x7c\x68\xd3\x65\x1b\xbc\x06\xe6\x9b\xe5\x7b\x69\x82\x5e\xa0\xdf\x7e\xeb\xbf\xbc\xec\xff\xe7\xec\xaa\x7f\xf1\x4c\x81\xdb\x90\xc4\xd5\x96\xf6\xca\xc6\x96\xf6\xee\xd0\x4b\x36\x2c\xef\x96\x34\x9c\xc7\x62\x24\xab\xeb\x37\x0e\xa3\xbb\x25\x96\xcc\x75\xd1\x9b\x6c\x06\x62\x58\x9c\xc8\x7a\x39\x26\x6e\x56\x06\x99\x7c\x74\x85\xd8\x4b\x98\xf7\xbd\x4d\x4f\xb8\xbf\xa6\xe8\x7a\x7f\xe4\x98\xe3\x61\x2f\x59\x05\x93\xfc\x60\x74\xdb\x05\x5b\x89\xc9\x0f\xc5\xb8\x1c\xa1\xf1\x03\x92\x64\x22\xb1\xa1\x5f\x0e\x7a\xf1\x67\x1b\xf5\x66\xb3\x67\x75\x33\x84\xfe\xb4\xc7\x6b\xf0\xcb\xa7\xc9\x9a\x1e\xbe\xfe\x36\xf0\x8b\x34\xb4\xe2\xfa\xa4\xa5\xd2\x4b\x42\x90\x2d\xa1\x54\xcd\x4a\x5d\xa4\x23\x56\xdf\xc4\xa4\xaa\x82\x3a\x45\xb9\x0d\x82\x5e\xbc\x58\x2a\xf2\x7a\xf6\x40\xf1\xeb\xca\xce\xa3\x66\x28\x25\x9f\x79\x42\x15\x44\xce\x95\xad\xe3\x2e\x52\x14\x96\x13\x12\x0f\xcb\xa7\xd6\x4f\x9d\xe4\xff\x01\x00\x00\xff\xff\x0e\xa3\x8b\x8f\x80\x0b\x00\x00"

func dataAwsSimpleBuildBuildRubyShBytes() ([]byte, error) {
	return bindataRead(
		_dataAwsSimpleBuildBuildRubySh,
		"data/aws-simple/build/build-ruby.sh",
	)
}

func dataAwsSimpleBuildBuildRubySh() (*asset, error) {
	bytes, err := dataAwsSimpleBuildBuildRubyShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/aws-simple/build/build-ruby.sh", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _dataAwsSimpleBuildTemplateJsonTpl = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x9c\x54\xcd\x8e\x9b\x30\x18\xbc\xf3\x14\x16\x52\xf6\x14\x20\xed\xae\xaa\xaa\xd7\x3e\xc6\x2a\x62\x0d\x78\x83\x15\xdb\x58\xfe\xec\x54\x59\xe4\x77\xef\x67\xfe\xa9\x36\x24\xdd\x5c\x8c\x3c\xc3\xcc\x78\xf8\x9c\x36\x22\xf8\x8b\x25\x57\xb9\xa6\xe5\x99\x99\xfc\xc2\x0c\xf0\x46\xc5\xbf\x48\x7c\x48\x7f\xa6\x87\x78\x1f\xf5\x9c\x0b\x35\x9c\x16\x82\x01\x42\xfd\x6b\xb8\x49\xff\x40\x4e\xcb\x92\x01\xe4\x67\x76\x45\x44\x39\x21\xf6\x4b\x14\x58\x69\x98\xbd\x85\x1a\x76\xea\xcd\x56\x08\x08\x77\xc2\x3c\xb6\x1e\x80\x6e\xdf\x8f\x41\xb4\x69\x2e\x3c\x64\xc4\xa4\x48\x78\x1d\xde\x6a\x77\xe4\xbd\x31\xa4\xe2\x86\x70\x85\x8f\x4e\x55\xd4\x22\x2b\xc7\x1d\x48\x0b\xc7\x45\x45\x76\x7e\x24\x0f\x2b\xca\xd9\xab\x66\xe1\xb4\x50\x33\x21\xe2\xfd\x0c\x70\x25\xb8\x0a\xd0\x6b\x2c\xcf\x41\x36\xd1\x24\xb3\x52\x67\x8d\xb5\x4d\x36\x1b\x24\x6d\x1b\x9c\x45\xd3\xe8\xf4\x37\xee\x5a\x66\x88\xf7\xf1\x71\x50\xf2\xfb\xdb\x9e\xef\x5c\xb0\xa5\x25\x34\xce\x94\x1d\x82\x9a\xc1\xd2\xfb\x6c\x89\x57\x0c\x2c\x57\x9d\x6b\x20\xfd\x47\x9a\x07\xc2\x6c\x15\x50\x56\x8f\x1e\xdd\x7b\xf2\xf4\x44\x0a\x0a\x35\x49\x33\x49\xb9\x4a\xa1\xfe\xa4\x8b\x1d\x61\xaa\x0a\xdf\x6b\xeb\x93\x6c\xd4\xb3\x23\x38\xa8\x05\x86\x90\xa8\x80\x29\x1c\xe0\x39\xdf\xa6\xc1\x79\xc3\x33\xf7\x1e\x0b\xda\x23\x4d\x26\x54\xeb\xd4\x9e\x3e\xbe\x54\x18\x94\x86\x6b\x1b\xa0\x6e\xdc\x12\xe3\x8a\x6b\x38\xfe\xa8\xd5\xad\xc7\x71\x8e\x3b\xce\x30\xc3\xd3\x85\x52\x54\x76\xda\x21\xcb\x24\x3d\x39\x52\x49\x3f\xb0\x75\x56\xc0\x8c\xad\xae\xdf\xad\x62\xd6\xf7\x74\xbb\x9d\x78\x75\x65\xb7\x14\x67\xe2\x1d\xc5\xe9\x9a\x6f\xa9\xf5\xa4\x7b\xd9\xba\x11\xc8\xa9\xe4\x7d\x1f\x3c\xf9\xfe\xed\xc7\xf3\xa1\x7a\x79\x99\x39\x5c\x81\xa5\x0a\x59\x63\x6d\xe5\x73\x2a\xa8\x39\xb1\x85\x0c\xd4\x79\x70\x1e\xeb\x76\x05\x0e\xaf\x5b\x94\x2a\x79\x3e\x62\x6d\x1b\x9e\x70\xae\xff\xcd\x8e\x2b\x4e\x11\x95\xfa\xb3\xc4\xfd\x7f\xd6\x31\x8a\x7c\xf4\x37\x00\x00\xff\xff\x16\x08\x5f\x5b\x65\x05\x00\x00"

func dataAwsSimpleBuildTemplateJsonTplBytes() ([]byte, error) {
	return bindataRead(
		_dataAwsSimpleBuildTemplateJsonTpl,
		"data/aws-simple/build/template.json.tpl",
	)
}

func dataAwsSimpleBuildTemplateJsonTpl() (*asset, error) {
	bytes, err := dataAwsSimpleBuildTemplateJsonTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/aws-simple/build/template.json.tpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _dataAwsSimpleDeployMainTfTpl = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xbc\x52\xc1\x8e\xdb\x20\x10\xbd\xe7\x2b\x46\x74\x8f\x5d\x27\xed\xb1\x52\xcf\xbd\xb5\x1f\x50\xad\x10\xc6\x24\x45\x8b\x01\xc1\x90\xca\xb2\xfc\xef\x1d\xa0\x16\xc1\xbb\xbd\x36\x39\xf9\xcd\x63\xde\xcc\x9b\xf7\x01\xbe\x29\xab\x82\x40\x35\xc1\xb8\xc0\x0f\x44\xf7\x11\x26\x07\xd6\x21\xa8\x49\x23\xcc\xc2\x26\x61\xcc\x72\x3a\xdd\x45\xd0\x62\x34\x0a\x98\xb6\xd7\x20\xb8\x9e\x18\xac\xdb\x03\x2c\x7e\x47\x2e\xa4\x54\x31\xf2\x57\xb5\xbc\x53\x8c\x4a\x06\x85\xff\x28\x06\x75\xd3\xce\x1e\x0a\x44\xe5\x56\xcc\xaa\xc0\x8f\x0f\x66\x7d\x60\x6a\x1b\x51\x58\xa9\x38\x2e\x3e\xd3\x61\x52\x57\x91\x0c\xc2\x57\x60\xf8\x79\x98\xb5\x0c\x8e\xc1\xe3\x8b\x98\x46\x4b\xd3\xf8\x34\x1a\x2d\x0f\xdd\xee\x5e\x72\xa9\xa7\xf0\x0e\xfc\x77\xed\x93\x0f\xee\xae\x27\x15\xca\xf4\x04\x9d\x00\xda\xf2\x59\xf5\x69\xa5\x87\x43\x6f\xca\xc6\x88\xd6\x6c\xe8\x69\x0d\x2f\xb4\x6a\x08\xe4\x5f\x47\xab\x38\x51\x68\x88\xa0\xa2\x4b\x41\x36\x7f\x53\xd0\xb8\xf0\x5b\x70\xc9\x33\x02\xbd\xaf\x93\x65\x0f\x6b\x9f\x75\xad\x1f\xdb\xf6\x5c\x5b\xee\xc7\x2c\x9a\x75\xc1\xa6\x57\xbf\xa9\x44\x35\x6d\x6f\x24\x17\x4b\x3f\x00\x5a\x1f\x9d\x74\xa6\x8e\xf7\xfc\xa9\x80\xd7\xe0\x66\xee\x5d\xc0\x02\x5e\x0a\x86\x6e\x47\x1a\x96\xad\xe5\xa3\x71\xf2\x35\x12\xf6\x93\x5d\x86\xf2\x3f\x5f\xd8\x0b\xd5\xb7\xac\xa6\xfe\x9b\xd8\x1b\x1b\xf7\x28\x3d\x1a\x48\x81\x83\xf6\x6b\xf7\x98\x75\xf1\xad\x4b\x5f\x2b\x77\x70\xbd\x7d\x0d\x1d\x79\xdc\xf5\xe9\xb2\x58\x88\x7b\xf2\x0f\x82\x3b\x5c\x4f\x92\xcf\xd3\x1f\x9d\x3a\xd7\x2d\x9f\xd6\xb7\x89\x18\x68\x9d\x21\x9f\xf3\x25\x3f\x46\x71\x23\x7f\xe1\x7b\x16\xe9\x82\xc1\xaa\x29\x2e\xa1\x4f\x08\x2c\x05\x53\x3d\xb8\x0b\x93\x0a\xf5\x17\xa2\xff\x72\x3e\x57\x89\x7d\xc7\xd2\xbc\x2e\xc0\x27\x1b\xb7\x73\x0e\xe8\x9f\x00\x00\x00\xff\xff\x5f\x73\x79\x4b\x5f\x04\x00\x00"

func dataAwsSimpleDeployMainTfTplBytes() ([]byte, error) {
	return bindataRead(
		_dataAwsSimpleDeployMainTfTpl,
		"data/aws-simple/deploy/main.tf.tpl",
	)
}

func dataAwsSimpleDeployMainTfTpl() (*asset, error) {
	bytes, err := dataAwsSimpleDeployMainTfTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/aws-simple/deploy/main.tf.tpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _dataAwsVpcPublicPrivateBuildBuildRubySh = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x56\x7f\x53\x1b\x39\x12\xfd\x7f\x3e\x45\xaf\xa1\xc2\x5d\x15\x33\x03\xd4\xed\xdd\xae\xbd\x6c\x1d\x10\xd8\xa5\x6a\xcb\xde\x32\xd9\xfb\x51\xd9\x1c\x35\x9e\x69\xdb\x0a\x63\x49\x91\x34\x06\x43\xfc\xdd\xef\xb5\xc6\xf8\x47\x48\x52\xf9\x07\xc6\x92\xba\xfb\x75\xf7\xeb\x27\xed\x7d\x97\x8f\x94\xce\x47\x85\x9f\x26\x89\xe7\x40\xa9\x21\x6d\x1a\xbd\xfa\x64\xe7\xf8\x41\xc5\x4f\xab\x2c\x8f\x0b\x55\xaf\x96\x83\x2b\x4a\x4e\x12\x7c\x19\xf7\x97\xbf\xd2\x53\x42\x44\xb5\x29\x8b\x9a\xbc\x69\x5c\xc9\x63\x55\xf3\xe9\xfe\xf1\x66\xb9\x56\x9a\xb5\x39\xdd\x3f\x91\x25\x2e\xa7\x86\x3a\x97\xc3\xe1\x60\x48\x45\xa0\xfd\xa7\x8d\xd1\xb2\xbb\xff\xd4\x9e\x5d\xf6\xe8\xb7\xc2\x07\xd8\x4f\x7c\xb7\x23\x66\x13\xc7\x96\x4c\x08\x86\xf2\x79\xe1\x72\x6c\xe4\x7e\xe1\xf1\x8f\x3e\x52\x88\xd8\x34\x9d\x1c\x25\xcb\x04\xe8\x2c\x1d\x44\x70\xd4\xd9\x7f\x3a\x3f\xbb\xf9\xf5\xf6\x66\xf0\xc7\xf0\xe2\x72\xd9\x91\x85\xdf\xae\xfb\x97\xfd\xc1\xb2\x73\x40\xc0\x90\x24\x86\x25\x05\x6c\xfc\xb3\x43\x27\x3f\xbf\x3a\x86\x3b\x38\x9d\xb0\xa3\x34\xb4\xf1\x7e\xa6\xbc\xe2\x79\xae\x9b\xba\xee\xd1\x32\x31\x75\x34\x68\xd3\x78\x2b\x27\xde\x11\x8c\x65\x2b\xd9\xa3\xb2\x36\x4d\x95\x96\x46\x8f\xd5\x84\xca\x42\x93\xd2\x81\xdd\x98\x1d\xd3\xbd\x0a\x53\x2a\x6c\xa0\xd2\xcc\x66\x85\xae\x3c\xa9\x31\xa9\x70\xe0\xc9\x07\x55\xd7\x38\x49\xd6\x19\xe4\xe9\x3d\x82\x50\xe7\xdf\x85\x0a\x4a\x4f\x68\x8c\x44\x76\xdc\x02\x13\x5c\xd8\x9a\x03\x67\x59\xd6\x49\x1a\x0d\x7b\x7a\xfb\x96\xd2\xf1\xaa\x38\x6a\x94\x47\x8b\x5c\x69\x1f\x0a\x5d\x72\x3e\x32\x26\xa4\x63\xa5\x95\x9f\x72\x45\xef\xde\xf5\xa8\x32\x28\xab\xaf\x19\x65\x3d\xca\xbe\x4f\x2a\xa3\xd1\x53\x89\x7b\x56\x55\x12\x56\x90\x02\x8b\x35\x5e\x05\xe3\x14\x7b\x02\x66\x6a\x6c\x55\x08\xaa\x18\xd8\x30\xf9\xa6\x32\x72\x34\x9d\x80\x35\x71\x93\x5f\x2c\x47\x10\x48\x30\x5d\x80\x20\xe3\x70\x5f\x38\x4e\x91\xa9\x65\x17\xe0\x36\x95\x72\x18\xbd\xb1\xaa\xaa\x54\x2c\xd1\xef\x36\xf4\x42\x0c\xad\x2d\xba\xe5\xd4\x29\x9f\xd6\x5c\xe4\xda\x54\x9c\xbd\xf7\x3b\x91\xc4\xee\xa5\xcd\xc8\xa9\xc9\x34\x8c\xcc\x43\xee\x9a\xd1\x22\xd5\x93\x1d\x9b\x3b\x5e\x20\xde\x9c\x52\xf9\xf2\xec\xe6\x68\xfb\xf4\xce\x76\xf3\x7c\xfd\x3b\x6b\x46\xa8\x6f\x93\x01\x65\xf7\x87\x23\x9c\x74\x5c\xce\xe3\x71\xfa\xfe\xef\xc7\x57\x3f\x9e\xff\x78\x71\x76\xf1\xb7\xa3\xf3\x93\xab\x7f\x24\x91\x13\x07\x15\x8f\x68\x1a\x82\xf5\x70\x63\xbc\x4f\x31\x64\x85\xd4\x2f\xb3\xd3\xc6\x2b\xa3\x6d\xe1\x3d\x6b\x30\x4c\x7c\xe6\x80\x91\xaf\x57\x28\xb8\xc6\x87\x05\xcd\x0a\xa5\x0f\xc0\xc4\x08\x34\x30\x53\xce\xa1\x8c\x47\xdb\x61\xf1\x59\xad\x7c\xc8\xaa\x8d\x65\x5c\xd8\xa6\xea\x97\x9a\x03\x8e\xbe\x19\xbc\x1e\x74\xc9\x16\xae\x98\x81\x42\x4e\x3d\x32\x49\x71\x08\xd9\x0a\x3e\x2a\xc0\x4c\x6d\x9b\x90\xf0\x83\x35\x2e\xd0\xf0\x8f\xf3\xff\xde\xfe\xeb\x72\x78\x73\x3d\xe8\x9f\x76\x9e\x9e\xe2\xe9\xdb\xe7\xd3\xcb\x65\xa7\xe5\xcd\x75\xdb\x66\xe1\xce\x10\x07\x0e\xe9\xf7\x67\x70\x87\xd4\x9f\x28\xfd\x70\x18\x19\x64\xc2\x14\x89\xda\xa2\xbc\x2b\x26\x48\x44\x78\xb4\x8a\xf3\xfa\xf2\xfc\xfa\xac\x7f\x7b\x35\x1c\xf4\xdf\x5c\xf6\x5f\x9f\x6a\xa3\xe3\xf4\x14\x65\x50\xf3\xaf\xd2\x6a\xf4\xe8\x68\x02\xad\x9a\xb1\x2b\x1b\xa7\x20\x39\xa3\x46\xd5\x55\xca\x02\x20\xc8\xef\x3f\x41\x76\x8c\x85\xfd\x90\xa2\x40\xf4\x88\xcf\xe3\x49\xfc\xfc\x32\x23\xa3\x8d\x04\x83\xa2\x68\x2f\x08\xd3\xd8\xd6\xb8\x2e\x0c\x7c\xdf\x7e\x4a\x35\xf6\xb7\x6b\xf4\x72\x25\x46\x8a\x66\x52\x87\x94\x1f\xe0\xd2\xd3\xba\x79\x2f\xea\x77\xde\xe8\xaa\x46\x53\xb7\x87\x6c\xc2\xb3\x75\xca\xa3\x76\x1f\x6c\xd4\x26\x75\x6a\xf5\xbf\x32\x65\xeb\xe9\x52\xfc\x97\xa1\x9d\x62\x1b\xbd\x44\x17\xb3\xbb\x4a\xc1\xc8\x52\xee\xdd\x3c\x17\xe9\xc2\xa4\xd9\x76\x2f\x14\x8e\x1e\x1f\x20\x20\x61\x66\xd7\x5b\x59\x98\x3c\x52\x7a\xf1\xc9\xf9\x5d\x95\xb0\xb5\x2a\x21\x09\xa8\x57\xe3\x3f\x81\x8c\x91\x94\x35\xc0\xab\x94\x2f\x46\x35\x57\xa9\xe4\x7c\x6f\x5c\x85\xb5\x09\x97\xc6\x53\xa7\x43\xbb\x8e\x6f\x38\x44\xe4\x68\xc6\x4c\x79\xa1\x98\xdf\x71\x8a\x19\xbb\xd7\x94\x0e\xd7\x66\xdd\xcf\xc1\xbb\x88\x52\x09\x2e\xc0\x53\x2c\x7a\xf4\x21\xe4\x9f\x2a\xd0\xdb\x43\xda\x3e\x34\xca\x41\x0b\x45\x5e\xb7\x06\x50\x0a\x1d\xa8\xc0\x7e\xe1\x8d\x16\xd0\xc4\x7a\xae\x9c\xd1\x33\x50\x89\xee\xa7\x22\xe5\xa0\x1a\xb4\x1d\xde\xa0\xa8\x15\xf1\x03\x97\x4d\x90\xa3\x1e\x1c\xbb\xc3\xb4\x36\xde\xc5\xbb\x15\x96\x87\x9b\x5f\xa0\x66\x7d\x48\x98\xe4\x8c\xae\x11\xa2\xf6\xc2\x65\x8c\x21\xfc\xd6\x0b\x38\xd3\xcc\xb8\x14\x80\xc0\x94\x38\x4a\x53\x28\x97\x5c\x0a\x98\x17\x6a\x95\x3f\xa3\x33\x6b\x59\xc7\xc2\x03\x82\x24\xa2\x7d\x33\x1e\xab\x52\xc1\x47\x46\xdd\xf4\x63\xdb\x4c\x8f\xbc\x52\x45\x07\xc7\x3e\xff\x9f\x80\xa0\xdf\xcf\xde\xfc\xda\xfb\x53\xe7\x07\xad\x92\xc4\x8a\xb4\x7f\x33\x71\xfd\x19\xab\x3d\x1a\xa0\xa0\x5d\x92\xd7\x80\x58\x63\x50\xb6\xca\x24\x37\x9b\x87\x2e\x3d\x6b\xdb\x57\x5c\x23\xb1\x3e\x12\x93\xbc\x1c\xcf\xcc\x9c\x91\x10\x90\xc3\x5d\x3c\x24\x85\x46\xd6\x90\x35\x82\x72\x73\x8b\xc4\xcd\xb6\x9d\xc9\xba\x4f\x39\x36\xa3\x82\xc0\x8d\x8b\xa6\x0e\x6d\x2f\xd9\x73\x7c\x5d\xe0\x72\x42\x5b\x2c\xee\x2b\x69\x12\x66\x4b\x46\x18\x9f\xfe\xb9\x80\x6b\xe8\xe9\x4a\x79\x2a\xda\x60\x3c\xc4\x40\x05\xf8\x8b\x97\x31\xfa\xae\x5a\x22\x54\x50\x04\x30\xc1\xb3\x48\x65\x78\x6e\x02\xdc\x21\xf9\xd0\x96\xcb\x34\x88\x06\xe6\xeb\x55\xbc\x2c\xc1\x2c\xd0\x4f\x3f\xf5\x7f\xb9\xee\xff\xe7\x62\xd0\xbf\x7a\x21\xe2\x6d\x4a\xe2\x6a\x47\xbe\x65\x61\x47\xbe\xf7\xe8\x17\xd6\x2c\x71\x2b\x82\x36\x4b\x33\x92\xf5\xf1\x5b\x87\xeb\xbc\x25\x96\xdc\xf5\xa2\x37\xf9\x1c\xc4\x30\xd8\x91\xef\xd5\x4d\x73\xbb\x36\xc8\xe5\x21\x16\xe2\x2c\xe1\x0d\xd0\xdb\xf6\x24\xca\xbf\xa6\xe8\x66\x7d\xec\x98\xe3\x66\x2f\x59\x27\x93\x7c\x63\x76\xbb\x0d\x5b\x8b\xc9\x37\xe5\xb8\xba\x85\xe3\xa3\x92\xe4\x52\x63\x4d\x3f\x1c\xf5\xe2\xcf\x36\xeb\xed\x61\xcf\x6d\x33\x82\xfe\xb4\xdb\x1b\xf0\xab\xd0\x64\x74\x0f\x2f\xc2\x2d\xfc\x22\x0d\xad\xb8\x3e\x6b\xa9\xcc\x92\x10\x64\x47\x28\xd3\x66\xad\x2e\x32\x11\xeb\x77\x32\xa5\x75\x49\x9d\xb2\xda\x05\x41\xaf\x5e\xad\x14\x79\x73\x27\x41\xf1\x6d\x6d\x16\x51\x33\xd2\x54\x9e\x7e\x42\x15\x64\xce\xb5\xb1\x71\x15\x25\x0a\xab\x9b\x13\x81\xe5\xf9\xf5\x5d\x27\xf9\x7f\x00\x00\x00\xff\xff\xdb\x24\x06\x87\x94\x0b\x00\x00"

func dataAwsVpcPublicPrivateBuildBuildRubyShBytes() ([]byte, error) {
	return bindataRead(
		_dataAwsVpcPublicPrivateBuildBuildRubySh,
		"data/aws-vpc-public-private/build/build-ruby.sh",
	)
}

func dataAwsVpcPublicPrivateBuildBuildRubySh() (*asset, error) {
	bytes, err := dataAwsVpcPublicPrivateBuildBuildRubyShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/aws-vpc-public-private/build/build-ruby.sh", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _dataAwsVpcPublicPrivateBuildTemplateJsonTpl = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x9c\x54\xcd\x8e\x9b\x30\x18\xbc\xf3\x14\x16\x52\xf6\x14\x20\xed\xae\xaa\xaa\xd7\x3e\xc6\x2a\x62\x0d\x78\x83\x15\xdb\x58\xfe\xec\x54\x59\xe4\x77\xef\x67\xfe\xa9\x36\x24\xdd\x5c\x8c\x3c\xc3\xcc\x78\xf8\x9c\x36\x22\xf8\x8b\x25\x57\xb9\xa6\xe5\x99\x99\xfc\xc2\x0c\xf0\x46\xc5\xbf\x48\x7c\x48\x7f\xa6\x87\x78\x1f\xf5\x9c\x0b\x35\x9c\x16\x82\x01\x42\xfd\x6b\xb8\x49\xff\x40\x4e\xcb\x92\x01\xe4\x67\x76\x45\x44\x39\x21\xf6\x4b\x14\x58\x69\x98\xbd\x85\x1a\x76\xea\xcd\x56\x08\x08\x77\xc2\x3c\xb6\x1e\x80\x6e\xdf\x8f\x41\xb4\x69\x2e\x3c\x64\xc4\xa4\x48\x78\x1d\xde\x6a\x77\xe4\xbd\x31\xa4\xe2\x86\x70\x85\x8f\x4e\x55\xd4\x22\x2b\xc7\x1d\x48\x0b\xc7\x45\x45\x76\x7e\x24\x0f\x2b\xca\xd9\xab\x66\xe1\xb4\x50\x33\x21\xe2\xfd\x0c\x70\x25\xb8\x0a\xd0\x6b\x2c\xcf\x41\x36\xd1\x24\xb3\x52\x67\x8d\xb5\x4d\x36\x1b\x24\x6d\x1b\x9c\x45\xd3\xe8\xf4\x37\xee\x5a\x66\x88\xf7\xf1\x71\x50\xf2\xfb\xdb\x9e\xef\x5c\xb0\xa5\x25\x34\xce\x94\x1d\x82\x9a\xc1\xd2\xfb\x6c\x89\x57\x0c\x2c\x57\x9d\x6b\x20\xfd\x47\x9a\x07\xc2\x6c\x15\x50\x56\x8f\x1e\xdd\x7b\xf2\xf4\x44\x0a\x0a\x35\x49\x33\x49\xb9\x4a\xa1\xfe\xa4\x8b\x1d\x61\xaa\x0a\xdf\x6b\xeb\x93\x6c\xd4\xb3\x23\x38\xa8\x05\x86\x90\xa8\x80\x29\x1c\xe0\x39\xdf\xa6\xc1\x79\xc3\x33\xf7\x1e\x0b\xda\x23\x4d\x26\x54\xeb\xd4\x9e\x3e\xbe\x54\x18\x94\x86\x6b\x1b\xa0\x6e\xdc\x12\xe3\x8a\x6b\x38\xfe\xa8\xd5\xad\xc7\x71\x8e\x3b\xce\x30\xc3\xd3\x85\x52\x54\x76\xda\x21\xcb\x24\x3d\x39\x52\x49\x3f\xb0\x75\x56\xc0\x8c\xad\xae\xdf\xad\x62\xd6\xf7\x74\xbb\x9d\x78\x75\x65\xb7\x14\x67\xe2\x1d\xc5\xe9\x9a\x6f\xa9\xf5\xa4\x7b\xd9\xba\x11\xc8\xa9\xe4\x7d\x1f\x3c\xf9\xfe\xed\xc7\xf3\xa1\x7a\x79\x99\x39\x5c\x81\xa5\x0a\x59\x63\x6d\xe5\x73\x2a\xa8\x39\xb1\x85\x0c\xd4\x79\x70\x1e\xeb\x76\x05\x0e\xaf\x5b\x94\x2a\x79\x3e\x62\x6d\x1b\x9e\x70\xae\xff\xcd\x8e\x2b\x4e\x11\x95\xfa\xb3\xc4\xfd\x7f\xd6\x31\x8a\x7c\xf4\x37\x00\x00\xff\xff\x16\x08\x5f\x5b\x65\x05\x00\x00"

func dataAwsVpcPublicPrivateBuildTemplateJsonTplBytes() ([]byte, error) {
	return bindataRead(
		_dataAwsVpcPublicPrivateBuildTemplateJsonTpl,
		"data/aws-vpc-public-private/build/template.json.tpl",
	)
}

func dataAwsVpcPublicPrivateBuildTemplateJsonTpl() (*asset, error) {
	bytes, err := dataAwsVpcPublicPrivateBuildTemplateJsonTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/aws-vpc-public-private/build/template.json.tpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _dataAwsVpcPublicPrivateDeployMainTfTpl = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xbc\x55\x4d\x6f\x13\x31\x10\xbd\xe7\x57\x8c\x5c\x4e\x88\x6e\x0a\xa7\xaa\x12\xb7\x4a\xdc\xe8\x85\x1b\x42\x2b\xaf\xd7\x09\x56\x1d\xdb\xf2\x47\xd0\x2a\xda\xff\xce\xd8\x8e\xe3\xfd\x48\x5b\x10\x12\xdb\x4b\xfb\xde\x8c\x67\xe6\x3d\x7b\x7a\x03\x5f\xb8\xe2\x96\x7a\xde\x43\x37\xc0\x93\xf7\xfa\x03\xf4\x1a\x94\xf6\xc0\x7b\xe1\xe1\x40\x55\xa0\x52\x0e\x9b\xcd\x91\x5a\x41\x3b\xc9\x81\x08\xb5\xb3\xb4\x15\x3d\x81\xd3\x38\x81\xe9\x2f\xd7\x52\xc6\xb8\x73\xed\x33\x1f\xae\x90\x8e\x33\xcb\xfd\x0b\xa4\xe5\x7b\xa1\xd5\x82\xc0\xd0\x56\xd1\x03\x4f\xf0\x34\xe1\x20\x16\x91\x42\x39\x4f\x15\xe3\xad\x1f\x4c\x0c\x87\x9e\xef\x68\x90\x1e\x3e\x03\xf1\x9f\x9a\x83\x60\x56\x13\x98\x66\x18\x2b\x8e\x38\x77\xeb\x42\xa7\xb0\xab\xd5\x38\x26\x74\x52\xb0\x17\xe9\xa3\x61\x2d\x13\xbd\xbd\x02\x9f\x63\x37\xc6\xea\xa3\xe8\xb9\x4d\x03\x22\xb4\x01\xa8\xfa\xc4\xc6\xde\x9d\x30\xb1\x99\xeb\x36\x12\x0c\xab\x4a\xcd\xc3\x2a\x9e\xc2\xb2\x66\x10\xbf\x59\x58\xc6\x31\x04\x9b\xb0\xdc\xe9\x60\x59\xb5\x20\x58\xe1\x87\x76\x6f\x75\x30\x04\x08\x97\x5d\xee\x2c\xca\x1c\x4f\x39\x9d\xf2\xaf\xe3\x78\x8b\xdc\x6d\x3e\xb4\x38\x9e\xaa\xe6\x11\x6b\xc5\xfc\x37\x52\xc8\xf1\x3d\xd6\x73\xe9\x40\x00\x9c\xdf\x6b\xa6\x65\xee\xef\xf6\x63\x02\x77\x56\x1f\x5a\xa3\xad\x4f\xe0\x5d\xc2\xbc\x2e\x48\xc5\xa2\xb6\x6d\x27\x35\x7b\x76\x88\x7d\x27\x77\x4d\xfa\xd9\xde\x91\x1f\xc8\x8f\xb1\x98\x50\x2f\x57\x23\x9e\x19\x72\xa5\xe0\xfd\xb5\x8a\xf7\x7f\x56\xf2\x6d\x35\xa9\x31\x13\x35\x61\xa1\xe7\x5f\x6a\xf9\xda\x78\xff\x28\x66\x2d\x16\x99\xf1\x3c\xdf\xff\xb4\x6f\xa5\x65\xba\x88\x2b\x01\x2f\xdf\x9b\x4a\xe6\x77\xea\x26\x09\x65\xcc\xe5\x43\xce\xe3\xce\xbd\x2b\xb2\xac\x5d\x6d\xb0\xb1\xa6\x24\x95\x2d\xe3\x66\x45\x62\x52\x61\x1a\x9c\xa0\x79\x7f\x4e\xc0\x8c\x1b\xf8\xf6\xf4\xf8\xf4\x80\x7b\xf4\x99\x83\x14\xce\x73\x85\xbe\x42\xd4\xcb\x01\xd3\x6a\x27\xf6\xc1\xc6\xd5\x81\xb1\x99\xc6\x7d\x91\xf5\x97\x5d\x95\x15\xe6\x37\x35\x52\x13\x77\x16\x37\xfe\xb2\x0b\xd7\x57\xbc\x52\x25\xbd\x26\x26\x53\x6e\xe0\x91\x1b\xa9\x07\xa0\xa8\x90\x07\xbd\xab\x33\x2f\x0c\x2b\xf8\xd4\x35\xdc\xcb\x73\xcf\xce\x3b\xe9\x20\x92\x47\xb3\x25\x5d\xe9\x19\x3c\x31\x33\xbe\x8c\xd9\x39\xab\x95\x9d\x82\xcb\x3f\x89\x45\xd1\x02\xe7\xc7\x14\xef\xfa\xdc\x58\x4c\x7f\xc5\xf5\x68\xe3\xc5\x44\x4f\xf7\xe5\x51\x7c\x5d\xad\xc9\x8b\x74\x3a\x78\x13\x3c\x90\x60\x65\x56\xe3\x48\x65\x48\xc1\x3f\xbd\x37\x0f\xdb\x6d\x2e\x14\xef\x53\x3c\xbd\x57\x2e\xf7\xb7\x8d\x7b\xfa\x77\x00\x00\x00\xff\xff\x11\x23\xa2\xbf\x89\x07\x00\x00"

func dataAwsVpcPublicPrivateDeployMainTfTplBytes() ([]byte, error) {
	return bindataRead(
		_dataAwsVpcPublicPrivateDeployMainTfTpl,
		"data/aws-vpc-public-private/deploy/main.tf.tpl",
	)
}

func dataAwsVpcPublicPrivateDeployMainTfTpl() (*asset, error) {
	bytes, err := dataAwsVpcPublicPrivateDeployMainTfTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/aws-vpc-public-private/deploy/main.tf.tpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _dataCommonDevVagrantfileTpl = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x56\x6d\x6f\xdb\x38\x12\xfe\xae\x5f\x31\x55\xdc\x37\x20\x92\xae\xc1\xe1\x3e\xb8\x4d\xd0\xbc\xb5\x09\xd0\x8b\x0b\x3b\x2d\x70\xb8\x5d\xb8\xb4\x34\x92\xd8\xca\x24\x97\xa4\x9c\xb8\x8e\xff\xfb\xce\x50\x72\xec\x6c\xba\xdd\x02\x71\x4c\x0f\x39\x6f\xcf\xcc\x3c\xe4\x1e\xbc\x47\x85\x56\x78\x2c\x60\xb6\x84\x91\xf7\x7a\x1f\x0a\x0d\x4a\x7b\xc0\x42\xfa\x27\xd1\x5e\xb4\x07\xd7\xb5\x74\x40\x7f\xbe\x46\xf8\x2c\x2a\x2b\x94\x2f\x65\x83\x50\xfd\x55\x17\x4a\x6d\xc3\xa9\x02\x17\xd8\x68\x33\x47\xe5\x41\x97\x64\xc2\xb3\x09\x61\x4c\x23\x73\xe1\xa5\x56\x99\x43\xbb\x90\x39\xa6\x70\xe9\xc1\xd5\xba\x6d\x8a\xe0\x74\x86\x50\x0b\x55\x24\xec\x1c\x8b\x14\xae\x35\xcc\x75\x21\xcb\x25\x9b\x25\x3b\x3b\xee\xf7\xa1\x75\x18\xbc\x1d\x1b\xc3\x82\x34\x8a\xfa\xed\x34\xd7\xaa\x94\x55\x6b\xf1\x45\x7c\x10\xbf\xe4\x8c\xee\x3a\xd1\x5d\x04\xd0\xad\xd2\xc5\x3c\x9d\xe9\x5b\x38\x84\xb8\x16\xae\x96\xb9\xb6\x26\x33\x16\x73\xe9\xf0\x3f\xff\x8e\x23\x3a\xb8\x07\x17\xda\x51\x02\xaa\x59\x82\x42\x7f\xa3\xed\xb7\x07\xea\xbd\x0c\x62\x63\xe5\x82\x70\x98\xf6\x82\x78\x1f\xa4\x19\x42\xbc\x5a\x31\x10\x53\x69\xa6\xa2\x28\x2c\x3a\x07\xeb\x75\x6f\x78\x82\xbe\x35\x20\xc0\x2d\x55\x4e\xf8\x95\xba\x29\xd0\x42\x69\xf5\x1c\x74\x6b\x81\xad\x48\x55\x41\x21\x29\x20\xaf\x2d\xa5\xaf\x21\x5b\x74\xd9\x3d\x88\xa1\x33\x30\xed\x0d\xb0\x4b\x23\x7c\x9d\x6e\x0c\x90\xc3\x7d\x88\x37\x9a\xf1\x3e\xe9\x02\xe8\x1b\xaa\x1b\xc5\x77\x2f\x85\xca\xea\xd6\xec\x48\xba\x20\xcf\x95\x98\x51\x99\x27\x93\x0b\x10\x15\x97\x92\xca\x7b\x23\x6c\xc1\x86\x9d\xa6\xf2\x7b\xcf\xcb\x3e\x7b\xca\xd5\xa0\x2a\x50\xe5\x12\x5d\xc8\xc0\x6d\x23\x75\xae\x4e\x7b\xed\x69\x67\xeb\x10\xbc\x6d\xb1\x73\xf4\x4e\xb7\xaa\x08\x7d\x01\x9b\xca\x75\xbf\x5e\xc8\x12\x84\x5a\xbe\xa4\x53\xab\xa7\xa1\xbb\x08\x11\x90\x8a\x96\x1b\x8d\x29\x49\x5c\x4a\x38\xc3\xd3\x35\x1d\xe3\x7d\x2a\x69\xa6\xa9\x1d\xb3\xed\xa9\x84\x80\x21\xf5\x46\x6b\x93\x9e\x92\xd4\x13\x58\x5c\x8c\x9f\x43\xc9\xc6\x02\x82\xb4\x78\x70\xd4\x58\xbd\x90\x8e\x23\x8c\x5d\x8d\x4d\xc3\x15\x57\x8d\x54\x48\x18\xe6\x05\xec\xad\x48\x61\x0d\xcf\x9e\xc1\x8c\x5a\xab\xff\x99\xcd\x85\x54\xa9\xab\xe3\x2e\x19\x82\x8a\xf3\xa1\xa0\x03\x04\x1f\xb4\x28\x40\x34\x4d\x28\x7f\x69\x45\xc5\xb3\xe3\xa0\x46\x8b\x21\x6f\x42\xe1\x01\xc0\xe9\x16\x92\xcd\x69\xc6\x85\xfb\x6d\xab\x1d\x10\xe1\xcc\x7b\xc9\x9d\x45\xf2\xb2\x5e\xff\x30\x82\x4b\xe5\x3c\x07\x30\x6e\x69\x9a\x67\xad\xa4\x89\x44\xb5\x90\x56\x2b\x56\xfd\xd5\xf4\x07\x2e\xb7\xd2\xf8\xa9\x25\x2b\x11\x79\x88\xa2\x5d\x09\x95\xe6\xcd\x9b\xc9\xe9\xf8\xf2\xe3\x75\xe4\xd0\x43\xc2\x54\xd3\xaa\x7e\x89\xd6\xe2\xad\x0c\x4b\x23\x0d\x96\x42\x36\xbd\xd8\x5b\x91\x53\xaf\xd0\x4a\xdb\x17\x2f\x61\xc5\x6d\xdc\xe8\x5c\x34\xd4\x87\xad\xcd\x91\xc7\xff\x70\xf0\x6a\x2b\xe6\x60\x94\x3e\x1c\x1c\xb0\x08\xf3\x5a\x43\x7c\x3e\x1e\x8f\xc6\x20\x3c\x0c\x56\x5b\xa5\xf5\x70\xb0\xea\xce\xae\x5f\xc3\x07\x41\xd3\xde\xe8\xca\x0d\xb9\x46\x34\x14\x68\x80\xfb\x88\x27\xcf\x66\xb4\x91\xb9\xa5\xa3\x2f\xb8\x03\x1f\x62\x53\x70\xf0\xaf\x68\x1d\x51\x74\x06\x9e\x87\xe0\x20\x1e\xac\x4e\x8e\x27\x17\xd3\xc9\xe8\xd3\xf8\xf4\x7c\x1d\xb3\xe0\xc3\xe5\xd5\xf9\xd5\x68\x1d\x3f\x07\x8a\x21\x22\x16\x63\xa3\x09\xde\x62\x3e\x04\xfe\xdf\xd2\xf0\xe4\x7a\x3e\x27\xe2\x83\x1b\xe9\x6b\xea\x01\x6f\xda\x10\x4a\xc5\xe4\x4a\x4b\xe6\xc6\x42\x3a\xd3\x88\x25\x16\x91\x46\x06\x01\x06\x6f\xe1\xe0\xe8\xd9\x2b\x0a\x27\x9c\xb4\x90\xf8\x2e\xde\x23\xc8\xa8\x11\x32\xd5\x36\xcd\x6b\x58\xdf\x7b\xa4\x53\xc3\x8d\x6d\x41\x63\x4b\x08\xdc\x92\xfd\x39\x31\x13\xcd\x64\xa4\x9b\x60\xb5\x43\xeb\xff\xac\xf1\x3b\xb9\x88\x7b\x0b\xff\x15\xdf\x10\xa8\x38\x34\xf8\xbe\x26\x14\xbf\xf4\x5c\x01\x34\xda\x5f\xa0\xd2\x34\xf3\x1d\x5b\x35\x81\xac\x98\x97\x89\x52\x59\x10\xa6\xa7\xb3\x4a\xb3\x71\xcf\x45\x70\x44\x61\xd6\x7a\x8e\x1b\x49\x96\xf2\xb4\xd8\x9c\xbd\x9d\xf6\x34\xc0\xfc\xc2\xfc\x13\xfa\x9c\xca\x43\x49\x52\x16\x52\x45\x44\x0c\x4f\xba\x0a\xc5\x9f\x1c\x9e\x5d\x4d\x08\xa2\x18\x32\xf4\x79\x46\x01\xf1\xa7\x98\x76\x0d\x0b\x47\x3b\x60\x50\x58\x2a\xda\x74\xc4\x8e\xe2\x1d\xb8\x96\x6e\x09\x8f\x08\x89\xf8\x27\x33\x64\x40\x63\xa7\xd0\x5f\x63\x0c\x02\x10\xc3\x7b\x61\x7d\x54\x4a\xea\xd4\x5b\xa3\xad\x87\xb3\xf3\x93\xcb\xe3\xab\xe9\xbb\xf1\xe8\xea\xfa\xfc\xea\xec\x50\x69\x25\x99\x7b\x44\xee\xe5\x82\x1a\x5a\x37\x10\x1f\x17\x81\x50\x85\xf1\x6c\xc1\x68\x27\x89\xf0\x99\x42\xb9\x1d\x5a\xc3\xfc\xa5\xaa\x34\x4d\xe3\x68\xe3\x94\x8e\x26\x44\xbe\xdd\x26\x3e\x12\xcb\x7e\x90\x93\x25\x98\xa5\xaf\x89\xfc\x9c\x2e\x3d\x51\x2f\x26\x34\xb8\x06\xad\x67\xeb\x3f\x90\x25\xdc\x84\x34\xd4\x6c\x88\x7a\x5a\x39\xce\x21\xa9\xbd\x37\x6e\xeb\xa4\x28\x12\xde\x27\xe4\xbb\x48\x97\xc1\x8f\x11\xc3\xbc\xb6\xd2\x25\x0d\x8a\x4c\xe9\x02\xd3\xaf\xee\x41\x60\xac\xf7\x58\x67\x66\x65\x55\x7b\xba\x86\x33\x66\x87\x44\x55\x7f\x97\x23\xbf\x42\x46\x67\xa3\x21\xdd\x6e\x56\xcc\x91\x20\x94\xdf\x11\x02\xa5\x2c\xd0\x06\x2a\x12\xf4\x46\x51\xd4\xda\x1b\xec\xc7\x9f\x4e\xfe\x37\xfd\x7c\x3e\x9e\x5c\x8e\xae\x0e\x99\xcd\xf9\xf4\x74\x73\x3a\xdc\xc4\x0c\x7f\x4f\x7b\x5c\x82\xc0\x7c\x83\xd5\xae\xe2\x3a\x54\xc1\xb5\x86\x4d\x86\xcb\x4e\xe4\xdf\x68\x58\x5c\x28\xc8\xaf\x55\xf9\x27\xf5\x99\x7d\xb7\x50\xd1\x50\xcd\xd1\xe6\xd4\xec\x44\x5c\x81\x78\x13\x9a\x48\xe2\x5c\xfe\xfd\x1b\x35\x5b\x23\x67\xe6\x8f\x84\xaf\xb8\xef\xb4\x7c\x55\x85\xe5\x4f\xea\xc7\x3a\x5c\x83\xaf\x2e\x2c\x39\xef\xc1\x6e\x52\x8f\x25\x6c\xf1\x11\x1c\x27\x74\x79\x36\x68\x37\xad\x57\xe1\xfc\x3e\xf4\x59\xb7\x05\x49\xa2\x74\x62\x65\xff\x5d\xe8\xbc\x33\xb2\x3b\xbd\xef\x29\x3d\xe2\x03\x7e\xab\xf1\x20\xb3\x05\xbe\x84\x74\x09\x17\xd7\xd7\x1f\x99\x4e\x6e\x88\x29\x84\xea\x9e\x18\x49\xff\x48\xb8\x7f\x54\x70\xcf\x80\x68\xe9\x49\xb3\x09\x83\xec\xf5\x13\x99\x24\x55\xa3\x67\x04\x52\x6b\x9b\x34\xa6\x8d\xb7\xf4\xa9\xdb\x19\x3d\xff\xe6\xc3\x38\xed\x5d\x8d\x4a\x7a\xe1\x71\x0f\x0f\xb3\x6c\xbb\x9f\xc5\x51\x7f\x0f\xfd\x19\x00\x00\xff\xff\x4a\xd0\x25\x41\x04\x0b\x00\x00"

func dataCommonDevVagrantfileTplBytes() ([]byte, error) {
	return bindataRead(
		_dataCommonDevVagrantfileTpl,
		"data/common/dev/Vagrantfile.tpl",
	)
}

func dataCommonDevVagrantfileTpl() (*asset, error) {
	bytes, err := dataCommonDevVagrantfileTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/common/dev/Vagrantfile.tpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"data/aws-simple/build/build-ruby.sh": dataAwsSimpleBuildBuildRubySh,
	"data/aws-simple/build/template.json.tpl": dataAwsSimpleBuildTemplateJsonTpl,
	"data/aws-simple/deploy/main.tf.tpl": dataAwsSimpleDeployMainTfTpl,
	"data/aws-vpc-public-private/build/build-ruby.sh": dataAwsVpcPublicPrivateBuildBuildRubySh,
	"data/aws-vpc-public-private/build/template.json.tpl": dataAwsVpcPublicPrivateBuildTemplateJsonTpl,
	"data/aws-vpc-public-private/deploy/main.tf.tpl": dataAwsVpcPublicPrivateDeployMainTfTpl,
	"data/common/dev/Vagrantfile.tpl": dataCommonDevVagrantfileTpl,
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
	"data": &bintree{nil, map[string]*bintree{
		"aws-simple": &bintree{nil, map[string]*bintree{
			"build": &bintree{nil, map[string]*bintree{
				"build-ruby.sh": &bintree{dataAwsSimpleBuildBuildRubySh, map[string]*bintree{
				}},
				"template.json.tpl": &bintree{dataAwsSimpleBuildTemplateJsonTpl, map[string]*bintree{
				}},
			}},
			"deploy": &bintree{nil, map[string]*bintree{
				"main.tf.tpl": &bintree{dataAwsSimpleDeployMainTfTpl, map[string]*bintree{
				}},
			}},
		}},
		"aws-vpc-public-private": &bintree{nil, map[string]*bintree{
			"build": &bintree{nil, map[string]*bintree{
				"build-ruby.sh": &bintree{dataAwsVpcPublicPrivateBuildBuildRubySh, map[string]*bintree{
				}},
				"template.json.tpl": &bintree{dataAwsVpcPublicPrivateBuildTemplateJsonTpl, map[string]*bintree{
				}},
			}},
			"deploy": &bintree{nil, map[string]*bintree{
				"main.tf.tpl": &bintree{dataAwsVpcPublicPrivateDeployMainTfTpl, map[string]*bintree{
				}},
			}},
		}},
		"common": &bintree{nil, map[string]*bintree{
			"dev": &bintree{nil, map[string]*bintree{
				"Vagrantfile.tpl": &bintree{dataCommonDevVagrantfileTpl, map[string]*bintree{
				}},
			}},
		}},
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

