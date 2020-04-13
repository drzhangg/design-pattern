package _1_simple_factory

import "fmt"

/**
简单工厂模式：
因为go中没有构造函数，所以一般会定义NewXXX函数来初始化相关类。NewXXX函数返回接口时就是简单的工厂模式，
也就是说golang的一般推荐做法就是简单工厂
*/

type API interface {
	Say(name string) string
}

func NewAPI(t int) API {
	switch t {
	case 1:
		return &hiAPI{}
	case 2:
		return &helloAPI{}
	default:
		return nil
	}
}

type hiAPI struct {
}

func (*hiAPI) Say(name string) string {
	return fmt.Sprintf("Hi, %s", name)
}

type helloAPI struct {
}

func (*helloAPI) Say(name string) string {
	return fmt.Sprintf("Hello, %s", name)
}
