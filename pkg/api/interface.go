package api

import "fmt"

type IServer interface {
	Run() error
}

func  NewServer() IServer {
	return nil
}

type MyInterface interface {
	Do()
}

//go的结构体不能传递引用的
var _ IServer = &FakeServer{}

type FakeServer struct {
	MyInterface  //指针
}

func (f *FakeServer)Run()error  {
	fmt.Printf("fakeServer run")
	return nil
}

func startFakeServer(fk FakeServer)  {
	fk.Run()
	fk.Do()
}
