package main

import (
	"fmt"
	"grpc_demo/service"

	"google.golang.org/protobuf/proto"
)

func main() {
	user := &service.User{
		Username: "mszlu",
		Age:      20,
	}
	//转换为protobuf
	marshal, err := proto.Marshal(user)
	if err != nil {
		panic(err)
	}
	newUser := &service.User{}
	err = proto.Unmarshal(marshal, newUser)
	if err != nil {
		panic(err)
	}
	fmt.Println(newUser.String())
}
