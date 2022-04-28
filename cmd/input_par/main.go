package main

import (
	"fmt"
	"log"

	"github.com/orsenkucher/f1/domain"
)

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}

func run() error {
	gdrs, err := domain.CreateGDRMap("gdr-params/gdr-params-smlo.dat")
	if err != nil {
		return err
	}
	pdrs, err := domain.CreatePDRMap("pdr-params/pdr_exp-draft.dat")
	if err != nil {
		return err
	}
	nuc := domain.Nucleus{
		Number: 50,
		Mass:   122,
	}
	gp := gdrs[nuc]
	pp := pdrs[nuc]
	res := fmt.Sprintf("%v\n%v\n%v\n%v\n%v\n%v\n%v", gp.E, pp.E, gp.G, pp.G, 1, gp.Z, pp.Z)
	// res := fmt.Sprintf("%v\n%v\n%v\n%v\n%v\n%v\n%v", gp.E, gp.Z, gp.G, pp.E, pp.Z, pp.G, 1.0)

	fmt.Printf("Printing input params for %+v\n", nuc)
	fmt.Println(res)

	return nil
}
