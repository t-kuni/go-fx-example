package a

type Test struct {
}

func NewTest() *Test {
	return &Test{}
}

func (t *Test) Hello() {
	println("Hello from a/b")
}
