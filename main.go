// learn project main.go
package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	//"time"
	"tracery/gs"
	"tracery/lib"
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

	APP_CONF, err := lib.NewConf(GS_HOME + CONF_PATH)
	if err != nil {
		fmt.Println(err)
	}

	for k, v := range APP_CONF.Stroes {
		fmt.Printf("key=%s,value=%v\n", k, v)
	}

	task_gs, err := lib.NewTask(gs.DEF_HANDLE)
	if err != nil {
		fmt.Errorf("Init handle thread failed %v \n", err)
		os.Exit(0)
	}

	if gs_svr, err := gs.NewGS("GS", APP_CONF.AssertInt("gs.port"), 100, task_gs); err != nil {
		lib.Run(gs_svr)
	}

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
