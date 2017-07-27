package sorting

import (
	"sort"
)

type SortedMap struct {
    Original map[string]int
    Keys []string
}

func (sm *SortedMap) Len() int {
    return len(sm.Original)
}

func (sm *SortedMap) Less(i, j int) bool {
    return sm.Original[sm.Keys[i]] > sm.Original[sm.Keys[j]]
}

func (sm *SortedMap) Swap(i, j int) {
    sm.Keys[i], sm.Keys[j] = sm.Keys[j], sm.Keys[i]
}

func SortedKeys(m map[string]int) *SortedMap {
    sm := new(SortedMap)
    sm.Original = m
    sm.Keys = make([]string, len(m))
    i := 0
    for key, _ := range m {
        sm.Keys[i] = key
        i++
    }
    sort.Sort(sm)
    return sm
}