package main

import "fmt"

type PayStrategy interface {
	Pay(account string, money int) bool
}

type PayContext struct {
	strategy PayStrategy
}

func (p *PayContext) Pay(account string, money int) bool {
	return p.strategy.Pay(account, money)
}

type Weixin struct {
}

func (w *Weixin) Pay(account string, money int) bool {
	fmt.Printf("Pay %d 元 to %s by weixin\n", money, account)
	return true
}

type Ali struct {
}

func (a *Ali) Pay(account string, money int) bool {
	fmt.Printf("Pay %d 元 to %s by Ali \n", money, account)
	return true
}

func NewPay(strategy PayStrategy) *PayContext {
	return &PayContext{strategy: strategy}
}

func main() {
	weixinPay := NewPay(&Weixin{})
	weixinPay.Pay("张三", 100)

	aliPay := NewPay(&Ali{})
	aliPay.Pay("马云", 200)
}
