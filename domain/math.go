package domain

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"path/filepath"
)

type Mather struct {
	Dir      string
	paramMap GDRMap
	nucleus  map[string]Nucleus
}

func NewMather(
	dir string,
	paramMap GDRMap,
	nucleus map[string]Nucleus,
) *Mather {
	return &Mather{
		Dir:      dir,
		paramMap: paramMap,
		nucleus:  nucleus,
	}
}

func (m *Mather) Math(groups []string) error {
	// err := m.dir()
	// if err != nil {
	// 	return err
	// }
	// NOTE: Use .dat file as data source,
	// but strip the filename
	for _, g := range groups {
		work := filepath.Dir(g + ".dat")
		err := m.math(work)
		if err != nil {
			return err
		}
	}
	return nil
}

func (m *Mather) math(g string) error {
	fmt.Println("mathin' group:", g)
	// copy key files
	cmd := exec.Command(
		"cp",
		"ma/KEYDATA_ma.DAT",
		"ma/KEYMODEL_ma.DAT",
		"ma/Inputpar_ma.dat",
		g,
	)
	out, err := cmd.CombinedOutput()
	fmt.Println(string(out))
	if err != nil {
		return err
	}

	// group := strings.TrimSuffix(g, filepath.Ext(g))
	group := g + "\\group"
	nuc := m.nucleus[group]
	if p, ok := m.paramMap[nuc]; ok {
		fmt.Println("found params for", group)
		bytes := []byte(fmt.Sprintf("%v\n%v\n%v\n%v\n%v\n%v\n%v", p.E1, p.Z1, p.G1, p.E2, p.Z2, p.G2, 1.0))
		err = ioutil.WriteFile("ma/Inputpar_ma.dat", bytes, 0644)
		if err != nil {
			return err
		}
	} else {
		fmt.Println("no params for", group)
	}

	// wolframscript -script FitRSF.wl
	cmd = exec.Command("wolframscript", "-script ../../ma/FitRSF.wl")
	cmd.Dir = g
	out, err = cmd.Output()
	// fmt.Println(string(out))
	fmt.Println("calculated", g, len(out))
	return err
}

// func (m *Mather) dir() error {
// 	return dir(m.Dir)
// }
