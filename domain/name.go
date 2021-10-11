package domain

import (
	"fmt"
	"log"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

// The data files naming convention is self-explanatory and includes:
// the type and multipolarity of the PSF XL={E1|E2|M1|1} (1 stands for E1+M1),
// if it is experimental or theoretical data, nuclide (Z,A),
// method used (NRF, OM, ARC/DRC, pg, pp, RM, photonuclear),
// NSR keynumber is added for photonuclear data.
// f{XL}_{exp|the}_Z_A_method[_NSRKeyNo].dat,
// e.g. fe1_exp_012_024_photoabs_1966Dol.dat, f1_exp_042_097_OM_3he_2.dat

type Name struct {
	PSF    string
	Number int // Z
	Mass   int // A
	Data   string
	Method string
	NSR    string
	Ext    string
}

const Regex = `^f(?P<PSF>.+)_(?P<Data>exp|the)_(?P<Z>\d{3})_(?P<A>\d{3})_(?P<Method>.+?)(_(?P<NSR>.+))*$`

func NewName(name string) Name {
	ext := filepath.Ext(name)
	name = strings.TrimSuffix(name, ext)
	re := regexp.MustCompile(Regex)
	group := ReGroup(re, name)
	number, err := strconv.Atoi(group["Z"])
	if err != nil {
		log.Fatalln("can't parse atomic number", err)
	}
	mass, err := strconv.Atoi(group["A"])
	if err != nil {
		log.Fatalln("can't parse mass number", err)
	}
	return Name{
		PSF:    group["PSF"],
		Data:   group["Data"],
		Number: number,
		Mass:   mass,
		Method: group["Method"],
		NSR:    group["NSR"],
		Ext:    ext,
	}
}

func ReGroup(re *regexp.Regexp, s string) map[string]string {
	match := re.FindStringSubmatch(s)
	res := make(map[string]string)
	for i, name := range re.SubexpNames() {
		if i != 0 && name != "" {
			res[name] = match[i]
		}
	}
	return res
}

func (n Name) String() string {
	name := fmt.Sprintf("f%s_%s_%03d_%03d_%s", n.PSF, n.Data, n.Number, n.Mass, n.Method)
	if n.NSR != "" {
		name += fmt.Sprintf("_%s", n.NSR)
	}
	return name + n.Ext
}
