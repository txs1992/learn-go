// struct and interface，结构体与接口。
package main

import "fmt"

type person struct {
	name string
	age  uint8
	size uint
	address
}

type address struct {
	province string
	city     string
}

func (p person) printAge() {
	fmt.Println("the person age is：", p.age)
}

func (p person) printName() {
	fmt.Println("the person name is：", p.name)
}

func (p person) printAddress() {
	fmt.Printf("这个人住在%s省%s市\n", p.province, p.city)
}

func printString(s fmt.Stringer) {
	fmt.Println(s.String())
}

func (p person) String() string {
	return fmt.Sprintf("the name is %s, age is %d", p.name, p.age)
}

func (addr address) String() string {
	return fmt.Sprintf("the addr is %s省%s市", addr.province, addr.city)
}

func main() {
	// p := person{"mt", 28, 170}
	var p person = person{age: 28, name: "mt", address: address{"安徽", "合肥"}}
	p.size = 170
	printString(&p)
	printString(p.address)

	p2 := personFactory("张三")
	fmt.Println(p2.name, p2.age)
	// fmt.Println(p, p.addr)
	// fmt.Println(p.name, p.age, p.size)
	// p.printAge()
	// p.printName()
	// p.printAddress()

	// 类型断言
	var s fmt.Stringer
	s = p2

	a, ok := s.(address)

	if ok {
		fmt.Println(a)
	} else {
		fmt.Println("s 不是一个 address")
	}
}

// 工厂函数
func personFactory(name string) person {
	return person{name: name}
}
