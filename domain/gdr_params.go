package domain

import (
	"fmt"
	"strconv"
	"strings"
)

type GDRParams struct {
	Nucleus
	E1, Z1, G1, E2, Z2, G2 float64
}

type GDRMap map[Nucleus]GDRParams

func CreateGDRMap(path string) (GDRMap, error) {
	res := make(GDRMap)
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
		e1, err := strconv.ParseFloat(fields[2], 64)
		if err != nil {
			return err
		}
		g1, err := strconv.ParseFloat(fields[3], 64)
		if err != nil {
			return err
		}
		z1, err := strconv.ParseFloat(fields[4], 64)
		if err != nil {
			return err
		}
		e2, err := strconv.ParseFloat(fields[5], 64)
		if err != nil {
			return err
		}
		g2, err := strconv.ParseFloat(fields[6], 64)
		if err != nil {
			return err
		}
		z2, err := strconv.ParseFloat(fields[7], 64)
		if err != nil {
			return err
		}
		nuc := Nucleus{Number: number, Mass: mass}
		res[nuc] = GDRParams{
			Nucleus: nuc,
			E1:      e1,
			Z1:      z1,
			G1:      g1,
			E2:      e2,
			Z2:      z2,
			G2:      g2,
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (gdr GDRMap) Print() {
	for k, v := range gdr {
		fmt.Printf("%+v\n", k)
		fmt.Println(v.E1, v.G1, v.Z1, v.E2, v.G2, v.Z2)
	}
}
