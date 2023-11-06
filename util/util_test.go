package util

import (
	"fmt"
	"testing"

	"github.com/samber/lo"
)

func TestCountryCode(t *testing.T) {
	code := ListCountryCode()
	lo.ForEach(code, func(v string, i int) {
		fmt.Println(v)
	})
}

func TestCountryMap(t *testing.T) {
	m := CountryMap()
	for k, v := range m {
		fmt.Printf("%s: %s\n", k, v)
	}
}
