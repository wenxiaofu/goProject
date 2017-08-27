package main

import (
	"bufio" //缓存IO
	"bufio" //缓存IO
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Student struct {
	Name    string
	Age     int
	Guake   bool
	Classes []string
	Price   float32
}

func (s *Student) ShowStu() {
	fmt.Println("show Student :")
	fmt.Println("\tName\t:", s.Name)
	fmt.Println("\tAge\t:", s.Age)
	fmt.Println("\tGuake\t:", s.Guake)
	fmt.Println("\tPrice\t:", s.Price)
	fmt.Printf("\tClasses\t: ")
	for _, a := range s.Classes {
		fmt.Printf("%s ", a)
	}
	fmt.Println("")
}

var xxx = map[string]string{}

//var xxx []map[string]interface{}

func readFile(filename string) (map[string]string, error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("ReadFile: ", err.Error())
		return nil, err
	}
	if err := json.Unmarshal([]byte(bytes), &xxx); err != nil {
		fmt.Println("Unmarshal: ", err.Error())
		return nil, err
	}

	return xxx, nil
}

//查看文件是否存在
func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

//写文件
func writeFile(filename string, content string) {
	w := bufio.NewWriter(f) //创建新的 Writer 对象
	n4, err3 := w.WriteString("bufferedn")
	fmt.Printf("写入 %d 个字节n", n4)
	w.Flush()
	f.Close()

}

func main() {
	//文件读取
	f, err3 := os.Create("./output3.txt") //创建文件

	//json解析
	st := &Student{
		"Xiao Ming",
		16,
		true,
		[]string{"Math", "English", "Chinese"},
		9.99,
	}
	fmt.Println("before JSON encoding :")
	st.ShowStu()

	b, err := json.Marshal(st)
	if err != nil {
		fmt.Println("encoding faild")
	} else {
		fmt.Println("encoded data : ")
		fmt.Println(b)
		fmt.Println(string(b))
	}
	ch := make(chan string, 1)
	go func(c chan string, str string) {
		c <- str
	}(ch, string(b))
	strData := <-ch
	fmt.Println("--------------------------------")
	stb := &Student{}
	//	stb.ShowStu()
	err = json.Unmarshal([]byte(strData), &stb)
	if err != nil {
		fmt.Println("Unmarshal faild")
	} else {
		fmt.Println("Unmarshal success")
		stb.ShowStu()
	}
	xxxMap, err := readFile("JsonModel.json")
	if err != nil {
		fmt.Println("readFile: ", err.Error())
	}

	fmt.Println(xxxMap)
}
