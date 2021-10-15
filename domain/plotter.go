package domain

import (
	"fmt"
	"os/exec"
)

type Plotter struct {
}

func NewPlotter() *Plotter {
	return &Plotter{}
}

func (p *Plotter) Plot(groups []string) error {
	for _, g := range groups {
		fmt.Println("plotting group:", g)
		cmd := exec.Command("python", "script/main.py", g)
		out, err := cmd.Output()
		fmt.Println(string(out))
		if err != nil {
			return err
		}
	}
	return nil
}
