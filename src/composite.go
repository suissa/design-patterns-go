package main

type Component interface {
    Search(string) bool
}

type File struct {
    name string
}

func (f *File) Search(keyword string) bool {
    return f.name == keyword
}

type Folder struct {
    components []Component
}

func (f *Folder) Add(c Component) {
    f.components = append(f.components, c)
}

func (f *Folder) Search(keyword string) bool {
    for _, c := range f.components {
        if c.Search(keyword) {
            return true
        }
    }
    return false
}