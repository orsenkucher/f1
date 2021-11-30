package domain

import "fmt"

type Nucleus struct {
	Number int
	Mass   int
}

type Result struct {
	Name    Name
	Content string
}

func (n Nucleus) String() string {
	return fmt.Sprintf("group_z%03d_a%03d", n.Number, n.Mass)
}
