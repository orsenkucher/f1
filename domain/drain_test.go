package domain

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v ./...
func TestSortRecords(t *testing.T) {
	rec := []Record{
		{E: "7.850"},
		{E: "8.430"},
		{E: "7.340"},
	}
	sortRecords(rec)
	fmt.Println(rec)
	assert.Equal(t, rec, []Record{
		{E: "7.340"},
		{E: "7.850"},
		{E: "8.430"},
	})
}
