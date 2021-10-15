package domain

import "log"

type Collector struct {
	Experiments map[Nucleus][]Result
}

func NewCollector() *Collector {
	return &Collector{
		Experiments: make(map[Nucleus][]Result),
	}
}

func (c *Collector) Collect(group Resource, items Items) {
	log.Println(group, len(items))
	for _, i := range items {
		name, err := NewName(i.Resource.Name)
		if err != nil {
			log.Println(err)
			continue
		}
		// log.Println(name)
		nuc := Nucleus{Number: name.Number, Mass: name.Mass}
		c.Experiments[nuc] = append(c.Experiments[nuc], Result{
			Name:    name,
			Content: i.Content,
		})
	}
}
