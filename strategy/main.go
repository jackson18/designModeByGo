package main

import "fmt"

// 定义一个策略接口
type PayStrategy interface {
	Pay(account string, money int) bool
}

// 定义一个上下文，并引用策略接口
type PayContext struct {
	strategy PayStrategy
}

// 上下文实现策略接口，实现根据上下文中的策略来具体调用
func (p *PayContext) Pay(account string, money int) bool {
	return p.strategy.Pay(account, money)
}

// 具体实现，并实现策略接口
type Weixin struct {
}

func (w *Weixin) Pay(account string, money int) bool {
	fmt.Printf("Pay %d 元 to %s by weixin\n", money, account)
	return true
}

// 具体实现，并实现策略接口
type Ali struct {
}

func (a *Ali) Pay(account string, money int) bool {
	fmt.Printf("Pay %d 元 to %s by Ali \n", money, account)
	return true
}

// 上下文对外暴露入口
func NewPay(strategy PayStrategy) *PayContext {
	return &PayContext{strategy: strategy}
}

func main() {
	weixinPay := NewPay(&Weixin{})
	weixinPay.Pay("张三", 100)

	aliPay := NewPay(&Ali{})
	aliPay.Pay("马云", 200)
}
