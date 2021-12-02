package domain

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

type Drain struct {
	Root       string
	GroupName  string
	GroupFiles []string
	Nucleus    map[string]Nucleus
	meta       Meta
}

func NewDrain(root, groupName string, meta Meta) *Drain {
	return &Drain{
		Root:      root,
		GroupName: groupName,
		meta:      meta,
		Nucleus:   map[string]Nucleus{},
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
			err = d.drainData(&data, e)
			if err != nil {
				return err
			}
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
		d.GroupFiles = append(d.GroupFiles, groupName)
		d.Nucleus[groupName] = k
	}

	log.Println("drained")
	return nil
}

type Data struct {
	Name          string
	Records       []Record
	NeutronEnergy string
	Element       Element
}

type Element struct {
	Symbol string
	Number int
	Mass   int
}

type Record struct {
	E       string
	DE      string
	F       string
	DFMinus string
	DFPlus  string
}

func (d *Drain) drainData(data *[]Data, r Result) error {
	name := fmt.Sprintf("F%s %s %s",
		r.Name.PSF,
		r.Name.Method,
		strings.TrimSpace(r.Name.NSR),
	)
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
		if len(fields) != 4 && len(fields) != 5 {
			// panic(r.Name.String())
			log.Println("malformed dat file:", r.Name.String())
		}

		record := Record{
			E:       fields[0],
			DE:      fields[1],
			F:       fields[2],
			DFMinus: fields[3],
		}

		// symmetric uncertainty
		if len(fields) == 4 {
			record.DFPlus = fields[3]
		}

		// asymmetric uncertainty
		if len(fields) == 5 {
			record.DFPlus = fields[4]
		}

		records = append(records, record)
	}

	sortRecords(records)
	*data = append(*data, Data{
		Name:          name,
		Records:       records,
		NeutronEnergy: d.FindMeta(r.Name).NeutronEnergy,
		Element: Element{
			Symbol: d.FindMeta(r.Name).Element,
			Number: r.Name.Number,
			Mass:   r.Name.Mass,
		},
	})
	return nil
}

func (d *Drain) FindMeta(name Name) MetaItem {
	id := ID{Number: name.Number, Mass: name.Mass}
	meta, ok := d.meta[id]
	if !ok {
		log.Printf("meta not found for %v %v\n", id.Number, id.Mass)
	}
	return meta
}

func sortRecords(records []Record) {
	sort.Slice(records, func(i, j int) bool {
		ri, rj := records[i], records[j]
		ei, err := strconv.ParseFloat(ri.E, 64)
		if err != nil {
			panic(err)
		}
		ej, err := strconv.ParseFloat(rj.E, 64)
		if err != nil {
			panic(err)
		}
		return ei < ej
	})
}

func lines(s string) (lines []string, err error) {
	scanner := bufio.NewScanner(strings.NewReader(s))
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err = scanner.Err()
	return
}
