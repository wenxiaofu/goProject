package main

import "fmt"

func sixteen2t(num int32) int32 {
	return num/10*16 + num%10
}

func main() {
	fmt.Println(sixteen2t(39))
}
