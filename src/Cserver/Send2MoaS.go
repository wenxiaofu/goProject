package Cserver

import (
	"NetMassage"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"net"
	"os"
	"time"
)

func checkError(err error) {
	if err != nil {
		fmt.Println("Error: %s", err.Error())
		os.Exit(1)
	}
}

//发送登录请求，每次都需要获取一下数据库里面的用户名和密码数据,这个maping这么取数据，遍历key和遍历values
func GetUserbyID(db *sql.DB, ip string) map[string]string {
	start := time.Now()
	fmt.Println("200.200.169.212")
	useranpwd := make(map[string]string)
	rows, _ := db.Query("SELECT username,password FROM users where server = ?", ip)
	defer rows.Close()
	for rows.Next() {
		var username string
		var password string
		if err := rows.Scan(&username, &password); err != nil {
			fmt.Println(err)
		}
		useranpwd[username] = password

		//fmt.Printf("name:%s ,id:is %d\n", name, id)
	}
	end := time.Now()
	fmt.Println("方式1 query total time:", end.Sub(start).Seconds())
	return useranpwd
}

func GetUsersByIp(db *sql.DB, ip string) ([]string, []string) {
	start := time.Now()
	fmt.Println("200.200.169.212")
	user := make([]string, 0)
	pwd := make([]string, 0)
	//使用切片可以支持append,append是追加到了最后面
	rows, _ := db.Query("SELECT username,password FROM users where server = ?", ip)
	defer rows.Close()
	for rows.Next() {
		var username string
		var password string
		if err := rows.Scan(&username, &password); err != nil {
			fmt.Println(err)
		}
		user = append(user, username)
		pwd = append(pwd, password)

		//fmt.Printf("name:%s ,id:is %d\n", name, id)
	}
	end := time.Now()
	fmt.Println("方式1 query total time:", end.Sub(start).Seconds())
	return user, pwd
}

//只接受massage

func Send2MoaS(user string, pwd string) {
	//在这里面构建认证和登录数据
	//构建认证数据
	authmassage := NetMassage.Authbyte(user)
	//构建登录数据
	loginmassage := NetMassage.LoginMassage(user, pwd)
	//构建业务数据库
	yewu_massage := NetMassage.CreateW()
	//创建连接,此处可以是连接池
	conn, err := net.Dial("tcp", "200.200.169.212:6800")
	checkError(err)
	defer conn.Close()
	//发送登录和认证请求
	conn.Write(authmassage)
	buft_w := make([]byte, 1024)
	conn.Read(buft_w)
	//发送登录请求
	conn.Write(loginmassage)
	conn.Read(buft_w)
	//发送业务请求
	conn.Write(yewu_massage)
	//获取响应值，这个地方怎么可以不定义变量
	conn.Read(buft_w)
	//增加一个打印
	fmt.Print(buft_w)
	//	ch <- id
}
