package main

import (
	"errors"
	"testing"
)

func TestIsPresent(t *testing.T) {
	a := Of(5)
	b := OfEmpty()

	if !a.IsPresent() {
		t.FailNow()
	}

	if b.IsPresent() {
		t.FailNow()
	}
}

func TestGet(t *testing.T) {
	a := Of(5)
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
	a := Of(5)
	b := OfEmpty()

	if a.IfPresentError(errors.New("")) != nil {
		t.FailNow()
	}

	if b.IfPresentError(errors.New("")) == nil {
		t.FailNow()
	}
}

func TestIfNotPresentError(t *testing.T) {
	a := Of(5)
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
	a := Of(5)

	a.IfNotPresent(fError2)
	a.IfNotPresentPanic("test fails")
}

func fError1(i interface{}) {
	panic("test fails")
}

func fError2() {
	panic("test fails")
}
