package domain

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type Drain struct {
	Root      string
	GroupName string
}

func NewDrain(root, groupName string) *Drain {
	return &Drain{
		Root:      root,
		GroupName: groupName,
	}
}

func (d *Drain) Drain(collector *Collector) error {
	if _, err := os.Stat(d.Root); !os.IsNotExist(err) {
		err := os.RemoveAll(d.Root)
		if err != nil {
			return err
		}
	}
	if _, err := os.Stat(d.Root); os.IsNotExist(err) {
		err := os.Mkdir(d.Root, os.ModePerm)
		if err != nil {
			return err
		}
	}

	log.Println(len(collector.Experiments))
	for k, v := range collector.Experiments {
		path := filepath.Join(d.Root, k.String())
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			return err
		}

		var data []Data
		for _, e := range v {
			path := filepath.Join(path, e.Name.String())
			bytes := []byte(e.Content)
			err = ioutil.WriteFile(path, bytes, 0644)
			if err != nil {
				return err
			}
			drainData(&data, e)
		}

		js, err := json.Marshal(data)
		if err != nil {
			return err
		}

		groupName := filepath.Join(path, d.GroupName)
		err = ioutil.WriteFile(groupName, js, 0644)
		if err != nil {
			return err
		}
	}

	return nil
}

type Data struct {
	Name    string
	Records []Record
}

type Record struct {
	E  string
	DE string
	F  string
	DF string
}

func drainData(data *[]Data, r Result) error {
	name := fmt.Sprintf("%s %s", r.Name.Method, r.Name.NSR)
	lines, err := lines(r.Content)
	if err != nil {
		return err
	}

	var records []Record
	for _, l := range lines {
		l = strings.TrimSpace(l)
		if strings.HasPrefix(l, "#") || l == "" {
			continue
		}

		fields := strings.Fields(l)
		if len(fields) != 4 {
			// asymmetric uncertainty
			log.Println("malformed dat file:", r.Name.String())
			// panic(r.Name.String())
		}

		records = append(records, Record{
			E:  fields[0],
			DE: fields[1],
			F:  fields[2],
			DF: fields[3],
		})
	}

	*data = append(*data, Data{
		Name:    name,
		Records: records,
	})
	return nil
}

func lines(s string) (lines []string, err error) {
	scanner := bufio.NewScanner(strings.NewReader(s))
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err = scanner.Err()
	return
}
