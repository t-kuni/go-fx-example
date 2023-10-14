package services

type IDummyService interface {
	Hello()
}

type DummyServiceImplA struct {
}

func NewDummyServiceImplA() IDummyService {
	return &DummyServiceImplA{}
}

func (t *DummyServiceImplA) Hello() {
	println("Hello from DummyServiceImplA")
}

type DummyServiceImplB struct {
}

func NewDummyServiceImplB() IDummyService {
	return &DummyServiceImplB{}
}

func (t *DummyServiceImplB) Hello() {
	println("Hello from DummyServiceImplB")
}
