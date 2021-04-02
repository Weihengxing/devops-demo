package api

import (
	"fmt"
	"net/http"
)

type IServers []IServer

func NewServer(servers ...IServer) IServers {

	//return &FakeServer{}
	//fakeServer:=new(FakeServer)
	//return fakeServer
	is :=make(IServers,0)
	is=append(is,servers...)
	return servers
}

func (i IServers)RunAll()<-chan error  {
	errors:=make(chan error,len(i))
	for _,server:=range i {
		_server:=server
		go func(is IServer) {
			if err:=_server.Run();err!=nil{
				errors <-err
			}
		}(_server)
	}
	return errors
}

type IServer interface {
	Run() error
	Listen(addr string)
	Register(string,http.Handler)error
}



type MyInterface interface {
	Do()
}

//go的结构体不能传递引用的
var _ IServer = &FakeServer{}

type FakeServer struct {
	MyInterface  //指针
}

func (f *FakeServer) Listen(addr string) {
	panic("implement me")
}

func (f *FakeServer) Register(s string,handler http.Handler) error {
	panic("implement me")
}

func (f *FakeServer)Run()error  {
	fmt.Printf("fakeServer run")
	return nil
}

func startFakeServer(fk FakeServer)  {
	fk.Run()
	fk.Do()
}
