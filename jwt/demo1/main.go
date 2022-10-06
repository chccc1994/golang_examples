package main

import (
	"fmt"
	"jwt_demo1/routers"
)

func main() {
	fmt.Println("hello world")
	// tk, err := GenToken("abc123")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(tk)
	// var u1 = &MyClaims{}

	// u1, err = ParseToken(tk)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(u1.Username)

	routers.InitRouter()
}
