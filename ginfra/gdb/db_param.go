package gdb

import (
	"github.com/SyaibanAhmadRamadhan/gocatch/garray"
)

type OrderBy struct {
	Column      string
	IsAscending bool
}

type OrderByParams []OrderBy

func (p OrderByParams) FilterDifferent(refColumns []string) (res OrderByParams) {
	for _, v := range p {
		if garray.Contains(refColumns, v.Column) {
			res = append(res, v)
		}
	}

	return
}