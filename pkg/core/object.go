package core

import "encoding/json"

type Kind = string

type IObject interface {
	UUID() string
	GetKind( )Kind
	Get(string) interface{}
	Set(string,interface{})
}



type Object = map[string]interface{}

func (o Object) GetKind() Kind {
	panic("implement me")
}

func FromBytes(b []byte)(Object,error)  {
	o:=make(Object)
	if err:=json.Unmarshal(b,&o);err!=nil {
		return nil,err   //nil=0x00
	}
	return o,nil
}

//好用的方法
func Iter(n int64) []struct{} {
	return make([]struct{},0,n)
}

//sized 可被计算长度都可以放在stack
//内存逃逸
type None struct {}

func example() {
	none :=&None{}
	_ = none

	channel:=make(chan struct{})

	channel <- struct{}{}

	for _,v:=range Iter(100){
			_=v
	}
}