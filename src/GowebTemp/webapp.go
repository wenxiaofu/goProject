package main

import (
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Note struct {
	Title       string
	Description string
	CreateOn    time.Time
}

type EditeNote struct {
	Note
	Id string
}

var noteStore = make(map[string]Note)
var id = 0
var templates map[string]*template.Template

func init() {
	if templates == nil {
		templates = make(map[string]*template.Template)
	}
	//了解template.Musthttp://www.jianshu.com/p/bee02c18b221
	templates["index"] = template.Must(template.ParseFiles("templates/index.html", "templates/base.html"))
	templates["add"] = template.Must(template.ParseFiles("templates/add.html", "templates/base.html"))
	templates["edit"] = template.Must(template.ParseFiles("templates/edit.html", "templates/base.html"))
	templates["Moa_index"] = template.Must(template.ParseFiles("templates/Moa_index.html", "templates/base.html"))
	templates["running"] = template.Must(template.ParseFiles("templates/running.html", "templates/base.html"))
}

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

func getNotes(w http.ResponseWriter, r *http.Request) {

	renderTemplate(w, "index", "base", noteStore)
}

func addNote(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "add", "base", nil)
}

func saveNote(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	title := r.PostFormValue("title")
	desc := r.PostFormValue("description")
	note := Note{title, desc, time.Now()}
	id++

	k := strconv.Itoa(id)
	noteStore[k] = note
	http.Redirect(w, r, "/", 302)
}

func editNote(w http.ResponseWriter, r *http.Request) {
	var viewModel EditeNote
	vars := mux.Vars(r)
	k := vars["id"]

	if note, ok := noteStore[k]; ok {
		viewModel = EditeNote{note, k}
	} else {
		http.Error(w, "Could not find the resource to edit,", http.StatusBadRequest)
	}

	renderTemplate(w, "edit", "base", viewModel)
}

func updateNote(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	k := vars["id"]
	var noteToUpd Note
	if note, ok := noteStore[k]; ok {
		r.ParseForm()
		noteToUpd.Title = r.PostFormValue("title")
		noteToUpd.Description = r.PostFormValue("description")
		noteToUpd.CreateOn = note.CreateOn
		delete(noteStore, k)
		noteStore[k] = noteToUpd

	} else {
		http.Error(w, "Could not find the resource to edit,", http.StatusBadRequest)
	}
	http.Redirect(w, r, "/", 302)
}

func deleteNote(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	k := vars["id"]
	if _, ok := noteStore[k]; ok {
		delete(noteStore, k)
	} else {
		http.Error(w, "Could not find the resource to edit,", http.StatusBadRequest)
	}

	http.Redirect(w, r, "/", 302)
}

type RunTest struct {
	Ip        string
	NumofUser int32
	RunTimes  int32
	Rate      int32
	Serv_op   []int32
	Contrl_op []int32
	Jcontext  string
	//这里应该可以直接用json传递，后面引入js以后再考虑
}

var InitRunTest RunTest

var servername = []string{"1haele", "2", "3"}
var serverop = []int{1, 2, 3, 4, 5, 5}

//这里先传一个默认的一些参数，noteStore
func moa_index(w http.ResponseWriter, r *http.Request) {
	InitRunTest.Ip = "200.200.169.212"
	InitRunTest.NumofUser = 100
	InitRunTest.RunTimes = 12
	InitRunTest.RunTimes = 1
	InitRunTest.Serv_op = []int32{1, 2, 3, 4, 5, 6, 7, 8, 9}
	InitRunTest.Contrl_op = []int32{11, 23, 432}
	InitRunTest.Jcontext = "{\"name\":\"wxf\"}"
	renderTemplate(w, "Moa_index", "base", InitRunTest)
}

type RunData struct {
	//正在运行的用户数
	Runningnum int32
	//已经结束的用户数
	FlishedUser int32
	//完成的事物数
	SamplerNum  int64
	AverageTime float32
	//还没有容错机制，这个忽略不用
	FieldUser int32
}

//解析前段传过来的数据，然后跳转到running.html页面
func moa_runnig(w http.ResponseWriter, r *http.Request) {

	//实时返回运行
	renderTemplate(w, "running", "base", "Runningdata")
}

//提供一个刷新运行状态的方法

var fleshtime int64

func moa_running_status_reflesh(w http.ResponseWriter, r *http.Request) {
	fleshtime++
	//get_runing_status,用于获取当前运行状态，用户数，成功了多少
	renderTemplate(w, "running", "base", fleshtime)
}

func main() {
	r := mux.NewRouter().StrictSlash(false)
	fs := http.FileServer(http.Dir("public"))
	r.Handle("/public/", fs)
	r.HandleFunc("/", getNotes)
	r.HandleFunc("/notes/add", addNote)
	r.HandleFunc("/notes/save", saveNote)
	r.HandleFunc("/notes/edit/{id}", editNote)
	r.HandleFunc("/notes/update/{id}", updateNote)
	r.HandleFunc("/notes/delete/{id}", deleteNote)
	r.HandleFunc("/moa", moa_index)
	r.HandleFunc("/moa/running", moa_runnig)
	r.HandleFunc("/moa/reflesh", moa_running_status_reflesh)

	server := &http.Server{
		Addr:    ":9090",
		Handler: r,
	}
	log.Println("Listening...")
	server.ListenAndServe()
}
