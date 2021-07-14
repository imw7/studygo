package main

import (
	"errors"
	"log"
	"net/http"
	"net/rpc"
)

// RPC 乘法运算，商和余数运算

// ArithService 注册用
type ArithService struct{}

// ArithRequest 声明参数
type ArithRequest struct {
	A, B int
}

// ArithResponse 返回给客户端的结果
type ArithResponse struct {
	// 乘积
	Pro int
	// 商
	Quo int
	// 余数
	Rem int
}

// Multiply 乘积
func (a *ArithService) Multiply(req ArithRequest, res *ArithResponse) error {
	res.Pro = req.A * req.B
	return nil
}

// Divide 商
func (a *ArithService) Divide(req ArithRequest, res *ArithResponse) error {
	if req.B == 0 {
		return errors.New("除数不能为0")
	}
	res.Quo = req.A / req.B
	return nil
}

// Remainder 取余数
func (a *ArithService) Remainder(req ArithRequest, res *ArithResponse) error {
	if req.B == 0 {
		return errors.New("除数不能为0")
	}
	res.Rem = req.A % req.B
	return nil
}

func main() {
	err := rpc.RegisterName("ArithService", new(ArithService))
	if err != nil {
		log.Fatal("Register error:", err)
	}

	rpc.HandleHTTP()
	err = http.ListenAndServe(":1234", nil)
	if err != nil {
		log.Fatal(err)
	}
}
