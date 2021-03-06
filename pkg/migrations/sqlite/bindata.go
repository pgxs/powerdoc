// Code generated by go-bindata. (@generated) DO NOT EDIT.

// Package sqliteMigrations generated by go-bindata.// sources:
// 1_create_users_table.down.sql
// 1_create_users_table.up.sql
package sqliteMigrations

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
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
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

// ModTime return file modify time
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

var __1_create_users_tableDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\xf0\x74\x53\x70\x8d\xf0\x0c\x0e\x09\x56\x28\x2d\x4e\x2d\x2a\xb6\xe6\x02\x04\x00\x00\xff\xff\x2c\x02\x3d\xa7\x1c\x00\x00\x00")

func _1_create_users_tableDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__1_create_users_tableDownSql,
		"1_create_users_table.down.sql",
	)
}

func _1_create_users_tableDownSql() (*asset, error) {
	bytes, err := _1_create_users_tableDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1_create_users_table.down.sql", size: 28, mode: os.FileMode(420), modTime: time.Unix(1592612455, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __1_create_users_tableUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x90\x41\x4b\xc4\x30\x10\x85\xef\xf9\x15\x73\x6c\xc1\xc3\xae\x22\x08\x7b\x8a\xee\x28\xc5\x58\x24\x64\xc1\x3d\x85\x61\x33\xac\x81\xb4\x95\x24\x45\x7f\xbe\x48\x5a\x6b\xd5\x77\x9c\xef\x7b\x73\x78\x77\x1a\xa5\x41\x30\xf2\x56\x21\x8c\x89\x63\x12\x95\x00\x00\x18\xbd\x83\x75\x9a\xd6\xe0\x03\x6a\x78\xd6\xcd\x93\xd4\x47\x78\xc4\xe3\x45\x51\x13\xc7\x9e\x3a\x5e\xd4\xd3\x2b\xc5\xea\x7a\x53\x17\xfe\x46\x29\xbd\x0f\xd1\xfd\xe2\xdb\xcb\x9b\x49\x58\x97\xbf\x85\xab\xf9\x01\x77\xe4\xc3\x5f\xbe\xdd\xcc\x42\xa0\x94\x6d\x18\xce\x67\x76\xd6\xf7\x96\x32\x38\xca\x9c\x7d\xc7\x85\x9f\x22\x53\x66\xf7\x05\xa6\xac\xb9\xe3\xc0\xff\x73\xd8\xe3\xbd\x3c\x28\x03\xed\x41\x29\x51\xef\xc4\x34\x58\xd3\xee\xf1\x05\xbc\xfb\xb0\x3f\xba\x43\x5f\x26\x84\x6a\x39\xd6\xbb\xcf\x00\x00\x00\xff\xff\x0a\x53\xe0\x62\x63\x01\x00\x00")

func _1_create_users_tableUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__1_create_users_tableUpSql,
		"1_create_users_table.up.sql",
	)
}

func _1_create_users_tableUpSql() (*asset, error) {
	bytes, err := _1_create_users_tableUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1_create_users_table.up.sql", size: 355, mode: os.FileMode(420), modTime: time.Unix(1592653395, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
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
	"1_create_users_table.down.sql": _1_create_users_tableDownSql,
	"1_create_users_table.up.sql":   _1_create_users_tableUpSql,
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
// AssetDir("foo.txt") and AssetDir("nonexistent") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
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
	"1_create_users_table.down.sql": &bintree{_1_create_users_tableDownSql, map[string]*bintree{}},
	"1_create_users_table.up.sql":   &bintree{_1_create_users_tableUpSql, map[string]*bintree{}},
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}
