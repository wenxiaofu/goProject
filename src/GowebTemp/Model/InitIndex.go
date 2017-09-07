package Model

import (
	"net/http"
)

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

func moa_index_init(w http.ResponseWriter, r *http.Request) RunTest {

	InitRunTest.Ip = "200.200.169.212"
	InitRunTest.NumofUser = 100
	InitRunTest.RunTimes = 12
	InitRunTest.RunTimes = 1
	InitRunTest.Serv_op = []int32{1, 2, 3, 4, 5, 6, 7, 8, 9}
	InitRunTest.Contrl_op = []int32{11, 23, 432}
	InitRunTest.Jcontext = "{\"name\":\"wxf\"}"
	renderTemplate(w, "Moa_index", "base", InitRunTest)
}
