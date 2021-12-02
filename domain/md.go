package domain

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

type Markdown struct {
	meta     Meta
	filename string
}

func NewMarkdown(filename string, meta Meta) *Markdown {
	return &Markdown{
		meta:     meta,
		filename: filename,
	}
}

func (md *Markdown) Write(groups []string, nucs map[string]Nucleus) error {
	sort.Strings(groups)
	var s string
	for _, group := range groups {
		n := nucs[group]
		id := ID{Number: n.Number, Mass: n.Mass}
		s += fmt.Sprintf(`$$\ce{^{%v}_{%v}%s}$$`, n.Mass, n.Number, md.meta[id].Element)
		group = strings.ReplaceAll(group, "\\", "-")
		s += fmt.Sprintf("![[plot/%s.png]]\n", group)
	}
	if err := ioutil.WriteFile(md.filename, []byte(s), 0644); err != nil {
		return err
	}
	return nil
}
