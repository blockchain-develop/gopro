package _interface

import (
	"fmt"
	"testing"
)

type Instance struct {
	n           int
	v           interface{}
}

type InstaceState struct {
	instance     Instance
	state        int
}

func work1(data interface{}) {
	instance := data.(Instance)
	fmt.Printf("instance n: %d\n", instance.n)
}

func TestInterface(t *testing.T) {
	instance := Instance{
		n : 100,
		v: nil,
	}
	work1(instance)

	instanceState := &InstaceState{
		instance: instance,
		state: 1,
	}
	work1(instanceState.instance)
}
