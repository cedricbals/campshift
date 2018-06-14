// Code generated by go-bindata. DO NOT EDIT.
// sources:
// templates/index.html
// templates/layout.html

package main


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
	info  fileInfoEx
}

type fileInfoEx interface {
	os.FileInfo
	MD5Checksum() string
}

type bindataFileInfo struct {
	name        string
	size        int64
	mode        os.FileMode
	modTime     time.Time
	md5checksum string
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
func (fi bindataFileInfo) MD5Checksum() string {
	return fi.md5checksum
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _bindataTemplatesIndexhtml = []byte(
	"\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x54\xc1\x6e\x9c\x30\x10\xbd\xe7\x2b\x46\xa8\xc7\x1a\xb4\x6c\x1b\x45" +
	"\xd1\xee\x5e\x9a\xaa\xca\xa5\xad\xfa\x07\x06\xcf\x82\x1b\x63\x23\xdb\x2c\xac\x10\xff\x5e\x19\xc8\xee\x42\xbc\x09" +
	"\xe9\x21\x17\xe4\xc1\xf3\x9e\xdf\x3c\xcf\xb8\x6d\x35\x95\x19\xc2\x27\x46\x8f\x70\xbf\x85\xf0\x81\x5a\xda\x75\x37" +
	"\x1b\xc6\x0f\x90\x0a\x6a\xcc\x36\x48\x84\x4a\x9f\xa0\x11\xf7\xc3\xa2\x26\xfb\x4a\x08\x17\xd7\x64\x15\x7d\x85\x22" +
	"\x21\x5f\xa0\x3c\x92\x5b\x28\x1b\xb2\x0e\x76\x37\x00\x00\x97\xf8\x11\xb0\x17\xd8\xf4\x1f\xa2\x55\x0d\xdc\x62\x61" +
	"\x48\x8a\xd2\xa2\x76\x14\x77\xa0\x51\x50\xcb\x0f\x38\x32\xcc\x59\x4e\x2a\x72\xce\x18\x4a\xa8\xc9\xdd\x45\x66\x9f" +
	"\x9d\x54\xd6\x2a\x39\xfd\xd9\x6f\xf0\x22\x9b\x8a\x09\xc0\xe8\x74\x1b\x84\x61\x44\x8d\x41\x6b\x22\x5e\xd0\x0c\x4d" +
	"\x44\xb5\x56\x35\x11\xb8\xb7\xa1\x39\x64\x01\x44\xb3\x13\xa2\xf9\x11\x9b\x88\xf1\x83\x5f\x70\x5f\xea\x6a\xae\x31" +
	"\x8f\x9f\xf7\x2d\x36\x96\xc4\x8d\x80\x7e\x31\x38\x11\xec\xda\x16\xc2\x9f\xb4\xc0\xae\xdb\x44\x79\x3c\x07\xaf\x27" +
	"\x60\x1f\xf6\x81\xda\x01\xbb\x5e\x24\xf2\x03\x5d\xd5\x3c\xcb\xff\xc3\xd6\x8b\x65\xdb\xc2\xd8\xad\x99\x56\x55\xf9" +
	"\x28\x19\x36\x9f\xc7\xc0\x35\xaf\x6b\xe2\xf0\x87\x8b\xb8\xcc\x7e\x69\x86\x1a\xba\xce\x5b\x76\x61\xc9\x2d\x0c\xb5" +
	"\x7b\xda\x6e\x9e\xdd\x9b\x2c\x32\xd7\xa6\xf1\xe0\xb8\xeb\x90\xde\xef\xf1\x70\xe7\xf8\xd4\xe3\x69\x78\x52\xce\x9d" +
	"\xe6\x41\xe8\x77\x69\x35\x47\x73\xa6\xb8\x7a\xfe\xdb\x6a\xe7\x88\x71\xe4\x4a\xb2\x06\xad\x2a\xc9\x90\x81\xc9\x29" +
	"\x53\x35\x29\x18\x24\x19\xa9\x73\x6e\x11\x4a\xad\x32\x4d\x8b\x82\x24\xaa\xf1\x10\xfa\x88\x4f\x10\x63\x35\x2f\xd1" +
	"\x71\xb9\xae\xfb\xa6\x84\x72\x66\x07\xbb\x59\xe1\xaf\x71\x4d\x1f\x84\xb1\x85\xae\x43\x7d\x70\xb2\xba\x78\x56\x52" +
	"\x25\xe0\x6f\x65\x2c\xdf\x1f\x09\xed\xcb\x7e\x83\xcd\x7b\xcf\x8d\x70\xf7\x4c\x2b\xab\x20\xd1\x48\x9f\x48\xad\x34" +
	"\x33\xe7\xc1\x7c\x79\xd7\x57\xa9\xcb\x8b\x0b\x8c\x61\xaf\xa4\x25\xb2\x92\xdc\xaa\x71\x54\xd1\xa4\x9a\x97\x96\x2b" +
	"\xd9\x93\x96\xef\x53\xeb\x48\x07\xa2\x3f\x68\x4a\x25\x0d\x4f\x04\xfe\x46\x6d\x94\x34\xcb\x44\x2e\x49\x99\x34\xd5" +
	"\x2a\x5e\xe2\xe8\xe2\x77\xc1\x69\x7f\x4c\xfb\xea\xbd\xef\xc2\x3b\xf5\xbe\xb2\xed\xd9\x7a\x39\xa1\x28\xd9\xf3\x14" +
	"\x9e\xa3\x31\xad\x6d\x51\xb2\xae\xfb\x17\x00\x00\xff\xff\xd7\x0d\x39\x01\x2f\x07\x00\x00")

func bindataTemplatesIndexhtmlBytes() ([]byte, error) {
	return bindataRead(
		_bindataTemplatesIndexhtml,
		"templates/index.html",
	)
}



func bindataTemplatesIndexhtml() (*asset, error) {
	bytes, err := bindataTemplatesIndexhtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name: "templates/index.html",
		size: 1839,
		md5checksum: "",
		mode: os.FileMode(436),
		modTime: time.Unix(1528969697, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}

var _bindataTemplatesLayouthtml = []byte(
	"\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x54\xc1\x8e\x9b\x30\x10\xbd\xe7\x2b\x46\xbe\xb6\xc6\xc9\x6e\x15\x55" +
	"\x11\xa4\xaa\x7a\xea\x61\xab\x4a\x5d\xa9\x52\x6f\x06\x26\x60\x75\x6c\x23\x3c\x9b\x80\x22\xfe\xa9\xdf\xd0\x2f\xab" +
	"\x9c\x10\x2d\x4a\xa1\x87\xe6\x10\x06\xe6\xcd\x9b\xc7\xf3\x0c\x69\xcd\x96\xf6\xab\x55\x5a\xa3\x2e\xf7\x2b\x00\x80" +
	"\x94\x8c\xfb\x09\x75\x8b\x87\x4c\xd4\xcc\x4d\xd8\x29\x75\xf0\x8e\x43\x52\x79\x5f\x11\xea\xc6\x84\xa4\xf0\x56\x15" +
	"\x21\x7c\x38\x68\x6b\xa8\xcf\x3e\x5a\xcd\xa6\x78\xf3\xed\x93\x80\x16\x29\x13\x81\x7b\xc2\x50\x23\xb2\xf8\x7f\xd2" +
	"\x2f\x2f\xce\xb0\xdf\x3d\xae\xd7\x8b\xac\x6c\x98\x70\xff\x03\x89\x49\x57\xd8\xc2\xc3\x7a\xf3\x1e\xbe\x92\x76\x2f" +
	"\xae\x4a\xd5\x35\x3b\xe9\x7f\xcf\x02\xdc\x37\x98\x09\xc6\x8e\x63\x63\x31\x0a\xd4\x21\x20\x07\x75\x05\x26\x31\x31" +
	"\x92\x58\x64\x0d\x4e\x5b\xcc\xc4\xd1\xe0\xa9\xf1\x2d\x0b\x28\xbc\x63\x74\x9c\x89\x93\x29\xb9\xce\x4a\x3c\x9a\x02" +
	"\xe5\xe5\xe6\x2d\x18\x67\xd8\x68\x92\xa1\xd0\x84\xd9\x26\x59\x8b\xfd\x2a\x55\x57\xb3\x57\x69\xee\xcb\x1e\x0a\xd2" +
	"\x21\x64\x22\xba\x21\x83\x76\x01\x2e\x51\xee\xa9\x14\x11\x53\x9a\xe3\x0d\x92\x57\xb2\x47\x22\x7f\x82\xbc\x92\xb9" +
	"\x67\xf6\x56\x06\x6e\x4d\x83\x37\x81\x13\x70\x94\xa5\x8d\xc3\x16\x9a\x5e\x6e\xa1\xc9\xe5\x1a\x3a\xda\x35\xb9\xdc" +
	"\x8e\xe8\xfb\x8a\x03\x61\x07\xf1\x4f\x9e\x5a\xdd\x5c\xa3\xc2\x13\x18\x46\x1b\x64\x81\x8e\xb1\x9d\x94\xde\x97\x9f" +
	"\xe4\x66\x0b\xc1\xee\x4e\xf2\xf1\x21\x3a\xad\xd9\x1c\xf1\x0e\x7f\xa9\x31\xb6\x7a\xad\x39\xbc\x10\x81\x65\xb9\x15" +
	"\x10\xda\x22\x13\x49\xa2\x46\xfb\x8d\xd5\x15\x06\x15\xbd\x4d\xc2\xb1\x9a\x63\x9a\x76\x47\xcd\xf5\x5f\xfa\xfe\x81" +
	"\x94\xa6\xf0\x6e\x01\xbe\x24\xb3\x5d\x96\x39\xb2\xaa\xf3\x19\x92\xef\xd7\xf8\x73\xcc\xc0\x30\x2c\x69\x52\xa5\x39" +
	"\x2e\xa4\x42\xa3\xdd\xb2\xb4\xf8\x9b\xf4\x49\x9e\xb4\x71\xc9\x33\xda\xe6\x49\x77\x30\x0c\xbf\x7f\x2d\xbf\x94\x9a" +
	"\x67\x9e\xd1\x32\xf7\xa8\xde\xdc\x1c\x89\x2b\x23\x3b\x8a\xe7\x7d\x09\x1f\xba\xf1\x14\xf7\xcf\xd1\x8e\xa6\xf5\x55" +
	"\xab\xad\x4d\x55\xbd\x99\x0c\xdb\x2b\xe5\x18\x8e\x97\xd5\xec\xe0\xce\x0c\x75\x47\xbb\xbb\x29\xb5\xb9\x7c\x77\x5b" +
	"\xc1\x89\xd1\xe7\x33\xf4\x06\xa9\x84\x61\x98\xe9\x97\xaa\xb8\x79\xb1\xaf\xba\x7c\xff\xfe\x04\x00\x00\xff\xff\x44" +
	"\x1c\xa9\x32\x06\x05\x00\x00")

func bindataTemplatesLayouthtmlBytes() ([]byte, error) {
	return bindataRead(
		_bindataTemplatesLayouthtml,
		"templates/layout.html",
	)
}



func bindataTemplatesLayouthtml() (*asset, error) {
	bytes, err := bindataTemplatesLayouthtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name: "templates/layout.html",
		size: 1286,
		md5checksum: "",
		mode: os.FileMode(436),
		modTime: time.Unix(1528963507, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}


//
// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
//
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, &os.PathError{Op: "open", Path: name, Err: os.ErrNotExist}
}

//
// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
// nolint: deadcode
//
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

//
// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or could not be loaded.
//
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, &os.PathError{Op: "open", Path: name, Err: os.ErrNotExist}
}

//
// AssetNames returns the names of the assets.
// nolint: deadcode
//
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

//
// _bindata is a table, holding each asset generator, mapped to its name.
//
var _bindata = map[string]func() (*asset, error){
	"templates/index.html":  bindataTemplatesIndexhtml,
	"templates/layout.html": bindataTemplatesLayouthtml,
}

//
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
//
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, &os.PathError{
					Op: "open",
					Path: name,
					Err: os.ErrNotExist,
				}
			}
		}
	}
	if node.Func != nil {
		return nil, &os.PathError{
			Op: "open",
			Path: name,
			Err: os.ErrNotExist,
		}
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

var _bintree = &bintree{Func: nil, Children: map[string]*bintree{
	"templates": {Func: nil, Children: map[string]*bintree{
		"index.html": {Func: bindataTemplatesIndexhtml, Children: map[string]*bintree{}},
		"layout.html": {Func: bindataTemplatesLayouthtml, Children: map[string]*bintree{}},
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
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
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