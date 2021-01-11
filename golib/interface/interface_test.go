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

type Shape interface {
	Lenght() uint64
	Width() uint64
}

type Rect struct {

}

func (rect *Rect) Lenght() uint64 {
	return 0
}

func (rect *Rect) Width() uint64 {
	return 0
}

func TestInterfaceFuncCall(t *testing.T) {
	{
		var shape Shape
		shape = &Rect{}
		fmt.Printf("xxx: %d, %d\n", shape.Lenght(), shape.Width())
	}
	/*
	{
		var shape Shape
		shape = Rect{}
		fmt.Printf("xxx: %d, %d\n", shape.Lenght(), shape.Width())
	}
	*/
}
