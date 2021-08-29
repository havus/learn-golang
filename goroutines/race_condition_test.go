package learn_goroutines

import (
	"fmt"
	"testing"
	"time"
	"strconv"
	"sync"
)

func TestRaceCondition(t *testing.T) {
	counter := 0

	for i := 0; i < 1000; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				counter++
			}
		}()
	}

	time.Sleep(4 * time.Second)
	fmt.Println("Finish looping at " + strconv.Itoa(counter))
	time.Sleep(1 * time.Second)
	fmt.Println("Finish looping at " + strconv.Itoa(counter))
	// === RUN   TestRaceCondition
	// Finish looping at 82801
	// Finish looping at 82801
	// --- PASS: TestRaceCondition (5.01s)
}

// Mutual exclusion
func TestMutex(t *testing.T) {
	counter := 0
	var mutex sync.Mutex

	for i := 0; i < 1000; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				mutex.Lock()
				counter++
				mutex.Unlock()
			}
		}()
	}

	time.Sleep(4 * time.Second)
	fmt.Println("Finish looping at " + strconv.Itoa(counter))
	time.Sleep(1 * time.Second)
	fmt.Println("Finish looping at " + strconv.Itoa(counter))
	// === RUN   TestRaceConditionWithMutex
	// Finish looping at 100000
	// Finish looping at 100000
	// --- PASS: TestRaceConditionWithMutex (5.01s)
}

type BankAccount struct {
	RWMutex sync.RWMutex
	Money int
}

func (account *BankAccount) AddBalance(amount int) {
	// Final balance  3643769 without locking
	// Final balance  4951000 with locking
	account.RWMutex.Lock()
	account.Money += amount
	account.RWMutex.Unlock()
}

func (account *BankAccount) GetBalance() int {
	account.RWMutex.RLock()
	balance := account.Money
	account.RWMutex.RUnlock()

	return balance
}

func TestUseRWMutex(t *testing.T) {
	myAccount := BankAccount{ Money: 1000 }

	fmt.Println("Beginning balance ", myAccount.GetBalance())

	for i := 0; i < 1000; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				myAccount.AddBalance(1)
				fmt.Println("Current balance ", myAccount.GetBalance())
			}
		}()
	}

	time.Sleep(3 * time.Second)
	fmt.Println("Final balance ", myAccount.GetBalance())
}
