package domain_test

import (
	"regexp"
	"testing"

	"github.com/orsenkucher/f1/domain"
	"github.com/stretchr/testify/assert"
)

// go test -v ./...
func TestRegex(t *testing.T) {
	re := regexp.MustCompile(domain.Regex)
	group := domain.ReGroup(re, "fe1_exp_012_024_photoabs_1966Dol.dat")
	assert.Equal(t, group["NSR"], "1966Dol.dat")
	t.Log(group)

	group = domain.ReGroup(re, "fe1_exp_012_024_photoabs_1966Dol_abc_xyz")
	assert.Equal(t, group["NSR"], "1966Dol_abc_xyz")
	t.Log(group)

	group = domain.ReGroup(re, "fe1_exp_012_024_photoabs")
	assert.Equal(t, group["NSR"], "")
	t.Log(group)

	match := re.MatchString("fe1_exp_012_024_photoabs_1966Dol_abc_xyz")
	t.Log(match)
	assert.True(t, match)
}

func TestName(t *testing.T) {
	name := domain.NewNameMust("fe1_exp_012_024_photoabs_1966Dol.dat")
	assert.Equal(t, name, domain.Name{
		PSF:    "e1",
		Data:   "exp",
		Number: 12,
		Mass:   24,
		Method: "photoabs",
		NSR:    "1966Dol",
		Ext:    ".dat",
	})
}

func TestNameString(t *testing.T) {
	name := "fe1_exp_012_024_photoabs_1966Dol.dat"
	assert.Equal(t, name, domain.NewNameMust(name).String())
	name = "fe1_exp_012_024_photoabs.dat"
	assert.Equal(t, name, domain.NewNameMust(name).String())
}
