package class_struct

import (
	"fmt"
	"testing"
)

type CCC struct {
	c string
}

func (c *CCC) print() {
	fmt.Println(c.c)
}

type DDD CCC

func (d *DDD) print1() {
	fmt.Println(d.c)
}

func TestFuncCall(t *testing.T) {
	d := &DDD{
		c: "ccc",
	}
	d.print1()
}

func TestSlice(t *testing.T) {
	cs := make([]*CCC, 0)
	cs = append(cs, &CCC{
		c: "ccc",
	})

	test := func(ds []*DDD) {
		for _, d := range ds {
			d.print1()
		}
	}
	test(cs)
}
