package main
// usage: go run src/test.go -name="dsb"
import (
	//参数化引用包
	"flag"
	"fmt"
)

var name string

func init () {
	flag.StringVar(&name, "name", "everyone", "the greeting object")
}

func main() {
	flag.Parse()
	fmt.Printf("hello %s\n", name)
}
