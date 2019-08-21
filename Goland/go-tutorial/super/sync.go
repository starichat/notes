/**
一个线性程序中，程序执行顺序只由程序的逻辑来决定。
**/

package main

import "sync"

var (
	mutex   sync.Mutex
	balance int
)

func Deposit(amount int) {
	mutex.Lock()
	balance = balance + amount
	mutex.Unlock()
}

func Balance() int {
	mutex.Lock()
	b := balance
	mutex.Unlock()
	return b
}

func Withdraw(amount int) bool {
	mutex.Lock()
	defer mutex.Unlock()
	Deposit(-amount)
	if Balance() < 0 {
		Deposit(amount)
		return false // insufficient funds
	}
	return true
}

func main() {

	Deposit(100)
	Withdraw(100)
}
