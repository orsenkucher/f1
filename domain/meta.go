package domain

import (
	"strconv"
	"strings"
)

type Meta map[ID]MetaItem
type MetaItem struct {
	NeutronEnergy string
	Element       string
}

func ParseMeta(path string) (Meta, error) {
	res := make(Meta)
	err := readLines(path, func(line string) error {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "#") || line == "" {
			return nil
		}
		fields := strings.Fields(line)
		mass, err := strconv.Atoi(fields[0])
		if err != nil {
			return err
		}
		number, err := strconv.Atoi(fields[2])
		if err != nil {
			return err
		}
		element := fields[1]
		energy := fields[3]
		if strings.Contains(energy, "#") {
			return nil
		}
		if strings.Contains(energy, "*") {
			return nil
		}
		id := ID{Number: number, Mass: mass}
		res[id] = MetaItem{
			Element:       element,
			NeutronEnergy: energy,
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}
