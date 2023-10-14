//go:generate mockgen -source=$GOFILE -destination=${GOFILE}_mock.go -package=$GOPACKAGE

package services

type IDummyService interface {
	Hello() string
}

type DummyServiceImplA struct {
}

func NewDummyServiceImplA() IDummyService {
	return &DummyServiceImplA{}
}

func (t *DummyServiceImplA) Hello() string {
	return "Hello from DummyServiceImplA"
}

type DummyServiceImplB struct {
}

func NewDummyServiceImplB() IDummyService {
	return &DummyServiceImplB{}
}

func (t *DummyServiceImplB) Hello() string {
	return "Hello from DummyServiceImplB"
}
