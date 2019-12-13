package main

// usage: go run test.go -name="dsb"
import (
	//参数化引用包
	"flag"
	//lib567 is alias
	lib567 "test_go/lib"
)

var name string

func init() {
	flag.StringVar(&name, "name", "everyone", "the greeting object")
}

func main() {
	flag.Parse()
	//fmt.Printf("hello %s\n", name)
	lib567.Hello(name)
	lib567.Hello("dad")
}
