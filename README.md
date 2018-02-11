[![Go Report Card](https://goreportcard.com/badge/gojp/goreportcard)](https://goreportcard.com/report/gojp/goreportcard)

Golang implements the zero value principle. Therefore, it can be tricky to differentiate empty from nil values.

goptional is a lightweight library to deal with this issue.

# Examples

```go
// Create an optional with an initial value
opt := optional.Of(5)

// Create an empty optional
opt := optional.OfEmpty()

// Get its value and a eventual error if the optional is empty
v, err := opt.Get()

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