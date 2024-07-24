package p2418

import "sort"

type peoples struct {
	names   *[]string
	heights *[]int
}

func (p peoples) check() bool {
	return p.names == nil || p.heights == nil || len(*p.names) != len(*p.heights)
}

func (p peoples) Len() int {
	if p.check() {
		panic("implement me")
	}
	return len(*p.names)
}

func (p peoples) Less(i, j int) bool {
	if p.check() {
		panic("implement me")
	}
	return (*p.heights)[i] > (*p.heights)[j]
}

func (p peoples) Swap(i, j int) {
	if p.check() {
		panic("implement me")
	}
	(*p.names)[i], (*p.names)[j] = (*p.names)[j], (*p.names)[i]
	(*p.heights)[i], (*p.heights)[j] = (*p.heights)[j], (*p.heights)[i]
}

func sortPeople(names []string, heights []int) []string {
	p := peoples{&names, &heights}
	sort.Sort(p)
	return names
}
