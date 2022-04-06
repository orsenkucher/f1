package domain

import (
	"fmt"
	"strconv"
	"strings"
)

type Deform struct {
	Nucleus
	Beta2ef float64
}

type DeformMap map[Nucleus]Deform

func CreateDeformMap(path string) (DeformMap, error) {
	res := make(DeformMap)
	err := readLines(path, func(line string) error {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "#") || line == "" {
			return nil
		}
		fields := strings.Fields(line)
		mass, err := strconv.Atoi(fields[1])
		if err != nil {
			return err
		}
		number, err := strconv.Atoi(fields[0])
		if err != nil {
			return err
		}
		beta2, err := strconv.ParseFloat(fields[3], 64)
		if err != nil {
			return err
		}
		nuc := Nucleus{Number: number, Mass: mass}
		res[nuc] = Deform{
			Nucleus: nuc,
			Beta2ef: beta2,
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (d DeformMap) Print() {
	for k, v := range d {
		fmt.Printf("%+v\n", k)
		fmt.Println(v.Beta2ef)
	}
}
