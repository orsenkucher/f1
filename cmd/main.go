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
	groups, err := traverser.Traverse(domain.GroupFn(domain.Kek))
	if err != nil {
		return err
	}
	log.Println(groups)
	return nil
}
