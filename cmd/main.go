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
	traverser.Traverse(func(group domain.Resource, items domain.Items) {
		log.Println(group, len(items))
		for _, i := range items {
			name, err := domain.NewName(i.Res.Name)
			if err != nil {
				log.Println(err)
				continue
			}
			log.Println(name)
		}
	})
	return nil
}
