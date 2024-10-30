package main

import "github.com/vietpham102301/go-design-patterns-practice/structural_patterns/composite"

func main() {
	file1 := &composite.File{"File1"}
	file2 := &composite.File{"File2"}
	file3 := &composite.File{"File3"}

	folder1 := &composite.Folder{
		Name: "Folder 1",
	}

	folder1.Add(file1)
	folder2 := &composite.Folder{
		Name: "Folder2",
	}
	folder2.Add(file2)
	folder2.Add(file3)
	folder2.Add(folder1)

	folder2.Search("rose")
}
