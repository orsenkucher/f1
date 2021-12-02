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
	allowList, err := domain.ParseAllowList("allowlist.txt")
	if err != nil {
		return err
	}
	collector := domain.NewCollector(allowList)
	if err := traverser.Traverse(collector.Collect); err != nil {
		return err
	}
	meta, err := domain.ParseMeta("meta.txt")
	if err != nil {
		return err
	}
	drain := domain.NewDrain("groups", "group.json", meta)
	err = drain.Drain(collector)
	if err != nil {
		return err
	}
	md := domain.NewMarkdown("samples.md", meta)
	err = md.Write(drain.GroupFiles, drain.Nucleus)
	if err != nil {
		return err
	}
	plotter := domain.NewPlotter("plot")
	err = plotter.Plot(drain.GroupFiles)
	if err != nil {
		return err
	}
	return nil
}
