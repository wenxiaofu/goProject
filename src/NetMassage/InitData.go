package main

import (
	//"com.sangfor.moa.protobuf"
	//"encoding/json"
	"fmt"
	//"github.com/golang/protobuf/proto"
	"reflect"
)

type User struct {
	Name string
	Age  int
	Id   string
}

//使用interface可以传入任意类型
func initmassage(pb interface{}) {
	a := reflect.ValueOf(pb)
	fmt.Println(a.Type())
	s := reflect.ValueOf(pb).Elem()
	fmt.Println(s)
	fmt.Println("第一个参数的类型：", s.Field(0).Type())
	ddd := reflect.ValueOf(s.Field(0))
	fmt.Println("*int的反射类型，是：", ddd.Type())
	//ddd_e := ddd.Elem()
	//fmt.Println("ddddd------------", ddd_e)

	dd := s.Field(0)
	fmt.Println("dd 的类型是", dd.Type())
	fmt.Println(ddd.CanSet())
	//ss := s.Field(0).Elem()
	typeOfT := s.Type()
	//fmt.Println(typeOfT)
	//s.FieldByName("DomainName").SetString("ddddddddd")
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Println(s.Field(i).IsValid())
		//	s.Field(0).Elem().SetInt(1222)
		//	s.Field(0).SetInt(122)
		fmt.Println(f.Kind())
		// switch f.Kind() {
		// case reflect.Int32:
		// 	f.SetInt(12)
		// }
		fmt.Printf("%d: %s %s = %v\n", i,
			typeOfT.Field(i).Name, f.Type(), f.Interface())
	}
	//return s
}

func main() {
	type T struct {
		Age      *int
		Name     string
		Children []int
	}
	t := T{nil, "someone-life", nil}
	//s := reflect.ValueOf(&t).Elem()
	// fmt.Println(reflect.TypeOf(s))
	// f := s.Field(0)
	// fmt.Println(f.Kind())
	// s.FieldByName("Age").SetInt(123)              // 内置常用类型的设值方法
	// sliceValue := reflect.ValueOf([]int{1, 2, 3}) // 这里将slice转成reflect.Value类型
	// s.FieldByName("Children").Set(sliceValue)
	initmassage(&t)
	//pb_auth := com_sangfor_moa_protobuf.PB_AthLoginReq{}
	//var auth com_sangfor_moa_protobuf.PB_AthLoginReq
	// pb_auth := com_sangfor_moa_protobuf.PB_AthLoginReq{
	// 	DomainName:    proto.String("wxf"),
	// 	LoginAccount:  proto.String("user"),
	// 	ClientVersion: proto.String("9.9.9"),
	// 	//		Cdt:           proto.Int64(3),
	// 	//	Did: proto.Int64(406429),
	// }
	//tonydon := &User{"TangXiaodong", 100, "0000123"}

	//initmassage(&pb_auth)
	// s := reflect.ValueOf(&pb_auth)
	// ss := s.Elem()
	// typeOfT := ss.Type()
	// fmt.Println(typeOfT)
	// for i := 0; i < ss.NumField(); i++ {
	// 	f := ss.Field(i)
	// 	fmt.Printf("%d: %s %s = %v\n", i,
	// 		typeOfT.Field(i).Name, f.Type(), f.Interface())
	// }

	//initmassage(auth)
}
