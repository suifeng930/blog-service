// Code generated for package configs by go-bindata DO NOT EDIT. (@generated)
// sources:
// configs/config.yaml
package configs

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

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _configsConfigYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x52\xcb\x6e\xdb\x30\x10\xbc\xeb\x2b\x16\xe8\xb9\xb2\x1e\x91\x1f\x3c\x35\x76\x1c\xd4\x45\xdc\x1a\x95\x82\x1c\x8b\xb5\xb5\x96\x15\x50\x22\x4d\xae\x1c\xb9\x5f\x5f\x50\xb2\xad\x04\xe9\x8d\x98\x59\x0e\x77\x66\x98\x92\x39\x91\x11\x1e\xc0\xef\xa6\x5e\xab\x9c\x04\xe4\xb4\x6d\x0a\x0f\xe0\x3b\xb3\xde\x28\xc3\x02\xa6\x41\x10\xb8\x09\xc2\x3c\x2b\x2b\x52\x0d\x0b\x18\x3b\xe4\xc5\x94\x4c\xef\xa1\x7b\xad\x9d\xd6\x03\xed\xb1\x91\xbc\x50\x35\x53\xcb\x1f\xef\x5c\xb8\x0d\x16\x94\x96\x7f\x49\x40\xe8\xd0\x35\xb6\xef\x11\x07\x3d\xa9\x22\xc5\x13\x6d\x90\x0f\x02\x2c\x2b\x83\x05\x8d\xa4\x2a\x6c\xcf\x3d\x96\x92\x7e\x62\x45\x02\x50\xeb\x01\x5a\xb6\x2c\xc0\x97\xca\x39\x78\xd6\x52\x61\xfe\x59\xa4\xe9\xf0\x61\xa0\xcb\xe0\xd9\x48\x01\x07\x66\x2d\x46\xa3\x30\x9a\xf8\x81\x1f\xf8\xa1\x70\xd6\x47\x96\x91\xcb\xdd\x6d\x7e\x55\x61\x41\x6b\x6c\xfb\x65\x13\xf8\x92\xac\xe7\x1f\xc9\x7b\x29\xd5\xdb\xb2\x65\xeb\xc2\x00\xf8\x0a\xfe\xab\x2e\x86\x23\xdd\xce\xba\x2e\xbc\x07\x64\x9c\xa3\xa5\x2e\xb8\x79\x76\xd6\x24\xa0\x3a\xdb\xa3\x74\x9a\x96\x4c\xdd\x99\x34\x4a\xb1\x07\xb0\x41\x6b\xdf\x94\xc9\x05\x84\x51\x7c\x97\x8c\x27\x53\xd7\x94\xb2\xec\x80\xeb\xd2\x71\x1c\x8c\x3b\xb1\x3e\x9f\xad\x54\xc5\x1f\x4b\xe6\x54\xee\xc8\x03\xc8\x70\x2b\x69\x63\x68\x5f\xb6\x17\xce\x03\x58\x1c\xd0\x58\x62\x01\x0d\xef\xa7\xdd\x3b\xc6\x76\xcd\x0a\xc8\x4c\x43\x7d\x43\xab\x5c\xd2\x42\xd5\xb5\x1d\x4a\xfb\xa5\xa9\xbe\x40\x71\xe0\xfd\x78\xc9\x9c\x8b\x94\x76\xc6\x69\x51\x9e\x9f\x77\xaf\x67\x0f\x60\x65\x6d\x43\xe6\xd3\x2a\xcb\x56\x97\x86\x04\x4c\xa2\x20\xf0\x96\x15\x96\x52\xdc\xec\xd8\x8a\xb5\x7f\x3c\xfa\x3b\x55\xb9\x7d\xba\x8f\x78\x37\x4e\x2e\xa1\xf4\xce\x66\xe3\x59\x7c\x37\x49\xa2\xd9\xb7\x61\xf0\x1a\x50\xf7\x6a\x9a\x3e\x09\xe0\xde\xc0\xa3\x32\xd5\x7f\xaf\x64\xea\x5a\x53\x14\x24\xe1\x2c\x0a\xe3\x68\x72\x65\xff\x05\x00\x00\xff\xff\x39\x3d\xe2\x81\x20\x03\x00\x00")

func configsConfigYamlBytes() ([]byte, error) {
	return bindataRead(
		_configsConfigYaml,
		"configs/config.yaml",
	)
}

func configsConfigYaml() (*asset, error) {
	bytes, err := configsConfigYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "configs/config.yaml", size: 800, mode: os.FileMode(420), modTime: time.Unix(1608389888, 0)}
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
	"configs/config.yaml": configsConfigYaml,
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
	"configs": &bintree{nil, map[string]*bintree{
		"config.yaml": &bintree{configsConfigYaml, map[string]*bintree{}},
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
