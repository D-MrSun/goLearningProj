package classpath

import (
	"io/ioutil"
	"path/filepath"
)

//目录形式的类路径
type DirEntry struct {
	absDir string
}

//定义newDirEntry形式的函数来模仿Java中的构造函数
func newDirEntry(path string) *DirEntry {
	absDir, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &DirEntry{absDir}
}

func (self *DirEntry) readClass(className string) ([]byte, Entry, error) {
	fileName := filepath.Join(self.absDir, className)
	data, err := ioutil.ReadFile(fileName)
	return data, self, err
}

func (self *DirEntry) String() string {
	return self.absDir
}
