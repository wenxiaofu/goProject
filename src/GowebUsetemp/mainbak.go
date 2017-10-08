package main

import "net/http"

import (
	"html/template"
	"io"
	"log"
	// "os"
)
//这里的viewModel就是传输的数据
func renderTemplate(w http.ResponseWriter, name string, template string, viewModel interface{}) {
    tmpl, ok := templates[name]
    if !ok {
        http.Error(w, "The template does not exist.", http.StatusInternalServerError)
    }
    err := tmpl.ExecuteTemplate(w, template, viewModel)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }

}

type HttpHander struct{}
type Mstring string

func getNotes(w http.ResponseWriter, request *http.Request) {
	t, _ := template.ParseFiles("static/index.html")
    t.ExecuteTemplate(w, name, data)
	// t.Execute(w, nil)
	// io.WriteString(response, "hello")
}
func func addNote(w http.ResponseWriter, r *http.Request) {
    renderTemplate(w, "add", "base", nil)
}
func main() {

	mux := http.NewServeMux()
	// mux.Handle("/", HttpHander{})
	//获取当前路径
	// wd, err := os.Getwd()
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	//func StripPrefix(prefix string, h Handler) Handler
	// 给定url 删除前缀
	mux.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("./static"))))
	mux.Handle("/css/", http.StripPrefix("/css", http.FileServer(http.Dir("./static/css"))))
	mux.Handle("/js/", http.StripPrefix("/js", http.FileServer(http.Dir("./static/js"))))
	mux.Handle("/fonts/", http.StripPrefix("/fonts", http.FileServer(http.Dir("./static/fonts"))))
	mux.Handle("/images/", http.StripPrefix("/images", http.FileServer(http.Dir("./static/images"))))
	log.Println("ddd" + "css")
	//mux.Handle("/css/", http.StripPrefix("/css", http.FileServer(http.Dir(wd))+"/css"))
	mux.HandleFunc("/index", getNotes)
	//http.HandleFunc("/", sayHello)
	http.ListenAndServe(":8080", mux)

}

// func (HttpHander) ServeHTTP(response http.ResponseWriter, request *http.Request) {

// 	io.WriteString(response, "ServeHTTP")
// }
func sayHello(response http.ResponseWriter, request *http.Request) {

	io.WriteString(response, "sayHello")
}

//http://qq466862016.iteye.com/blog/2274404文章地址
