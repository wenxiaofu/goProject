package main

import (
	"Cserver"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "moatest:moatest@tcp(200.200.169.215:3306)/moatest?charset=utf8")
	if err != nil {
		fmt.Println(err)
	}
	//关闭数据库，db会被多个goroutine共享，可以不调用
	defer db.Close()
	useranpwd := Cserver.GetUserbyID(db, "200.200.169.212")
	users, pwd := Cserver.GetUsersByIp(db, "200.200.169.212")
	fmt.Println(users[0])
	fmt.Println(pwd[0])
	// useranpwd1 := make(map[string]string)
	// start := time.Now()

	// //查询没有数据的时候，下面的方法会报错，这里先不管了，一般都是有数据的
	// rows, _ := db.Query("SELECT username,password FROM users where server = ?", "200.200.169.212")

	// for rows.Next() {
	// 	var username string
	// 	var password string
	// 	if err := rows.Scan(&username, &password); err != nil {
	// 		fmt.Print(err)
	// 	}
	// 	useranpwd[username] = password

	// 	//fmt.Printf("name:%s ,id:is %d\n", name, id)
	// }
	// end := time.Now()
	// fmt.Println("方式1 query total time:", end.Sub(start).Seconds())

	for country := range useranpwd {
		fmt.Println("Capital of", country, "is", useranpwd[country])
	}

	Cserver.Send2MoaS(users[0], pwd[0])

}
