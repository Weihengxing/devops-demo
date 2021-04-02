package service

import (
	"github.com/Weihengxing/devops-demo/pkg/core"
	"github.com/Weihengxing/devops-demo/pkg/store"
)

type IService interface {
	Create(object core.IObject)error
	Update(src core.IObject,target core.IObject)error
	Delete(uuid string)error
	Query(c store.Condition) ([]core.IObject,error)
	Range(c store.Condition)error
}

var _ IService=&FakeService{}

type FakeService struct {}

func (f FakeService) Range(c store.Condition) error {
	panic("implement me")
}

func (f FakeService) Create(object core.IObject) error {
	panic("implement me")
}

func (f FakeService) Update(src core.IObject, target core.IObject) error {
	panic("implement me")
}

func (f FakeService) Delete(uuid string) error {
	panic("implement me")
}

func (f FakeService) Query(c store.Condition) ([]core.IObject, error) {
	panic("implement me")
}
