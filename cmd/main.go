package main

import (
	"fmt"
	"log"

	"github.com/orsenkucher/f1/domain"
)

func main() {
	fmt.Println(">>> f1 <<<")
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}

func run() error {
	traverser := domain.NewTraverser("all_exp")
	collector := domain.NewCollector()
	traverser.Traverse(collector.Collect)
	drain := domain.NewDrain("groups", "group.json")
	err := drain.Drain(collector)
	if err != nil {
		return err
	}
	plotter := domain.NewPlotter()
	err = plotter.Plot(drain.GroupFiles)
	if err != nil {
		return err
	}
	return nil
}
