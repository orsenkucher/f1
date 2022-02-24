package domain

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"path/filepath"
)

type Mather struct {
	Dir     string
	gdrMap  GDRMap
	pdrMap  PDRMap
	nucleus map[string]Nucleus
}

func NewMather(
	dir string,
	gdrMap GDRMap,
	pdrMap PDRMap,
	nucleus map[string]Nucleus,
) *Mather {
	return &Mather{
		Dir:     dir,
		gdrMap:  gdrMap,
		pdrMap:  pdrMap,
		nucleus: nucleus,
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
	if gp, ok := m.gdrMap[nuc]; ok {
		var bytes []byte
		if pp, ok := m.pdrMap[nuc]; ok {
			fmt.Println("found pdr and gdr params for", group)
			fmt.Println(gp, pp)
			bytes = []byte(fmt.Sprintf("%v\n%v\n%v\n%v\n%v\n%v\n%v", gp.E, gp.Z, gp.G, pp.E, pp.Z, pp.G, 1.0))
		} else {
			fmt.Println("found only gdr for", group)
			bytes = []byte(fmt.Sprintf("%v\n%v\n%v\n%v\n%v\n%v\n%v", gp.E, gp.Z, gp.G, 1.0, 1.0, 1.0, 1.0))
		}
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
