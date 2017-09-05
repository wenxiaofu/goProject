package main

import (
	//	"NetMassage"
	"com.sangfor.moa.protobuf"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang/protobuf/proto"
	"reflect"
)

//插入原始数据
func insertOrgmodel(db *sql.DB, name string, serv_op int32, contrl_op int32, jcontext string, pcontext []byte) {
	result, err := db.Exec(
		"INSERT INTO orgmodel (name, serv_op,contrl_op,jcontext,pcontext) VALUES (?, ?,?,?,?)",
		name,
		serv_op,
		contrl_op,
		jcontext,
		pcontext,
	)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}

//插入用户模板数据
func InsertUsermodel(db *sql.DB, name string, serv_op int32, contrl_op int32, jcontext string, pcontext []byte) {
	result, err := db.Exec(
		"INSERT INTO orgmodel (name, serv_op,contrl_op,jcontext,pcontext) VALUES (?, ?,?,?,?)",
		name,
		serv_op,
		contrl_op,
		jcontext,
		pcontext,
	)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}

//读取模板数据,根据操作码和服务码进行查询
func GetOrgData(db *sql.DB, serv_op int32, contrl_op int32) (jcontext string, pcontext []byte) {

	db.QueryRow("select jcontext, pcontext from orgmodel where serv_op = ? and contrl_op = ?", 2, 1).Scan(&jcontext, &pcontext)
	return jcontext, pcontext
}

func main() {
	db, err := sql.Open("mysql", "moatest:moatest@tcp(200.200.169.215:3306)/moatest?charset=utf8")
	if err != nil {
		fmt.Println(err)
	}
	a := reflect.ValueOf(db)
	fmt.Println("ddddddddddddddd", a)
	//关闭数据库，db会被多个goroutine共享，可以不调用
	defer db.Close()
	//插入默认的值

	//先插入一个author数据
	// fmt.Println("db connect is %s", db)
	// jsont := NetMassage.Auth2json("12200000001")
	// pcon := NetMassage.Auth2Pb("12200000001")
	// result, err1 := db.Exec(
	// 	"INSERT INTO orgmodel (name, serv_op,contrl_op,jcontext,pcontext) VALUES (?, ?,?,?,?)",
	// 	"authdata",
	// 	2,
	// 	1,
	// 	jsont,
	// 	pcon,
	// )
	// if err1 != nil {
	// 	fmt.Println(err1)
	// }

	//fmt.Print(result)
	//读取数据，发送请求通过服务码和操作码进行查询
	serv_op := int32(2)
	contrl_op := int32(1)
	jcontext := "test"
	pcontext := []byte{}
	fmt.Println("-----------------------")
	GetOrgData(db, serv_op, contrl_op)
	jcontext, pcontext = GetOrgData(db, serv_op, contrl_op)
	fmt.Println(jcontext, pcontext)
	fmt.Println("111111111111111")

	rows3 := db.QueryRow("select serv_op, contrl_op, jcontext, pcontext from orgmodel where serv_op = ? and contrl_op = ?", 2, 1)

	rows3.Scan(&serv_op, &contrl_op, &jcontext, &pcontext)
	fmt.Println(serv_op, contrl_op, jcontext, pcontext)
	//将bcontext进行解码,转存到数据库的pb读出来可以直接转，那么可以直接保存pb就可以了
	var umData com_sangfor_moa_protobuf.PB_AthLoginReq
	err = proto.Unmarshal(pcontext, &umData)

	if err != nil {
		fmt.Println("Error2: ", err)
		return
	}

	// 输出结果
	fmt.Println(*umData.DomainName, *umData.LoginAccount, *umData.ClientVersion)
}
