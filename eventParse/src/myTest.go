package main

import (
	"fmt"
)

type a1 struct {
	name string
	age  int
	flag string

	Event Event
}

// func (a a1) Decode(data string, data1 string) {
// 	a.Event.Decode(data, data1)
// }
// func (a a1) Dump() {

// }

type Event interface {
	Decode(data string, data1 string)
	Dump()
}

type QueryEvent struct {
	Schema string
	Query  string
}

func (e *QueryEvent) Decode(data string, data1 string) {
	e.Schema = data
	e.Query = data1
	fmt.Printf("Decode === Schema: %s\n", data)
	fmt.Printf("Decode === Query: %s\n", data1)
}

func (e *QueryEvent) Dump() {

	fmt.Printf("Dump === Schema: %s\n", e.Schema)
	fmt.Printf("Dump === Query: %s\n", e.Query)

}

type TableEvent struct {
	Tbale string
	Field string
}

func (e *TableEvent) Decode(data string, data1 string) {
	e.Tbale = data
	e.Field = data1
	fmt.Printf("Decode === table: %s\n", data)
	fmt.Printf("Decode === filed: %s\n", data1)
}

func (e *TableEvent) Dump() {

	fmt.Printf("Dump === table: %s\n", e.Tbale)
	fmt.Printf("Dump === filed: %s\n", e.Field)

}

type Shape interface {
	area() float64
	circumference() float64
}

// 结构体正方形，属性边长
type square struct {
	length float64
}

// 方法area，由正方形结构体实现
func (s *square) area() float64 {
	sarea := s.length * s.length
	fmt.Println("正方形的面积为：", sarea)
	return sarea
}

// 方法circumference，由正方形结构体实现
func (s *square) circumference() float64 {
	scircumference := s.length * 4
	fmt.Println("正方形的周长为：", scircumference)
	return scircumference
}

// func getarea(len float64) Shape {
// 	s := &square{
// 		length:4,
// 	}
// 	fmt.Println("正方形的面积为：",s.area())
// 	fmt.Println("正方形的周长为：",s.circumference())
// 	return s
// }

func getarea(len float64) Shape {
	s := &square{
		length: 4,
	}
	s.area()
	s.circumference()
	// fmt.Println("正方形的面积为：",s.area())
	// fmt.Println("正方形的周长为：",s.circumference())
	return s
}
func DealEvent(flag, data, data1 string) Event {
	var s Event
	

	if flag == "schema" {
		// 接口动态类型
		s = &QueryEvent{}
		
	} else {
		s = &TableEvent{}
		
	}
	// 接口的动态调用
	s.Decode(data, data1)
	return s

	// s = &QueryEvent{
	// 	Schema: "schema0001",
	// 	Query:  "Query001",
	// }

	// a111 := a1{
	// 	name: "zhangsan",
	// 	age:  123,
	// 	Event:  s,
	// }
	// hahha := a111.Event.(*QueryEvent)
	// fmt.Println(hahha)
	// hahha.Decode("abc","ccc")
	// hahha.Dump()

}
func main() {
	// a11 := a1{
	// 	name:    "asdd",
	// 	age:     12,
	// 	address: "asdads",
	// }

	// // tetet := a11.Event
	// // fmt.Println(tetet)
	// var a Event
	// // b := a.(QueryEvent01)
	// b.Dump
	/*
		第一种调用方式
	*/
	// var shape Shape
	// shape = &square{
	// 	length : 4,
	// }
	// shape.area()
	// shape.circumference()

	/*
		第二种调用方式
	*/
	// getarea(4)

	/*
		接口类型断言
	*/
	// var s Event

	// s = &QueryEvent{
	// 	Schema: "schema0001",
	// 	Query:  "Query001",
	// }
	// test := s.(QueryEvent)
	// test001 := s.(*QueryEvent)
	// fmt.Println(test001)
	// test001.Decode("122", "12331")
	// test001.Dump()

	// a111 := a1{
	// 	name: "zhangsan",
	// 	age:  123,
	// 	Event:  s,
	// }
	// hahha := a111.Event.(*QueryEvent)
	// fmt.Println(hahha)
	// hahha.Decode("abc","ccc")
	// hahha.Dump()

	/*
		接口类型断言其二
	*/
	var s Event

	s = &QueryEvent{
		Schema: "schema0001",
		Query:  "Query001",
	}
	// test := s.(*QueryEvent)
	test001 := s.(*QueryEvent)
	fmt.Println(test001)
	// test001.Decode("122", "12331")
	// test001.Dump()
	// var s Event
	a111 := a1{
		name:  "zhangsan",
		age:   123,
		flag:  "schema",
		Event: s,
	}
	tete := a111.Event.(*QueryEvent)
	fmt.Println(tete)
	// DealEvent("schema","test001", "test0002")
	// DealEvent("table","test001", "test0002")
	
	
	

}
