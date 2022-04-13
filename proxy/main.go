package main

import "fmt"

type UserDao interface {
	addUser() bool
}

type UserDaoImpl struct {
}

func (u *UserDaoImpl) addUser() bool {
	fmt.Println("userDaoImpl execute...")
	return true
}

type Proxy struct {
	u UserDaoImpl
}

func (p *Proxy) addUser() bool {
	fmt.Println("proxy before execute ....")
	p.u.addUser()
	fmt.Println("proxy after execute ...")
	return true
}

func main() {
	u := UserDaoImpl{}
	p := Proxy{u: u}
	p.addUser()
}
