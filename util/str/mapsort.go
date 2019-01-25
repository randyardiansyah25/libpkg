package strutils

import (
	"sort"
)

type MapSort struct{}

func (s *MapSort) SortString(m map[string]interface{}) map[string]interface{} {
	nm := make(map[string]interface{}, 0)
	keys := make([]string, 0)
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		nm[k] = m[k]
	}
	return nm
}
