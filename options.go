package test_go_version

// MyOptions struct updated to V2.
type MyOptions struct {
	v           int
	v2			int
}

// NewMyOptions returns a new MyOptions updated to V2 struct with default value.
func NewMyOptionsDefault() *MyOptions {
	return &MyOptions{}
}


// NewMyOptions returns a new MyOptions updated to V2 struct.
func NewMyOptions(v int) *MyOptions {
	return &MyOptions{v: v}
}

