[![Go Report Card](https://goreportcard.com/badge/gojp/goreportcard)](https://goreportcard.com/report/gojp/goreportcard)

goptional is a lightweight library to provide a container for optional values in golang.

## Initialization

```go
// Create an optional with an initial value
opt, _ := optional.Of(5)

// An optional created with a nil value returns an error
_, err := optional.Of(nil)

// Create an empty optional
opt := optional.OfEmpty()
```

## Features


```go
// Get its value and an eventual error if the optional is empty
v, err := opt.Get()

// Get the optional value regardless of its emptiness
v := opt.GetValue()

// Test if the optional is not empty
opt.IsPresent()

// Execute a function if the optional is not empty
opt.IfPresent(f)

// Execute a function if the optional is empty
opt.IfNotPresent(f)

// Returns an error if the optional is not empty (otherwise returns a nil value)
err := opt.IfPresentError(errors.New("present")

// Returns an error if the optional is empty (otherwise returns a nil value)
err := opt.IfNotPresentError(errors.New("not present")

// Panic if the optional is not empty
opt.IfPresentPanic("present")

// Panic if the optional is empty
opt.IfNotPresentPanic("not present")
```