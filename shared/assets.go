package shared

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
	"os"
	"time"
	"io/ioutil"
	"path"
	"path/filepath"
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

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindata_file_info struct {
	name string
	size int64
	mode os.FileMode
	modTime time.Time
}

func (fi bindata_file_info) Name() string {
	return fi.name
}
func (fi bindata_file_info) Size() int64 {
	return fi.size
}
func (fi bindata_file_info) Mode() os.FileMode {
	return fi.mode
}
func (fi bindata_file_info) ModTime() time.Time {
	return fi.modTime
}
func (fi bindata_file_info) IsDir() bool {
	return false
}
func (fi bindata_file_info) Sys() interface{} {
	return nil
}

var _groups_yml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x5c\xef\xae\xdc\x36\x76\xff\xee\xa7\x20\x8c\x02\xb5\x81\xeb\xd9\x74\xb3\xd8\x0f\xe9\xa7\x38\xc9\x62\x53\xc4\xcd\xd6\x76\x9a\x16\x08\x30\xe0\x48\x9c\x19\xee\xd5\x88\xb3\xa2\x74\xc7\x93\x4f\x79\x8d\x02\xdb\x97\xcb\x93\xf4\xfc\x23\x45\x4a\xd4\xcc\xfd\x63\x77\xb7\x40\x81\xc0\xb9\x23\x52\xe4\x39\x3f\x1e\x1e\x9e\x7f\x94\x6d\x7d\xaf\xdb\xca\x7c\xf1\x4c\xa9\xe1\xd8\xdb\x03\xfd\xa5\xd4\xde\x34\xc7\x2f\xd4\xf3\xf7\x7b\xa3\xee\x74\x33\x18\xe5\xb6\xaa\x87\x1f\xdc\x47\x6d\xad\x69\x6a\x55\xb9\xae\x33\xfe\xe8\xda\xda\xab\xde\x51\x7b\x3b\x1c\x36\xa6\xc3\xde\xde\x54\xdc\xb0\xd7\x3d\x35\x1d\x5c\xbb\x73\x5e\xb9\x8e\xff\xaa\xd5\xb1\x73\x95\xf1\x5e\xed\xb5\x57\x1b\x63\x5a\xa5\xab\xde\xde\x99\xd5\x73\xa2\xa0\x3f\x1f\xcd\x17\x30\xc5\xd0\xf6\xa6\x8b\xd4\xad\x8d\x87\xff\xe9\x3e\x27\x93\xdb\xbe\x91\x26\x1c\xf8\xce\xd6\xc6\xa7\x14\xc3\x1c\x95\x6e\xaa\xa1\x81\x0e\xb5\xda\x76\xee\xa0\xde\x20\x19\x5f\xbf\xfe\x47\xaf\x2c\xce\xd1\xea\x06\xa7\xeb\xbc\x79\xb5\xeb\xb4\x6d\xa1\x1b\xbd\x79\x6b\xcc\xd1\xb6\x3b\xe5\xcf\xbe\x37\x87\x05\xea\x1a\x07\xa3\xaf\x8b\xf8\x51\xd3\x7b\x1c\x89\x91\xb4\x4c\x57\x35\x00\x78\x6d\x4f\x73\xdc\x00\xeb\x00\x66\x8d\xd3\x08\x90\xde\x74\x77\xa6\xbb\x01\xd2\xd4\x0f\xef\xbf\x52\xfe\x68\x2a\x0b\xa8\xd7\xf8\x40\xb7\xea\xdb\x77\xdf\x7f\x8d\xac\x6e\x5d\x07\x2c\x17\x89\x7a\xa6\x3d\x0c\xd2\x7b\x24\xa8\x33\x3b\xe0\xbc\x9b\xd1\x26\xcf\xc3\x3b\xaa\xef\x74\x75\xeb\x27\x2b\x19\x3a\xf1\x78\xd6\xb5\x5e\x01\x3e\x1e\x68\xf1\x16\x44\x27\x21\x37\x2e\x29\x08\x55\x07\x38\xaf\xd4\x57\x7b\x53\xdd\x52\x8f\xc6\xed\x40\x6c\x1a\x22\x19\x24\xa0\x03\x20\x5a\xa6\x1e\x46\x54\x7a\xe3\x06\x12\x13\x0f\x82\x02\x23\xe8\x9d\xf1\x0b\x50\x9f\x74\xd7\x02\x50\x33\x5e\xe4\xf9\x65\x5e\xa4\xd3\x27\xe6\x20\xcc\xb2\xc0\xc1\xc1\xcf\xa9\x87\x67\x97\x29\x17\x54\xfe\x5e\x56\x61\x80\xb9\x66\x4c\xe0\xc3\xc8\x45\x67\x8e\x0e\xc4\x6f\xc2\xc6\xaf\xbf\xfc\x95\x7a\x89\x70\xfe\xfa\xcb\x7f\xb3\x86\xd8\xeb\x3b\x50\x32\x15\xed\x8a\x94\xa5\x46\x7b\xde\x23\x17\x19\x7c\x4f\x24\x6b\x60\xc7\x74\x9d\xeb\x44\xeb\xd0\x44\x07\x7d\x56\x3b\xd3\x9a\x0e\xb6\xcb\x8d\xf2\x43\xb5\x47\x5d\x80\x7c\x02\x35\xb5\xf5\xb7\xb0\xb7\x34\x4c\x06\x78\xd4\xc3\xb1\xb1\x15\x6e\xab\x5b\x73\x5e\xa9\xff\x74\x03\xe8\x8c\x16\x66\x33\x77\xb4\x53\x79\x92\x11\xff\xcd\x19\xb0\xfc\x80\x32\xa7\x91\xa4\x4d\x63\x0e\xea\x64\xfb\xbd\x3a\x83\x1a\x51\xfa\xc8\xa3\x21\xb0\x38\xb8\x39\x36\xee\x7c\x80\x81\xd2\xf5\x10\x1d\xc4\xeb\x52\x58\x92\x05\xf4\x3b\xd7\x34\x0e\x90\xf0\xf3\x1d\x1d\x5a\xe2\x3a\x00\x8f\xc7\x46\x9f\xa7\x0b\x81\x98\x26\xda\x39\xbc\x17\x5e\xf3\xbc\x22\xf8\x18\xd6\x83\x5a\x1e\xb1\x28\xe3\x68\x27\xdb\x34\xe3\x24\xa0\xe4\x7e\x36\x9d\x53\x7a\x8b\x14\xfe\xf6\xf3\xcf\x12\x58\x57\xea\x07\x8f\x03\x83\x9e\x64\x85\x09\x9d\x45\xa3\xc3\x70\x30\xda\x87\x3e\x28\x49\x07\xff\x74\xdc\x0b\x95\x38\x3d\x13\xc1\x52\xb5\xee\x35\xd0\xd2\x0d\x55\x3f\x74\xe5\x43\xe5\xd9\x06\xf6\xd9\xae\x83\x5f\xf5\x7a\xdb\x0c\x7e\x2f\x6a\x85\xfe\x36\x39\xb6\xf2\x0c\xb5\xb7\x1e\x37\x2a\xa2\x57\x21\x44\xd5\x4c\xd0\x03\xbe\x86\x28\xd9\x68\xe0\x09\x0f\x3a\x1e\xa7\x56\x1a\xe0\x38\x75\xb6\x37\x74\x74\xa2\x1c\x22\x60\x91\x67\x82\x0b\x48\x3b\xa1\xb0\xc6\x01\xba\x01\xa4\x0e\xc5\xa4\x01\xb1\x41\xc8\x4d\x67\x1d\x9c\xb1\x32\x5d\x59\x56\x7a\xd7\xc3\xd1\x74\x98\x8b\x4a\x68\x90\x29\xb3\x53\x93\xda\x52\xed\x03\xf4\xd8\x70\xa4\xbf\x38\xf8\x97\x93\x73\x3d\x9e\xe6\x46\x04\x07\x0e\x2c\xd8\x32\xc8\x21\xee\x8f\x17\x76\x65\x56\x2a\x60\xfc\x92\x17\x27\xf2\xfd\xda\x54\x7a\x08\x6b\x8e\x00\xa3\x26\xf2\xae\x19\x7a\x39\x33\x6f\x70\xe1\x3d\x10\xd7\xd1\x8c\xc1\x22\x09\x6b\xa2\x5b\xc0\x13\xc4\x0a\xd4\x15\xb2\x93\xc8\xcb\xc6\xf4\x3d\xc9\x34\x8b\x0d\x42\x47\x93\xc0\xfc\xc3\xa1\x0c\xd7\x38\xd0\x0c\xb0\x64\x0e\x26\x01\xd0\xaa\x3a\xbb\x11\xcc\x3a\xd3\xd0\x9e\x05\x1e\x8f\x38\xf3\x09\x6d\x9a\x5c\x28\x52\x8a\x47\x98\xf5\x01\x67\x0f\xab\x38\x22\x9b\x09\x4e\x0e\x68\x0e\x20\x99\x19\xba\x43\x91\x18\xe5\xf4\x46\x56\x07\x14\x4a\x63\x6f\x4d\x73\x4e\xf7\x14\xfc\x11\x1e\x3a\xd4\xd4\x60\xc7\xe1\xf0\x5a\xfd\xf4\xbc\x45\xdd\xd3\xdc\xfc\xf4\x9c\xa8\xf9\x67\xb5\x77\x27\x43\xa6\x88\xde\x70\x13\x4f\x8e\x9a\xd1\xdf\x9a\x53\x32\x6a\x0a\xe8\x4e\x0f\x3b\x83\x86\x11\xa8\x8a\x12\x96\x99\x55\x29\x9d\xc4\xac\x14\x03\x29\x47\x85\x2c\xa1\x54\x0a\x6f\x46\x9c\x48\x1d\x11\xe3\xca\x1d\x51\xd1\xa3\xca\xed\x9d\xbb\x45\xe6\x2a\x77\x38\x36\xa6\x37\x25\xb5\x02\x6c\xd9\xed\x79\x1c\x28\xd8\x64\x30\x08\x69\x60\xd4\x77\x42\xa2\x28\x39\x5c\x38\x4b\xba\xa6\x01\x0b\x91\xb5\x3d\x36\xc3\xb0\xbd\xeb\x40\xdd\x0b\x3a\x22\x80\x35\x9e\x11\x89\xd8\xd0\xba\xcb\xb6\x5b\x44\x6b\x6b\x5b\x8b\x3a\x62\x6e\x49\xa6\xad\x02\x56\xdc\xb5\x9a\x35\x4e\xaf\x0f\xc7\x14\xd6\xc8\x7f\x3d\x43\x48\x14\xa6\x98\x92\x62\x49\xaa\x6f\xb7\x13\x39\x21\x09\x02\x88\x60\x53\xaa\x2d\xac\xf7\xc1\xb6\x03\xaa\x2d\x07\xb3\xb3\xc8\xdf\x11\x98\x74\xe8\x31\x4c\xbf\xfe\xf2\x5f\x3e\x33\x70\x79\x8b\x56\xb4\xcb\x50\x7c\x71\x23\xd6\x76\xbb\x35\xd0\xa3\x12\xdd\x8d\xdd\x7e\x76\x2d\xac\x74\x67\xe8\x0c\x21\x6b\x38\xdd\x06\x78\x94\x43\xdb\xd0\xf4\xf8\x82\x77\x07\x6e\x03\x6d\xe8\x0b\x68\x3e\x03\x31\x69\x41\x31\xe3\x8e\x44\x28\x85\x9e\x65\x49\x0c\x04\x5f\x74\x6b\x92\x41\x43\x63\xa4\x4f\x64\x84\x9c\x8b\xaa\xb1\x30\x96\x17\xa5\x2e\xef\xc3\x01\xda\x0c\x41\xc5\x86\xe9\x60\x31\x41\xd7\x83\xe2\xf4\x78\xde\xab\xaf\x4a\x9a\x0e\x64\x48\xdb\x46\x83\x85\x81\x93\xea\xba\xe6\x55\xc9\x8e\xc3\x92\x5a\x0b\x72\x15\x5f\xcf\xb8\x1f\x07\x4d\xa4\xa8\x8a\xbb\x2e\xe3\x7b\x68\x07\xb4\x35\xc7\x57\x80\x17\x77\x60\x7b\x3b\x41\x24\x85\x83\x6d\x27\x1a\x38\x63\x6b\x14\xae\x16\xa5\x73\x63\x5b\x96\xc7\xb8\x99\x66\x2b\x02\xfc\xc1\x19\x0d\x66\x44\x1f\x14\xe7\x38\x29\xac\xbe\x06\x03\xa5\xcd\xe6\xbe\x89\x1a\xf6\x87\x7f\xfd\xf6\x3f\xd4\xd0\xd8\x83\xed\xd5\x3b\x38\x0d\xc8\xf8\xaf\x5d\x35\xa0\x25\x76\xc9\x0a\x66\x77\x0f\xc6\x00\x59\xd8\x83\xac\x7b\x9c\x63\x64\x3f\xe1\xba\x04\x38\x6f\xf2\xaa\x33\xe8\x6e\x66\xa0\x53\xcb\x57\xdc\x50\xc4\x1d\x2d\x83\x22\xb8\x32\x5a\xee\x21\x2e\x08\x58\xbe\x28\xc1\xbe\x66\x0b\xae\x82\xfd\x62\xea\xa2\x41\x54\x0f\x9d\xde\xd8\xc6\xf6\x67\xda\x31\xee\x00\xb0\xcd\x75\xb7\x3c\xcf\x2d\x86\xc4\xf2\xe9\x74\xeb\xb5\x4c\x8e\x07\x56\x8f\xe7\x20\x13\xfd\x67\x50\x12\xe8\x68\xc3\x44\x61\x7b\x93\x96\x0a\xcf\xd1\x1a\x3b\xca\x04\xec\x96\x83\x30\xac\x4a\x08\xcb\x1b\xa6\x5e\x1f\x36\x33\x0a\x63\xe3\x9b\xd7\x39\x95\xe3\xa1\x42\x9a\x03\xcf\x14\xb3\xd3\x9b\x33\x2a\xb4\x17\x6f\x5e\xbf\x4c\xe9\x7d\x14\xad\x25\x52\xc9\xce\x5b\xf7\x6e\x8d\x73\xae\xd1\xfd\xf2\x25\xa2\xa9\xdb\x7b\x07\xca\x58\xff\x01\xfb\x3c\x9a\x76\x52\x40\x81\xca\x44\x49\x91\xe7\xe7\x9f\xce\x0f\x1e\x29\x1d\xeb\xab\x92\x6c\x84\xb6\xd1\xae\x10\x85\x97\x36\xe1\x56\x0b\x5a\x86\x68\x2b\x0b\xca\x17\xcf\x5e\x24\x2b\xed\xed\xcf\x66\xed\xb6\x84\xa3\xfa\x0d\xe8\x84\x30\xe4\xb4\xed\xe5\x02\xd9\x28\xb6\x6b\xdb\xae\x79\x45\x1a\x57\xdd\x2e\x09\xf7\xb7\xed\x8f\xd8\xe7\x3b\xe8\xb2\xa8\x1d\xc3\x3e\xa0\x0d\x16\x7d\xd7\xd3\x1e\xfd\x6b\x66\x88\xa2\x3f\xb7\xea\x04\xf6\x1b\x4c\x81\x9e\xb8\xbc\x82\x51\x9c\xb4\x87\x6d\x6b\xf6\x3c\x75\x74\x09\x5b\x07\x06\x2c\xa9\x3d\x78\xb8\x37\xfa\xee\x1c\x5f\x00\x85\x87\xfa\xad\x42\x55\x81\x2a\x6c\x3b\x74\xe4\x09\xd5\x56\xef\x5a\xe7\x6d\x51\x25\x19\xdd\x35\xe7\xf5\xd2\x9e\xa6\xd6\x40\x1d\x2b\xe0\xce\x6c\x97\x1d\x9b\x40\x65\x67\xfe\x02\xce\x17\xaa\x25\x1d\x44\x67\x83\xb6\x84\xb8\x86\x15\x58\x2a\x03\x7a\x90\x97\x37\x78\xc1\x36\x33\xad\x1f\x3a\x31\x86\xc9\xb4\xb8\x38\x02\x1a\x2a\xad\xc3\x93\xc2\x91\x73\x44\xb0\xd0\x6b\xa3\xdf\x9d\x82\xe2\x87\xc3\x41\x77\xe7\x54\xe3\x51\xf4\x4e\x4c\xd5\x7a\x6e\x26\xd4\xfd\xc4\x4f\x2a\x5a\xa4\x53\xb3\x95\x3d\x67\x10\x89\x6a\x1f\x21\x13\x87\xd1\x88\x0b\x00\xbd\xde\xbc\x43\xa9\x4d\x50\x60\xdb\xee\x9a\xe3\xcb\xaf\x4a\x67\xf6\x82\x4b\x3c\x2a\x8c\x60\x1c\x41\xda\x77\xeb\xcd\x80\x36\xd7\x8c\x37\x6c\xff\xce\xed\x5e\x53\xeb\xa3\xd8\x64\xef\x04\xc7\xd1\x9d\x44\x31\x59\x58\xf3\xed\xbc\x52\xef\xc0\x81\x68\x52\xa7\x5d\xe4\x5e\x3c\xb5\xb0\xc8\x89\x05\x5e\x66\x29\xea\xd4\xa8\x28\xca\xda\xf4\x5f\x64\xc0\xc7\x33\x05\x27\xd9\x00\x34\x9f\xa3\xef\x35\x65\x09\x95\x75\x30\x17\xe0\x15\x53\xb3\x9d\x5f\x9b\x3b\x5b\x19\x96\xd0\xad\x46\x2b\x17\x0d\x22\x0d\x10\x57\xfd\x03\xf8\x1b\xcf\x8c\xab\x07\xc6\x13\x98\x4c\x78\x4b\x4e\x0a\x0e\xcf\x08\xa3\xd0\xfc\x89\x78\xed\xcc\x41\x1f\xd7\xc7\xce\xde\x81\x24\xac\xef\xac\x39\x15\x02\xd5\xd0\xe5\x4f\xdc\xe3\xdf\xa1\xc3\x13\x58\xa5\xa1\x8e\x6c\x5c\x1d\xcf\xaf\x5c\xfb\x8a\x25\xf5\x60\xc0\x12\x3c\x2b\x6c\x84\xad\x89\x54\xf8\x8f\x25\xae\xcf\x60\xeb\x76\x7a\x8d\x46\x26\x72\x76\x44\x6f\x70\xab\xc1\x8f\x99\xaf\x69\xd2\xa6\xde\x26\x11\xd4\x69\x40\x06\xfb\x29\xe9\x47\x3a\x12\x15\xb1\x05\x85\x49\x21\xcd\xe8\xe4\x01\x0b\x7f\x4a\x7a\x82\x4a\xe7\xe8\xdb\xd8\x21\x7f\xbb\xe4\xcf\xa0\xc7\x51\x51\x68\x8f\x0f\x6a\x52\x66\xd6\xb7\xe0\xe5\xf5\x99\x43\x20\x09\x1c\x41\x92\xc3\x12\x29\x3f\x21\x6c\x86\x7e\x1c\x58\xa2\x60\xce\xc2\x24\x75\xa7\xd1\xee\xae\x68\x8b\x89\x69\x72\x70\x07\xb2\x1a\x90\x4f\xe7\xba\xcc\x1f\x47\x71\xc3\x01\xc8\x4b\xc3\x6c\x0e\xfb\x0d\x64\xe1\xc3\xba\xc9\x2a\x9a\xf6\xce\x76\xae\xe5\x61\xf0\x15\x09\x8f\x70\x54\xd0\xa0\x63\xf6\x9d\xbc\x81\xad\xe0\xed\x75\x1a\x96\x36\x83\xb5\x76\x74\xa4\x80\x1d\x8d\xa1\xf0\xce\x36\xe7\xe4\x98\x06\xf7\xd9\xfb\x72\xc8\x03\xce\xea\xe3\x7a\xc0\xe0\xf9\x9a\x8c\xb3\xd9\x1a\x4f\x3b\x8c\x71\x0f\xd7\x36\xe7\x04\x52\x30\x94\x7e\x68\xed\x87\xdf\x7c\x07\x2e\xf7\x07\xd9\x76\x9e\x7d\x9b\x6e\x26\x1b\x68\x00\xe1\x22\xf0\x90\x00\x1d\x4e\x23\xc1\x6d\xf2\xdd\x36\xe7\x7c\x79\x25\x62\x57\xf0\x9c\x77\x8d\xdb\x80\xfb\x12\xec\x23\x76\x66\xae\x24\x07\xb1\x0f\x65\xb9\x26\x56\x5f\x12\xbc\xa9\x3a\x17\xb7\xe7\x18\x4c\x8e\xf4\x90\x31\x2a\xd1\x63\xb6\x6f\xd0\xe1\xa1\xf0\x39\xdb\x5b\x4c\x16\x9a\x64\xe2\xf5\xc0\x7f\x60\x0c\xec\xf6\x00\x1a\xca\x2f\x10\x23\xde\xa2\x20\xc2\xf2\xcb\x19\xc0\xc5\xbc\xdd\xed\x15\xce\xb0\xcb\xc7\x63\x2c\x09\x5b\x8d\xec\x8c\x29\x50\xb4\x12\x17\x42\xfe\x08\xc5\x32\x95\x6c\x50\x67\xb1\xfe\x62\x3c\x32\x72\x13\x43\x51\xf8\x0b\x76\x83\x3b\xcd\x94\x1c\xdb\xb5\xa3\xa6\x20\x1f\x72\x4f\x96\x49\x4e\xfe\x96\xac\xc0\xb6\x87\x75\x40\xbb\x0b\xfc\xe5\xae\x9f\x87\xa7\x57\xea\x8f\x76\xb7\x7f\xc0\x2c\xc9\x0c\xa0\x3a\xf3\x39\xae\x47\xc0\xe7\xb2\xbc\x96\x78\xc2\x1a\x06\x1a\x4c\x94\xec\x2b\x42\x9d\x9b\xff\x18\xae\x30\x82\x1c\x76\x49\x08\xa7\x51\xc1\xfe\xd7\x36\xc6\xb8\x38\xe5\x86\xf2\xfa\x25\x87\xb0\x61\x07\x13\x0b\x1e\xcf\x15\x7e\xe3\x06\x94\x4e\x07\x0a\x10\x73\xac\xd0\x82\xa9\x73\xc1\x2f\x19\x1b\x1e\x0d\x98\x71\xa7\x38\x79\x8b\x51\x4d\x90\xb2\xae\x85\x71\x1b\xef\xb2\xf0\x78\x1a\x5d\x69\x00\xef\xe8\xad\x90\x7e\x80\xbf\x69\xd2\x19\x95\xb0\xd5\x6a\x76\x45\x5e\x98\x15\x1c\xf4\xf8\xdb\x74\xfe\x25\xc9\x49\xe2\xa8\x70\x2b\x3d\xc0\x56\x5c\x45\x00\x87\x0c\xa4\x92\x2e\x94\x61\x2e\xc8\x2d\x77\x08\x41\xdf\xf1\x80\x9b\x9e\x50\x98\xdc\x93\xf5\x03\x94\x04\x6b\x22\x6e\x89\x93\x05\xd0\xb1\xfd\xd5\x93\x91\x5f\xf4\xf3\x2f\x31\x2b\x1d\x3e\x22\xb3\xe3\xca\x2c\x70\x4b\x1d\x1e\xc0\x2e\xf9\x51\xc2\x2d\x4e\xb2\xc8\x6e\xbe\xb1\x28\xda\x79\xef\x1d\x65\xcb\x36\x8d\x98\x0f\x3c\xd8\xa5\x48\xab\xa8\x7f\xd9\x8d\x3e\x04\x5b\x47\xec\xc4\x60\x40\xa4\x48\x1c\x12\x06\x0b\xe2\x7d\xe9\x75\x06\x78\xf6\x7e\xd8\x00\x4f\x92\x79\x74\xea\xb4\x6d\x67\x71\x85\x65\x1c\x96\xf9\x7a\x92\x3c\x96\x08\xb9\x17\x11\x53\x74\x0a\x62\x02\x1a\xc2\x7c\x58\x87\x3c\x30\xd2\xc2\xc6\xe4\xc4\x2e\x0a\x0f\x17\xca\x06\x92\x6c\xf5\xf4\xa8\x90\x37\xb1\x32\x06\x66\x32\x3e\xcb\xa3\xda\x18\x7a\x8a\x61\x66\x01\x79\x6f\xc5\x32\x3c\x58\x9c\x97\x0f\xa7\x82\xa5\x3f\x06\x20\xb2\x74\x9c\x4c\x26\x71\x66\xe8\x02\x5b\xac\xb3\xf8\x00\xad\x68\x7d\x4b\x55\x01\xf5\x9d\x06\x60\x77\x21\x89\xe4\xcd\x48\x63\x48\xb3\x08\x8b\xb5\x33\x1c\xc0\xa0\x9c\x2f\xe7\xc8\xc9\xbe\xa0\x4e\x07\xd2\xe3\x05\x92\x02\x15\x62\xaa\x1a\x72\xb8\x60\xd9\x70\xfb\x0f\x47\x04\x92\xfb\xc1\x6e\x2e\x5b\x15\xfb\x52\x2c\x68\x7f\xcf\x18\x10\xef\x97\x96\xa9\x48\xcb\xb8\x64\x41\x08\x5d\xce\x0e\x63\x56\x57\xb2\x17\x9d\xe9\xc1\x69\xa2\x01\xf9\x45\x0a\x57\x06\x9f\xe1\x4b\x98\x7d\x5c\x87\xc8\xb3\x1f\x79\x93\xb7\x06\x54\x03\x23\xa5\xd1\x8f\x19\x53\x99\x3b\x8a\x97\x53\x61\x02\x22\x91\xac\x7d\x14\x36\xc9\x2e\x83\xc6\xa4\x68\x4b\x44\x98\x72\x02\xd3\x19\x61\x0f\x6c\xed\x6e\x60\xe1\x5b\xa8\xec\xb1\x33\xc1\x46\x40\xf9\x71\x84\x34\x33\x23\x97\x40\x1d\xf3\x74\x1a\x5c\xcd\xc3\x51\x22\xff\xe2\x89\x45\xd4\xa9\x3b\xc6\x17\x71\xfd\x29\x24\x1c\x7d\x2f\x10\xb7\x5f\x7f\xf9\x2b\xcf\x7d\x83\x85\x36\x22\x25\x49\xac\x71\x0b\x9e\x06\x8c\x8b\xb2\x7b\xc6\x22\x95\x38\xe9\x8d\xda\x0c\x3d\x1a\x95\x7b\x4a\x7a\xc2\x3b\x80\x86\xad\x48\x17\xe0\xd1\x20\x40\x12\x0d\x64\x3c\x4e\xea\x30\x62\x0c\x6b\xea\x31\xc7\xc5\x63\xaf\x8c\x73\xd0\x28\xef\x78\x58\x65\x7e\x9e\x57\x27\xd3\x34\x0b\xc6\xb0\x41\x1f\xae\x10\x25\xc0\xc7\xf7\x16\xdd\x51\x04\xf3\xa2\x17\x92\x62\x1a\x6b\xd1\x9a\x97\x14\x25\x95\xb9\x9c\x8f\xe2\xbd\xe6\x79\xd3\xcf\x18\xc3\xe1\x6a\x31\x4b\x38\xcb\xd9\x37\x8d\x85\x7e\xe0\xae\x71\x98\x8f\x48\xfc\x2a\x50\x38\x8f\xf3\x05\x75\x8f\x2b\xbd\x2e\xbb\x09\xd8\xf4\x96\x3c\x84\x4c\x2d\xc6\x28\x3c\xed\x23\x20\x2d\xe8\xc2\x89\x06\x8d\x1c\x7e\x86\x42\xa2\x8f\xc0\x81\xae\xb0\x92\x43\x7d\x56\xd0\xf8\x68\x11\x78\x8e\xa6\xe2\x9f\x98\xa1\x21\x0f\x09\xa9\x12\xa3\x61\x1e\x82\x7c\x2b\xf2\x92\xaa\xff\x49\xf8\x66\xe2\x6a\x85\x9d\x72\xce\x15\xf3\xc4\x3f\x19\xcd\xc1\x4b\x85\x9b\x85\xca\xc8\x8f\x4e\x0f\xcd\x74\x8d\x1c\xc1\x67\x1e\x00\xfb\xf1\x53\xe1\x33\x5a\x90\x97\x00\x9a\x13\x74\xfa\x44\x00\x5d\xa6\x27\x15\x2e\x5d\x61\xcc\x0a\xa4\x70\x2e\x5f\x9f\x10\x3f\x89\x97\x4e\xec\xf0\x7b\x82\xf9\xe9\xc5\xff\x22\x79\x17\x85\xef\x59\x6b\xfa\x93\xeb\x28\xe2\x43\x31\xa4\xb5\x9d\xa7\xf7\xb2\xb2\x1d\xea\xf5\x6d\x1b\x49\x4f\xf4\xed\x48\xbb\x8c\x8a\x59\x61\x3c\x41\x6e\x62\x88\x0a\xeb\x3c\x2a\x03\xc7\xab\x44\xa6\xac\x4f\xec\xfb\xcb\xf9\xa0\xc9\x98\xca\x87\xb0\xcf\x7e\x5a\x8e\xa6\xc8\x4d\x08\x5e\x11\x87\x0a\xcd\x07\xd0\xb1\xbd\x98\x90\x78\x0a\x85\xe3\x87\xc2\xd7\xaf\xd2\xba\x4d\x99\xa0\xbc\x98\x0c\x92\x1b\x2e\x94\x94\x44\x94\xbe\x87\x93\xe0\xb1\x30\x11\x77\x64\x24\x3d\x1d\x24\x89\x01\xfe\xef\x81\x04\x07\xef\x5a\x72\x84\xf3\xf3\x1a\x1a\xdf\x4a\xdb\x18\x03\x1d\xcb\x2b\x05\xc0\xa9\x9b\x58\x03\x9d\x70\x2c\xf7\x21\xf7\x98\x9c\xe6\x12\xf1\xc3\x8d\x10\x84\xeb\x3e\xe5\xa4\x61\x9b\x04\x81\x46\x7e\xe3\xb2\x05\xf3\x32\x47\x57\x92\x79\x54\xe9\x14\x90\x1e\x7a\xdb\xd8\x9f\xa5\xb8\xea\x1e\x90\xa6\x20\x2e\x58\xe8\xcf\xdc\x31\xf3\x9d\xc0\x49\x33\x5d\x2e\x70\xfc\x68\x96\xa4\x5e\xc6\x4f\x5e\x48\xfc\xa8\xb8\x11\x47\x73\x27\xd8\xed\x72\x3f\x24\x8b\x61\x2e\x14\x25\x93\x15\x99\x91\xc6\x76\xe5\xfd\x29\x0b\x3e\xd4\x03\xe9\x59\xba\x30\x52\xcf\xef\x89\xd4\xd9\xfd\x90\xab\x04\xc9\x0b\x39\x54\xd6\x3c\x95\xb4\xda\x60\x45\x5e\x46\x1a\x3f\x7a\x00\x69\xf2\x42\x1a\xa9\x7a\x02\x45\x3b\xd3\xa3\xcf\x93\x91\x24\xcf\x1e\x40\x13\xb8\x1b\xf2\x12\xba\x1b\x0f\x26\x2d\x86\x75\x78\x12\x4c\x21\x6e\x0c\xf9\x83\x0a\x2b\xf0\x95\xe5\x59\x59\xaa\x38\x50\x81\x75\xac\xee\xb4\x52\xef\xe8\x50\xd4\xf0\x1c\xcb\x26\x3c\x2a\xbb\x1a\x99\x7a\x83\x0c\x24\x84\x80\x62\xc0\x20\x58\x60\x00\x8e\xdb\xb8\x03\x0b\xe9\x90\x11\x1f\x2c\x37\x80\x2d\x9b\xe1\x23\xcf\x1e\x80\x8f\xbc\xe1\x39\x77\x54\xcf\x0b\x08\x1f\xbf\xfd\x52\x45\xb1\x46\xb6\x3e\x96\xb6\x08\x10\x51\x80\x65\xaa\x38\x9e\x22\x72\x4f\x57\x17\x09\x69\x41\x73\xfc\xad\x15\x46\x42\xd2\x5c\x77\xfc\xad\x55\x46\x42\xdc\xff\x6b\x8f\xff\x7b\xda\x23\x59\xbe\x8f\xa3\x48\x8a\x7a\x84\x03\x48\xe4\x04\x94\x62\x83\xd1\xb4\xdd\x50\x19\x9b\x57\xc6\x52\x90\xe2\xf7\xbf\xc3\xd0\xc0\xe7\xbf\xbd\xc1\x82\x2b\x80\x0f\xbd\x0f\xac\xe9\xa5\x62\x81\x1e\x93\xef\x18\x5a\xaf\xf6\xe0\x1f\xd1\x85\x9c\x24\xd0\x91\xd4\x23\x0a\xb5\x58\xd6\x67\x9b\x0c\xd8\x95\x02\xbb\xec\xe0\x7c\x1f\x59\xf1\x2a\xdc\x17\xf9\xfd\xef\x42\xbd\x6f\x34\xf2\x62\x3c\xb5\xda\xeb\x76\x67\xc6\x88\xea\x62\x3c\x09\xd6\xe0\x52\x71\x78\xe8\x71\x21\xe7\x9c\xd9\xf3\x6f\xbf\x7c\x73\x33\x2f\xd1\xbc\x49\xb2\x3a\x97\x72\xf2\xc4\xad\xdc\xb6\x98\x06\x90\x8c\xd4\xa5\x63\xa4\x98\xfa\xd5\xa6\x0e\xbb\x3a\xaf\xdc\xf0\x59\x88\x39\xbe\x18\xe2\x37\x89\x90\x8d\x74\x4b\x6d\x0f\x4b\x41\x19\xac\x3b\xdb\x61\x45\x54\x86\x95\x3c\xcb\x13\xd0\x7f\x19\x74\xdb\xdb\xfe\x5c\x04\x02\xe6\x0a\x2f\x49\xcd\x46\x0a\x48\xee\xa0\xac\xd4\x8f\x68\x40\x8f\x85\x48\x60\x8a\x63\x38\xb9\xe6\x32\x9f\xb8\x48\x61\x40\xf4\x22\x7a\xd5\x18\xba\xa6\x76\xb2\xb2\x15\x62\x37\x2e\xf1\xa1\x18\x7c\x78\x23\xc6\xba\xbc\xdd\xb5\x20\x96\xa0\x3d\x70\x8d\xa4\x6a\x84\x6e\x3f\x48\x61\x10\xe7\x7f\x3e\x57\xa1\x72\x9c\xe2\x89\x2f\x43\xa0\x5e\xa7\x35\x22\x81\x31\x20\xe4\xb6\x18\xb0\xa3\x11\x97\x65\x4e\x66\x5c\x28\x02\x96\x56\x9e\xa3\x08\xf1\x44\xb8\xc6\x8b\x56\xa1\x12\x71\xc0\xa8\x34\x0f\xf0\x4a\x86\xa3\xc2\xaf\x1b\xb5\x74\x55\x68\x63\xe4\xdf\xe5\x3d\x90\x54\xa3\x00\x95\x79\xe6\xc4\x8d\x7f\x97\x43\x98\x44\xc4\x1a\xdd\xa5\x62\x61\x1f\xb7\xa3\x30\x84\xba\xbe\x47\x83\xc3\xa5\xf2\x51\xf1\x24\x22\x88\xde\x60\x5a\xf3\x36\xbd\x96\xa7\x9b\x13\x4a\x38\x42\xb1\x2c\x5b\xef\x93\x10\xb8\xd4\xf5\x48\x71\x7e\x8d\x27\x51\x22\xca\xa8\x43\x59\x9a\x17\xf4\x71\xdf\xd9\xca\x63\xed\x82\x77\x1d\x45\xbe\xea\x10\x74\x08\x57\x0a\xf0\x19\xba\xaa\x17\xaf\xf0\xf1\x00\xe9\x75\x00\x7a\x8f\xae\xc2\x5e\xbf\x42\x3c\xc9\x57\xa1\x50\xe0\xde\xc0\x15\x45\x45\x44\xe9\x2e\x4c\xba\x84\xcb\xe2\x7c\xdd\xb6\xb0\x29\xda\xcc\xed\xa5\xeb\xba\xc8\xf7\x65\xae\xe1\x0c\xa6\x58\x54\xeb\x08\x80\x69\xcc\xa5\x75\xef\xf9\xe9\xd2\x95\x04\x7c\x3f\x02\x10\xef\x99\xb8\x23\xd1\xf0\xf5\xeb\x7f\x43\x9b\x60\xf5\x3d\xfd\x5c\x8d\x83\x61\x02\x80\xa2\x05\x72\xff\x57\x1e\x73\x55\xa4\x96\xda\x13\xf6\xa8\x29\x5b\x0a\xaa\xae\x10\x0e\x57\xea\x68\xdb\x76\xb2\xcf\x19\xe6\x77\xbd\xee\x07\xbf\x12\x66\x57\x4c\xe0\x0a\x89\x5d\xf1\x3b\x4b\xfc\xfc\xf4\x9c\xdb\x7f\x7a\x9e\xb1\x56\x9c\x7d\x9e\x91\x9f\x94\xb5\xe4\x43\x67\x62\x12\x14\x05\xad\xa2\xa5\x24\x71\x08\x2b\xc6\x0b\x4e\x53\xa5\x62\x3e\xec\xe1\x37\x18\x07\x72\x65\x48\x06\xbc\x49\xd2\x08\x89\x76\xe1\x2a\x05\x18\x0f\xef\x03\xaf\xd4\x1f\xc3\x65\x43\xb6\xd6\x3a\xc3\x11\x21\xa9\x5f\x00\x69\xc4\x1c\x62\xbc\x83\x13\x46\xc6\xa4\x84\xc8\xe3\x42\x4d\x45\x9c\x8f\xee\x90\x59\x5f\xaa\x14\x0a\x32\x17\x6e\x07\x7d\x11\xcd\xee\xba\x60\x77\x97\x4a\xf0\x12\xef\x5c\xc6\xf0\x61\x84\xe2\xde\x0e\x1e\xda\x64\x82\xf0\xf0\x9e\x33\x84\xee\x0b\x53\x70\xbe\x75\x32\x45\x78\x78\xcf\x29\x62\xf7\xcd\x39\xb8\x5a\x0b\xb3\xb1\xd7\x53\x17\xbc\xa8\xfb\xce\x25\xbd\x2f\x2a\x43\xb0\x26\xd7\x74\x1b\x92\xf4\xc7\xfa\x14\x2a\xf8\x60\xbc\x5c\x31\x0c\x87\x8b\xb3\xc2\x38\xdf\xc1\x30\xdf\xe0\x28\xa9\xc0\x90\x8a\x48\x93\x71\x1c\xd8\x97\x62\x18\xb9\xce\x7c\x7a\x19\xb2\xaf\x96\x43\x88\xae\x35\xd1\x2c\x80\x03\x65\x83\x16\x18\x56\x52\x04\xcb\x3d\x28\x14\x5d\xdd\xb6\xee\x04\x0a\x7f\xc7\x3a\x77\x52\x45\x21\xc3\x6b\x15\xaa\x9c\x43\x22\x9b\xcc\x90\x7f\x5a\x15\x8a\x4e\x9e\xc5\x6b\xde\x54\x03\x3d\xdf\xed\xd2\x50\x00\xa3\x94\x65\x48\xee\x7a\x4f\xaf\x79\x8f\x39\x86\xa4\x0c\xe4\x22\x8c\x7f\xb7\xd0\x5d\x91\x2a\x04\xf1\x24\x3a\x3f\x47\x34\x3e\xbd\x5e\xac\x92\x33\x3f\x2d\x5d\x19\x0f\x60\xed\x89\x57\xba\xdb\x2a\xee\x5f\x98\x45\xc5\x2b\x88\x64\xf5\x27\x50\x97\x03\x41\x81\x9b\x38\x19\x7d\xd2\x00\xde\x39\xb8\x7c\x5b\xca\xb3\x05\x26\xe6\xe1\x0b\xc9\x2a\xb0\xbf\xc7\x3a\x7f\xdc\xb5\xe2\x90\xa0\xd3\x12\x8b\xc9\x79\x88\x58\xbd\x25\xc5\x02\x7f\x80\x1e\xe6\x83\xc6\x5b\xc9\x62\x20\x04\xbf\xf7\x44\x95\x76\x9d\xc1\x4f\xe1\xa0\xa1\x30\x25\x41\x85\xaf\x7a\xd0\x70\xff\x00\x36\x8b\xb4\xb9\x8e\xab\x37\xaa\xce\x1c\xe4\x4b\x1d\xc9\xd5\xd9\x70\x9b\x06\x5f\xb6\xa1\x50\x02\xeb\x66\x16\x95\x72\xbd\xd7\x93\x9b\x61\xfc\x68\x01\xaa\x10\x7c\x92\xcf\x41\xd0\x31\x49\x5d\xd6\x56\xae\x69\x33\xd7\x5c\x02\x24\xbd\x6f\xe2\x81\x49\x46\x25\xa1\x69\xa8\xf2\x5c\x6a\x12\xdc\x64\x0c\x2e\x54\xbf\xb5\x47\x25\xa5\xa3\x18\x2b\x6b\xd0\x9e\x02\x43\xf2\x3c\xb9\xf4\x95\x70\xe3\xc1\x97\x59\xc3\xbb\x6b\x40\x75\x72\xf7\x07\x9b\xbe\x6c\xeb\xef\xb1\xe1\xa2\x9e\xcc\x38\x94\x02\x1e\xef\xe8\x8c\xe2\x4e\x81\x79\xdd\x22\xbc\xa2\x1a\xd8\xa0\x74\x69\xc8\x10\xf8\x24\x53\x51\xca\x57\x2e\x0a\x30\xb1\xb8\x86\x75\xaa\x86\x9e\xb7\x23\xd2\x3b\xb3\xa4\xf8\xd9\x45\xf2\x19\x50\x8b\xd5\xf3\x61\x8c\x10\x75\x88\x65\x5b\xad\x94\xc1\xbc\x22\x4c\x0d\x8a\x0f\x7f\x81\x25\x0f\x2b\x49\xe5\x84\xd7\xfc\x4d\xa9\x36\x0c\x27\x2b\x0e\xfb\xf5\xc8\x9f\x96\x31\x1f\x60\x20\xdb\xbe\x78\x59\xde\xa5\x91\x99\xb5\xdb\xfc\x19\xf3\x81\x93\xe3\x99\xf6\x00\xd9\x3e\xf1\xb2\x72\x28\x70\x72\xf1\xf3\x26\x7c\x01\xa2\xe6\xaa\x2a\xde\x35\x0d\x66\xd1\xd3\xeb\xfd\xae\x7d\x45\x57\x42\xc4\xc3\x21\xff\xee\x22\xec\x3c\x12\x92\x73\x70\x77\x93\x62\x26\x7a\x72\x11\x6a\xd6\x7d\xa3\x5e\xc0\x37\xe8\x1c\x98\x90\x13\xab\xa0\xb2\xa2\xbc\xb0\x25\xc2\x35\x8d\x95\xfa\x3a\x1f\x69\xa2\x27\xa7\xda\x29\xde\x2a\xa1\x35\x12\x17\x94\x7c\xe0\x00\xe2\xc6\x9c\x1d\x5f\x0d\xb7\x5d\x40\xcb\x44\xa5\x83\xaf\x5c\x01\xe7\xd8\xac\xd1\x81\x39\xaf\x37\xba\xaf\xe4\x2b\x31\x0f\x35\x3b\xe4\x55\xf6\x84\xb0\x50\x0e\xd3\xfa\x9e\xf6\x75\xd1\x3d\x9e\x7d\xcc\xe5\xa9\xa7\x7c\xf1\x40\x27\xb6\x28\x6e\x37\xa2\x2a\x69\x67\x54\xc2\x8d\xdb\xdd\x13\x1a\x24\xcc\x1d\x73\xfa\xe0\xf7\x45\x48\x68\xfc\x2c\xd6\xca\xd8\x5c\x9f\x72\xbc\xd3\x48\xed\x93\x20\xeb\xd0\xf6\xd9\xbc\xe1\x93\x02\x45\x3f\x21\xc8\x28\xd1\xc2\xe3\x96\xa3\x14\x1f\xf8\xba\xf1\xfc\x7a\x0f\x34\xbd\x83\x96\xd7\x14\x68\x48\xe7\x85\x06\x7b\x00\xb1\x48\x85\x92\x67\x98\x56\x5b\x69\x4a\x1c\xf7\xba\x25\x07\x94\xca\x38\x84\x2e\x5e\xb1\x9b\xf0\x11\x12\x0a\x6f\x86\x9a\x44\xf0\x87\xca\x12\xb3\x40\xa8\x2f\x52\x19\xbf\x49\x91\x50\x49\x49\x72\xb9\x8d\x75\x1d\x9d\x7c\x6d\x24\x33\xbe\x96\xb8\xfe\xa3\x36\x4b\x48\x24\x8c\xeb\x74\x23\x61\x65\xdd\xcd\xed\x13\xa9\x09\xa0\x38\x43\x5d\x5b\x6c\xa1\x9b\x40\x53\x75\x11\x05\x3b\x0d\xe2\xfb\x33\x98\x14\xde\x0d\x5d\xb5\xf4\x15\xc2\x8f\xb4\xff\xc4\x42\xaa\xf9\x03\x35\x74\x1d\x58\xee\x5d\x22\x59\x73\x8e\xaf\xef\x82\x69\xf1\x4e\x46\xdc\x66\xb6\xd0\x53\xaa\xe4\x38\xd1\xf5\x55\x60\xca\xb8\x3c\x7c\xbb\x27\x19\xf5\xc5\x59\x2f\xad\x84\x14\xd1\x17\xbf\xb1\x21\x6d\xe1\x2b\x1b\xd7\xf5\x4e\xcc\x3c\x4a\x5d\xb0\x0c\xba\xca\x6d\x34\x7e\x0a\x1b\xb4\x35\xa7\xec\x45\xac\xc2\x92\xcf\xdc\x70\x44\x8b\xbf\x06\x10\x15\xca\x58\x3d\x9f\xc6\x1d\x75\x88\x28\x49\x00\x23\x94\x91\x44\x6c\xc0\x6a\xe4\x8b\xfd\x28\x0d\x37\x6a\xc2\x15\x91\x34\xda\xbc\x18\x36\x11\x2a\x02\xd1\xde\x70\xa5\x11\x13\xcc\x68\x2a\xba\x9b\x9f\xa0\x7c\xd9\x89\x20\xd9\x22\x56\xd7\x63\x32\xf9\xa1\x3b\x98\x0d\xaf\x18\xd5\xb0\x6d\xb8\xc5\x8e\xe3\x7e\xda\x8d\xc6\x27\x5b\xa8\xef\x91\x23\x8f\x7c\x05\x4a\x65\xcb\x87\x34\xae\xd2\x53\xc0\x63\xcc\xaf\x17\xb4\x2b\x37\xb0\x7e\x9d\xc7\xc4\xe7\xa1\x9b\x87\xd3\x00\xd6\x13\x7e\x07\x02\x83\x55\x8f\xd2\xaa\x23\x0d\x38\x4c\x9e\x85\x83\xb1\x5f\x6d\x0d\x98\x28\x18\x70\xdb\x99\x34\xf7\xb8\x58\xf9\xfe\x51\xd7\x0b\x49\xa2\xaf\x9c\x45\x22\x27\x99\xd7\x87\x53\x58\x46\x4f\xee\x4f\x3c\x4a\xa8\xc9\xa7\x00\xe2\xc8\x7d\x10\x10\x37\xe7\x18\xbd\x90\xaf\x60\x90\x0f\xfb\x91\x79\xf9\x44\x68\xe7\x1c\x7d\x3c\xc4\xf1\xcb\x68\xf0\x0a\x85\x21\x3a\x63\x80\xb2\x7e\xed\x0d\x26\x83\xc1\x76\xab\x6e\x4d\xbf\x96\x90\xf1\x44\x8b\x73\xe3\x37\xa1\xed\x7a\xb0\x25\xb1\x68\x2b\xfc\xc2\xa6\x7c\x61\x03\x27\x55\x38\x2b\xf9\x21\x8e\x2a\x31\x5b\xd1\xc0\x7e\xb0\x60\x6f\x87\x4c\xdf\xdc\x8b\x2a\xc3\x3f\x65\xa3\x58\xe0\x18\xab\x13\x2f\xd1\x9d\x1a\xe1\x34\x96\xe4\x9f\xc6\x3b\xe0\x33\x92\x16\x3c\x83\x29\x4d\xf7\xf5\x95\x93\xeb\x77\x17\xe6\x2c\x11\x7a\x71\xd1\xfb\xbe\x49\xe2\xe7\x31\xa6\xee\x4b\x91\xf4\xaf\x93\x38\xf3\x03\x42\xea\xf2\xa1\x35\x36\x9f\xd2\xb0\x2d\xcc\x2d\xc1\xa6\x32\x54\x47\x3d\xbb\xa5\xc3\x8f\xae\x88\x18\x58\xec\xf1\x83\xa1\x31\x43\x06\xc7\x2f\x39\xc5\x23\x79\x57\xc9\x2a\x02\xf7\x3f\x01\x00\x00\xff\xff\xac\x03\xa2\x76\x8a\x5b\x00\x00")

func groups_yml_bytes() ([]byte, error) {
	return bindata_read(
		_groups_yml,
		"groups.yml",
	)
}

func groups_yml() (*asset, error) {
	bytes, err := groups_yml_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "groups.yml", size: 23434, mode: os.FileMode(420), modTime: time.Unix(1423625052, 0)}
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
	"groups.yml": groups_yml,
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
	Func func() (*asset, error)
	Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"groups.yml": &_bintree_t{groups_yml, map[string]*_bintree_t{
	}},
}}

// Restore an asset under the given directory
func RestoreAsset(dir, name string) error {
        data, err := Asset(name)
        if err != nil {
                return err
        }
        info, err := AssetInfo(name)
        if err != nil {
                return err
        }
        err = os.MkdirAll(_filePath(dir, path.Dir(name)), os.FileMode(0755))
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

// Restore assets under the given directory recursively
func RestoreAssets(dir, name string) error {
        children, err := AssetDir(name)
        if err != nil { // File
                return RestoreAsset(dir, name)
        } else { // Dir
                for _, child := range children {
                        err = RestoreAssets(dir, path.Join(name, child))
                        if err != nil {
                                return err
                        }
                }
        }
        return nil
}

func _filePath(dir, name string) string {
        cannonicalName := strings.Replace(name, "\\", "/", -1)
        return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

