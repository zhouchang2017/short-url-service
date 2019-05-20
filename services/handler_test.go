package services

import (
	"testing"
)

func TestRegUrl(t *testing.T) {
	urls := map[string]bool{
		"https://www.wewee.com/":                     true,
		"https://tailwindcss.com/docs/installation/": true,
		"fdakfasf":        false,
		"21312fda":        false,
		"web://fdasfasf": true,
	}

	for url, res := range urls {
		if r := isUrl(url); r != res {
			t.Errorf("当前%s\t,期望 %v,实际 %v\n", url, res, r)
		}
	}
}


