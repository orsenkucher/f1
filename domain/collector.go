package domain

import "log"

type Collector struct {
	allowList   AllowList
	Experiments map[Nucleus][]Result
}

func NewCollector(allowList AllowList) *Collector {
	return &Collector{
		allowList:   allowList,
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
		if !c.allowList.Has(name) {
			log.Printf("skipping: %v %v\n", name.Number, name.Mass)
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
