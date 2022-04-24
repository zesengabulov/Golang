package main

import (
	"fmt"
)

func a(myChannel chan bool) {
	for i := 0; i < 50; i++ {
		fmt.Print("a")
	}
	myChannel <- true
}

func b(myChannel chan bool) {
	for i := 0; i < 50; i++ {
		fmt.Print("b")
	}
	// time.Sleep(time.Millisecond)
	myChannel <- false
}

func main() {
	var myChannel chan bool = make(chan bool)
	go a(myChannel)
	go b(myChannel)
	// time.Sleep(time.Second)
	fmt.Println(<-myChannel)
	fmt.Println(<-myChannel)
	fmt.Println("end main()")
}
