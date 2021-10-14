package domain

import "log"

type Collector struct {
	Table Table
}

func NewCollector() *Collector {
	return &Collector{
		Table: Table{
			Groups: make(map[string]Group),
		},
	}
}

func (c *Collector) Collect(group Resource, items Items) {
	log.Println(group, len(items))
	for _, i := range items {
		name, err := NewName(i.Res.Name)
		if err != nil {
			log.Println(err)
			continue
		}
		log.Println(name)
		c.Table.Groups
	}
}
