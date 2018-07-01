package main

import "fmt"
//一种或者多种数据类型的集合
type Person struct{
	name string
	age int
}

func main(){
	var person Person
	person.name = "gxm" //访问数据成员用 .
	person.age = 28
	fmt.Printf("person name %s\n",person.name)
	fmt.Printf("persion age %d\n",person.age)

	var struct_pointer *Person //申明指针 
	struct_pointer = &person //取地址
	fmt.Printf("person name %s\n",struct_pointer.name)
	fmt.Printf("person age %d\n",struct_pointer.age)
}