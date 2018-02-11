package optional

import (
	"errors"
	"testing"
)

type point struct {
	x, y int
}

type complexPoint struct {
	time int
	p    point
}

func TestOfWithNilValue(t *testing.T) {
	_, err := Of(nil)

	if err == nil {
		t.FailNow()
	}

	_, err = Of(5)

	if err != nil {
		t.FailNow()
	}
}

func TestNilPointer(t *testing.T) {
	var p *int

	_, err := Of(p)

	if err == nil {
		t.FailNow()
	}
}

func TestArray(t *testing.T) {
	array := [...]int{1, 2, 3}
	nilArray := [3]int{}

	_, err := Of(array)

	if err != nil {
		t.FailNow()
	}

	_, err = Of(nilArray)

	if err == nil {
		t.FailNow()
	}
}

func TestSlice(t *testing.T) {
	slice := []int{1, 2, 3}
	var nilSlice []int

	_, err := Of(slice)

	if err != nil {
		t.FailNow()
	}

	_, err = Of(nilSlice)

	if err == nil {
		t.FailNow()
	}

	_, err = Of([]int(nil))

	if err == nil {
		t.FailNow()
	}
}

func TestStruct(t *testing.T) {
	var p point
	var c complexPoint

	_, err := Of(p)

	if err != nil {
		t.FailNow()
	}

	_, err = Of(c)

	if err != nil {
		t.FailNow()
	}

	_, err = Of(point{3, 2})

	if err != nil {
		t.FailNow()
	}

	_, err = Of(complexPoint{0, point{}})

	if err != nil {
		t.FailNow()
	}
}

func TestFunction(t *testing.T) {
	var f func()

	_, err := Of(f)

	if err == nil {
		t.FailNow()
	}
}

func TestIsPresent(t *testing.T) {
	a, _ := Of(5)
	b := OfEmpty()

	if !a.IsPresent() {
		t.FailNow()
	}

	if b.IsPresent() {
		t.FailNow()
	}
}

func TestGet(t *testing.T) {
	a, _ := Of(5)
	b := OfEmpty()

	v, err := a.Get()

	if v != 5 || err != nil {
		t.FailNow()
	}

	v, err = b.Get()

	if v != nil || err == nil {
		t.FailNow()
	}
}

func TestIfPresentError(t *testing.T) {
	a, _ := Of(5)
	b := OfEmpty()

	if a.IfPresentError(errors.New("")) == nil {
		t.FailNow()
	}

	if b.IfPresentError(errors.New("")) != nil {
		t.FailNow()
	}
}

func TestIfNotPresentError(t *testing.T) {
	a, _ := Of(5)
	b := OfEmpty()

	if a.IfPresentError(errors.New("")) == nil {
		t.FailNow()
	}

	if b.IfPresentError(errors.New("")) != nil {
		t.FailNow()
	}
}

func TestIfPresent(t *testing.T) {
	a := OfEmpty()

	a.IfPresent(fError1)
	a.IfPresentPanic("test fails")
}

func TestIfNotPresent(t *testing.T) {
	a, _ := Of(5)

	a.IfNotPresent(fError2)
	a.IfNotPresentPanic("test fails")
}

func fError1(i interface{}) {
	panic("test fails")
}

func fError2() {
	panic("test fails")
}
