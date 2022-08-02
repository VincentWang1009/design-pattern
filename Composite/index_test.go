package composite

import (
	"testing"
)

func TestInit(t *testing.T) {
	file1 := &File{name: "file1"}
	file2 := &File{name: "file2"}

	file1.Search("aaa")

	folder1 := &Folder{name: "folder1"}
	folder1.Add(file1)
	folder1.Add(file2)
	folder1.Search("bbb")
}
