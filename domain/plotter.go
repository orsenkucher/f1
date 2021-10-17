package domain

import (
	"fmt"
	"log"
	"os/exec"
	"sync"
)

type Plotter struct {
}

func NewPlotter() *Plotter {
	return &Plotter{}
}

func (p *Plotter) Plot(groups []string) error {
	var wg sync.WaitGroup
	wg.Add(len(groups))
	work := make(chan string)
	for i := 0; i < 10; i++ {
		go func() {
			for {
				w := <-work
				err := plot(w)
				if err != nil {
					log.Println(err)
				}
				wg.Done()
			}
		}()
	}

	for _, g := range groups {
		work <- g
	}
	wg.Wait()
	return nil
}

func plot(g string) error {
	fmt.Println("plotting group:", g)
	cmd := exec.Command("python", "script/main.py", g)
	out, err := cmd.Output()
	fmt.Println(string(out))
	return err
}
