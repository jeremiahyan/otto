// Code generated by go-bindata.
// sources:
// data/aws-vpc-public-private/deploy/main.tf
// data/aws-vpc-public-private/deploy/module/join.sh
// data/aws-vpc-public-private/deploy/module/main.tf
// data/aws-vpc-public-private/deploy/module/outputs.tf
// data/aws-vpc-public-private/deploy/module/variables.tf
// data/aws-vpc-public-private/deploy/variables.tf
// data/common/app-dev/main.sh.tpl
// data/common/app-dev/upstart.conf
// data/common/app-dev-dep/main.sh.tpl
// DO NOT EDIT!

package consul

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

var _dataAwsVpcPublicPrivateDeployMainTf = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xe4\x90\x41\x6e\xc3\x20\x10\x45\xf7\x3e\x05\x42\xdd\x9a\xc6\x89\x94\x76\xd3\xb3\x20\x02\xa3\x76\xda\x18\x2c\x06\xdc\x56\x91\xef\x5e\x6c\x88\x70\x94\x2a\xca\x3e\xde\xf1\xfe\xc3\xfe\xfe\x83\x77\x23\x1a\xf0\x8c\xab\x6f\xe2\xec\xd4\xb0\xf4\x28\xad\x81\x48\x7e\xc1\x2f\x7b\x63\xfc\xe9\x34\x2a\x2f\x52\x2c\x2b\x9f\xf8\x22\x12\x68\x0f\xe1\x5a\xac\xbc\x88\x1e\xde\xd1\xd9\x2a\xe5\x73\x0a\xa7\xa6\xe9\x9d\x89\x47\x60\x5c\x3b\x4b\xf1\xd8\x76\xe7\x16\xe4\xa2\xd7\x30\xdf\x11\xcf\xd9\xe1\xcd\x12\xa0\x35\xf0\x33\xf3\x2e\xbf\x7c\xf0\x38\xaa\x00\x2d\x0e\x0b\xdc\x88\x8d\xe8\xc4\x6b\xce\x54\x8f\xab\x6a\x3d\x96\x3e\xa9\x59\x6b\x55\x0f\x35\x4b\x44\xce\xe4\xfc\x67\xf1\x60\x21\xb4\x68\xaa\x51\x50\xf9\x5a\xf1\xc6\x41\x5f\x48\xe9\x2c\xd1\x94\xf0\xd3\xa1\x95\xca\x18\xbf\xea\xb5\xcf\xd1\x41\x51\x48\x0b\xc8\x0f\x47\xa1\xde\x5e\xd3\xe9\x52\x8c\x04\xfe\x5a\x9c\xe9\xbf\x2b\x6e\xef\x5f\x71\x7b\x63\xc5\x97\xc7\x5e\x71\x77\xff\x8a\xbb\x1b\x2b\xee\x1f\x6b\xc5\xbf\x00\x00\x00\xff\xff\xc6\x51\x33\x6f\x54\x04\x00\x00"

func dataAwsVpcPublicPrivateDeployMainTfBytes() ([]byte, error) {
	return bindataRead(
		_dataAwsVpcPublicPrivateDeployMainTf,
		"data/aws-vpc-public-private/deploy/main.tf",
	)
}

func dataAwsVpcPublicPrivateDeployMainTf() (*asset, error) {
	bytes, err := dataAwsVpcPublicPrivateDeployMainTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/aws-vpc-public-private/deploy/main.tf", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _dataAwsVpcPublicPrivateDeployModuleJoinSh = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x4c\x90\xb1\x4e\xc3\x30\x10\x86\x77\x3f\xc5\x5f\xb7\x23\x4d\x9b\x2e\x48\x54\x9d\x10\x03\x6b\x01\x31\x20\x86\x6b\x7c\x69\x0e\x5c\x3b\x8a\x2f\x12\x79\x7b\x9c\x90\x8a\x6e\x3e\xdf\x7d\x9f\x7f\xdf\x72\xb1\x39\x49\xd8\x9c\x28\x35\xc6\x70\xd5\x44\xd8\x77\x12\x95\x70\x46\x1d\x3b\x3c\xc6\x90\x7a\x8f\xaf\x28\xa1\x28\x0a\x6b\xcc\xeb\xf1\xf9\xe9\xe5\x50\x6e\x8d\xe9\x83\x8a\x47\xf5\x3f\x80\x55\xb9\x87\x8b\x06\xf0\xac\x98\x06\xd7\x87\xd2\xe4\xfa\xcf\x7b\xe3\x42\x4d\xe2\xd9\x15\x38\xb2\x76\xc2\x29\x13\xb5\x3e\x60\x35\x41\x36\x13\x52\xe3\x63\x2e\xb1\xf6\x8c\x2d\x3e\xf7\xd0\x86\x43\xee\x01\x4b\xbc\xb5\x8e\x94\xc7\x1b\xd8\x0b\x9d\xa5\xb2\xf8\xe6\x01\xd9\x3c\xbf\xa2\x0d\x29\x12\x0d\x09\x36\x44\x38\x6e\x7d\x1c\xd2\xc2\x4e\xf8\x35\x4e\xef\x1d\x42\x54\x74\x4c\x55\x03\x6a\x5b\x2f\x15\xa9\xc4\x70\x07\x4a\xa9\xbf\x8c\x4b\x10\x9d\xb3\x42\x23\x92\x52\xa7\x57\xc9\x4f\x6e\xed\xee\xf3\xb9\x96\xf1\x8f\xc9\x33\xb7\xd8\x19\x17\x03\x9b\xdf\x00\x00\x00\xff\xff\xb7\x5c\x3d\x0e\x56\x01\x00\x00"

func dataAwsVpcPublicPrivateDeployModuleJoinShBytes() ([]byte, error) {
	return bindataRead(
		_dataAwsVpcPublicPrivateDeployModuleJoinSh,
		"data/aws-vpc-public-private/deploy/module/join.sh",
	)
}

func dataAwsVpcPublicPrivateDeployModuleJoinSh() (*asset, error) {
	bytes, err := dataAwsVpcPublicPrivateDeployModuleJoinShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/aws-vpc-public-private/deploy/module/join.sh", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _dataAwsVpcPublicPrivateDeployModuleMainTf = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xac\x53\x3d\x8f\x1b\x21\x10\xed\xef\x57\x20\x74\x5d\xe2\xb5\x2f\x91\xd3\xa5\x4a\x9f\x26\xe5\xc9\x42\x18\xc6\x36\x39\x16\x10\x1f\x8e\x2d\x6b\xff\x7b\x06\x16\xef\x47\x6e\x93\xe8\xa4\xc3\x8d\xf7\xcd\x9b\x79\x0f\x66\xc6\x43\xb0\xc9\x0b\x20\x94\xff\x0a\x4c\x99\x10\xb9\x11\x40\x09\x15\xd6\x84\xa4\x29\xb9\x3d\x10\x3c\xbc\x55\xe4\x2b\xa1\x8f\xb7\x33\xf7\x0d\x7e\x74\xb4\xc0\x77\x3e\x8b\x57\x07\x99\x10\x3f\x35\xa1\xe5\x5a\xf7\xe1\x17\xb8\x32\xc3\x5b\x18\x53\x11\x59\x65\xa4\xe6\x87\xb4\x37\x10\x99\x92\x23\xa3\x87\x56\x4a\x56\xca\xd9\x09\x16\x40\x24\xaf\xe2\x95\x1d\xbd\x4d\x0e\xe9\x01\xf9\xcf\x98\x90\x3d\xcf\x83\x4d\xef\xbb\xc9\xf9\xbb\x52\xc0\x79\x75\xe6\x11\x98\x72\xa3\x48\xc5\x56\xca\xa1\x4a\x61\x45\x7e\x0c\xf5\xae\xf9\x7c\xaf\xae\xfb\x6a\xa4\xcf\x52\x46\xc2\xa5\xda\xea\xfa\x34\x8c\x1b\x10\x51\x59\x33\x49\x4e\x01\xfc\xfd\x7f\x2e\x92\xf6\xc9\xc4\x44\x87\xf8\xc9\x86\x38\x8d\x3f\xde\x02\xe8\x43\x33\x1a\xed\x46\xee\x9e\x87\x5c\x9d\x95\x9c\xc1\xff\x14\x5d\x20\x17\x03\xaf\xc8\x19\x9d\xbb\x77\xde\x9e\x55\xc0\x18\xd2\xe9\x41\x69\xa0\x93\x5b\xd4\xb9\x28\x65\x1c\x8f\xa7\xa6\xb5\x32\x69\xe8\xd6\x3f\xad\x32\x4d\x38\x8d\xb2\x12\x50\xc0\xf0\xf2\x0a\x48\x5f\xc7\xd6\xad\x83\xf0\xca\xc5\x81\xb6\x24\xe8\xa1\xb5\xd8\x03\xb8\x80\x98\xea\x2a\xa3\x95\xc9\xba\xcf\x03\x94\x0f\x15\x27\x34\x40\x3e\x5c\xc8\x1f\xf5\x3f\xce\x69\xf3\x68\xed\x5c\xb6\xcc\xb8\x94\x78\xff\x91\xbe\xab\xce\xd0\xdb\x7c\x0b\xe6\x13\xf5\x6a\x17\xcc\xff\x66\x03\xdf\xa3\xe8\xd7\xf7\xf8\x51\xcb\x91\x52\x8e\x1c\xac\x27\xdf\xfe\x96\x9a\xa7\x7d\xba\x0d\xf8\xdd\xaf\x42\x5d\xb7\x23\x3a\x9d\xce\xe9\xc1\xdb\x96\x39\xeb\xf3\x6c\x3c\x0d\x68\xb4\x77\xec\xcb\x76\xfb\x79\x3b\xe0\xf8\xfc\xd1\x0a\xab\xcb\x50\x4a\x37\x76\x50\x28\xe9\xd9\x5e\x5b\xf1\xd2\x6f\xd6\xa6\x29\xbf\xf5\x86\xee\xa6\xed\x7b\x4f\xfd\x28\xde\xae\x0f\xff\x92\xdf\x2c\xc8\x6f\x16\xa5\x57\x4f\x6f\x51\xee\x1e\x7e\x07\x00\x00\xff\xff\x58\x0b\x1e\x88\x24\x05\x00\x00"

func dataAwsVpcPublicPrivateDeployModuleMainTfBytes() ([]byte, error) {
	return bindataRead(
		_dataAwsVpcPublicPrivateDeployModuleMainTf,
		"data/aws-vpc-public-private/deploy/module/main.tf",
	)
}

func dataAwsVpcPublicPrivateDeployModuleMainTf() (*asset, error) {
	bytes, err := dataAwsVpcPublicPrivateDeployModuleMainTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/aws-vpc-public-private/deploy/module/main.tf", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _dataAwsVpcPublicPrivateDeployModuleOutputsTf = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x0c\xc0\x41\x0a\x05\x21\x08\x00\xd0\x7d\xa7\x10\xf9\xeb\x6e\xf0\xcf\x12\x52\x2e\x82\x30\x49\x6d\x16\xd1\xdd\x67\xde\x0c\xd7\x70\x40\x6a\x6d\xb1\x19\xc2\x49\xf0\xd9\x34\x82\xe1\x0f\xf8\x3b\xf4\x58\xe9\x62\x4e\x52\x39\xd7\x29\x16\x23\xeb\xea\x9b\x9c\x4b\xd7\x8b\xe9\xa6\x37\x00\x00\xff\xff\x7c\x4c\xbb\x14\x45\x00\x00\x00"

func dataAwsVpcPublicPrivateDeployModuleOutputsTfBytes() ([]byte, error) {
	return bindataRead(
		_dataAwsVpcPublicPrivateDeployModuleOutputsTf,
		"data/aws-vpc-public-private/deploy/module/outputs.tf",
	)
}

func dataAwsVpcPublicPrivateDeployModuleOutputsTf() (*asset, error) {
	bytes, err := dataAwsVpcPublicPrivateDeployModuleOutputsTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/aws-vpc-public-private/deploy/module/outputs.tf", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _dataAwsVpcPublicPrivateDeployModuleVariablesTf = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x91\x41\x4b\xc4\x30\x10\x85\xef\xfd\x15\x43\xef\xc5\x5f\xe0\x61\x59\x0f\xf6\x20\x14\x16\xbc\xee\x4e\x9b\xd1\x1d\xdd\x4e\x4a\x26\x2d\x8a\xf8\xdf\xcd\x04\x41\xac\x8d\x6c\x4e\x21\xef\x7d\x8f\x79\x93\x05\x03\x63\x7f\x21\xa8\x71\xe4\x1a\x3e\x2a\x48\xc7\x91\x0e\x81\xa7\xc8\x5e\xe0\x16\xea\xdd\x43\x0b\xd1\xc3\xac\x04\x4f\x3e\xc0\xde\x8b\xce\x97\xba\xfa\xac\xaa\x1f\x9a\xc5\xd1\x5b\x81\x6f\x4d\xcb\x68\x3c\x13\x08\x8e\xb4\x82\x5f\xe9\xbd\xc9\xcf\xdb\xfc\xe1\x70\x0f\xc9\xb2\x45\xf6\xa8\xe6\x3a\x9e\xbd\xc6\x7f\xe8\x6f\x1b\x64\xdb\x76\x42\x6a\x17\xae\x48\xc8\xb6\xdf\x09\x53\xe0\x05\x23\x35\x3c\x95\xfa\x77\xb6\x3e\x54\xe5\x67\xb1\x9b\x6d\x81\x45\x23\xca\xb0\xee\xa3\x73\x2f\x14\x1b\x76\xa5\x51\xb2\x0e\xed\xdd\x8a\x5b\xa6\xa1\x0c\x3d\x76\xfb\xbf\xc4\x8b\x67\x39\xa2\x73\xa5\xd2\x6d\x77\xb3\x4b\x2a\xa9\xda\xc8\xa7\x21\x7f\x3a\x18\x75\xb2\xa4\xaf\x00\x00\x00\xff\xff\xa5\x37\x94\xc2\x38\x02\x00\x00"

func dataAwsVpcPublicPrivateDeployModuleVariablesTfBytes() ([]byte, error) {
	return bindataRead(
		_dataAwsVpcPublicPrivateDeployModuleVariablesTf,
		"data/aws-vpc-public-private/deploy/module/variables.tf",
	)
}

func dataAwsVpcPublicPrivateDeployModuleVariablesTf() (*asset, error) {
	bytes, err := dataAwsVpcPublicPrivateDeployModuleVariablesTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/aws-vpc-public-private/deploy/module/variables.tf", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _dataAwsVpcPublicPrivateDeployVariablesTf = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xac\x91\x41\x6b\xb3\x40\x10\x86\xef\xfe\x8a\xc1\x9c\xfd\xf8\x28\xe4\xd8\x43\x48\xa1\xf5\x50\x08\x15\xda\xa3\xac\xeb\x18\x97\x6e\x76\x65\x67\x55\x4a\xe9\x7f\xef\xb8\xf1\xd0\x48\x57\x0a\xc9\xe0\x6d\x9e\xf7\x51\xdf\xd9\x64\x37\x98\x64\x03\x3b\x29\x91\x08\x72\xd3\xd8\xe4\x36\xce\x64\x10\x4e\x89\x4a\x23\xa4\x62\xa4\x52\x84\x17\x94\xef\xf8\x91\xc2\x67\x02\x3c\x35\x92\x74\xaa\xf3\xca\x1a\xb8\x87\x74\xfe\x02\x06\xa0\xb1\x0e\x76\x6f\x45\x9a\x7c\x2d\x2d\x84\xd2\xa1\x5f\xb1\x14\x01\x58\xb1\x38\x3c\x32\x1a\x49\xbf\x84\x25\x8c\x2d\x3a\x84\x91\x1f\xa5\x35\xd8\x0e\x9d\xf0\xf8\x2f\x88\xae\xec\x86\x9b\x7e\x44\xc3\x3e\x0d\x84\xde\x2b\x73\xa4\xab\x95\x3f\x2b\x3a\xa9\x58\xbb\xcf\x39\x78\x0b\x5a\xf4\x46\xb6\xfc\x5f\xbe\x85\xbd\x35\xd4\xeb\x74\xa6\x1b\xd1\x6b\x3f\x91\xec\xc8\xea\xe6\xee\xff\x76\x5b\x89\x45\x77\x95\xa0\xc9\x57\xb6\x96\x7c\xac\xff\xe2\x09\x66\x0c\x02\xf6\xbb\xa1\x27\x74\x7f\x30\x04\xec\xd2\xc0\x97\x2d\x8d\x38\xe1\x4a\x7a\x3a\x7e\x40\x2e\x93\xd4\x57\x06\x7d\xd6\x39\x35\xf0\x39\x23\xf9\xc3\x79\x0b\x67\x78\x61\x18\x3a\x59\xaa\x3a\x92\x7c\x3d\xec\x21\x7f\x98\x12\xdf\x01\x00\x00\xff\xff\x92\x93\x40\xa1\x94\x03\x00\x00"

func dataAwsVpcPublicPrivateDeployVariablesTfBytes() ([]byte, error) {
	return bindataRead(
		_dataAwsVpcPublicPrivateDeployVariablesTf,
		"data/aws-vpc-public-private/deploy/variables.tf",
	)
}

func dataAwsVpcPublicPrivateDeployVariablesTf() (*asset, error) {
	bytes, err := dataAwsVpcPublicPrivateDeployVariablesTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/aws-vpc-public-private/deploy/variables.tf", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _dataCommonAppDevMainShTpl = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x84\x54\x5d\x4f\xe3\x38\x14\x7d\xf7\xaf\xb8\xa4\x68\x05\xda\x4d\xd2\x22\x96\x5d\x41\x41\xbb\x82\x61\x84\x84\xe8\x68\x3a\xa3\x79\x40\xa8\x32\x89\xd3\x78\x70\xec\x4c\xec\x14\x3a\x25\xff\x7d\xae\x1d\xa7\xd3\x96\xf2\xf5\x52\xdb\xe7\x1e\x9f\x7b\xee\x89\x7b\x3b\xf1\x3d\x97\xf1\x3d\xd5\x39\xd1\xcc\x40\xc8\x08\x51\x6c\x6f\x1f\x16\xb0\xfb\x1f\x1c\x9c\xfd\x31\x80\x67\x10\x6a\x3a\x65\x15\x84\x06\x94\x31\x0a\xce\x20\x4e\xd9\x2c\x96\xb5\x10\x27\xd0\x10\x25\x1c\x9c\x25\xb9\x82\xe0\xd6\x22\xee\xb0\x36\xb0\x47\x24\xe5\xd5\xe9\xee\x5e\xf9\x98\xee\x13\xd2\x83\x0b\xf5\x28\x85\xa2\x29\x50\x99\x02\xde\x56\x97\x70\xae\xa4\xae\x05\x20\x8e\x25\x46\x55\x9c\x69\xc2\x33\xd8\x81\x44\x15\x85\x45\x85\x33\xfc\xe9\x20\x67\xcb\x4b\x9d\xac\x13\x30\x39\x93\x04\xf0\x4f\x09\x08\xae\xa4\x36\x54\x08\x2e\xa7\x9e\x32\x8a\xa2\xa0\x3d\x65\xa0\xeb\x54\x01\x2d\x4d\x38\xc5\x0e\xeb\x32\xa5\x86\x41\x38\xdf\x7a\xca\x5b\x1e\x3c\x86\x5a\xfe\xe4\xa5\x03\x25\x29\xc4\xa6\x28\xbb\x82\x47\x0b\xcc\x8d\x29\xf5\x71\x1c\xa7\x22\x42\x03\x4d\x45\xe7\x11\x6a\x8e\x0b\x6e\x92\x9c\x09\x91\xc7\xad\xec\xb8\x1f\xfd\x1d\x1d\x4c\x50\x58\xfd\x34\xa1\x45\x7a\x74\x18\x21\x2b\x84\x23\xdf\x56\xd4\xdd\x81\xbc\xee\xc2\x2d\xfb\x4e\x60\x92\x17\x2a\x85\x3f\x9f\xfc\xf9\xda\x59\xb1\x34\x29\xae\x75\x15\x0b\x95\x50\xe1\xc6\xba\x0d\xfb\x80\x5e\x43\x58\x42\xcc\x4c\xe2\x01\x51\xfa\x0a\xa4\x90\xe6\x7d\x0e\xcd\xaa\x19\x4f\x30\x37\x16\xd2\x83\x6f\x15\x47\x7b\x71\x38\x90\x09\x3a\xd5\x80\x81\xa1\x60\x58\x51\xaa\x8a\x56\x73\xc8\xb8\x60\x6e\xfe\x85\x9a\x31\xe0\xd6\x71\x44\x94\x82\x22\x83\xf3\x9a\x1a\x1c\x35\xba\xed\x2f\x9e\xb4\x2c\xc3\x21\x7c\x18\x5d\x12\xf6\x84\x34\x06\xce\x47\x37\xe3\xaf\xd7\x93\xcb\xeb\xff\x3f\x8e\x4f\x83\xd0\x2a\xb0\xf9\xbc\x57\xca\x68\x1c\x05\xfa\x8b\x33\xa6\xa1\x4d\xdf\x4a\x0f\x01\xb1\x14\xbe\x91\xd6\xcf\xfe\xd1\xe1\x21\xbc\xb8\x6d\xd3\xdc\x97\x72\x56\x1b\xef\x1c\xf2\xfd\x8f\x57\x73\xed\x21\x6d\xe0\x0d\x45\xe9\xdc\x6c\xb2\xef\x2e\x50\x67\x13\xd7\xa5\x03\x60\x8a\x64\xd6\xf2\x73\xc9\x3b\xe9\x6e\x77\xfd\x86\x8b\x9b\xf1\xb6\xf4\xa7\x52\x17\x54\xff\x80\x4c\x55\xef\x7d\x09\x2b\x59\xf7\x55\x0e\xd6\x7e\xc8\xad\xa7\xa7\x5d\x8c\x07\x07\xff\x44\x7d\xfc\x1f\xf4\xfe\x3d\xea\xf7\x03\xfb\x06\x58\x53\x56\xeb\x36\x0d\xeb\x94\xb8\x5e\xfc\x22\x4a\xe3\x41\x3f\xdc\x12\xa9\x65\xc3\x88\xe8\x0a\x2b\xe6\x1c\x21\x19\xb7\x8f\xc7\x17\x0c\x0f\x26\xa5\x75\xd1\x33\x74\xd5\xab\x9b\xf0\xfc\x0c\xa6\xaa\x99\xad\xf9\x9d\xc5\x6e\x10\x36\x7e\xc4\x5a\x86\xde\x64\x7c\x5a\x57\xd6\xb3\x64\x6d\x58\xc7\xb0\x58\xa0\x47\xe5\x24\x71\x90\x68\xdc\x6e\xdf\xd0\x82\x41\xd3\x04\xc4\x26\x74\x38\xbc\x18\x9d\xfb\x9c\xfa\xb2\xe8\xbb\x56\x92\x2c\xb0\xab\xc0\xef\x04\xc8\xe4\x9a\x0c\x24\xd6\xe2\x2a\x78\x93\xf8\xaf\x16\x6b\x30\x5f\x88\xbd\xbd\xf3\x6b\x1b\xf8\xe0\x15\x4d\x9f\xec\xc7\xd0\x34\x08\x6c\x48\x43\x50\x12\xd9\x92\xec\x35\x7d\x9b\x43\x5a\x3d\x5c\x7f\x12\x96\x47\x6f\x89\x6e\x49\xd1\xe9\xcf\xcc\x3d\xec\xbe\x18\xae\x30\x5d\x1a\xd4\x03\x9d\x03\x3e\xe7\x26\xc7\x45\x46\xb9\xd0\x91\x13\xd8\xda\x5d\xb9\x12\xf2\x2b\x00\x00\xff\xff\xa2\x13\xfe\x15\x82\x06\x00\x00"

func dataCommonAppDevMainShTplBytes() ([]byte, error) {
	return bindataRead(
		_dataCommonAppDevMainShTpl,
		"data/common/app-dev/main.sh.tpl",
	)
}

func dataCommonAppDevMainShTpl() (*asset, error) {
	bytes, err := dataCommonAppDevMainShTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/common/app-dev/main.sh.tpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _dataCommonAppDevUpstartConf = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x6c\x52\x6f\x6b\xdb\x3e\x18\x7c\xef\x4f\x71\x3f\xb7\xfc\xba\x41\x53\x77\xd9\x9f\x37\x23\x81\x2e\x63\x21\xd0\x26\x61\x69\x60\xd0\x95\x45\x96\x1e\x3b\x22\xaa\x64\x24\x39\x09\x74\xfd\xee\x7b\x1c\xb9\x2f\x06\x03\x63\xc4\x3d\xa7\xbb\x7b\xce\x56\x14\xa4\xd7\x4d\xd4\xce\x22\x9f\x38\x1b\x5a\x03\x51\x93\x8d\x79\x96\x85\x28\x7c\x04\x0f\x7c\x6b\x0d\xed\xc9\xe0\x61\xf8\xfe\xc3\xc7\x47\x1e\xb8\xe6\x6f\xfc\xbf\x34\xc8\x3c\x85\x46\x1c\x6c\x96\x9d\xe1\x7e\xab\x03\xf8\x89\x0e\x62\xef\xb4\xc2\xba\x49\x82\x9e\x06\x27\x92\xb6\x35\xe2\x96\xd0\x78\x27\x29\x04\xb4\x0d\x4b\x6e\x64\xca\x60\x48\xec\x69\x93\x59\xe7\x9f\x84\x01\x1d\x75\xc4\x35\x66\xf3\xfb\x4e\xf9\x64\xdf\xf3\x0e\xda\x18\x58\x17\xf1\x24\xfc\x8e\x0f\x8a\x20\x02\x2a\xa1\x0d\x29\x94\x6d\x64\xa1\x2a\x66\xbb\x8e\x15\x74\x6d\x59\xeb\x24\x92\x96\xce\x00\x5d\xe1\x01\x83\x0a\x79\x41\x51\x16\x81\xfc\x5e\x4b\x2a\x92\x78\x8e\xc7\xcf\x5d\x42\xcb\x3c\xe0\x0a\xff\xa0\xf0\xa4\xd2\x19\xbf\xcf\x70\x27\x76\x84\xd0\x7a\xea\x36\x6e\x03\xe7\x60\x4f\xd7\x7a\x4c\x96\xeb\x70\x89\x92\xa4\xe8\xd0\xbe\x63\x29\x2c\x4a\xe3\xe4\x0e\x02\x41\x6e\x49\xb5\x86\x3c\x9b\x79\x12\x8a\xe5\xe8\xd8\x38\xae\x6a\xba\xb8\xbb\xf9\xb1\xfc\xbe\x98\xac\x46\x1b\xdb\xf5\xb4\x49\x5e\x53\x8a\xa9\xba\xb6\x34\x5a\x62\xb6\x64\xf4\xcb\x6c\xfe\x75\xb4\xd1\x15\x07\xab\x74\x0d\x8a\xdb\x6b\xfc\x46\xed\xa9\x41\xae\x2d\x5f\x10\x4a\xf9\x9c\x21\x71\xd8\xe1\xe2\x99\x6b\xd7\x36\x72\xe0\x32\x44\xff\xe6\x7c\x78\xf9\xe9\x2d\x5e\x2e\x4e\xfa\x74\x24\x89\xa2\x0d\xbe\xe0\x80\xc2\x14\xa5\xb6\xfd\xba\xe9\xd7\xc0\xcf\x53\x21\x83\xe4\x34\x50\xda\x8f\x52\x7d\x89\x74\xa5\xf2\x57\x06\xdf\x54\xa3\xf3\x2e\x59\x8f\x9c\x3f\x4f\x16\xf3\xd5\xfa\xf6\xd7\xb7\xdb\x9b\xe9\xea\xa5\x47\xc7\xe3\x62\x2f\x3a\xb7\xfa\x55\x82\x8f\x18\x8e\xff\x7f\x97\x91\x55\xe8\x3f\xd6\x9f\x00\x00\x00\xff\xff\x72\xe2\xa9\xb0\xab\x02\x00\x00"

func dataCommonAppDevUpstartConfBytes() ([]byte, error) {
	return bindataRead(
		_dataCommonAppDevUpstartConf,
		"data/common/app-dev/upstart.conf",
	)
}

func dataCommonAppDevUpstartConf() (*asset, error) {
	bytes, err := dataCommonAppDevUpstartConfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/common/app-dev/upstart.conf", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _dataCommonAppDevDepMainShTpl = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x7c\x91\x41\x4b\xc3\x40\x14\x84\xef\xfb\x2b\xc6\xad\x88\x82\xcd\xaa\x14\x0f\x5a\x8a\x50\x2f\x5e\x54\xf4\xe0\xa1\x14\xd9\x26\x9b\x64\x75\x93\x17\xb2\x9b\x82\xc4\xfc\x77\x5f\x93\xe8\xa5\xd8\xe3\xec\x7c\xf3\x18\x66\x27\x47\x6a\x63\x4b\xb5\xd1\x3e\x17\xde\x04\x4c\x8d\x10\x64\x4e\xcf\xd0\xe2\xf8\x0e\x57\x8b\x93\x4b\x7c\xc3\x51\x96\x99\x1a\xd3\x00\x0a\x81\xb0\x80\x4a\xcc\x56\x95\x8d\x73\xb7\xe8\x04\xb9\x1e\x37\x71\x4e\x90\xab\x1d\xb1\xe6\xac\xdc\x59\x62\x82\xb7\xda\x06\x83\x90\x1b\x78\x53\x6f\x6d\x6c\x90\x5a\x67\x38\x04\xb9\xa4\x32\xb5\x59\x53\xdb\x32\x43\x4c\xa5\x6f\xdc\x2f\x73\x83\xb6\x85\xae\xaa\xf7\xb8\x47\xa2\xd7\xe1\xf9\x51\x17\x06\x5d\x27\x45\xac\x03\xe6\xf3\xfb\xa7\x25\x16\x2a\x14\x95\x1a\x63\xd1\x87\xa7\x52\xb4\x02\x90\xe3\x8b\xe4\x4b\x2c\xf9\xa1\xe4\x2c\x2b\x79\xf0\xf0\xf9\xc0\x06\x9d\x79\x66\x57\xeb\x51\x57\x54\x07\xf9\x4f\xa7\x67\xf6\x38\xca\x60\x27\x3a\xc1\x95\x78\x3e\xc4\x79\x41\x09\x2e\xae\x67\x33\xec\xf7\x63\xdf\x37\x09\xa1\xd8\xee\x9b\x50\x26\xc4\x6a\x18\x23\x4a\xfe\xac\x43\xa5\x87\xa3\xbc\xf4\x8b\x71\xa4\x93\x71\xc9\x08\x0f\x01\xd6\x83\x3e\xf5\x17\x6c\xca\x1f\xc0\x22\xd5\xd6\xf9\xa8\x2f\x38\xcc\x5d\xf7\x11\xf1\x13\x00\x00\xff\xff\xcc\x36\x71\x96\x05\x02\x00\x00"

func dataCommonAppDevDepMainShTplBytes() ([]byte, error) {
	return bindataRead(
		_dataCommonAppDevDepMainShTpl,
		"data/common/app-dev-dep/main.sh.tpl",
	)
}

func dataCommonAppDevDepMainShTpl() (*asset, error) {
	bytes, err := dataCommonAppDevDepMainShTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/common/app-dev-dep/main.sh.tpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"data/aws-vpc-public-private/deploy/main.tf": dataAwsVpcPublicPrivateDeployMainTf,
	"data/aws-vpc-public-private/deploy/module/join.sh": dataAwsVpcPublicPrivateDeployModuleJoinSh,
	"data/aws-vpc-public-private/deploy/module/main.tf": dataAwsVpcPublicPrivateDeployModuleMainTf,
	"data/aws-vpc-public-private/deploy/module/outputs.tf": dataAwsVpcPublicPrivateDeployModuleOutputsTf,
	"data/aws-vpc-public-private/deploy/module/variables.tf": dataAwsVpcPublicPrivateDeployModuleVariablesTf,
	"data/aws-vpc-public-private/deploy/variables.tf": dataAwsVpcPublicPrivateDeployVariablesTf,
	"data/common/app-dev/main.sh.tpl": dataCommonAppDevMainShTpl,
	"data/common/app-dev/upstart.conf": dataCommonAppDevUpstartConf,
	"data/common/app-dev-dep/main.sh.tpl": dataCommonAppDevDepMainShTpl,
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
		"aws-vpc-public-private": &bintree{nil, map[string]*bintree{
			"deploy": &bintree{nil, map[string]*bintree{
				"main.tf": &bintree{dataAwsVpcPublicPrivateDeployMainTf, map[string]*bintree{
				}},
				"module": &bintree{nil, map[string]*bintree{
					"join.sh": &bintree{dataAwsVpcPublicPrivateDeployModuleJoinSh, map[string]*bintree{
					}},
					"main.tf": &bintree{dataAwsVpcPublicPrivateDeployModuleMainTf, map[string]*bintree{
					}},
					"outputs.tf": &bintree{dataAwsVpcPublicPrivateDeployModuleOutputsTf, map[string]*bintree{
					}},
					"variables.tf": &bintree{dataAwsVpcPublicPrivateDeployModuleVariablesTf, map[string]*bintree{
					}},
				}},
				"variables.tf": &bintree{dataAwsVpcPublicPrivateDeployVariablesTf, map[string]*bintree{
				}},
			}},
		}},
		"common": &bintree{nil, map[string]*bintree{
			"app-dev": &bintree{nil, map[string]*bintree{
				"main.sh.tpl": &bintree{dataCommonAppDevMainShTpl, map[string]*bintree{
				}},
				"upstart.conf": &bintree{dataCommonAppDevUpstartConf, map[string]*bintree{
				}},
			}},
			"app-dev-dep": &bintree{nil, map[string]*bintree{
				"main.sh.tpl": &bintree{dataCommonAppDevDepMainShTpl, map[string]*bintree{
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
