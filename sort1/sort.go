package sort1

import (
	"sort"
)

type SortedMap struct {
    M map[string]int
    S []string
}

func (sm *SortedMap) Len() int {
    return len(sm.M)
}

func (sm *SortedMap) Less(i, j int) bool {
    return sm.M[sm.S[i]] > sm.M[sm.S[j]]
}

func (sm *SortedMap) Swap(i, j int) {
    sm.S[i], sm.S[j] = sm.S[j], sm.S[i]
}

func SortedKeys(m map[string]int) *SortedMap {
    sm := new(SortedMap)
    sm.M = m
    sm.S = make([]string, len(m))
    i := 0
    for key, _ := range m {
        sm.S[i] = key
        i++
    }
    sort.Sort(sm)
    return sm
}