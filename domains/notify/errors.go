package notify

type Errors interface{}

type errors struct{}

func newErrors() Errors {
	return &errors{}
}
