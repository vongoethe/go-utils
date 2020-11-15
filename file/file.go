package file

import (
	"io/ioutil"
	"os"
)

// Exist 判断文件是否存在
func Exist(path string) bool {
	_, err := os.Stat(path)
	if err != nil && os.IsNotExist(err) {
		return false
	}

	return true
}

type FS interface {
	Exist(path string) bool
	WriteFile(path string, data []byte) error
	ReadFile(path string) ([]byte, error)
	MkDir(path string) error
	Delete(path string) error
}

type LocalFS struct{}

func (l LocalFS) MkDir(path string) error {
	return os.MkdirAll(path, os.ModePerm)
}

func (l LocalFS) Delete(path string) error {
	return os.RemoveAll(path)
}

func (l LocalFS) Exist(path string) bool {
	return Exist(path)
}

func (l LocalFS) WriteFile(path string, data []byte) error {
	return ioutil.WriteFile(path, data, os.ModePerm)
}

func (l LocalFS) ReadFile(path string) ([]byte, error) {
	return ioutil.ReadFile(path)
}
