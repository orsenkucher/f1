package domain

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type Drain struct {
	Root string
}

func NewDrain(root string) *Drain {
	return &Drain{
		Root: root,
	}
}

func (d *Drain) Drain(collector *Collector) error {
	if _, err := os.Stat(d.Root); os.IsNotExist(err) {
		err := os.Mkdir(d.Root, os.ModePerm)
		if err != nil {
			return err
		}
	}

	log.Println(len(collector.Experiments))
	for k, v := range collector.Experiments {
		path := filepath.Join(d.Root, k.String())
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			return err
		}

		for _, e := range v {
			path := filepath.Join(path, e.Name.String())
			bytes := []byte(e.Content)
			err = ioutil.WriteFile(path, bytes, 0644)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
