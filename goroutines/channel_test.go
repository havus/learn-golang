package learn_goroutines

import (
	"os"
	"os/exec"
	"runtime"
	"fmt"
	"testing"
	"time"
	"strconv"
)

func CallClear() {
	var clear map[string]func() // create a map for storing clear funcs

	clear = make(map[string]func()) // Initialize it
	clear["darwin"] = func() { 
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}

	value, ok := clear[runtime.GOOS] // runtime.GOOS -> linux, windows, darwin etc.
	if ok { // if we defined a clear func for that platform:
			value()  // we execute it
	} else { // unsupported platform
			panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}





func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Hello world"
		fmt.Println("finished send data to channel")
	}()

	data := <- channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}

// CHANNEL AS PARAMS

func GimmeResponse(channel chan string) {
	time.Sleep(1 * time.Second)
	channel <- "This is the response"
}

func TestChannelAsParamater(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GimmeResponse(channel)

	data := <- channel
	fmt.Println(data)
	fmt.Println("data")

	time.Sleep(5 * time.Second)
}

// CHANNEL IN OUT
func OnlyIn(channel chan<- string) {
	channel <- "Hello only in!"
}

// func OnlyOut(channel <-chan string) string {
// 	return <- channel
// }
func OnlyOut(channel <-chan string) {
	time.Sleep(2 * time.Second)
	data := <- channel
	fmt.Println(data)
}

func TestChannelOutIn(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	fmt.Println("waiting 5 seconds")
	time.Sleep(5 * time.Second)
}

// CHANNEL BUFFER
func TestBufferChannel(t *testing.T) {
	channel := make(chan string, 2)
	defer close(channel)

	fmt.Println("length", len(channel))
	fmt.Println("capacity", cap(channel))

	go func() {
		channel <- "Hello"
		channel <- "World"
	}()
	go func() {
		fmt.Println(<- channel)
		fmt.Println(<- channel)
	}()

	fmt.Println("waiting 3 seconds")
	time.Sleep(3 * time.Second)
}

// CHANNEL RANGE
func TestRangeChannel(t *testing.T) {
	channel := make(chan string)
	
	go func() {
		defer close(channel)

		for i := 0; i < 10; i++ {
			channel <- "In looping " + strconv.Itoa(i)
		}
	}()

	for data := range channel {
		fmt.Println("Receiving data <- ", data)
	}

	fmt.Println("waiting 3 seconds")
	time.Sleep(3 * time.Second)
}

// CHANNEL SELECT
func TestMultipleChannel(t *testing.T) {
	channel 			:= make(chan string)
	secondChannel := make(chan string)
	
	go GimmeResponse(channel)
	go GimmeResponse(secondChannel)

	var counter int = 0

	for {
		select {
		case data := <- channel:
			fmt.Println("data from channel 1", data)
			counter++
		case data := <- secondChannel:
			fmt.Println("data from channel 2", data)
			counter++
		default:
			CallClear()
			fmt.Println("waiting the data")
		}

		if counter == 2 {
			break
		}
	}
 
	fmt.Println("waiting 3 seconds")
	time.Sleep(3 * time.Second)
}
