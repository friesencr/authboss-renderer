// Code generated by go-bindata.
// sources:
// templates/confirm_email.html.tpl
// templates/confirm_email.txt.tpl
// templates/layout.html.tpl
// templates/login.html.tpl
// templates/recover.html.tpl
// templates/recover_complete.html.tpl
// templates/recover_email.html.tpl
// templates/recover_email.txt.tpl
// templates/register.html.tpl
// DO NOT EDIT!

package abrenderer

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
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
	name    string
	size    int64
	mode    os.FileMode
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

var _confirm_emailHtmlTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\xcc\xc1\x0d\x82\x50\x0c\x06\xe0\xfb\x9b\xe2\x0f\x03\xd8\x70\xaf\xbd\xb8\x80\x2b\xd4\xa6\xa4\x44\xa4\xa4\xc2\xc1\x90\xb7\xbb\x0b\xb0\xc0\xc7\x31\xca\x23\xd7\x69\xae\x0f\x7e\x79\x14\xd4\x2c\x8f\x75\x67\x8a\x51\x1a\xbf\x0a\x24\x8d\x37\x79\x2e\xae\x5f\x87\x2d\xb3\xbd\xc1\x8a\x28\x9f\xee\xc3\x79\xde\x7a\x1f\x40\x12\x5e\xce\xa4\x82\x3d\x61\x97\xdc\x26\xed\x1f\x00\x00\xff\xff\xd9\xba\x4f\x59\x6c\x00\x00\x00")

func confirm_emailHtmlTplBytes() ([]byte, error) {
	return bindataRead(
		_confirm_emailHtmlTpl,
		"confirm_email.html.tpl",
	)
}

func confirm_emailHtmlTpl() (*asset, error) {
	bytes, err := confirm_emailHtmlTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "confirm_email.html.tpl", size: 108, mode: os.FileMode(436), modTime: time.Unix(1517803040, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _confirm_emailTxtTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x24\xc8\xc1\x0d\x82\x40\x10\x05\xd0\xbb\x55\xfc\x0a\xac\xc5\x02\xb8\xac\xeb\xa0\x13\xc7\xff\xc9\xec\x10\x42\x08\xbd\x7b\xe0\xfa\x1e\x61\x6d\x18\xba\x96\x1d\x8d\x2f\x2c\x6d\x94\xa1\x3e\x86\x59\x11\xda\x9c\x6f\x84\xf3\x0b\x67\x09\xbb\xd6\xc4\x33\xb5\x0d\x4b\x94\xd0\xc5\xd9\xf3\x77\x79\xeb\x5d\x2b\x6b\xe2\xc4\xe3\xb8\x9f\xe7\xed\x1f\x00\x00\xff\xff\xd9\x6e\x6b\x96\x5c\x00\x00\x00")

func confirm_emailTxtTplBytes() ([]byte, error) {
	return bindataRead(
		_confirm_emailTxtTpl,
		"confirm_email.txt.tpl",
	)
}

func confirm_emailTxtTpl() (*asset, error) {
	bytes, err := confirm_emailTxtTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "confirm_email.txt.tpl", size: 92, mode: os.FileMode(436), modTime: time.Unix(1517803055, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _layoutHtmlTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb2\x51\x74\xf1\x77\x0e\x89\x0c\x70\x55\xc8\x28\xc9\xcd\xb1\xe3\xb2\x81\x50\x0a\x0a\x0a\x0a\x36\x19\xa9\x89\x29\x10\x26\x98\x5b\x92\x59\x92\x93\x6a\x57\x5d\x9d\x94\x93\x9f\x9c\xad\xa0\x04\xe6\x2a\x29\xe8\xd5\xd6\x56\x57\xa7\xe6\xa5\xd4\xd6\xda\xe8\x43\x54\x40\x34\xeb\x23\x74\xdb\x24\xe5\xa7\x54\x22\x0c\x82\x9b\x90\x58\x5a\x92\x91\x94\x5f\x5c\x8c\x6c\x08\x54\x33\x44\x87\x8d\x3e\xc4\x35\x80\x00\x00\x00\xff\xff\x8e\x27\x12\xb2\xa5\x00\x00\x00")

func layoutHtmlTplBytes() ([]byte, error) {
	return bindataRead(
		_layoutHtmlTpl,
		"layout.html.tpl",
	)
}

func layoutHtmlTpl() (*asset, error) {
	bytes, err := layoutHtmlTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "layout.html.tpl", size: 165, mode: os.FileMode(436), modTime: time.Unix(1517806846, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _loginHtmlTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x92\xdd\x8a\x9d\x30\x14\x85\xef\xcf\x53\x6c\xf6\xfd\xe8\x0b\xa8\x50\x68\x2f\x0a\xfd\x19\xda\xa1\xb7\x25\xc6\xed\x24\x4c\x7e\x64\x27\x9e\x1f\x42\xde\xbd\xa4\x6a\xab\xcc\xe1\xdc\x08\x2e\x97\xdf\x5a\xc9\xde\x29\x5d\x74\x54\x50\x8d\x46\x04\xf5\x3b\xcc\x52\x52\x08\x39\xa7\x54\x95\x07\xb9\x21\xe7\x53\x33\x7a\xb6\x20\x64\xd4\xde\xb5\x98\x92\xf5\xb3\x8b\x93\x88\x8a\x06\x40\xe3\x5f\xb5\xc3\x9c\x11\x2c\x45\xe5\x87\x16\x9f\xbf\xff\x7c\xc1\xee\x04\x00\xb0\xd1\x89\xd9\xf3\x4a\x6d\x7a\x86\xba\xdb\xd8\xc5\xd5\x68\x37\xcd\x11\xe2\x6d\xa2\x16\x23\x5d\x23\x82\x34\x22\x84\x16\x4b\xf0\x93\xf4\x2e\xb2\x37\x08\x4e\x58\x6a\x91\xac\xd0\x06\x61\x32\x42\x92\xf2\x66\x20\x6e\xf1\xd3\xd3\x22\x9e\x85\x99\xa9\x54\xac\x26\xd6\x56\xf0\xed\xf3\xc7\x5f\x45\xca\x19\xbb\x25\x77\x1f\xb8\x26\x4e\x22\x84\x8b\xe7\xe1\x61\xea\x7f\xd3\x21\xf8\x79\x93\xef\xe0\x17\xba\xd2\xc3\x40\x6e\xc3\xa4\x54\xc9\xc0\xe3\x37\x61\x4b\xa7\x5d\xdf\xa2\xbe\xf8\x37\x72\x45\xae\xb7\xdb\xd3\x23\x54\x41\xf9\xcb\x0f\xb2\x64\x7b\xe2\x9c\x0f\x6c\xa9\x48\xbe\xf5\xfe\xba\xd1\xd9\xfe\x43\x46\x9e\x09\x3b\xd8\x7e\x84\xaf\x74\xb8\xf1\x7e\x8e\xd1\xbb\x15\x13\xe6\xde\xea\x88\xdd\x97\x32\xca\xa6\x5e\xbe\xed\x0f\xb4\x2f\x22\xfd\xf9\x6f\x0f\x01\x8a\x69\x7c\xb7\x0e\xbc\x18\xca\x42\x74\xab\x19\x3e\x48\x59\x2c\x4d\x2d\x0e\x63\xdf\x53\x5f\x75\x88\x8f\xb1\x8b\x63\xe5\x2e\x2f\xf7\xc0\x4d\x5d\xa6\xd7\x9d\xfe\x04\x00\x00\xff\xff\xd7\xc4\xf9\xc5\xd8\x02\x00\x00")

func loginHtmlTplBytes() ([]byte, error) {
	return bindataRead(
		_loginHtmlTpl,
		"login.html.tpl",
	)
}

func loginHtmlTpl() (*asset, error) {
	bytes, err := loginHtmlTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "login.html.tpl", size: 728, mode: os.FileMode(436), modTime: time.Unix(1517807284, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _recoverHtmlTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x92\x31\x4f\xc3\x40\x0c\x85\xf7\xfe\x0a\xeb\x54\xd6\x64\x47\x97\x2c\x65\x61\x81\x0a\x2a\x56\x74\x4d\x9c\xe6\x44\xce\x77\x72\x9c\xd2\xea\xc8\x7f\x47\xd7\x54\x55\x0b\x52\x2b\x06\xa6\x24\xf6\x7b\xcf\xf1\x27\xeb\xc6\xb3\x03\x53\x89\xf5\x54\xa8\x18\x9d\x1f\x48\x82\x91\x16\x6b\x50\x8c\x95\xdf\x22\xab\x71\x54\xe0\x50\x5a\x5f\x17\x6a\xf9\xfc\xba\x52\xe5\x0c\x00\x40\x5b\x0a\x83\x80\xec\x03\x16\x4a\x70\x27\x0a\xc8\x38\x4c\x31\x59\x60\xeb\x0c\xef\x1f\x1f\x92\x37\x74\xa6\xc2\xd6\x77\x35\x72\x6a\x8a\x95\x0e\xe1\x52\xb2\x35\xdd\xf0\xc3\xf9\x96\x4a\xa9\x97\x97\x7a\xcd\x90\x4f\x43\x63\x9c\x07\x5b\xc3\x7d\x71\x11\x10\xe3\xa7\x95\x16\x32\x64\xee\x4f\x5f\x73\x64\xee\x6c\x2f\x49\x6c\xa9\xc6\x1d\x64\x90\xcc\x49\xc0\x86\x36\x78\x52\x8c\xa3\xee\x83\xa1\x32\xc6\x6c\x1c\x75\x7e\x78\x9f\x66\xc6\x88\x74\x30\x9c\x3f\xae\x6f\x5f\x79\x6a\x2c\xbb\xf7\xab\x14\x16\x93\x08\x6e\xd1\x38\x86\x2d\x6f\x42\xa9\x7e\x51\x81\x2f\x08\x6c\x49\x1a\x38\xfd\xd2\x5d\xaf\xfe\xc2\xaa\xfa\x1f\x58\xad\xad\x6b\xa4\xb3\x63\xd9\xf5\xdc\x3c\x19\x87\x97\xab\xa7\xea\xca\x7f\x20\x4d\xeb\x4e\x39\xeb\x41\xc4\xd3\x31\xa8\x1f\xd6\xce\x8a\x2a\x5f\xa6\x3b\xd5\xf9\xd4\x3d\x27\xa3\x0d\xb4\x8c\x4d\xa1\xf2\xce\x6f\x2c\xa9\x72\x61\xa8\xc2\x4e\xe7\xa6\x9c\xe9\x3c\x1d\x7f\x39\xfb\x0e\x00\x00\xff\xff\x48\x32\x12\xea\x04\x03\x00\x00")

func recoverHtmlTplBytes() ([]byte, error) {
	return bindataRead(
		_recoverHtmlTpl,
		"recover.html.tpl",
	)
}

func recoverHtmlTpl() (*asset, error) {
	bytes, err := recoverHtmlTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "recover.html.tpl", size: 772, mode: os.FileMode(436), modTime: time.Unix(1517803087, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _recover_completeHtmlTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x92\xb9\x4e\x03\x31\x10\x86\xfb\x3c\xc5\xc8\xa2\xde\xed\x91\xbd\x4d\x7a\x88\x20\x3d\x72\xd6\xb3\x59\x0b\x5f\x1a\x7b\x93\x20\xcb\xef\x8e\x9c\xcd\xc5\x25\x1a\xa8\xec\xb9\xbe\xd1\xaf\x7f\xf8\xe0\xc9\x82\xec\x93\xf6\x4e\xb0\x9c\xad\x9f\x5c\x0a\x32\x8d\xa8\x80\x11\xf6\x7e\x87\xd4\xf6\xde\x06\x83\x09\x59\x29\x0c\x2c\xa6\xd1\x2b\xc1\x56\x8f\xcf\x6b\xd6\x2d\x00\x00\xb8\x76\x61\x4a\x90\xde\x02\x0a\x36\x6a\xa5\xd0\x31\x70\xd2\xa2\x60\xc9\xbf\xd6\x60\x27\xcd\x84\x95\xdf\x1c\x13\x95\xd3\x7e\x33\x1b\x64\x8c\x7b\x4f\xea\x3c\x7d\x8d\x83\x91\x3d\x8e\xde\x28\x24\xc1\x56\x97\xf4\x89\x5b\x69\x7c\x43\x67\x66\xce\x7b\x9d\x46\x68\x90\x28\x96\x72\x8a\xee\x90\xc8\xe8\x98\xe0\x5e\x80\x76\x0a\x0f\xd0\xc0\x75\x41\x6d\x23\xe9\xb6\x78\xe9\x2b\x85\xc7\x20\x5d\x97\x73\x53\x0a\x6f\x8f\xff\x79\x47\xce\xe8\x54\x1d\xb8\x7d\x7e\x17\xd3\x7b\x37\x68\xb2\x2f\x3f\x88\x5a\xce\x65\xf8\x3b\x71\x5f\x16\xfe\x83\xc8\x8f\x6e\xe7\xdc\x1c\x22\x0d\x0f\xd2\x62\xb5\xf8\xea\x7a\xcd\xae\x3f\x3b\xbf\x99\x52\xf2\xee\x04\x8a\xd3\xc6\xea\xc4\xba\xa7\xf9\xe6\x78\x3b\x57\x6f\x95\x73\x09\x23\xe1\x20\x58\x6b\xfc\x56\x3b\xd6\x2d\xa5\xeb\xd1\xf0\x56\x76\x0b\xde\xd6\x43\xee\x16\xef\x01\x00\x00\xff\xff\x61\x7e\x06\xd4\xd0\x02\x00\x00")

func recover_completeHtmlTplBytes() ([]byte, error) {
	return bindataRead(
		_recover_completeHtmlTpl,
		"recover_complete.html.tpl",
	)
}

func recover_completeHtmlTpl() (*asset, error) {
	bytes, err := recover_completeHtmlTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "recover_complete.html.tpl", size: 720, mode: os.FileMode(436), modTime: time.Unix(1517803102, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _recover_emailHtmlTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\xcc\xd1\x09\xc2\x40\x0c\x06\xe0\xf7\x9b\xe2\xa7\x03\x18\xfa\x1e\x33\x83\xb8\x41\x0c\x91\x88\xc5\x94\xd8\x13\xe4\xb8\xdd\x5d\xc0\x05\x3e\x8e\x55\xae\x6e\xf9\xf1\xc2\x37\x7b\x41\xcd\xb2\xbf\x0e\xa6\x58\xa5\xf1\xad\x40\xd2\x78\x97\xcb\xe6\xfa\x76\xd8\xf6\xb0\x27\x58\x11\xe5\xf7\xf3\x32\xc6\x69\xce\x05\x24\xe1\xe5\x4c\x2a\x38\x12\xf5\x97\xdb\xa5\xfd\x02\x00\x00\xff\xff\x9d\xc9\x73\xb1\x6c\x00\x00\x00")

func recover_emailHtmlTplBytes() ([]byte, error) {
	return bindataRead(
		_recover_emailHtmlTpl,
		"recover_email.html.tpl",
	)
}

func recover_emailHtmlTpl() (*asset, error) {
	bytes, err := recover_emailHtmlTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "recover_email.html.tpl", size: 108, mode: os.FileMode(436), modTime: time.Unix(1517803126, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _recover_emailTxtTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x24\xc8\xc1\x0d\x82\x40\x10\x05\xd0\xbb\x55\xfc\x0a\xac\xc5\x02\xb8\xac\xeb\xa0\x13\xc7\xff\xc9\xec\x10\x42\x08\xbd\x7b\xe0\xfa\x1e\x61\x6d\x18\xba\x96\x1d\x8d\x2f\x2c\x6d\x94\xa1\x3e\x86\x59\x11\xda\x9c\x6f\x84\xf3\x0b\x67\x09\xbb\xd6\xc4\x33\xb5\x0d\x4b\x94\xd0\xc5\xd9\xf3\x77\x79\xeb\x5d\x2b\x6b\xe2\xc4\xe3\xb8\x9f\xe7\xed\x1f\x00\x00\xff\xff\xd9\x6e\x6b\x96\x5c\x00\x00\x00")

func recover_emailTxtTplBytes() ([]byte, error) {
	return bindataRead(
		_recover_emailTxtTpl,
		"recover_email.txt.tpl",
	)
}

func recover_emailTxtTpl() (*asset, error) {
	bytes, err := recover_emailTxtTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "recover_email.txt.tpl", size: 92, mode: os.FileMode(436), modTime: time.Unix(1517803143, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _registerHtmlTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x53\x4d\xaf\xdb\x20\x10\x3c\x3f\xff\x0a\x84\xde\xd9\xbe\x3f\x61\x2e\xe9\xa5\x97\x2a\xaa\xaa\x5e\x2b\x62\xd6\x31\x2a\x5f\x5a\x70\x93\x08\xf1\xdf\x2b\xec\xc4\x8e\x13\xb5\x4d\x2f\x36\xac\x66\x76\x67\x56\x03\xeb\x1d\x1a\x22\xba\xa8\x9c\x6d\x69\x4a\xc6\x8d\x36\x7a\x11\x07\x90\x84\x22\x1c\x55\x88\x80\x34\x67\x4a\x0c\xc4\xc1\xc9\x96\x7a\x17\x22\xe5\xd5\x1b\xd3\xe2\x00\x9a\xf4\x0e\x0b\xaf\xf6\xa8\x8c\xc0\xcb\xe7\x4f\x39\x53\x9e\x52\x54\x51\x03\xb9\xaf\x7e\xb0\x66\x62\x14\xaa\xb2\x7e\x8c\xc4\x0a\x03\x4f\x5c\x12\x2f\x1e\x5a\x1a\xe1\x1c\x29\xf9\x25\xf4\x38\x41\x4e\x2a\x0e\x77\xdd\xbe\x97\x7a\xce\x29\xd5\xe5\x03\x56\x16\xa2\xd7\xa2\x83\xc1\x69\x09\x93\xa2\x67\x05\x94\x34\x9c\x1d\x90\x34\xbc\x7a\x4b\xe9\xdd\x2b\x49\x3e\xda\x0d\xe2\x36\x08\x10\xc3\x72\x7b\x07\x44\xad\x42\x2c\x60\x65\x25\x9c\x49\x4d\x0a\xb9\x00\x50\xd8\x23\x2c\x88\x9c\x59\xf0\xc2\xf2\x49\x17\x6b\xa6\xf3\x3c\xf0\x2a\x72\xfb\xdb\xee\xd0\x8b\x10\x4e\x0e\x25\xe5\xfb\xeb\xe9\x4f\x1b\x5b\x90\xd7\x5d\xad\xf7\xcd\x0a\xf6\x4b\xf9\xde\xf6\xd6\xe0\xac\xbf\xbe\x75\x78\xd5\xc0\x56\x79\xe7\x6c\xaf\xd0\xfc\x58\x1d\xec\xe6\x0a\xf9\x97\x93\x27\xe6\xdf\x1d\x3d\xb6\x7d\xc1\xd9\xe3\x84\xff\x70\x38\x0b\x9d\x05\x85\xf1\x60\xd4\x1a\xc8\xaf\xb7\x87\xb1\x8c\x67\x82\x0c\x08\x7d\x4b\x1b\xca\x77\xc2\x76\xa0\x59\x23\x78\xf5\xd0\x66\x50\x52\x82\xa5\x6b\xf2\xcf\x01\xfb\x2f\xc2\x40\x09\xe7\x12\xf6\xa9\xfa\xcd\xfd\x04\x3b\x67\xb6\x62\x4d\x79\xa5\xbc\xfa\x1d\x00\x00\xff\xff\x32\x06\xee\xd5\xad\x03\x00\x00")

func registerHtmlTplBytes() ([]byte, error) {
	return bindataRead(
		_registerHtmlTpl,
		"register.html.tpl",
	)
}

func registerHtmlTpl() (*asset, error) {
	bytes, err := registerHtmlTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "register.html.tpl", size: 941, mode: os.FileMode(436), modTime: time.Unix(1517805610, 0)}
	a := &asset{bytes: bytes, info: info}
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
	if err != nil {
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
	"confirm_email.html.tpl": confirm_emailHtmlTpl,
	"confirm_email.txt.tpl": confirm_emailTxtTpl,
	"layout.html.tpl": layoutHtmlTpl,
	"login.html.tpl": loginHtmlTpl,
	"recover.html.tpl": recoverHtmlTpl,
	"recover_complete.html.tpl": recover_completeHtmlTpl,
	"recover_email.html.tpl": recover_emailHtmlTpl,
	"recover_email.txt.tpl": recover_emailTxtTpl,
	"register.html.tpl": registerHtmlTpl,
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
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"confirm_email.html.tpl": &bintree{confirm_emailHtmlTpl, map[string]*bintree{}},
	"confirm_email.txt.tpl": &bintree{confirm_emailTxtTpl, map[string]*bintree{}},
	"layout.html.tpl": &bintree{layoutHtmlTpl, map[string]*bintree{}},
	"login.html.tpl": &bintree{loginHtmlTpl, map[string]*bintree{}},
	"recover.html.tpl": &bintree{recoverHtmlTpl, map[string]*bintree{}},
	"recover_complete.html.tpl": &bintree{recover_completeHtmlTpl, map[string]*bintree{}},
	"recover_email.html.tpl": &bintree{recover_emailHtmlTpl, map[string]*bintree{}},
	"recover_email.txt.tpl": &bintree{recover_emailTxtTpl, map[string]*bintree{}},
	"register.html.tpl": &bintree{registerHtmlTpl, map[string]*bintree{}},
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
