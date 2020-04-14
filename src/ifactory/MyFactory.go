package ifactory

// 整一个工厂接口
type Hclass interface {
	Do()
}

// factoryRestore字典存储着多少结构体来注册
var factoryRestore = make(map[string] func() Hclass)

// 注册函数, 保存在factoryRestore里
func Register(name string, factory func() Hclass) {
	factoryRestore[name] = factory
}

// 生成不同结构体类型的接口
func Create(name string) Hclass{
	if f,ok := factoryRestore[name]; ok{
		return f()
	} else {
		panic("hinimix said that: 'name not found'")
	}
}