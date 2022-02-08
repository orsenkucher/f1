package domain

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"sync"
)

type Plotter struct {
	Dir string
}

func NewPlotter(dir string) *Plotter {
	return &Plotter{
		Dir: dir,
	}
}

func (p *Plotter) Plot(groups []string) error {
	err := p.dir()
	if err != nil {
		return err
	}

	var wg sync.WaitGroup
	wg.Add(len(groups))
	work := make(chan string)
	for i := 0; i < 10; i++ {
		go func() {
			for {
				w := <-work
				err := p.plot(w)
				if err != nil {
					log.Println(err)
				}
				wg.Done()
			}
		}()
	}

	// NOTE: Use .json file as data source
	for _, g := range groups {
		work <- g + ".json"
	}
	wg.Wait()
	return nil
}

func (p *Plotter) plot(g string) error {
	fmt.Println("plotting group:", g)
	cmd := exec.Command("python", "script/main.py", g, p.Dir)
	out, err := cmd.Output()
	fmt.Println(string(out))
	return err
}

func (p *Plotter) dir() error {
	return dir(p.Dir)
}

func dir(name string) error {
	if _, err := os.Stat(name); !os.IsNotExist(err) {
		err := os.RemoveAll(name)
		if err != nil {
			return err
		}
	}
	if _, err := os.Stat(name); os.IsNotExist(err) {
		err := os.Mkdir(name, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}
