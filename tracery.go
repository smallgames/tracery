// learn project main.go
package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"tracery/gs"
	"tracery/lib"
)

var (
	GS_HOME   = os.Getenv("GS_HOME")
	CONF_PATH = "/conf/app.conf"
)

func main() {
	if len(strings.TrimSpace(GS_HOME)) < 1 {
		GS_HOME, _ = os.Getwd()
	}
	fmt.Println("GS_HOME:", GS_HOME)

	conf, err := lib.NewConf(GS_HOME + CONF_PATH)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	if gs_svr, err := gs.NewGS(conf.AssertInt("gs.port"), 100); err == nil {
		go gs_svr.Run()
	} else {
		fmt.Errorf("GS startup failed %v \n", err)
		os.Exit(0)
	}

	ch := make(chan os.Signal)
	for {
		signal.Notify(ch, syscall.SIGINT, syscall.SIGUSR1, syscall.SIGUSR2, syscall.SIGHUP)
		select {
		case sig := <-ch:
			fmt.Println("Server receives the signal(%v) will now exit.", sig)
			switch sig {
			case syscall.SIGHUP:
				fmt.Println("SIGHUP")
				os.Exit(0)
			case syscall.SIGINT:
				fmt.Println("SIGINT")
				os.Exit(0)
			case syscall.SIGQUIT:
				fmt.Println("SIGQUIT")
				os.Exit(0)
			case syscall.SIGUSR1:
				fmt.Println("SIGUSR1")
				os.Exit(0)
			case syscall.SIGUSR2:
				fmt.Println("SIGUSR2")
				os.Exit(0)
			
			SIGABRT   = Signal(0x6)
	SIGALRM   = Signal(0xe)
	SIGBUS    = Signal(0xa)
	SIGCHLD   = Signal(0x14)
	SIGCONT   = Signal(0x13)
	SIGEMT    = Signal(0x7)
	SIGFPE    = Signal(0x8)
	SIGILL    = Signal(0x4)
	SIGINFO   = Signal(0x1d)
	SIGIO     = Signal(0x17)
	SIGIOT    = Signal(0x6)
	SIGKILL   = Signal(0x9)
	SIGPIPE   = Signal(0xd)
	SIGPROF   = Signal(0x1b)
	SIGSEGV   = Signal(0xb)
	SIGSTOP   = Signal(0x11)
	SIGSYS    = Signal(0xc)
	SIGTERM   = Signal(0xf)
	SIGTRAP   = Signal(0x5)
	SIGTSTP   = Signal(0x12)
	SIGTTIN   = Signal(0x15)
	SIGTTOU   = Signal(0x16)
	SIGURG    = Signal(0x10)
	SIGUSR1   = Signal(0x1e)
	SIGUSR2   = Signal(0x1f)
	SIGVTALRM = Signal(0x1a)
	SIGWINCH  = Signal(0x1c)
	SIGXCPU   = Signal(0x18)
	SIGXFSZ   = Signal(0x19)
			default:
				fmt.Printf("default sig=%v\n", sig)
				os.Exit(0)
			}
		}
	}
}
