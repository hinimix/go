package lib5

import (
	"fmt"
	"hinimix.com/ifactory"
)

// 实现工厂的接口
type class2 struct {

}

func (s *class2) Do() {
	fmt.Println("class2")
}

// init在运行时候初始化注册到工厂
func init() {
	ifactory.Register("class2", func() ifactory.Hclass {
		return new(class2)
	})
}
