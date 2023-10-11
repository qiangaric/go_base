package main

import "fmt"

type s1 struct {
	name string
	age  int
}

func (s *s1) f1() {
	fmt.Printf("name = %s, age = %d \n", s.name, s.age)
}

type s2 struct {
	*s1
	interface1
}

type interface1 interface {
	interface2
	f2()
}

type interface2 interface {
	f3()
}

type s3 struct {
}
// 构造函数
// func newS3() interface1 {
// 	return &s3{}
// }

func (s *s3)f2() {
	fmt.Println("调用了s3.f2")
}
func (s *s3)f3() {
	fmt.Println("调用了s3.f3")
}

func main() {
	ss3 := &s3{}
	s := &s2{
		s1: &s1{
			name: "老张", age: 18,
		},
		interface1: ss3,
	}
	s.f1()
	s.f2()
	s.f3()
}
