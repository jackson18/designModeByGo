package main

import "fmt"

// 定义一个接口
type IPerson interface {
	SetName(name string)
	Say()
	DoOther()
}

// 结构体Person，类似于模板
type Person struct {
	name string
	sub  IPerson
}

// 模板方法一
func (p *Person) SetName(name string) {
	p.name = name
}

// 模板方法二
func (p *Person) Say() {
	fmt.Println(p.name, " say...")
}

// 模板方法三
func (p *Person) DoOther() {
	if p.sub == nil {
		return
	}

	// 延伸到子类去实现
	p.sub.DoOther()
}

// 具体实现，匿名组合Person
type Man struct {
	Person
}

// 重写DoOther方法
func (m *Man) DoOther() {
	fmt.Println("man is busying....")
}

func main() {
	p := &Person{}
	p.sub = &Man{}
	p.SetName("superMain")
	p.Say()
	p.DoOther()
}
