package main

import (
	"fmt"
	"time"
	"sync"
)

/*
银行账户案列
1：创建银行账户类
2：存取款需要保证并发安全（不允许并发访问余额）
3：查询余额和打印流水可以并发操作
4：创建账户实例，并发执行存取款，查询余额，打印流水操作
*/

type Account struct {
	money int
	mu    sync.Mutex
}

//存钱
func (a *Account) SaveMoney(n int) {
	fmt.Println("save money start...")
	<-time.After(time.Second * 3)

	a.mu.Lock()
	a.money += n
	a.mu.Unlock()

	fmt.Println("save money finish...")
}

//取钱
func (a *Account) GetMoney(n int) {
	fmt.Println("get money start...")
	<-time.After(time.Second * 3)
	a.mu.Lock()
	a.money -= n
	a.mu.Unlock()
	fmt.Println("get money finish...")
}

//查询余额
func (a *Account) QueryBalance() {
	fmt.Println("query money start...")
	<-time.After(time.Second * 3)
	fmt.Println("当前的余额是: ", a.money)
	fmt.Println("query money finish...")
}

func (a *Account) PrintHistory() {
	//fmt.Println("print log start...")
	<-time.After(time.Second * 3)
	fmt.Println("打印流水-", a.money)
	//fmt.Println("print log start...")
}

//取钱
func main() {

	var wg sync.WaitGroup

	account := new(Account)
	account.money = 27000
	account.mu = sync.Mutex{}

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			account.SaveMoney(100)
			wg.Done()
		}()
	}

	//
	for j := 0; j < 3; j++ {
		wg.Add(1)
		go func() {
			account.GetMoney(100)
			wg.Done()
		}()
	}

	//
	for x := 0; x < 3; x++ {
		wg.Add(1)
		go func() {
			account.QueryBalance()
			wg.Done()
		}()
	}

	//
	for y := 0; y < 3; y++ {
		wg.Add(1)
		go func() {
			account.PrintHistory()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("main over")

}
