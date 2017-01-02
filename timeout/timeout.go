/*
 * Command timeout example #1.
 *
 * Based on http://stackoverflow.com/questions/11886531/terminating-a-process-started-with-os-exec-in-golang
 */

package main

import (
	"fmt"
	"os/exec"
	"time"
	"log"
)


func someCommand(c chan string) {
	cmd := exec.Command("sleep", "5")
	s, err := cmd.Output()
	if err != nil {
		log.Fatalf("%s\n", err)
	}
	c <- string(s)
	
}

func timeout(c chan bool, s time.Duration) {
	time.Sleep(time.Second * s)
	c <- true
}

func main () {

	finished := make (chan string, 1)

	go someCommand(finished)
	//go timeout(timedout, 15)
	
	select {
	case <- time.After(time.Second * 15):
		fmt.Println("Timed out")
	case <- finished:
		fmt.Println("Done")
	}
}