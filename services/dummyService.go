//go:generate mockgen -source=$GOFILE -destination=${GOFILE}_mock.go -package=$GOPACKAGE

package services

import (
	"go.uber.org/fx"
)

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

func NewDummyServiceImplB(lc fx.Lifecycle) (IDummyService, error) {
	//lc.Append(fx.Hook{
	//	OnStart: func(ctx context.Context) error {
	//		return errors.Wrap(types.NewBasicBusinessError("error from NewDummyServiceImplB#OnStart", nil))
	//	},
	//})

	//return nil, errors.Wrap(types.NewBasicBusinessError("error from NewDummyServiceImplB#OnStart", nil))

	return &DummyServiceImplB{}, nil
}

func (t *DummyServiceImplB) Hello() string {
	return "Hello from DummyServiceImplB"
}
