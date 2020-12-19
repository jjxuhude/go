package main

import (
	json2 "encoding/json"
	"fmt"
	"reflect"
)

func main() {
	//test1()
	//test2()
	//test3()
	//test4()
	//test5()
	//test6()
	test7()
}

func test7() {
	type User struct {
		Name string
		Age  int
	}
	type ColorGroup struct {
		ID     int
		Name   string
		Colors []string
		Users  []User
	}
	group := &ColorGroup{
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon", "粮食"},
		Name:   "Reds",
		Users:  []User{{"user01", 11}, {"user02", 22}},
	}
	group.Users = append(group.Users, User{"user04", 444})
	jsonData, _ := json2.Marshal(group)
	json := string(jsonData)
	fmt.Println(json)
}

func test6() {
	userIns := &user{"徐华德", 18}
	userRef := reflect.ValueOf(userIns)
	m1 := userRef.MethodByName("Method1")
	params := []reflect.Value{reflect.ValueOf("方法m1")}
	resutl := m1.Call(params)
	fmt.Println(resutl)
}

func test5() {
	var a = 5
	aRef := reflect.ValueOf(&a)
	aRef.Elem().SetInt(10)
	fmt.Println(a)

}

func test4() {
	var a *int
	aRef := reflect.ValueOf(a)
	fmt.Println("isNil:", aRef.IsNil())
	fmt.Println("isValid:", aRef.IsValid())
}
func test3() {
	var size int = 1024
	sizeRef := reflect.ValueOf(size)
	var size1 = sizeRef.Interface().(int)
	fmt.Println(size1)

	user1 := user{"jjxuhuade", 18}
	userRef := reflect.ValueOf(user1)

	fmt.Println("字段数量：", userRef.NumField())
	name := userRef.Field(0)
	fmt.Println("获取索引0的值：", name)
	fmt.Println("获取字段的类型：", name.Kind())
	fmt.Println("获取字段的值：", userRef.FieldByName("age"))
	fmt.Println("NumMethod:", userRef.NumMethod())
	for i := 0; i < userRef.NumMethod(); i++ {
		method := userRef.Method(i)
		params := []string{"徐华德"}
		in := make([]reflect.Value, len(params))
		in[0] = reflect.ValueOf(params[0])
		fmt.Println(method.Call(in))
	}
}

func test1() {
	i := 3
	fmt.Println(reflect.TypeOf(i))
	v := reflect.ValueOf(i)
	fmt.Println(v)
	fmt.Println(v.Interface())
	fmt.Println(v.Kind())
	fmt.Println(v.Type())
}

type user struct {
	name string
	age  int
}

type address struct {
	addr string
	tel  int
}

func (u user) Method1(name string) string {
	return "m1 " + name
}

func (u user) Method2(name string) string {
	return "m2 " + name
}

func test2() {
	var a int
	ins := reflect.TypeOf(a)
	fmt.Println(ins.Name(), ins.Kind())

	user1 := user{"jjxuhuade", 18}

	userIns := reflect.TypeOf(user1)
	fmt.Println(userIns.Name(), userIns.Kind())
	useInsPtr := reflect.TypeOf(&user1)
	fmt.Println("prt:", useInsPtr, useInsPtr.Elem())

	for i := 0; i < userIns.NumField(); i++ {
		field := userIns.Field(i)
		fmt.Println("字段：", field.Name, field.Type)
	}
	if name, ok := userIns.FieldByName("name"); ok {
		println("fieldByName:", name.Name)
	}
}
