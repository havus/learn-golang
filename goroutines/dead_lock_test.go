package learn_goroutines

import (
	"fmt"
	"testing"
	"time"
	"sync"
)

type UserBalance struct {
	sync.Mutex
	Name string
	Balance int
}

func (userBalance *UserBalance) Lock() {
	userBalance.Mutex.Lock()
}

func (userBalance *UserBalance) Unlock() {
	userBalance.Mutex.Unlock()
}

func (userBalance *UserBalance) Change(amount int) {
	userBalance.Balance += amount
}

func Transfer(fromUser *UserBalance, toUser *UserBalance, amount int) {
	fromUser.Lock()
	fromUser.Change(-amount)
	fmt.Println(fromUser.Name, " Locked")

	time.Sleep(1 * time.Second)

	toUser.Lock() // this locking is root cause deadlock
	toUser.Change(amount)
	fmt.Println(toUser.Name, " Locked")

	time.Sleep(1 * time.Second)
	fromUser.Unlock()
	toUser.Unlock()
	fmt.Println("Unlocked both")
}

func TestDeadLock(t *testing.T) {
	andi := UserBalance{
		Name: "andi",
		Balance: 10000,
	}
	budi := UserBalance{
		Name: "budi",
		Balance: 10000,
	}

	go Transfer(&andi, &budi, 1000)
	go Transfer(&budi, &andi, 2000)

	time.Sleep(3 * time.Second)

	fmt.Println("andi.Balance = ", andi.Balance)
	fmt.Println("budi.Balance = ", budi.Balance)
}
