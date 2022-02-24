package domain

import (
	"fmt"
	"strconv"
	"strings"
)

type PDRParams struct {
	Nucleus
	E, Z, G float64
}

type PDRMap map[Nucleus]PDRParams

func CreatePDRMap(path string) (PDRMap, error) {
	res := make(PDRMap)
	err := readLines(path, func(line string) error {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "#") || line == "" {
			return nil
		}
		fields := strings.Fields(line)
		if len(fields) <= 5 {
			fmt.Println("Skipping PDR line:")
			fmt.Println("\t", fields)
			return nil
		}

		number, err := strconv.Atoi(fields[0])
		if err != nil {
			return err
		}
		mass, err := strconv.Atoi(fields[1])
		if err != nil {
			return err
		}
		e1, err := strconv.ParseFloat(fields[3], 64)
		if err != nil {
			return err
		}
		g1, err := strconv.ParseFloat(fields[4], 64)
		if err != nil {
			g1 = 1.0
			// return err
		}
		z1 := 1.0
		// z1, err := strconv.ParseFloat(fields[4], 64)
		// if err != nil {
		// 	return err
		// }
		nuc := Nucleus{Number: number, Mass: mass}
		res[nuc] = PDRParams{
			Nucleus: nuc,
			E:       e1,
			Z:       z1,
			G:       g1,
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (pdr PDRMap) Print() {
	for k, v := range pdr {
		fmt.Printf("%+v\n", k)
		fmt.Println(v.E, v.G, v.Z)
	}
}
