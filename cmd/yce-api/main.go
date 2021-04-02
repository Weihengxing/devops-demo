package main

import (
	"fmt"
	"github.com/Weihengxing/devops-demo/pkg/api"
	"github.com/Weihengxing/devops-demo/pkg/api/gin"

	//dd "github.com/laik/devops-demo/pkg/api"
)

//需要提供一个用户的CRUD

func main() {


	fmt.Println("yce.api")
	server:=api.NewServer(gin.NewGin())
	for err:=range server.RunAll() {
		panic(err)
	}


}
