package main
import (
	"fmt"
	"net/http"
	"strings"
	"log"
)
func Rsp_web(w http.ResponseWriter,r *http.Request) {
	r.ParseForm() //解析参数
	fmt.Println(r.Form)//服务端打印客户端上传的信息
	fmt.Println("path",r.URL.Path)
	for k,v := range r.Form{
		fmt.Println("key",k)
		fmt.Println("val",strings.Join(v,""))
	}
	fmt.Fprintf(w,"hello astadd")
	
}
func main(){
	http.HandleFunc("/",Rsp_web)//设置访问路由
	err := http.ListenAndServe(":9090",nil)
	if err != nil {
		log.Fatal("listenandserve: ",err)
	}
}