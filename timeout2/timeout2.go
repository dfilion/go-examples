/*
 * Command timeout example #2.
 *
 * Based on http://stackoverflow.com/questions/11886531/terminating-a-process-started-with-os-exec-in-golang
 */

package main

import (
	"fmt"
	"os/exec"
	"time"
)

func TimedExec(cmd *exec.Cmd, secs time.Duration) {
	err := cmd.Start()
	if err != nil {
		panic(err)
	}

	donech := make(chan error, 1)
	go func() {
		donech <- cmd.Wait()
	}()

	select {
	case <-time.After(secs * time.Second):
		cmd.Process.Kill()
		fmt.Printf("Timeout")
	case <-donech:
		fmt.Printf("Done")
	}
}

func main() {

	cmd := exec.Command(`sleep`, `10`)

	fmt.Printf("Should print 'Done': ")
	TimedExec(cmd, 20)

	fmt.Printf("\nShould print 'Timeout': ")
	cmd = exec.Command(`sleep`, `10`)
	TimedExec(cmd, 5)
	fmt.Printf("\n")
}
