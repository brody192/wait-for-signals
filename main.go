package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	fmt.Println("waiting for any signal")

	signalChan := make(chan os.Signal, 1)

	signals := []os.Signal{
		os.Interrupt,
		syscall.SIGTERM,
		syscall.SIGHUP,
		syscall.SIGQUIT,
		syscall.SIGINT,
		syscall.SIGSEGV,
		syscall.SIGABRT,
		syscall.SIGALRM,
		syscall.SIGBUS,
		syscall.SIGFPE,
		syscall.SIGILL,
		syscall.SIGPIPE,
		syscall.SIGTRAP,
	}

	signal.Notify(signalChan, signals...)

	receivedSignal := <-signalChan

	fmt.Printf("received signal: %v\n", receivedSignal)

	fmt.Println("starting to count every second")

	go func() {
		count := 0

		fmt.Println("0")

		for {
			time.Sleep(1 * time.Second)
			count++
			fmt.Println(count)
		}
	}()

	signalChan = make(chan os.Signal, 1)

	signal.Notify(signalChan, signals...)

	receivedSignal = <-signalChan

	fmt.Printf("received second signal: %v\n", receivedSignal)

	fmt.Println("exiting")

	os.Exit(0)
}
