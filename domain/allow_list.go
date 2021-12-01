package domain

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type AllowList []Allow
type Allow struct {
	Number int // Z
	Mass   int // A
}

func ParseAllowList(path string) (AllowList, error) {
	var res AllowList
	err := readLines(path, func(line string) error {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "#") || line == "" {
			return nil
		}
		fields := strings.Fields(line)
		number, err := strconv.Atoi(fields[0])
		if err != nil {
			return err
		}
		mass, err := strconv.Atoi(fields[1])
		if err != nil {
			return err
		}
		res = append(res, Allow{
			Number: number,
			Mass:   mass,
		})
		return nil
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func readLines(path string, f func(line string) error) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		err := f(line)
		if err != nil {
			return err
		}
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

func (al AllowList) Has(name Name) bool {
	for _, a := range al {
		if a.Number == name.Number && a.Mass == name.Mass {
			return true
		}
	}
	return false
}
