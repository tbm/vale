// Code generated by go-bindata.
// sources:
// rule/Annotations.yml
// rule/ComplexWords.yml
// rule/Editorializing.yml
// rule/GenderBias.yml
// rule/Hedging.yml
// rule/Litotes.yml
// rule/PassiveVoice.yml
// rule/Redundancy.yml
// rule/Repetition.yml
// rule/Uncomparables.yml
// rule/Wordiness.yml
// DO NOT EDIT!

package rule

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

var _ruleAnnotationsYml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x1c\xc6\xb1\x0a\xc2\x30\x10\x06\xe0\xfd\x9e\xe2\x28\x48\x27\x5f\x20\xb3\x15\x1c\xb4\x4b\x87\xac\xa1\xfe\x0d\xc1\x78\x01\xff\x53\xf2\xf8\x42\xa6\xef\x7b\x83\x4c\x19\x41\xa7\xf9\xc4\x59\x2b\x0e\xd7\x62\xea\xe8\x3e\x09\xba\xc3\x9e\x0c\x8a\x5e\xe8\xb0\x1d\x52\xb2\xb5\x0f\xf6\x44\x04\x3d\x52\x25\xa4\xe2\x87\x1a\x94\xdf\x9c\x41\x2f\xcd\xc4\xdb\x0b\xc6\x20\xaa\x67\x8d\x31\x0e\xaf\xb7\x78\x5f\xc6\xb6\xf5\xb2\x8e\x3c\xd6\x6d\x91\x7f\x00\x00\x00\xff\xff\x31\x1d\x63\x53\x80\x00\x00\x00")

func ruleAnnotationsYmlBytes() ([]byte, error) {
	return bindataRead(
		_ruleAnnotationsYml,
		"rule/Annotations.yml",
	)
}

func ruleAnnotationsYml() (*asset, error) {
	bytes, err := ruleAnnotationsYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "rule/Annotations.yml", size: 128, mode: os.FileMode(420), modTime: time.Unix(1494777592, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _ruleComplexwordsYml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\x56\xc1\x92\x23\x2b\x0e\xbc\xf7\x57\x54\xbc\x88\x17\xde\x3d\xcc\xf4\xdd\x97\x39\xec\x97\xc8\xa0\xaa\xd2\x34\x20\x06\x09\x97\x3d\x5f\xbf\x01\x88\x72\x77\xdf\xec\x94\x00\x41\xa6\x52\x85\x0f\xc5\xe4\xe5\xba\x48\xbd\x89\x92\x56\x25\x4e\x6f\x11\x45\x60\xc3\xeb\xf2\x3f\x4e\x42\x1e\xcb\x52\x85\xd2\xb6\xfc\x2b\x0b\x25\x51\x04\xbf\xf0\xba\x5c\xfe\x95\xcb\x5b\xa0\xf4\x71\x5d\x76\xd5\x7c\x7d\x7f\x3f\x8e\xe3\x67\x0e\x40\x29\x40\xda\x2a\x6c\xf8\x73\xe3\xfb\xfb\xce\x87\xf2\xfb\xc1\xc5\x4b\xdd\x36\x94\x76\x84\xbc\x0b\xc5\x1c\xb0\xa3\x3f\xdd\x1a\xdf\x68\x4b\x5c\xd0\x81\xe0\x75\xd1\x52\xf1\x2d\xe0\x1d\xc3\x75\x39\xa0\x24\x4a\xdb\x9b\x1c\x90\xaf\x6f\xcb\x72\x81\x9c\x0b\x3f\x28\x82\xe2\x7f\x7e\x5d\xc3\xf3\xbf\xbf\x2e\xd7\xe5\x9f\x0b\xdc\xb8\xea\xe5\x9f\xb7\x65\x81\x9b\x60\xd2\x86\x25\x4e\x78\x59\xb8\x2c\x97\xc4\xba\xec\x58\xd0\x12\x6a\xf2\x90\x1c\xb6\x9c\x1c\x30\xe9\x73\xe0\xce\x61\xc0\x02\xda\x03\x92\x11\xfd\x52\xf3\x19\x4a\x5a\x67\x48\x0b\x8a\xcc\x00\xc7\x0c\xe9\xd9\xf0\x8d\x97\x83\x74\xff\x14\x08\x24\x7b\x8b\x38\x28\xe5\xb9\xb4\x0a\x7b\x39\x9e\xcf\x9c\xe2\xd1\xf7\xb5\x74\xc7\xf4\x09\xa5\xb4\x85\xbe\xa9\x9c\xb9\xa5\xf6\xd3\xc1\xfb\x89\xd4\x59\x6c\xa1\x6d\xb7\xbd\xf1\x01\xce\x1e\xc2\xfd\xa9\x84\x32\xee\x09\x5b\x41\x7c\xc1\xa5\x83\x1b\xda\xa2\x5b\xb5\x27\xf0\x9e\x1a\x3f\x10\x5a\x38\x72\xc1\xb9\xa9\x16\x98\x19\xed\xf2\x2d\xec\x49\x5c\x9d\xef\x30\x60\xc4\x1e\x79\x72\x35\xf4\x37\xb4\x87\x5b\x94\x3b\x1f\xf8\x68\x3f\x67\xa8\x8a\x46\x63\xca\xed\x90\x36\x2b\xcf\x47\x12\xa1\x5b\x18\x65\x87\xc0\x07\xda\x85\xfd\x1d\x92\xc2\x86\x5c\xfb\x29\x3b\x86\xbc\xd6\x30\x63\x24\x7d\x85\x62\x30\x68\xdb\x0a\x6e\xf6\x40\xca\x0a\x06\x53\x71\x05\x56\x1d\xe4\x43\xb2\x53\x43\xc0\x3b\x59\x32\x82\x9c\x28\x3b\x03\x41\x84\xb6\x64\xfc\xd1\x9d\xfc\x4c\x51\x2c\x09\x94\xee\x38\xf8\xe2\xf2\x1d\x97\x71\x43\x26\x87\x32\xd6\x73\xee\x4d\x30\x12\x23\x06\xe2\x49\x24\xc5\x5c\xf8\x8e\x33\x92\xfc\xf7\xd7\x49\x4a\x8e\xf2\x2c\xf4\x91\x71\x92\x9d\x33\x94\xf9\x9a\x01\xa1\x8c\x93\x7a\x3b\x8e\x04\x71\x58\x14\x28\x4d\xea\xf8\x8e\x96\xb4\x52\xf2\xcb\xd9\x3e\x22\x24\x3a\xdb\xa3\xbd\xf1\x80\x75\xae\x8d\x88\x3a\x21\x8c\xb9\x9f\xa8\xc5\xf4\x53\x75\xe7\x42\x7f\x5f\xdc\x75\xf8\x86\x01\x74\xe8\xbc\xfd\x30\x4c\x94\x8f\x29\xfd\x0e\x39\x04\xb1\x06\xe3\x6c\xc2\x4b\x83\x7b\xc7\x21\xc0\xed\x7c\xa6\x83\xcb\xc7\xa2\xbc\xa1\xee\x58\x2c\x21\x46\xb4\xa2\x6f\xb8\xd9\xa5\x5b\x0f\x62\x12\x5b\x95\xe1\x79\xa2\x9c\xec\xb1\x32\x14\x3d\xd1\x62\x22\x5a\xb9\xc4\x51\x00\x25\x17\xaa\x71\xed\x38\x39\x1c\x17\x26\x8f\xf0\xc2\xba\x45\x7d\x75\x21\xc7\x69\xc5\x72\x5e\xaf\xef\x05\x07\x94\x79\x9d\x24\xf8\xa7\x62\xd2\x2f\x4d\xde\x60\x0e\xe4\xad\xde\x88\x65\x3b\x4f\x1e\xf6\x3c\xc4\xd4\xca\x93\x19\x68\xbc\x8c\x7e\x80\x13\xbb\x63\xc2\x2f\x5c\x79\x8c\x6d\x8b\xd3\xdc\x76\x3e\x4c\x20\xa7\xdc\x3c\xb6\xa7\xe8\x1c\x21\x9c\x58\x13\xbd\x2d\x72\x3b\xb3\xbc\xf0\x61\x21\x07\x24\xf3\x90\x83\x64\xb7\xa0\x62\x89\x34\x2a\xf0\xe8\x5a\xab\x9c\x3a\x9b\x19\x85\x5a\xef\x0f\x9f\xb9\x81\x1f\x09\x3b\x94\x38\x3b\xba\x49\x34\xb0\xe9\x61\x87\xe9\x44\x67\x7b\x77\x09\x27\xa5\x54\x5f\x92\x31\x5c\x30\xd2\x2c\x5a\x9a\x82\xfa\xca\x53\x66\x18\xe8\x15\x9f\x02\xc3\x50\xdd\xf9\xf0\xf8\x78\xf5\x0d\xc6\x1c\xb8\x93\x54\xed\xf2\x98\x7a\x61\x4d\xce\xd4\xe7\xe3\x67\xa1\xf8\x99\xc2\x35\xe9\x10\xc0\x49\x02\x26\x8f\x70\xe7\xf2\xb9\x63\x30\xd5\x78\xce\x9c\xbe\x68\xc0\x7f\x2a\x29\x98\x09\xae\x40\xe5\x44\xef\x10\x4c\xb9\xf8\xa7\x9a\xa7\xe1\xc3\x85\x2a\x2f\x0b\x4a\xc1\x36\x7f\x64\xf4\x34\xf6\xde\x6b\xb1\x23\x57\x70\x14\x48\xbf\x9b\xdd\x8a\x11\xc2\xd0\xd7\xc1\xd1\x86\xd1\x4a\x09\x82\xf5\x73\x1f\x68\xa8\x2f\x32\x27\xe1\x6b\xf9\xac\x65\x5e\xd5\x16\x93\xc7\x66\x57\x83\x64\x81\x38\xce\xa1\xe4\xb8\x14\x74\xfd\x0e\x47\xe1\xb4\x19\xec\xc9\x41\xf3\xc5\x9e\xdd\xbc\x76\xc0\xa4\x74\xce\xdd\xd6\xab\x63\x60\x9d\x2d\x4e\x8a\x91\xfe\x9a\xb9\x90\xa8\x31\xf0\x1b\x39\x43\xf1\x56\x7a\x21\xf9\xe8\x70\x20\xb0\x16\xef\x16\xd2\xc7\xf5\xe8\x03\x28\x9a\xb0\xbc\x06\x78\x04\x4a\xd3\xf2\x3e\x10\xcd\x8f\xa4\xe6\xcc\x66\x18\x11\x75\x67\xcf\x81\xb7\xe7\xa0\xb9\xfd\x1d\x11\xf6\xb4\x3e\xbf\x59\x77\xe4\x44\x3a\xc8\x77\x3b\xba\x0f\xeb\x1a\x50\x67\x07\xd6\xa0\x94\x07\xe3\x11\xd2\xe0\x2a\xa1\x43\x91\x93\x2c\x07\x53\x84\x89\xd5\x8e\x38\x3b\xa2\x2b\xc9\x66\xe3\xb9\x01\xdf\x7e\xa3\x6b\x53\xa8\xdb\x13\x99\xab\x6d\x6c\xca\xe1\x5b\xa0\x39\x23\x6f\x34\x5b\xa5\xdb\xa6\x25\x64\xa5\x58\xe3\xb0\x55\xb1\xd7\x8f\x2c\xe3\x09\x72\x6b\x74\xf3\x0c\x43\xb8\x4c\x0e\x4f\x67\xcd\x2c\x62\xdf\x0c\x7c\x0c\xd6\x72\xc1\x3b\x59\xad\x08\x25\x90\xd9\xf8\xc4\x87\x94\x6e\xb8\xb2\x7d\xb4\xe5\x42\x5c\x48\x27\x9f\x90\x3e\x0c\x66\x57\x87\x13\xcd\x6f\x98\xe6\x68\xe4\xf1\xab\xf5\xf2\xba\xce\x13\x6a\x71\xbb\xcd\x99\xb9\xa4\xe0\x6b\xce\xc7\xe9\x87\x5d\xd4\xa2\x63\xf6\x8f\xd3\x9a\x39\xbb\x71\x5f\x8b\x0e\xb8\xd1\xf3\x83\xd7\x1f\xba\xe3\x8f\xe9\xa1\xa0\x67\xb8\xde\xcc\xec\x67\xc0\xe6\x6e\xfb\x1e\x3a\x33\xb4\x8d\xf6\xd1\x2a\x01\xa6\xef\x4b\x5d\x57\x72\x34\x1b\x3e\x71\xdd\x86\x58\x86\xc3\x7e\xf3\x30\x2d\x90\xc4\xf8\x90\x09\x56\xa5\x40\x7f\xcf\xc6\x9a\xfa\x19\x30\xbe\xa0\xff\x07\x00\x00\xff\xff\x4d\x4b\x11\x7e\x0b\x0c\x00\x00")

func ruleComplexwordsYmlBytes() ([]byte, error) {
	return bindataRead(
		_ruleComplexwordsYml,
		"rule/ComplexWords.yml",
	)
}

func ruleComplexwordsYml() (*asset, error) {
	bytes, err := ruleComplexwordsYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "rule/ComplexWords.yml", size: 3083, mode: os.FileMode(420), modTime: time.Unix(1494777592, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _ruleEditorializingYml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\x91\x41\x72\xdc\x3a\x0c\x44\xf7\x3a\x05\xca\x55\xbf\xbc\xfa\x17\xd0\x36\x27\x81\xa4\x1e\x0e\xca\x14\x30\x01\x40\x79\x92\xd3\xa7\xc6\x12\x25\x27\xbb\xd7\x04\xd8\x20\x9a\x78\x26\x74\x89\x91\xf0\x94\x48\xe8\x8c\x61\x45\x04\x17\x8c\xf4\xf6\xc3\x34\x64\x81\x93\x63\xb5\x4d\xb4\xd0\xfb\x7f\xf1\xfe\x36\x48\x51\x73\xcc\x1c\x18\x29\xbd\x61\xa8\xd8\x50\x47\xfa\x64\x57\xd1\x32\xa4\x7d\x40\x63\x1c\x88\xfe\x27\x9e\xb3\x71\xad\xbf\x76\xf1\xc8\x4e\x0e\x62\xd2\xb6\x4e\xf0\xaf\x83\x89\x43\xe6\xb3\x71\xae\x60\xef\x6c\xeb\xa3\x22\x71\x48\x44\x40\x53\xce\x56\x3c\x67\x60\x11\x2d\xdf\x74\xad\xd0\x3c\x54\x3a\xd6\x7e\xf7\xc6\xd2\x5d\x6f\xe6\xd9\x94\x4f\xdb\x7b\x2b\xf8\x02\xd1\x84\x23\xf2\x72\x94\xf8\xfb\xa9\x92\xaf\x23\x59\x1f\xe6\xc9\x9a\x94\x46\x0e\xae\xf2\x1b\xbd\x1c\x77\x6b\x75\xa1\x09\xa4\x96\x58\x28\xef\xbc\x3f\xa7\xb2\x97\x3e\x71\xb5\xe8\x69\xa8\x25\x4f\x17\xe3\xba\x60\xd3\x26\xd6\xe2\xa8\xd9\x8d\x66\x6b\x1e\xfb\xa0\x9f\x4d\x72\x27\x47\xe5\x94\xad\x3b\x3b\x56\xf6\x8f\xd3\x31\xb0\xc1\xb9\xee\x2c\x45\xe5\x26\x33\x6b\x9f\x1d\x6d\x8a\xd7\x1a\x57\xa2\xd1\xfc\xe1\x12\x57\x00\x29\x7a\x80\xe5\xd9\x95\xce\xe5\xdb\x8f\x35\xfd\x37\xd1\xa6\x29\x67\xf2\x1b\xfb\x6b\x8f\x83\x63\x5f\x6e\x83\xef\xd5\x4f\xc9\xbb\xb5\x24\xa6\xc5\xda\x94\xc3\x9f\x00\x00\x00\xff\xff\x04\x28\xfc\x00\x96\x02\x00\x00")

func ruleEditorializingYmlBytes() ([]byte, error) {
	return bindataRead(
		_ruleEditorializingYml,
		"rule/Editorializing.yml",
	)
}

func ruleEditorializingYml() (*asset, error) {
	bytes, err := ruleEditorializingYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "rule/Editorializing.yml", size: 662, mode: os.FileMode(420), modTime: time.Unix(1494782626, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _ruleGenderbiasYml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x95\x41\x8e\xeb\x36\x0c\x86\xf7\x3e\x05\x37\xc5\x9b\x59\xf4\x02\xd9\xcc\xa2\xc7\x78\x78\x0b\xc6\xa2\x65\x76\x64\xd2\x20\xa5\x71\x03\xcc\xe1\x0b\x59\x4e\x9a\x22\xd6\xf3\xc6\xb1\xfe\x0f\xd4\x4f\x91\x62\xe8\x9f\x4c\x12\xfc\x02\x5e\xae\x9e\x39\x97\xcc\x2a\xc3\x42\xee\x18\xe9\x02\x7f\xa9\x38\x07\x32\x28\xce\x12\xe1\xc7\x1f\xfe\x03\x58\x3c\x13\x06\xd0\x69\xff\x1e\x38\x8a\x1a\x8d\xe8\x74\x81\x6c\x85\x86\x44\x5f\x94\x2e\x40\x66\x6a\x83\x6f\xb8\x5e\x06\x80\xb7\x8f\x0b\xa6\xb2\x08\x7e\xef\xaf\xe2\xef\x17\x78\x3c\xd1\x30\x14\xcc\xf4\xcc\x51\x03\xf9\x8c\xf3\x06\x2e\x9a\x67\xb2\xef\x09\xeb\xeb\x3d\xa1\x84\x3b\x2b\x98\xf9\x8b\xa0\x2e\x9d\xa1\x59\x25\x16\xba\x3c\xa3\x6d\x69\x00\x40\xb6\xca\xff\x44\xfa\x25\xdf\x9b\xb6\x1f\x77\x13\x2b\x27\xcd\x6f\xfe\x5e\x39\x19\x67\xed\xa1\x87\xd8\xc0\x92\x67\x35\x72\x7f\x4a\xe4\xf1\x34\x71\x00\x18\x71\x21\xc3\x4e\xb8\x26\x82\xae\x64\x98\xef\x71\xc7\xf9\x37\x56\x9b\xd8\x38\x95\x58\xb7\x3f\x45\x17\x5a\xae\x54\xc1\x5a\xce\x3b\x39\x00\x04\xed\xe6\xb6\x47\x1c\x99\x2c\x52\xdb\x20\x18\x4e\xf9\x3c\xfc\x21\xd2\x61\x65\x62\xa3\x6e\xd4\x2a\x4e\x1c\xe7\x27\xda\x67\xea\xb9\x38\xc4\x06\x1a\xf9\xdc\x8b\x3b\xb1\x79\xfe\xf3\x46\x68\xe0\xb9\x04\x92\xa3\x7e\x11\xed\x8a\xb1\x63\x67\x43\xcf\x04\xa3\xa6\x44\xe3\xe3\xc0\x63\xb9\x9d\xd6\x70\x7f\xe8\x8b\xec\xa6\x52\x3b\x28\x61\xb8\x41\xc2\xed\x46\xf6\x8a\xb7\xf5\x83\x4a\xfc\x49\xa7\x11\x47\x2d\x96\x49\x8b\xef\xa0\x84\xa4\x16\x4e\xc1\x6b\xe1\x14\xea\xd5\x5c\x50\x30\xee\x81\x17\xe4\xd4\x3d\xe4\x2a\xc2\x88\x66\x4c\xe6\x3b\x2c\x80\x12\x60\xe3\xe9\xd5\xc9\x5c\xfc\x5a\xc5\x3b\x70\xf0\x24\x5a\xe2\x7c\x62\xc7\xb3\xa9\xc4\x43\x6f\xf0\x27\xcb\xb9\x71\x98\x4b\x8d\x55\xf5\x46\xae\xba\x9d\x9d\xd7\x83\x34\x72\x2d\x36\xd2\xe1\x7a\xc1\x70\x7e\x74\x55\x2c\x13\x8e\xb9\x18\xed\xa1\xa9\x65\x18\xd9\xd2\x4b\xf9\xee\xe2\xa6\x0b\x49\x85\x39\x84\xd4\x6b\x51\x96\x4c\xb6\x50\x60\xb4\xdb\x00\x20\xb4\xf5\x9a\x1e\xe0\x6f\x2d\x26\x98\xd8\x8f\x6e\xd3\xe5\x5a\xc2\x8e\xe3\xce\xa2\x3c\xcd\xb5\x26\x56\x4a\xa8\xac\x0b\x8a\xcf\xbc\xbe\x24\x57\x56\xcf\x18\x59\xe2\x00\xb0\x2a\xe5\xce\x44\x69\xe2\xce\x24\x1e\x7b\xb9\x34\x11\x74\x9a\x78\xbc\x5f\x23\xa3\xb5\x3f\x51\x32\x8d\xb3\xf0\xc8\x28\x0d\x76\x4c\xd4\x4d\x7f\x17\x57\x32\x57\x01\xb5\xf6\x09\x2b\xe9\x9a\x6a\x0f\x39\xd9\x57\xd7\x99\x6b\x0a\x7c\x77\xe4\x99\x36\xb4\xf0\xf6\x71\x21\xf7\xf7\x8f\xff\xa7\x3b\xa5\x3a\x2c\x00\x73\xfd\x07\x43\xa9\x29\x67\xe3\x6b\xd7\xd5\x2e\xfe\x37\xf0\x86\x7a\xcb\x39\xf7\x06\xf3\x2e\xee\xf7\x69\xaf\x16\x04\xad\x83\xe0\x85\x6c\xcb\x0f\xca\x47\x26\xc9\xec\xf9\xa7\xff\x7a\xb2\xfb\x58\x3e\xf6\x55\xfb\xec\x76\x4e\x15\x9b\xc1\x7f\x03\x00\x00\xff\xff\x43\x3a\xb5\xad\x9c\x07\x00\x00")

func ruleGenderbiasYmlBytes() ([]byte, error) {
	return bindataRead(
		_ruleGenderbiasYml,
		"rule/GenderBias.yml",
	)
}

func ruleGenderbiasYml() (*asset, error) {
	bytes, err := ruleGenderbiasYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "rule/GenderBias.yml", size: 1948, mode: os.FileMode(420), modTime: time.Unix(1494777592, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _ruleHedgingYml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\x8f\x41\xae\xdb\x30\x0c\x44\xf7\x3e\x05\x11\xa0\x48\xbb\xe8\x05\xbc\xe9\xa2\xab\x1c\x43\x8e\xc7\x32\x61\x89\x14\x48\xc9\x49\x6e\x5f\xd8\x72\xdb\xbf\xf2\xbc\xb1\xf0\xc0\xc1\xbb\x42\x66\x1f\x09\x6f\xf6\x0a\x79\x62\xc8\x70\x0f\x11\x23\xdd\x7e\xab\x38\xcf\x30\x32\x64\xdd\x59\x22\xdd\xbf\xf9\xfd\x36\x70\x14\x35\x3c\x83\x63\xa4\x6a\x0d\x43\xc2\x8e\x34\xd2\x2b\x98\xb0\xc4\xa1\xea\x06\xf1\x71\x20\xfa\x49\xa1\x14\x04\xa3\xaa\x34\xa1\x17\x16\x5b\x98\xd2\xa7\x83\xd3\x12\xec\xf8\x3c\x68\x13\x7d\x9d\xe5\x53\x5b\x9a\xcf\xb4\x98\x66\xca\x1f\x2a\xca\x52\x49\x17\xda\x19\xfd\xcd\xaa\x05\x4b\x4b\x97\xe6\x41\x13\x12\x63\xc7\x45\xb3\xb6\xa9\x5e\xb9\xae\x2c\xdb\x99\x59\x28\x42\x60\x21\xf5\x5f\x72\xa8\xb5\xb0\xb0\xca\xd9\x24\xde\x70\x19\x33\xc7\xb5\x1b\xb2\x1a\x48\x8d\x12\xdc\xcf\xa2\xc0\xd6\x50\xae\xac\xee\xfc\x77\x4c\x31\x78\xcb\xe1\x3f\xea\xf4\x0f\x1c\xc8\xdf\xfd\xc7\xaf\x0e\x9a\xf1\x5a\x43\xf7\x57\x3d\xce\x38\xc6\x27\xcc\xb1\x2f\x68\xf2\xe5\x94\xe6\x2d\x1c\x43\xff\x04\x00\x00\xff\xff\xf4\x7e\x73\x23\xab\x01\x00\x00")

func ruleHedgingYmlBytes() ([]byte, error) {
	return bindataRead(
		_ruleHedgingYml,
		"rule/Hedging.yml",
	)
}

func ruleHedgingYml() (*asset, error) {
	bytes, err := ruleHedgingYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "rule/Hedging.yml", size: 427, mode: os.FileMode(420), modTime: time.Unix(1494782621, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _ruleLitotesYml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\x90\x41\x72\xe4\x20\x0c\x45\xf7\x3e\x85\x36\x53\xe9\xbe\x82\x36\x5e\xcc\x49\x64\xf3\x4d\x93\x80\x70\x49\xd0\x4e\xdf\x7e\xaa\x1d\xdb\x99\x15\x5f\xef\x15\x48\x02\xdf\x0d\x1a\x9c\xc9\xfb\xe4\x2d\xb5\xde\x52\xd5\xa1\xc0\x5d\x22\x98\xfe\x56\xf5\x14\x60\xd4\x3d\x69\xa4\x8f\x3f\xfe\x41\x49\xbd\x41\x02\xd5\x65\xaf\x87\x14\xb5\x1a\x66\x71\x30\x35\xeb\x18\x32\x9e\xc8\x4c\x9b\x98\x26\x8d\x83\x6f\xb2\xf2\x40\xa4\x95\x4a\xd2\x6a\x4c\x45\x3e\xab\xfd\x90\x6a\x21\xa9\xd8\x8b\x09\xdf\xcd\xe4\x2c\x77\x79\x1b\xb9\xdd\x47\xf2\x22\x39\x33\x65\xb1\x88\x9d\x37\xba\x8d\x2c\xf7\x91\x26\x09\x4c\xb1\xd6\x70\x60\x99\x67\xac\xed\x36\x32\xc2\x7d\x64\x32\x7c\x62\x6e\xb8\x6c\xce\x75\xbb\xe4\x6a\x78\x42\x7f\xed\x0c\x6b\x92\x94\xa9\xeb\x11\x4f\x71\x7c\xc1\x75\xf3\x67\xdf\xf3\xde\x22\x29\x33\xad\xe2\x7e\xa1\x87\x3c\xf1\x9e\x77\xfe\xba\xf9\xfd\x60\x45\xf4\xc5\xb4\x60\x3b\xea\x9a\x03\x41\x6b\x8f\x0f\xa6\x56\x2b\xbd\x6a\xd7\x78\x38\x43\x41\x99\x60\x4c\x4b\xb5\x58\xdb\x81\xbd\xcf\x33\x10\x78\x6f\x79\x35\x6b\x0f\xd0\x64\x29\x3e\x1a\xbc\x31\x85\x5e\xa6\xff\x8c\x4b\x01\x53\x48\xcb\x02\x83\x9e\x2f\x75\x95\x4d\x0c\x4c\xfb\x71\xc1\x39\x43\x94\x69\x3f\x2e\xb8\x48\x49\x39\xc9\x7b\x98\x23\x5d\x2a\xa7\xaf\xf7\x9e\xe9\xeb\xf7\x89\x35\x43\x5c\xb4\x31\x9d\xe9\x54\x8e\x0c\x77\x7e\x87\xa5\xe7\x83\x6e\x56\x35\x32\xed\xe3\x0f\xff\x02\x00\x00\xff\xff\x69\x1e\x9e\x3b\x8b\x02\x00\x00")

func ruleLitotesYmlBytes() ([]byte, error) {
	return bindataRead(
		_ruleLitotesYml,
		"rule/Litotes.yml",
	)
}

func ruleLitotesYml() (*asset, error) {
	bytes, err := ruleLitotesYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "rule/Litotes.yml", size: 651, mode: os.FileMode(420), modTime: time.Unix(1494777592, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _rulePassivevoiceYml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\x55\x5d\xce\xe3\x3a\x08\x7d\xef\x2a\xa2\x4f\xba\xea\xfd\xd1\xdd\x40\xb7\x32\x9d\x07\x37\x21\xb1\x55\xc7\xf6\x80\x1d\x7f\x1a\x75\xf1\xa3\x36\x1c\x92\x27\x0e\x04\x30\x1c\x83\x43\xdf\x95\xd2\x24\xb7\x81\xbe\x83\x54\x4a\x23\x5d\x56\x12\x71\x0b\xdd\x86\xaf\xeb\x5f\x72\x1d\x82\x0c\xc5\x89\x84\x8d\x86\x2d\x87\x91\xbe\x2e\x61\x49\x99\x69\x74\x42\xb7\xa1\x72\xa3\x4b\xa4\x8d\xe2\x6d\x90\xb6\x2c\x24\x35\xe4\x74\x61\xd7\x6f\x97\x61\xf8\x7f\xb8\x3f\xfe\x76\xeb\xcb\x31\xbd\x3a\x31\xbd\x1e\x14\xd2\xf2\x0a\xf2\x7a\x10\xa5\x57\x77\x6f\xf0\xcf\xfd\x71\x97\x7f\x2f\x35\x3f\x29\xc9\x1e\x75\xfd\x71\xef\x3f\xff\xa3\xe9\xfa\xd1\x5c\x7f\x7f\xfa\xc0\x07\xb9\xaa\x60\xcc\x2b\x29\xb4\x8f\x4b\x03\x4a\x70\x13\x02\x52\x19\x26\xc8\x09\x71\xa1\x56\xc0\x48\xfa\x39\xe6\xae\xa6\xcc\x00\x6d\xf1\x15\x30\xa9\x1f\x23\x80\xb3\x9b\x46\x27\x15\x9a\x95\xcc\xa7\xb8\x16\x22\x10\x27\x43\x1a\x64\xd1\xa3\xb3\x88\xd1\x67\xd1\x44\x63\x6c\x69\xd9\x11\x5a\x1f\x33\x22\x98\x8a\xa2\xb6\xcb\x89\x9c\x1e\x35\x85\x4d\x6b\x9c\x72\xda\xe3\x26\x76\xda\xdd\xc4\xe4\x56\xf5\xe3\xb0\x11\xac\x2d\x3d\x77\xd4\xf6\x23\xc9\x81\xa2\xd9\xc5\x08\xa8\x79\x67\xd2\xa3\xe6\xa0\x12\x34\xce\x46\xe3\x6c\xd5\xcf\x99\x4f\xe4\xcf\x99\x69\x41\x5d\x73\xe6\xc5\x8a\x78\x2b\xd9\x6e\x66\xce\x2c\xee\x69\x8a\x11\x34\xdb\x55\xcc\x9c\x7f\xeb\xf7\x23\x89\xa5\x3e\xa5\x5a\xd8\x62\x16\x46\x7d\x9e\x1c\x4f\x8a\xa2\x82\xa3\x4a\xaf\x8d\x79\x34\xe1\x1b\xef\x96\x27\x88\x7f\x26\xb0\xf0\x4c\x01\x00\xd9\xa3\xd3\xb1\x8b\x2e\xa8\x85\x9c\x06\x46\x72\x98\x05\xd0\x16\x69\x86\xc1\xbe\xa8\x0c\x8b\xaf\xf0\xc2\xed\xaf\x6e\xda\x9b\x5c\xc9\xa9\xff\xaa\xfe\x6b\x10\x29\x28\x6c\x0d\x52\x8d\xc3\x15\xb5\xe5\x8d\xd8\x26\xea\xad\xd8\x98\xbc\x95\x23\xe0\xa3\x79\x23\xac\xa0\xa5\x82\xaa\x0b\x67\xd0\x5e\x74\x0a\x7f\x35\xa5\x82\xc9\xed\x4e\x1c\x4c\x82\x5c\x0e\x18\x71\x6e\x26\x77\x9a\x05\x87\x88\xee\xbd\x60\x70\x05\x5b\x2f\xa0\x08\x9b\x2e\x04\x17\x6f\xb5\x8b\x77\x9b\x41\x2d\x57\x7c\x36\xa0\x0d\x8b\xc7\xb2\x8b\xcf\x15\xc0\xd2\xd9\x56\x88\xd7\xfe\xc4\xee\x53\x22\x06\x41\x22\x6a\x8e\x01\x16\xeb\x67\x3d\xde\x1a\xc9\x3a\x67\x72\x4c\xb3\xd8\x61\x05\x55\x16\xeb\xaf\xe0\xf1\x90\x82\xc4\xc5\x8e\x28\xf6\xe0\x48\x31\xb2\xa5\x1c\x4c\x16\xe5\x56\x6a\xc6\xfe\x4a\xcd\xa0\xa0\x9e\xee\xe3\xad\x6c\x07\x6e\xe3\xd3\x20\x92\xd5\xc3\x78\xb2\x81\x1d\x33\x99\xa5\x1b\x39\x3d\xdb\xf3\x21\xdd\xd8\xee\x6d\x05\xd0\xd8\x63\xee\xea\xf1\x1a\x56\x9f\x4f\x98\xed\x61\x3b\x8d\x65\xf5\xdc\x74\x2d\x2a\x08\xae\x38\xa7\x72\xb6\x2e\x5b\x9a\x88\x0f\x06\x5a\xb1\xbd\x6f\x05\xa3\xd4\x35\xbf\x95\xdf\x43\xf5\xe6\xf7\x56\x8e\x04\xc7\x4f\xaa\x67\x48\x06\xc0\x7b\xd3\x6d\x43\x3a\x1f\x93\xd0\x3f\xc4\xfe\x09\x00\x00\xff\xff\x73\xa8\xee\x60\x8b\x07\x00\x00")

func rulePassivevoiceYmlBytes() ([]byte, error) {
	return bindataRead(
		_rulePassivevoiceYml,
		"rule/PassiveVoice.yml",
	)
}

func rulePassivevoiceYml() (*asset, error) {
	bytes, err := rulePassivevoiceYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "rule/PassiveVoice.yml", size: 1931, mode: os.FileMode(420), modTime: time.Unix(1494777592, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _ruleRedundancyYml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x53\xc1\x6e\x1b\x47\x0c\xbd\xfb\x2b\x88\x00\x45\x22\xa0\x09\xd0\x1e\x75\x31\x1c\x07\x4d\x05\x24\x85\x60\x3b\x39\xf9\x42\xed\xbc\x5d\xb1\x9e\xe1\x2c\x38\x5c\x45\x42\xb7\xff\x5e\x8c\x66\xe5\x38\x05\xda\xdb\x90\x33\x7c\x7c\xf3\xf8\x88\xa3\x43\x43\x59\x13\x8e\x52\x1c\xda\xe1\x2a\xa1\x14\x1e\xb0\xa6\x57\xaf\x7f\x2a\xaf\x49\x0a\x19\xc2\xa4\x81\xd5\x5f\x5d\xc9\xa0\xd9\xd0\x71\xc1\x9a\xdc\x26\x5c\x45\x1c\x10\xd7\x04\xb3\x6c\x57\x9e\x9f\xa0\x65\x7d\x45\xf4\x96\xde\x5c\xaf\x39\x04\x84\x19\x47\x37\x5e\xd1\x2e\xeb\x54\xbe\xdf\x1c\x58\x3b\xcc\xa3\x49\xb6\x15\x69\x76\xe9\x70\xb9\x1c\xd8\xf7\xb0\x99\x4b\x41\xda\x45\xcc\xec\xce\xdd\x7e\xde\x45\x68\x98\x93\x1c\xe7\x2e\xa7\x9d\x28\xe6\x2e\xc7\xc8\xbb\x6c\xec\xf5\xac\x8a\xce\xe7\x2e\xe7\x11\x35\xb3\x22\xcf\x03\x2a\xd2\x05\xd7\x61\x26\x15\x70\x9f\xdb\x61\x45\x6e\x3c\x20\x9c\xce\x2f\x1e\xc3\x5f\xbf\xfc\xfc\xeb\xdf\x74\xcd\x8f\xef\xae\xd3\xe3\xbb\x6b\x12\x25\xdf\x83\x52\x36\x15\x1d\x7e\x7c\x34\x5e\x1e\xb1\x93\xca\xb0\xf7\xff\xb8\x5e\x30\xb8\x77\x98\xe6\xac\xff\xff\x0c\x07\x3c\xb7\xba\x79\xff\x99\x92\x94\x22\x11\xff\x8e\xcb\x92\xb8\xa7\x9d\xf1\x93\xe8\x40\xe5\x54\x1c\xe9\x9c\xe6\x5d\xc9\x71\x72\xc4\x53\xfd\xb4\xa2\xab\xf3\xb4\xd3\x8c\x52\xa0\x2e\x1c\x57\xad\xfa\xf6\x81\x1c\xa5\xf1\xe6\x68\xe0\x70\x6a\x2e\xb8\x10\x60\x4b\x08\x34\x4c\x9a\x58\x97\x84\x4b\x2f\x9d\x70\xa4\xd1\x72\xf1\x3d\x8a\x2c\x54\x1e\x3e\x53\xe2\x6e\x2f\xda\xb8\xee\x38\x86\xb7\x7b\x70\x40\x68\xb1\x09\xfa\x4a\x27\xe5\x04\xf5\xb9\x4c\x29\xb1\x9d\x1a\x93\xdb\x0f\x14\xa4\x74\xe7\x73\x07\x2e\x20\xd6\x40\xa1\x82\x37\x76\x9d\x58\x17\x41\x6c\x79\xd2\x86\xd7\xc5\x5c\x50\x49\x1c\x25\x89\x9f\xbe\xe7\x02\xf5\xcf\x55\x39\x8d\x11\x17\x1d\x02\x8a\x5b\x3e\x55\x3f\x46\x49\xa2\xd5\x33\xd0\x61\x8a\x3d\xc2\xdc\x4b\x8c\x08\x73\x99\xec\xdc\x62\xb5\xd4\x4f\x31\xd0\x98\x4b\x91\x5d\x6c\x2d\x6e\xb7\x1b\xda\x68\xc0\xb1\x3d\x98\xcc\xa0\x4e\x6e\xd0\xf0\xa6\xac\xae\xcf\x59\x68\x20\x43\x99\xa2\x2f\xa1\xc3\x48\x9a\x80\xb0\x1c\x40\xfc\x8d\x1b\x5c\xcf\xf1\xfc\x0b\x38\xb4\x34\xe1\x7a\x51\x8e\x2f\xeb\x7b\x03\x68\x90\xbe\x45\x1f\xb7\xf7\x2f\x47\xfd\xf1\xcb\x86\xa4\x36\xe8\x79\x59\x9f\xdf\x37\x5f\xe9\x20\xb6\x6c\x9a\xa8\xc2\xa8\x07\xa2\xe8\xd0\x52\x9b\xfb\xf7\x7f\x90\x4e\x69\xb7\xec\xc5\x9f\xb9\x5a\xef\xe5\xa6\x7c\x6a\xf3\x18\xe3\xc2\x32\xe6\xae\x51\x92\x00\xf5\xe7\x7f\x2a\xfb\x64\x1c\x49\xb4\x3a\xa6\x6b\xfc\x14\xdf\x48\xf4\x50\x7d\xb6\x58\x5d\xa7\x18\xcf\xf3\x3c\x64\x69\xb3\x1b\xb9\x38\xed\xa5\x78\xb6\xd6\x60\xfb\xe1\x37\xea\xb3\x25\x6e\x18\xdb\xcd\x0f\x04\xc7\x78\xb6\xc4\xdd\xfd\xd7\xed\x39\xbe\xbb\xf9\x4c\x09\xe9\x52\x7c\x77\x53\x15\xd1\x60\x39\x35\x05\xee\x36\xdb\xba\x4d\x23\x2e\x92\x14\xee\x41\x7b\x3e\xa0\x11\xba\xbf\xf9\xf4\x40\xce\xf1\xa9\x2c\xe1\x8b\x35\x28\xfb\x3c\x39\xc5\x3c\x85\x65\xe0\x17\x47\x20\x50\x56\xe2\x18\xa9\xaa\xd0\x2a\x6b\xda\x06\xe3\x30\xb1\x83\x8a\x4f\x55\x9e\xe5\x06\xc7\x11\x9d\x23\x54\x80\xd1\x64\x19\xee\xa4\x75\x58\x67\xb4\x24\xc5\xf9\xa9\xa5\xbf\x6c\x6f\xa9\xcb\x01\xcf\xd2\x1e\x24\xc7\x6a\x2c\x1c\xc7\x98\x4b\x55\xf2\x9f\x00\x00\x00\xff\xff\x2b\x6c\x38\x4c\xa7\x05\x00\x00")

func ruleRedundancyYmlBytes() ([]byte, error) {
	return bindataRead(
		_ruleRedundancyYml,
		"rule/Redundancy.yml",
	)
}

func ruleRedundancyYml() (*asset, error) {
	bytes, err := ruleRedundancyYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "rule/Redundancy.yml", size: 1447, mode: os.FileMode(420), modTime: time.Unix(1494777592, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _ruleRepetitionYml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x1c\xc9\x4b\x0a\x84\x30\x0c\x06\xe0\x7d\x4e\x91\x29\x0c\x59\x0c\x73\x81\x5c\xc5\x07\x14\xfa\x23\xc5\xda\x4a\x12\xc4\xe3\x0b\xae\x3f\xdc\x81\x5e\x5c\xd9\x70\x22\x6a\xd4\xd1\xe9\x80\x7b\xde\xa0\x9c\xe4\xeb\xc2\xd5\x5f\xcc\x81\xf2\x49\xd4\x70\xa1\x29\xc3\x6c\x18\xc5\xd8\xd1\x5d\x89\xf9\xcf\x32\xad\xb3\x2f\x3f\xa1\x27\x00\x00\xff\xff\xfa\x4a\x06\xdd\x53\x00\x00\x00")

func ruleRepetitionYmlBytes() ([]byte, error) {
	return bindataRead(
		_ruleRepetitionYml,
		"rule/Repetition.yml",
	)
}

func ruleRepetitionYml() (*asset, error) {
	bytes, err := ruleRepetitionYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "rule/Repetition.yml", size: 83, mode: os.FileMode(420), modTime: time.Unix(1494777592, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _ruleUncomparablesYml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x92\x51\x8e\x9c\x30\x10\x44\xff\x39\x85\xb5\x52\x34\x49\xa4\x5c\x80\x9f\x5c\x64\x7f\x1a\x28\xb3\xad\xb5\xbb\x99\x6e\x9b\x1d\x24\x1f\x3e\x62\x80\x59\xe5\xef\x95\x29\xe8\x72\x35\x78\x14\xc8\xe4\x7d\xc0\x83\xbd\x40\x46\x74\x19\xee\x34\xa3\x0f\x6f\xb7\x1f\x7e\x0b\xec\x41\xb4\x84\x51\xf3\x42\x46\x43\xc2\x5b\xc7\xb3\xa8\x61\x24\x47\x1f\x8a\x55\x74\x09\x2b\x52\x1f\x60\xa6\xd6\x19\x7d\xf5\x5d\x08\x7f\xc2\xfb\xf0\xf3\x6f\x4f\x83\x6b\xaa\x05\x69\x6b\x59\xbd\xb4\xac\x86\x96\xe0\xde\x12\xc8\x4b\x5b\x61\x5b\xbb\x57\x2e\x68\x89\x6c\xde\x7d\x78\x14\x43\xde\x89\x65\x34\x90\xb3\xcc\x69\x6b\x9f\x2c\x53\xd0\xd8\x32\xa7\x69\x6b\x1f\x64\x53\xda\xda\x6c\xa0\x92\xb6\xe6\x6a\x25\x68\xfc\xf5\x3e\xbc\xfb\xef\xae\xe8\x27\xc4\x8f\x14\x57\x80\x43\x4c\xb8\x57\x3a\xc5\x7e\xa5\x84\x97\x30\xc3\x58\x0e\x86\x15\x62\x79\xf2\x84\x55\x79\x7a\x22\xa4\xb0\x1d\xee\x5b\xa4\xe4\xb8\x3d\x39\x52\xa1\x74\xd2\xaa\xc6\xe7\x07\x23\xcb\x79\xcc\x13\x2e\xca\x8b\xba\xf3\x90\x0e\x0b\x0b\x56\x2e\xf4\x2d\x23\xcb\xf5\x3a\x9b\x61\xd5\xf1\xf5\x30\x5f\x81\x32\x09\x47\xf8\x91\x54\x25\x6d\x4f\xd8\x77\x93\xb5\xca\x71\xbc\xc0\xe2\x75\x99\x05\xb6\xa0\xd4\x33\xc1\x7f\xf3\x17\x43\x84\xd1\xb7\x64\x19\x79\x39\x9d\x7b\xeb\x35\x91\x1d\xa2\x50\x61\x15\xb2\x63\x9a\xd7\x18\x79\x64\x9c\xe3\x6e\xfb\x4f\x70\x94\x51\x85\x84\xb3\x56\xbf\xd4\xde\xde\x6b\x42\x95\xc1\xf6\xd5\x9c\x82\xa3\x5a\xbe\xf8\x5e\x2f\x0f\xaf\x30\x3f\x53\xbc\xca\xff\xfa\xd0\x84\xee\x5f\x00\x00\x00\xff\xff\x48\x92\xd0\xc4\xaf\x02\x00\x00")

func ruleUncomparablesYmlBytes() ([]byte, error) {
	return bindataRead(
		_ruleUncomparablesYml,
		"rule/Uncomparables.yml",
	)
}

func ruleUncomparablesYml() (*asset, error) {
	bytes, err := ruleUncomparablesYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "rule/Uncomparables.yml", size: 687, mode: os.FileMode(420), modTime: time.Unix(1494777592, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _ruleWordinessYml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x56\xcb\xb2\xdc\x3c\x0a\xde\xfb\x29\xd8\x4c\xe5\x9c\xc5\xbc\x80\x36\x67\x31\x8f\x31\x35\x0b\xda\xc6\x36\x89\x2c\x14\x5d\xba\x8f\xa7\xf2\xf0\x7f\x81\xe4\x4b\xf2\x27\x55\xff\xa6\x1b\x3e\xb0\x2e\x08\x3e\xa0\xcf\x42\x61\xca\x0e\x72\x7d\xe4\xc2\xa5\x16\x96\x30\x6c\x94\x33\x2e\xe4\xe0\x3f\x12\x32\x4f\x94\xa0\x66\x0e\x0b\x7c\xf9\x57\xfe\x02\x1c\x72\x21\x9c\x40\x66\xd3\x07\x5e\x82\x24\x1a\x31\x93\x83\x92\x2a\x0d\x9e\x9e\xe4\x1d\xbc\x30\x05\x0e\xcb\x90\x5f\x18\xdd\x00\xf0\xf6\xe1\x16\x7e\xd2\x8f\x05\x9f\xf4\x0e\x89\x33\x41\x11\x07\x5e\xd7\x2a\xd2\x1c\x62\xa2\x27\x4b\xcd\x3f\x62\x62\x49\xef\xe6\xf0\xa0\x59\x12\x0d\x00\xa8\x1e\x1e\xd3\x42\xef\x1f\xb0\xe1\x57\x49\x5c\x76\x90\xd9\xc1\x26\xb9\xfc\xea\x10\xea\xf6\xa0\xd4\xcc\x18\x76\x33\x6f\x7b\x62\x3b\xb9\xeb\xa2\xa2\x13\x7d\xaf\x58\xa8\x7f\xe0\x80\x82\xd4\x65\x35\xcb\x93\x52\x26\xbf\x03\x6f\x11\xc7\xe2\x60\xad\xc9\xf6\xf1\x1e\x70\x4c\x92\xb3\xeb\xff\x1d\x94\x19\x10\x72\x9d\x26\x0a\xae\xff\xfb\xfd\xb2\x95\x95\x2c\x48\xfa\x77\xa2\x4e\xff\x9b\xf6\xef\xc2\x1b\x41\xa2\x51\xd2\xe4\xfa\xbf\x59\xf4\x7e\xea\x70\xdd\xb4\x41\x81\x9e\x7a\xe2\x4c\x7e\x92\xcd\x60\x09\x8b\xae\x0f\x9e\x03\x65\x5b\x3d\xf3\xc6\x1e\x53\x0b\x31\x06\xc0\x18\x13\x8d\x8c\x0f\x4f\xbf\x8b\x51\x00\xca\x85\x37\x2c\x34\x39\xc0\x87\x54\xdb\x2e\xec\x80\x61\x6a\x47\xe8\xc7\x4d\x04\x1c\x00\x97\x44\xb4\x51\x28\xae\x89\x6a\xc9\x1a\x69\x2c\xc5\x56\x86\xd9\x22\xc7\xc1\x84\xd3\x4c\x18\xda\xf1\xda\xb1\x14\x4b\x94\xab\x2f\x06\x3e\x68\xc4\x9a\x09\x64\x6e\x46\x99\x61\xa7\xe2\xf4\xa7\x01\x51\xaf\x1d\x29\xa9\x56\x00\xc1\xa3\xee\x36\x61\x21\xd7\xe4\x6e\xf0\x1e\x34\xa6\xfa\x4e\xfe\x85\x7b\x6e\xb0\x06\x28\x26\xca\x14\x8a\x99\x1d\x04\x79\xfd\x03\x13\x67\x88\xc2\xa1\xe8\x6d\x9a\xf1\x27\x78\x00\x78\x60\xa6\x49\xcd\x96\x85\x10\x31\x15\x90\xe0\x3a\x2e\xe1\x74\x91\x60\x5b\x69\x4c\xa0\xac\x58\xce\x3b\x37\x0f\x1e\x21\xd0\x48\x39\x73\xd9\xdd\x25\xaa\xf1\x0c\xcd\x9f\x17\x18\x71\xd3\xda\xb2\x98\xa2\xe7\xff\xa3\x16\xb5\xeb\x0a\x4d\x77\x8f\x00\xf8\x48\x35\x16\xa0\x30\x69\xea\x4f\xff\xa5\xe9\x7f\x1d\xb3\xd4\x1d\x31\xa5\x1d\xa4\x16\x4b\x8d\x27\xfa\x6a\xab\xd9\x33\x75\xd5\xb6\xf4\x92\x09\x26\x79\x05\xd7\xe4\x03\x9b\xee\xa0\x6d\x2d\x5b\xf4\x54\x08\x72\x49\x18\x16\xcb\xdf\x2e\xdd\xac\x7e\x87\x4c\x11\x93\xbd\xe8\x21\x99\x3d\x8c\x64\xbc\x62\xd7\x3f\xd3\x4c\x6f\xb7\x60\x9a\x94\x70\xcc\x6b\xaa\x63\xb1\x00\x3c\x99\x5e\xdd\x41\xc5\xbb\x35\x00\x87\xa7\xa6\xfb\xd2\x23\x74\xa9\x74\xf3\xa3\xcf\x48\x89\x35\xc9\xb3\xbb\x29\xcd\xa3\x70\xa8\x64\x6f\x7c\x28\x03\xc0\x44\x39\x72\xa1\x5f\x5f\x08\x7d\x59\x3b\xb7\x4c\x9c\x31\x46\xc2\x04\x73\x92\x0d\x32\x2f\x6b\x71\x17\xaa\x1e\x22\x1b\x29\x33\xc2\x8c\xec\x5d\xd7\xd5\x50\xed\xe9\xfe\xf8\xfa\x53\x4d\x47\x78\xf4\xa8\xd2\xf8\xae\xa1\x3f\x9b\x8d\x6e\xda\xe7\xaf\x95\xbd\x7e\x4c\x1b\xa5\x85\xc2\xb8\x43\xe6\x52\x7b\x58\x4e\x50\x1d\x3e\x47\x8a\x05\x5e\xab\x52\x5c\x0d\x9e\x8c\xfc\x14\xcd\x99\x9f\x17\x8b\x16\x91\x83\x53\x5a\x8f\xe9\xe1\xe6\x72\xc5\x9a\x2d\xcc\xb3\x56\x69\x4b\x92\xb9\x71\xcb\x4c\x17\x42\x0d\x91\x64\x27\x9e\x6a\xba\xd2\xef\xbc\xd2\x82\x65\x25\x65\xb8\x85\x54\x70\x1d\x18\x00\x56\xcc\xf6\x1d\x3e\xd8\x5b\x29\x8d\x18\x6e\xf0\x88\x11\x47\x6d\x22\xda\x67\x7e\x36\x49\x8c\x92\x4a\x0d\xa7\x55\xaa\xd7\xe8\xaf\xe2\x27\x23\x30\x2a\x1c\x16\x67\xc2\x00\xc0\x73\x23\x02\xce\x10\xa4\xf4\xc5\x95\xec\x79\x56\x40\x3d\x02\xa0\x56\x13\xcd\xd5\x6b\x64\x82\x1e\xb4\xeb\x56\x69\xe6\xd0\x12\xa4\xdc\x7d\x2e\xe8\xe6\xc6\x9b\x96\xc8\xe9\x62\x6a\x37\x06\xa0\x79\x96\x54\xec\xd8\x46\xaf\x1c\xe0\x41\xe5\x45\xfa\x64\x5d\x68\xa8\x67\xaa\x16\xca\xab\xa7\x37\x83\xbe\x9c\x5d\x20\x3b\x90\xb9\x1c\xfe\xd6\x74\x3a\x5c\x73\xc5\xf3\x44\x92\x26\x7b\x80\x6b\xc3\x2c\x1b\x1d\xae\x2a\x1b\x0d\x77\x93\xd5\xc6\xdf\x09\xec\x56\x1e\x37\x2f\x77\x54\x53\x83\xf5\x9b\xb7\x0f\xf7\xa4\xb4\x6b\x93\xb7\x0a\xaa\xa5\x26\xa5\x09\x91\x70\x39\xd1\xd3\x38\xdc\x56\xe6\xf9\xc2\x03\xf1\xb2\x3e\x24\xad\xd2\x2b\x23\xe9\x9e\xc7\x45\xd4\xe3\xc9\x23\x87\x63\xb4\x68\xc4\xd6\x6e\x55\xe0\xa5\x59\x00\xbd\x72\xfb\xa9\x63\xc4\x44\xa1\xb1\xa5\xe7\xb9\x40\x8d\xce\x84\x01\x60\xc3\x49\x1b\xfa\x4c\x89\xc2\xd8\xe6\x1d\xd3\x12\xf5\x99\x67\xc3\x6f\xbf\x75\xe8\x56\xfe\xbc\x65\xf5\xc6\x9f\x03\x40\x90\x40\xbd\xb7\x39\x53\x0c\x2b\x2d\x2f\xa2\x64\xb6\xfa\xd0\x85\x6a\xd0\x26\xdf\xcd\x51\x72\xe6\x87\xd7\x8c\xdc\x0e\x79\x00\x7d\x05\x9b\xa4\x0c\x4d\x05\xc3\xd8\x3d\x54\xd6\x1b\x44\x4a\xb3\xa4\xcd\x3a\x45\xce\x94\xb3\x52\x5f\x1b\x5b\x4c\x6d\x2e\x05\xb9\x71\xb2\x5c\x43\x43\xf4\x38\x92\x7e\x67\xc9\xe1\xda\x5f\xc3\x77\xed\xf7\xdf\x68\x87\x24\x5e\xc7\x08\xa7\x75\x43\x59\xbb\x2e\xa3\x6f\x77\xff\x5d\x17\x4e\x84\x13\xfb\xfd\x8c\xf9\x15\xfd\x01\x5a\xc6\xb5\xac\x6a\x29\xa7\x58\xd4\x73\xf7\x51\x4d\x15\xc5\xea\x23\xd3\xf7\x6a\x6b\xeb\x71\xe7\x36\x2e\xe4\x3a\x2a\x87\x59\x99\x9d\x6d\xc8\x9d\x92\x79\xcc\x33\x8f\xac\x1f\xf6\xc1\xe9\xed\xc3\xc9\xfc\xfe\x71\x9b\x19\x8b\xbe\x27\x8e\x8d\xe2\xda\xc0\x63\x10\x07\xed\xb6\xe3\x28\x55\x4f\x3d\xf6\xb1\x5a\xad\x2b\xc1\xf7\xaa\x5d\x47\x34\xc2\xca\xec\xaf\xb5\x3f\x78\x17\x9a\x97\xce\x5b\x4a\x2e\x30\x49\x7d\x14\x78\xd4\x23\xb9\x4d\xef\x3c\x6c\x1c\x34\x61\x9f\xd6\x16\x1b\x36\x39\xab\x74\x18\x55\xd3\x10\x7c\xa5\xb3\xee\x0c\xef\x90\xba\x69\x57\x78\xfb\x70\x73\xc2\x8d\x7e\xb4\x16\xf2\xde\x38\x66\x00\xa8\xc1\x2a\xdd\xc6\x24\x79\x72\x66\xe9\x83\x9c\x19\xcc\xa1\xb0\xd7\x68\xae\x6d\x25\x54\xae\x50\x4c\x6d\x3a\x0f\x28\x97\xcf\x95\x3c\xc4\x9a\xa2\x74\x2e\xb9\xe1\x03\x1c\x17\x07\x49\x9a\xbb\xf7\x40\xbc\xb8\xac\xbf\xd4\xcb\x91\x6f\xdd\xa4\x53\x40\xaf\xa3\x6b\x20\xe8\xb6\x1c\xed\xd6\xc7\x47\xba\xbe\x51\x86\x99\x8d\x34\xac\xbb\x9d\xf3\x4d\xeb\x75\xb3\xa4\xe1\xaf\x00\x00\x00\xff\xff\xec\x9b\x3b\x9b\x28\x0d\x00\x00")

func ruleWordinessYmlBytes() ([]byte, error) {
	return bindataRead(
		_ruleWordinessYml,
		"rule/Wordiness.yml",
	)
}

func ruleWordinessYml() (*asset, error) {
	bytes, err := ruleWordinessYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "rule/Wordiness.yml", size: 3368, mode: os.FileMode(420), modTime: time.Unix(1494777592, 0)}
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
	"rule/Annotations.yml": ruleAnnotationsYml,
	"rule/ComplexWords.yml": ruleComplexwordsYml,
	"rule/Editorializing.yml": ruleEditorializingYml,
	"rule/GenderBias.yml": ruleGenderbiasYml,
	"rule/Hedging.yml": ruleHedgingYml,
	"rule/Litotes.yml": ruleLitotesYml,
	"rule/PassiveVoice.yml": rulePassivevoiceYml,
	"rule/Redundancy.yml": ruleRedundancyYml,
	"rule/Repetition.yml": ruleRepetitionYml,
	"rule/Uncomparables.yml": ruleUncomparablesYml,
	"rule/Wordiness.yml": ruleWordinessYml,
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
	"rule": &bintree{nil, map[string]*bintree{
		"Annotations.yml": &bintree{ruleAnnotationsYml, map[string]*bintree{}},
		"ComplexWords.yml": &bintree{ruleComplexwordsYml, map[string]*bintree{}},
		"Editorializing.yml": &bintree{ruleEditorializingYml, map[string]*bintree{}},
		"GenderBias.yml": &bintree{ruleGenderbiasYml, map[string]*bintree{}},
		"Hedging.yml": &bintree{ruleHedgingYml, map[string]*bintree{}},
		"Litotes.yml": &bintree{ruleLitotesYml, map[string]*bintree{}},
		"PassiveVoice.yml": &bintree{rulePassivevoiceYml, map[string]*bintree{}},
		"Redundancy.yml": &bintree{ruleRedundancyYml, map[string]*bintree{}},
		"Repetition.yml": &bintree{ruleRepetitionYml, map[string]*bintree{}},
		"Uncomparables.yml": &bintree{ruleUncomparablesYml, map[string]*bintree{}},
		"Wordiness.yml": &bintree{ruleWordinessYml, map[string]*bintree{}},
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

