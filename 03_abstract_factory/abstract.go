package main

import "fmt"

/**
抽象工厂模式：
抽象工厂模式提供用于创建已发布对象家族的接口，所生成的对象都是有关联的
*/

// OrderMainDAO 为订单主记录
type OrderMainDAO interface {
	SaveOrderMain()
}

// OrderDetailDAO 为订单详细记录
type OrderDetailDAO interface {
	SaveOrderDetail()
}

//抽象模式工厂接口
type DAOFactory interface {
	CreateOrderMainDAO() OrderMainDAO
	CreateOrderDetailDAO() OrderDetailDAO
}

//RDBMainDAP 为关系型数据库的OrderMainDAO实现
type RDBMainDAO struct {
}

func (*RDBMainDAO) SaveOrderMain() {
	fmt.Println("rdb main save")
}

type RDBDetailDAO struct {
}

func (*RDBDetailDAO) SaveOrderDetail() {
	fmt.Println("rdb detail save")
}

//RDBDAOFactory 是RDB 抽象工厂实现
type RDBDAOFactory struct{}

func (*RDBDAOFactory) CreateOrderMainDAO() OrderMainDAO {
	return &RDBMainDAO{}
}

func (*RDBDAOFactory) CreateOrderDetailDAO() OrderDetailDAO {
	return &RDBDetailDAO{}
}

type XMLMainDAO struct {
}

func (*XMLMainDAO) SaveOrderMain() {
	fmt.Print("xml main save\n")
}

//XMLDetailDAO XML存储
type XMLDetailDAO struct{}

// SaveOrderDetail ...
func (*XMLDetailDAO) SaveOrderDetail() {
	fmt.Print("xml detail save")
}

//XMLDAOFactory 是RDB 抽象工厂实现
type XMLDAOFactory struct{}

func (*XMLDAOFactory) CreateOrderMainDAO() OrderMainDAO {
	return &XMLMainDAO{}
}

func (*XMLDAOFactory) CreateOrderDetailDAO() OrderDetailDAO {
	return &XMLDetailDAO{}
}
