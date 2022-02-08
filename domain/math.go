package domain

import (
	"fmt"
	"os/exec"
	"path/filepath"
)

type Mather struct {
	Dir string
}

func NewMather(dir string) *Mather {
	return &Mather{
		Dir: dir,
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

	// wolframscript -script FitRSF.wl
	cmd = exec.Command("wolframscript", "-script ../../ma/FitRSF.wl")
	cmd.Dir = g
	out, err = cmd.Output()
	fmt.Println(string(out))
	return err
}

// func (m *Mather) dir() error {
// 	return dir(m.Dir)
// }
