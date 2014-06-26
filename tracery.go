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

type A struct {
}

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
				fmt.Println("get sighup\n") //Utils.LogInfo是我自己封装的输出信息函数
				os.Exit(0)
			case syscall.SIGINT:
				os.Exit(0)
			case syscall.SIGQUIT:
				fmt.Println("收到退出信号，准备退出")
				os.Exit(0)
			case syscall.SIGUSR1:
				fmt.Println("usr1\n")
				os.Exit(0)
			case syscall.SIGUSR2:
				fmt.Println("usr2\n")
				os.Exit(0)
			default:
				fmt.Printf("get sig=%v\n", sig)
				os.Exit(0)
			}
		}
	}
}
