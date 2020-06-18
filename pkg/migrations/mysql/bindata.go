// Package mysqlMigrations Code generated by go-bindata. (@generated) DO NOT EDIT.
// sources:
// 1_create_users_table.down.sql
// 1_create_users_table.up.sql
package mysqlMigrations

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

	info := bindataFileInfo{name: "1_create_users_table.down.sql", size: 28, mode: os.FileMode(420), modTime: time.Unix(1592299585, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __1_create_users_tableUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\xd0\xc1\x4a\x03\x31\x10\x80\xe1\xfb\x3e\xc5\x1c\x77\xc1\x43\x45\x3c\xf5\x14\xdb\x11\x16\xd3\x28\x21\x0b\xf6\x14\x86\x66\xa8\x81\x6c\x56\x92\x14\x7d\x7c\x91\xb4\xae\xab\x9d\x63\xbe\xf9\x73\x98\x8d\x46\x61\x10\x8c\x78\x90\x08\xa7\xcc\x29\x37\x6d\x03\x00\x70\xf2\x0e\x96\xe3\x63\x81\x17\xdd\xef\x84\xde\xc3\x13\xee\x41\x0c\xe6\xd9\xf6\x6a\xa3\x71\x87\xca\xdc\xd4\x2a\x73\x8a\x34\xf2\x5c\x1d\xde\x28\xb5\xf7\xab\xae\xfa\x3b\xe5\xfc\x31\x25\xf7\xc7\x6f\x57\x97\x85\x65\xfc\xb3\x70\x77\x71\x1e\xc9\x87\xff\x3e\x7f\x10\x28\x17\x1b\xa6\xe3\x91\x9d\xf5\xd1\x52\x01\x47\x85\x8b\x1f\xb9\xfa\x21\x31\x15\x76\xdf\x70\x9e\xa5\x3b\x0e\x7c\xdd\x61\x8b\x8f\x62\x90\x06\xd4\x20\x65\xd3\xad\x9b\xf3\xed\x7a\xb5\xc5\x57\xf0\xee\xd3\xfe\x6a\xa7\x58\xaf\x09\xed\xfc\xd8\xad\xbf\x02\x00\x00\xff\xff\x23\xef\xad\xda\x6e\x01\x00\x00")

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

	info := bindataFileInfo{name: "1_create_users_table.up.sql", size: 366, mode: os.FileMode(420), modTime: time.Unix(1592464855, 0)}
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
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
