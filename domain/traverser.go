package domain

import (
	"io/fs"
	"io/ioutil"
	"path/filepath"
)

type Traverser struct {
	Root string
}

func NewTraverser(root string) *Traverser {
	return &Traverser{Root: root}
}

type TraverseFn func(group Resource, items Items)

// Traverser will provide group name, dat files name
// and their contents to callback.
func (t *Traverser) Traverse(f TraverseFn) error {
	var files files
	files, err := ioutil.ReadDir(t.Root)
	if err != nil {
		return err
	}

	res := files.filterDirs().asRes(t.Root)

	for _, r := range res {
		items, err := r.items()
		if err != nil {
			return err
		}
		f(r, items)
	}
	return nil
}

type files []fs.FileInfo

type Resources []Resource
type Resource struct {
	Name string
	Path string
}

func (ff files) filterDirs() files {
	var res files
	for _, f := range ff {
		if f.IsDir() {
			res = append(res, f)
		}
	}
	return res
}

func (ff files) filterDats() files {
	var res files
	for _, f := range ff {
		if filepath.Ext(f.Name()) == ".dat" {
			res = append(res, f)
		}
	}
	return res
}

func (ff files) asRes(root string) Resources {
	var res Resources
	for _, f := range ff {
		name := f.Name()
		res = append(res, Resource{
			Name: name,
			Path: filepath.Join(root, name),
		})
	}
	return res
}

type Items []Item
type Item struct {
	Res Resource
	Val string
}

func (i Item) String() string {
	return i.Res.Path
}

func (r Resource) items() (Items, error) {
	var files files
	files, err := ioutil.ReadDir(r.Path)
	if err != nil {
		return nil, err
	}

	res := files.filterDats().asRes(r.Path)

	var items Items
	for _, r := range res {
		bytes, err := ioutil.ReadFile(r.Path)
		if err != nil {
			return nil, err
		}
		items = append(items, Item{
			Res: r,
			Val: string(bytes),
		})
	}
	return items, nil
}
