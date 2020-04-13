package main

/**
构造者模式：
使用简单对象构造复杂对象

Builder模式将复杂对象的构造与其表示分离，以便相同的构造过程可以创建不同的表示
*/

type Speed float64

const (
	MPH Speed = 1
	KPH       = 1.60934
)

type Color string

const (
	BlueColor  Color = "blue"
	GreenColor       = "green"
	RedColor         = "red"
)

type Wheels string

const (
	SportsWheels Wheels = "sports"
	SteelWheels         = "steel"
)

type Builder interface {
	Color(Color) Builder
	Wheels(Wheels) Builder
	TopSpeed(speed Speed) Builder
	Builder() Interface
}

type Interface interface {
	Drive() error
	Stop() error
}

func main() {

}