package learn_context

import (
	"fmt"
	"testing"
	"context"
	"runtime"
	"sync"
	"time"
)

func TestContext(t *testing.T) {
	background := context.Background()
	fmt.Println(background)

	todo := context.TODO()
	fmt.Println(todo)
}

func TestContextWithValie(t *testing.T) {
	contextA := context.Background()

	contextB := context.WithValue(contextA, "key_b", "value is B")
	contextC := context.WithValue(contextA, "key_c", "value is C")

	contextD := context.WithValue(contextB, "key_d", "value is D")
	contextE := context.WithValue(contextB, "key_e", "value is E")

	contextF := context.WithValue(contextC, "key_f", "value is F")

	fmt.Println(contextA)
	fmt.Println(contextB)
	fmt.Println(contextC)
	fmt.Println(contextD)
	fmt.Println(contextE)
	fmt.Println(contextF)

	fmt.Println("\nhierarki:", contextF)
	fmt.Println(contextF.Value("key_b")) // <nil>
	fmt.Println(contextF.Value("key_c"))
	fmt.Println(contextF.Value("key_f"))
}

func CreateCounterUseGroup(ctx context.Context, group *sync.WaitGroup) chan int {
	destination := make(chan int)

	go func() {
		defer close(destination)
		defer group.Done()

		counter := 0

		for {
			select {
			case <-ctx.Done():
				return
			default:
				destination <- counter
				counter++
				time.Sleep(1 * time.Second)
			}
		}
	}()

	return destination
}

func CreateCounter(ctx context.Context) chan int {
	destination := make(chan int)

	go func() {
		defer close(destination)

		counter := 0

		for {
			select {
			case <-ctx.Done():
				return
			default:
				destination <- counter
				counter++
				time.Sleep(1 * time.Second)
			}
		}
	}()

	return destination
}

func TestWithCancel(t *testing.T) {
	group 				:= sync.WaitGroup{}
	parentContext := context.Background()
	ctx, cancel 	:= context.WithCancel(parentContext)

	fmt.Println("Total goroutines:", runtime.NumGoroutine())

	group.Add(1)
	destination := CreateCounterUseGroup(ctx, &group)

	for n := range destination {
		fmt.Println("counter n:", n)

		if n == 5 {
			break
		}
	}

	fmt.Println("Total goroutines finish at:", runtime.NumGoroutine())
	
	cancel()

	group.Wait() // group will not done before we cancel context
	fmt.Println("Total goroutines after cancel:", runtime.NumGoroutine())
}

func TestWithTimeout(t *testing.T) {
	parentContext := context.Background()
	ctx, cancel 	:= context.WithTimeout(parentContext, 3 * time.Second)

	defer cancel()

	fmt.Println("Total goroutines:", runtime.NumGoroutine())

	destination := CreateCounter(ctx)

	for n := range destination {
		fmt.Println("counter n:", n)
	}

	fmt.Println("Total goroutines finish at:", runtime.NumGoroutine())
	time.Sleep(1 * time.Second)
	fmt.Println("Total goroutines after cancel:", runtime.NumGoroutine())
}

func TestWithDeadline(t *testing.T) {
	parentContext := context.Background()
	ctx, cancel 	:= context.WithDeadline(parentContext, time.Now().Add(3 * time.Second))

	defer cancel()

	fmt.Println("Total goroutines:", runtime.NumGoroutine())

	destination := CreateCounter(ctx)

	for n := range destination {
		fmt.Println("counter n:", n)
	}

	fmt.Println("Total goroutines finish at:", runtime.NumGoroutine())
	time.Sleep(1 * time.Second)
	fmt.Println("Total goroutines after cancel:", runtime.NumGoroutine())
}