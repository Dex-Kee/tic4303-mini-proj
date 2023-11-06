package util

import "github.com/biter777/countries"

func ListCountryCode() []string {
	all := countries.All()
	codes := make([]string, len(all))
	for i, v := range all {
		codes[i] = v.Alpha2()
	}
	return codes
}

func CountryMap() map[string]string {
	all := countries.All()
	m := make(map[string]string)
	for _, v := range all {
		m[v.String()] = v.Alpha2()
	}
	return m
}
