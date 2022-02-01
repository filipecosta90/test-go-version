package test_go_version

// MyOptions struct.
type MyOptions struct {
	v           int
}

// NewMyOptions returns a new MyOptions struct.
func NewMyOptions(v int) *MyOptions {
	return &MyOptions{v: v}
}

