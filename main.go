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

	signal.Notify(signalChan,
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
	)

	receivedSignal := <-signalChan

	fmt.Printf("received signal: %v\n", receivedSignal)

	fmt.Println("sleeping for 3 seconds")

	time.Sleep(3 * time.Second)

	fmt.Println("exiting")

	os.Exit(0)
}
