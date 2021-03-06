package domain

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
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
	deform     DeformMap
	drainers   []Drainer
}

func NewDrain(
	root, groupName string,
	meta Meta,
	deform DeformMap,
	drainers ...Drainer) *Drain {
	return &Drain{
		Root:      root,
		GroupName: groupName,
		meta:      meta,
		deform:    deform,
		Nucleus:   map[string]Nucleus{},
		drainers:  drainers,
	}
}

type Drainer interface {
	Drain(groupName string, data []Data) error
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

		groupName := filepath.Join(path, d.GroupName)

		for _, drainer := range d.drainers {
			err := drainer.Drain(groupName, data)
			if err != nil {
				return err
			}
		}

		// TODO: maybe ref
		// groupName += ".json"
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
	Deform        string
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

		// filter zero values
		e, err := strconv.ParseFloat(record.E, 64)
		if err != nil {
			return err
		}
		f, err := strconv.ParseFloat(record.F, 64)
		if err != nil {
			return err
		}
		if math.Min(e, f) == 0 {
			log.Println("skipping zero values:", r.Name.String())
			continue
		}

		// set errors to const value
		record.DFPlus = "0.5E-07"
		record.DFMinus = record.DFPlus

		records = append(records, record)
	}

	sortRecords(records)
	*data = append(*data, Data{
		Name:          name,
		Records:       records,
		NeutronEnergy: d.FindMeta(r.Name).NeutronEnergy,
		Deform:        fmt.Sprintf("%f", d.FindDeform(r.Name).Beta2ef),
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

func (d *Drain) FindDeform(name Name) Deform {
	nuc := Nucleus{Number: name.Number, Mass: name.Mass}
	deform, ok := d.deform[nuc]
	if !ok {
		log.Printf("deform not found for %v %v\n", nuc.Number, nuc.Mass)
	}
	return deform
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

type JsonDrainer struct{}

func NewJsonDrainer() *JsonDrainer {
	return &JsonDrainer{}
}

type DatDrainer struct{}

func NewDatDrainer() *DatDrainer {
	return &DatDrainer{}
}

func (d *JsonDrainer) Drain(groupName string, data []Data) error {
	groupName += ".json"
	js, err := json.Marshal(data)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(groupName, js, 0644)
	return err
}

func (d *DatDrainer) Drain(groupName string, data []Data) error {
	groupName += ".dat"
	dat := ""
	for _, group := range data {
		for _, rec := range group.Records {
			line := fmt.Sprintf("%v %v %v", rec.E, rec.F, rec.DFPlus)
			dat += line + "\n"
		}
	}
	bytes := []byte(dat)
	err := ioutil.WriteFile(groupName, bytes, 0644)
	return err
}
