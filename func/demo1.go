package main

import "fmt"

// func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	...
// }
// ServeHTTP不是一个独立的功能。函数名称前面的括号是Go定义这些函数将在其上运行的对象的方式。
// 所以，本质上ServeHTTP是一个类型处理程序的方法，可以使用类型处理程序的任何对象来调用，比如h。
// 他们也被称为接收者。这里是定义他们的方法有两种。如果你想修改接收器，使用如下的指针：
type Mutatable struct {
	a int
	b int
}
type User struct {
	ageA int
	ageB int
}

func (m Mutatable) StayTheSame() {
	m.a = 5
	m.b = 7
}

func (m *Mutatable) Mutate() {
	m.a = 5
	m.b = 7
}

func main() {
	m := &Mutatable{0, 0}
	fmt.Println(m)
	m.StayTheSame()
	fmt.Println(m)
	m.Mutate()
	fmt.Println(m)

	// 报错
	// u1 := User{1, 1}
	// u1.Mutate()
	// Mutate()
}
