package composite

import "fmt"

type Component interface {
	Search(string)
}

type Folder struct {
	Components []Component
	Name       string
}

func (f *Folder) Search(keyword string) {
	fmt.Printf("recursive search reached keyword, folder: %s, %s\n", keyword, f.Name)
	for _, component := range f.Components {
		component.Search(keyword)
	}
}

func (f *Folder) Add(c Component) {
	f.Components = append(f.Components, c)
}

type File struct {
	Name string
}

func (f *File) Search(keyword string) {
	fmt.Printf("searching for keyword %s in file %s\n", keyword, f.Name)
}

func (f *File) GetName() string {
	return f.Name
}
