package lib5

import (
	"fmt"
	"hinimix.com/ifactory"
)

// 实现工厂的接口
type class1 struct {

}

func (s *class1) Do() {
	fmt.Println("class1")
}

// init在运行时候初始化注册到工厂
func init() {
	ifactory.Register("class1", func() ifactory.Hclass {
		return new(class1)
	})
}
