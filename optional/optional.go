/*
	A container for optional values.
*/

package optional

import "errors"

type optional struct {
	v     interface{}
	empty bool
}

// Create an optional with an initial value
func Of(in interface{}) optional {
	o := optional{in, false}

	return o
}

// Create an empty optional
func OfEmpty() optional {
	o := optional{nil, true}

	return o
}

// Get the optional value with an eventual error if the optional is empty
func (v *optional) Get() (interface{}, error) {
	if v.IsPresent() {
		return v.v, nil
	}

	return nil, errors.New("optional is empty")
}

// Get the optional value regardless of its emptiness
func (v *optional) GetValue() interface{} {
	return v.v
}

// Test whether the optional is not empty
func (v *optional) IsPresent() bool {
	return !v.empty
}

// If the optional is not empty, execute a function taking the optional value as the input parameter
func (v *optional) IfPresent(f func(interface{})) {
	if v.IsPresent() {
		f(v.v)
	}
}

// If the optional is empty, execute a function
func (v *optional) IfNotPresent(f func()) {
	if !v.IsPresent() {
		f()
	}
}

// If the optional is not empty, returns an error. Otherwise, returns a nil value
func (v *optional) IfPresentError(e error) error {
	if v.IsPresent() {
		return e
	}

	return nil
}

// If the optional is empty, returns an error. Otherwise, returns a nil value
func (v *optional) IfNotPresentError(e error) error {
	if !v.IsPresent() {
		return e
	}

	return nil
}

// If the optional is not empty, panic
func (v *optional) IfPresentPanic(i interface{}) {
	if v.IsPresent() {
		panic(i)
	}
}

// If the optional is empty, panic
func (v *optional) IfNotPresentPanic(i interface{}) {
	if !v.IsPresent() {
		panic(i)
	}
}
