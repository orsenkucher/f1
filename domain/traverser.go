package domain

import (
	"io/ioutil"
	"path/filepath"
)

type Traverser struct {
	Root string
}

func NewTraverser(root string) *Traverser {
	return &Traverser{Root: root}
}

type Grouper interface {
	Group(name string, items []string) Group
}

func (t *Traverser) Traverse(g Grouper) (Groups, error) {
	files, err := ioutil.ReadDir(t.Root)
	if err != nil {
		return nil, err
	}

	var res Groups
	for _, f := range files {
		if !f.IsDir() {
			continue
		}

		name := f.Name()
		items, err := t.items(name)
		if err != nil {
			return nil, err
		}

		res = append(res, g.Group(name, items))
	}

	return res, nil
}

func (t *Traverser) items(dir string) ([]string, error) {
	files, err := ioutil.ReadDir(filepath.Join(t.Root, dir))
	if err != nil {
		return nil, err
	}

	var res []string
	for _, f := range files {
		if f.IsDir() {
			continue
		}

		name := f.Name()
		if filepath.Ext(name) != ".dat" {
			continue
		}

		res = append(res, name)
	}

	return res, nil
}

type GroupFn func(string, []string) Group

func (f GroupFn) Group(name string, items []string) Group {
	return f(name, items)
}

// TODO: kek
func Kek(name string, items []string) Group {
	return Group{}
}
