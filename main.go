// learn project main.go
package main

import (
	"fmt"
	"learn/lib"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
	//"regexp"
)

var (
	GS_HOME = os.Getenv("GS_HOME")

	CONF_PATH = "/conf/app.conf"

	APP_LOG  *lib.Log
	APP_CONF *lib.Conf
)

func main() {

	if len(strings.TrimSpace(GS_HOME)) < 1 {
		GS_HOME, _ = os.Getwd()
	}
	fmt.Println("GS_HOME:", GS_HOME)

	APP_LOG, err := lib.NewLog(GS_HOME+"/log/app.log", lib.LOG_DEBUG)
	if err != nil {
		fmt.Println("c log e:", err)
	}

	(*APP_LOG).Info("start load Configure...")

	dir, _ := os.Getwd()
	APP_CONF, err := lib.NewConf(dir + CONF_PATH)
	if err != nil {
		fmt.Println(err)
	}

	t, err := lib.NewTask()
	fmt.Println(t)
	if err != nil {
		fmt.Errorf("Init handle thread failed %v \n", err)
		os.Exit(0)
	}

	lib.Fork(t, func() {
		for {
			select {
			case <-time.After(time.Second * 5):
				fmt.Println(APP_CONF)
			}
		}
	})

	ch := make(chan os.Signal)
	for {
		signal.Notify(ch, syscall.SIGINT, syscall.SIGUSR1, syscall.SIGUSR2, syscall.SIGHUP)

		select {
		case sig := <-ch:
			APP_LOG.Info("Server receives the signal(%v) will now exit.", sig)
			break
		}
		//switch sig {
		//default:
		//	fmt.Println("get sig=%v\n", sig)
		//case syscall.SIGHUP:
		//	fmt.Println("get sighup\n") //Utils.LogInfo是我自己封装的输出信息函数
		//case syscall.SIGINT:
		//	os.Exit(1)
		//case syscall.SIGUSR1:
		//	fmt.Println("usr1\n")
		//case syscall.SIGUSR2:
		//	fmt.Println("usr2\n")
		//}
	}
}
