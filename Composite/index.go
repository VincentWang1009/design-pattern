package composite

import "fmt"

type Component interface {
	Search(string)
}

type File struct {
	name string
}

func (f *File) Search(s string) {
	fmt.Printf("Search for keyword %s in file %s\n", s, f.name)
}

func (f *File) GetName() string {
	return f.name
}

type Folder struct {
	components []Component
	name       string
}

func (f *Folder) Search(key string) {
	fmt.Printf("searching fro keyword %s in folder %s\n", key, f.name)
	for _, v := range f.components {
		v.Search(key)
	}
}

func (f *Folder) Add(c Component) {
	f.components = append(f.components, c)
}
