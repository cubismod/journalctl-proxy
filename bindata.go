package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
)

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

var _assets_index_html = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x57\x7b\x6f\xdb\xc8\x11\xff\x5f\x9f\x62\x8e\x77\x05\x49\xc0\x22\xed\x5c\xda\xa6\x12\x29\x20\x67\xab\xad\x8b\x38\x17\x9c\x14\xa4\x41\x1a\x24\x2b\x72\x44\x6e\xbd\xdc\x65\x76\x47\x92\xd5\xc0\xdf\xbd\x58\xbe\x44\x4a\xb2\x81\xf8\x0f\x6b\x1f\xf3\x9b\xc7\x6f\x1e\x24\xa3\x9f\x6e\x7e\xbf\x5e\x7e\x7c\x37\x87\x9c\x0a\x31\x1b\x45\xf6\x07\x04\x93\x59\xec\xa0\x74\x66\xa3\x11\x40\x94\x23\x4b\x67\x23\x00\x80\xa8\x40\x62\x90\xe4\x4c\x1b\xa4\xd8\x79\xbf\xfc\xfb\xf8\x95\xd3\xbf\xca\x89\xca\x31\x7e\xdb\xf0\x6d\xec\xfc\x7b\xfc\xfe\xf5\xf8\x5a\x15\x25\x23\xbe\x12\xe8\x40\xa2\x24\xa1\xa4\xd8\xb9\x9d\xc7\x98\x66\x38\x40\x4a\x56\x60\xec\x6c\x39\xee\x4a\xa5\xa9\x27\xbc\xe3\x29\xe5\x71\x8a\x5b\x9e\xe0\xb8\xda\x5c\x00\x97\x9c\x38\x13\x63\x93\x30\x81\xf1\x55\x70\xd9\xaa\x22\x4e\x02\x67\xff\x55\x1b\x2d\x99\x48\x48\x44\x61\x7d\x32\xaa\xaf\x0d\xed\xed\x06\xaa\xbf\x95\x4a\xf7\xf0\xbd\xd9\x00\xac\x95\xa4\xf1\x9a\x15\x5c\xec\x27\xe0\x5e\xab\x8d\xe6\xa8\xe1\x2d\xee\xdc\x0b\x68\x76\x17\x50\x28\xa9\x4c\xc9\x12\x9c\x0e\x71\x86\xff\x0f\x27\x70\xf5\x6b\xf9\xd0\x5e\x3c\x8e\x9a\x85\x65\x0f\x75\xcf\x50\xca\x4d\x29\xd8\x7e\x02\x6b\x81\x0f\x07\x3d\x4c\xf0\x4c\x8e\x39\x61\x61\x26\x90\xa0\x24\xd4\x4f\xe9\xca\xaf\xfa\x7e\x0b\x7c\x18\x67\x5a\xed\x26\x70\x75\x02\x20\xb6\x12\xd8\x13\x5e\x29\x9d\xa2\x1e\x27\x4a\x28\x3d\x81\x9f\x5f\xbe\x7c\xf9\xd7\x3f\xb3\xe9\x99\x6b\xc1\x4a\x83\x13\x68\x57\xa7\x9e\x50\x21\x3e\xa5\x8c\xd8\x98\x72\x2c\x30\x76\x53\xa6\xef\xdd\xcf\x3d\x53\x00\x2b\x96\xdc\x67\x5a\x6d\x64\x3a\x81\x9f\x5f\xbc\x7a\xc1\x7e\xfd\xcb\xb4\x77\xdd\x3a\xb1\x7e\xb5\x7e\xb5\x7e\x71\x64\x21\x0a\x9b\x5c\xd9\x02\x0c\xeb\x0a\xb4\x4b\x9b\xb4\x2e\x9b\x89\xe6\x25\xb5\xe9\x1c\xb5\x5a\xa5\xa1\xca\x3d\x88\x21\x55\xc9\xa6\x40\x49\x41\x86\x34\x17\x68\x97\xe6\xb7\xfd\x92\x65\x6f\x59\x81\x9e\x6b\xa5\x5c\xff\xd3\xe5\xe7\x36\xaa\xf5\x46\x26\xc4\x95\x04\x52\x59\x26\x70\x69\x43\xf3\xfc\x5e\x50\x02\x09\x92\x8d\xd6\x28\xa9\xba\x84\xb8\x32\x15\x58\x26\x0c\x52\x50\x91\xd1\x49\xf3\x35\x78\x43\xe9\x38\x86\x9a\x29\xff\x0c\x0e\x62\x70\x05\xcf\x72\x72\x3b\x0d\x28\x0c\x3e\x21\x59\xa9\xe9\x48\x6b\x38\x6b\x18\xa9\xb7\x75\xb5\xb4\xfc\x44\xf9\x55\xaf\x31\xa0\xd4\xea\x61\x1f\x85\xf9\x55\x77\x9f\xf2\xed\xac\x33\x1c\x19\x14\x98\x10\x28\x99\xe4\x4c\x66\x18\x3b\x86\x98\xa6\x37\xdc\x10\x4a\x2e\xb3\xa5\x5a\xa0\xb6\x2d\xe9\xf9\xce\xac\x97\xd4\x48\x95\x96\xc0\x19\x13\x22\x0a\x9b\xf5\x41\x69\x58\x6b\xed\x9d\xac\x36\x44\x4a\x5a\x33\x82\x27\xf7\xb1\x33\x20\xde\x99\x2d\xab\x2d\xdc\xa9\x14\xa3\xb0\x96\xed\xfc\x0d\x3b\x87\xeb\x02\xb1\xa1\x8e\xda\x48\xa0\x2a\x9f\xd8\x51\x5b\xd4\x6b\xa1\x76\xe3\xfd\x04\x4c\xa2\x95\x10\x53\xc8\xd1\x92\x3c\x81\xbf\x5d\x6e\xf3\x29\xa8\x0d\x09\x2e\x71\x5c\x01\x26\xc0\x36\xa4\xa6\x0e\xf0\x34\x76\x6a\x79\xd4\x5d\x80\x51\xdd\x4e\xf5\x44\x72\xae\x2e\x2f\xff\xd4\x8b\x3d\xa2\xc3\x9c\x6c\x4f\x74\x7f\x5b\x89\xcc\xde\xdd\xde\x44\x21\xe5\xa7\x17\xef\x25\xa7\xf3\x37\x4b\x5e\xa0\x21\x56\x94\xe7\xaf\xef\xd0\x18\x96\xe1\xf1\x65\x14\xf6\xcd\xdb\xdb\x81\x7b\x11\x55\xad\x14\x85\xf5\x6f\x47\x6a\x15\x62\x4b\x6b\xc7\xf0\x51\xc7\x8d\x0e\xdd\x60\x54\x72\x8f\x04\x31\xc8\x8d\x10\xd3\xd1\xa0\x0b\x9b\x1a\xea\xf5\xe1\xb7\x0d\xea\xfd\xa2\x3a\x56\xda\x73\x6b\x01\xd7\x9f\x0e\x60\x35\xcb\x4f\xa3\xaa\x7b\x0b\x6a\x50\x1e\x33\x7b\x99\x80\xe7\x43\x3c\xeb\xf5\x6a\xeb\x43\x55\xa6\xe6\x0f\xfc\xb6\x41\x63\x9d\x61\x3b\xc6\x09\xd6\x48\x49\xee\xb9\xa1\xe0\x86\xc6\xad\xd0\xc1\x93\x63\xb8\x2d\xfc\x0e\x7b\xa4\x33\x20\x7c\x20\xef\xe0\x8f\x7d\x26\x68\xf0\x2c\x3b\xb6\xb2\x40\xad\x07\x6a\x02\x53\x0a\x4e\x9e\xfb\x1f\xe9\xfa\xfe\x60\x60\x56\xd2\x71\xf5\x13\x90\xe6\x85\xe7\xf7\xe7\xa5\x9d\x26\xb5\x84\x9d\x22\xae\x0f\x2b\x8d\xec\xbe\x67\xb5\x16\xf9\xa9\x82\x57\xdd\x6a\x3e\x70\xca\x3d\xf7\xfd\xdb\xdb\xe5\xb1\xa9\xa3\xf8\xec\x3c\x6c\x2d\x37\xee\x41\x35\x19\xa7\x67\x30\x75\x5b\xf7\x13\x94\x68\x64\x84\xcd\x84\xf5\xdc\x5a\xc0\xf5\x87\xe0\xfa\x34\xd8\x32\xb1\xb1\xb6\x7a\x96\x5b\x93\xc1\x19\x93\x0d\x8a\x4b\x89\x7a\x89\x0f\x34\x44\x0e\x65\xeb\x72\x0a\x58\x59\xa2\x4c\xaf\x73\x2e\x52\xaf\x86\xfb\x3d\xb1\xc7\xd1\xf1\xea\xd1\xef\x25\xaf\x7b\x04\x3c\x39\xef\x7a\x3c\x5a\xbe\xeb\x06\x18\xb2\x5b\x9f\x05\x89\x50\x06\xbd\xab\xcb\xcb\xcb\x0b\x70\x3f\x28\x7d\x0f\x89\x2a\x4a\x81\x84\x7d\x6e\x1e\x0f\x19\x3c\x34\x13\xee\xe0\x03\xae\x16\xd5\xde\xfb\xba\x33\x93\x30\xfc\xe5\xfb\x8e\xcb\x54\xed\x02\xa1\x12\x56\x91\x92\x2b\x43\x8f\xe1\xce\x84\xbf\x7c\x6f\x42\xaf\xc8\x7d\xfc\xda\xaf\xc5\xc6\x17\x25\x55\x89\x36\x69\x5d\x80\x1e\x0e\x9d\xb6\xc9\x55\x02\x03\xa1\x32\xef\xeb\x27\x2b\xfd\x19\xae\x95\x94\x58\x8b\xdb\x11\xb4\x12\xdc\xe4\x98\xc2\x8e\x53\xde\xa6\x61\x02\x67\xac\x77\xb1\x9d\x73\xa4\xa8\x07\xd6\xd0\x97\x2d\x4a\x3a\xf5\x87\x80\xcb\x44\x15\x5c\x66\x77\x1d\xe8\x5f\x8b\xdf\xdf\x06\xa5\x7d\x05\xad\x51\xd5\xa3\xd1\x1f\xf4\x41\x0d\xd5\x6a\xf7\x4c\x91\x92\x1e\x16\x68\x8d\x29\x79\xfa\x1c\x26\x3d\x8b\xd1\x5c\x69\x4e\xfb\x1f\x06\x9a\xbd\x21\x2c\x52\x3b\xf9\x7f\x18\xab\x91\x09\xe2\x05\x76\x0f\x87\x1f\xd6\x70\x48\xc3\xf3\xb8\x73\xa6\x53\x3b\x82\x6f\xda\x4a\xbd\x61\x84\x5e\x95\x91\x5b\x49\xde\x51\xc6\x82\x2f\x5f\xfe\x98\xbf\x7e\xb3\xbc\xbd\x9b\x7f\xb1\xff\x16\xcb\xd7\x77\xef\xfc\xd0\x76\xc5\x50\x79\xc9\xd3\x41\x9b\x9f\xe8\x79\x77\x7b\xd3\x8f\xa2\xc7\xde\xf3\xb8\xc5\xc7\xc5\x72\x7e\x77\xf3\xc5\xce\xc1\xbe\x82\x13\x0a\x07\x6a\x0e\x51\x06\xa4\xde\x28\xfb\xa5\xb1\x20\xcd\x65\x36\x9c\xca\x0d\x8b\xcf\x7a\x70\x37\x5f\x2c\x5e\xff\x63\x7e\x86\xca\xf6\x15\xe3\xfc\xbb\xea\x6f\xfb\xdb\xd4\x73\x5b\x19\xd7\xef\xe3\xb5\xda\x0d\x06\x5d\xc9\xd3\x81\x5f\xc7\xf7\x3d\xb6\x9e\x95\x3b\x21\xe5\x59\xe9\x26\xfa\x61\x22\xab\xe7\xf3\x50\xa9\xda\x0d\xd4\xb4\x21\x05\xf5\x62\xa9\xbc\xef\xa4\xca\xc9\xf1\xf9\x3f\xab\x17\xb5\x0b\x58\x61\xce\xb6\xdc\x7e\x2c\x38\xa6\x50\x8a\x72\xe7\xd1\xef\x8f\x98\x66\x31\x3a\x7d\x0b\x8e\xc2\xe6\x83\x21\x0a\xeb\x0f\xdc\xff\x07\x00\x00\xff\xff\xe6\x6d\x8b\xc7\xf1\x0e\x00\x00")


func assets_index_html() ([]byte, error) {
	return bindata_read(
		_assets_index_html,
		"assets/index.html",
	)
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		return f()
	}
	return nil, fmt.Errorf("Asset %s not found", name)
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
var _bindata = map[string]func() ([]byte, error){
	"assets/index.html": assets_index_html,
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
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func func() ([]byte, error)
	Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"assets": &_bintree_t{nil, map[string]*_bintree_t{
		"index.html": &_bintree_t{assets_index_html, map[string]*_bintree_t{
		}},
	}},
}}
