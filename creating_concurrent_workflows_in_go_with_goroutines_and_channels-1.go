package main

import (
	"fmt"
	"time"
)

func myfun1(mychnl1 chan string) {

	for v := 0; v < 4; v++ {
		time.Sleep(100 * time.Millisecond)
		mychnl1 <- "GeeksforGeeks"
	}
	close(mychnl1)
}
func myfun2(mychnl2 chan string) {

	for v := 0; v < 4; v++ {
		time.Sleep(100 * time.Millisecond)
		mychnl2 <- "tutorialspoint"
	}
	close(mychnl2)
}
func main() {
	fmt.Println("!...Main Go-routine Start...!")
	c1 := make(chan string)
	c2 := make(chan string)
	go myfun1(c1)
	go myfun2(c2)
	for {
		res, ok := <-c1
		if ok == false {
			fmt.Println("Channel1 Close ", ok)
			break
		}
		fmt.Println("Channel1 Open ", res, ok)
	}
	for {
		res, ok := <-c2
		if ok == false {
			fmt.Println("Channel2 Close ", ok)
			break
		}
		fmt.Println("Channel2 Open ", res, ok)
	}
	time.Sleep(3500 * time.Millisecond)
	fmt.Println("\n!...Main Go-routine End...!")
}
