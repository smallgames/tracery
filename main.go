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

const (
	KB = 1 * 1024
	MB = KB * 1024
	GB = MB * 1024
	TB = GB * 1024
)

var (
	GS_HOME = os.Getenv("GS_HOME")

	CONF_PATH = "/conf/app.conf"
)

func main() {
	if len(strings.TrimSpace(GS_HOME)) < 1 {
		GS_HOME, _ = os.Getwd()
	}
	fmt.Println("GS_HOME:", GS_HOME)

	//log, err := lib.NewLog("app", GS_HOME+"/log/app.log", lib.LOG_DEBUG)
	//if err != nil {
	//	fmt.Println("c log e:", err)
	//}

	//(*log).Info("start load Configure...")

	conf, err := lib.NewConf(GS_HOME + CONF_PATH)
	if err != nil {
		fmt.Println(err)
	}

	task_gs, err := lib.NewTask("iGS")
	if err != nil {
		fmt.Errorf("Fork thread failed %v \n", err)
		os.Exit(0)
	}

	//(*log).Info("start load gs...")

	if gs_svr, err := gs.NewGS(conf.AssertInt("gs.port"), 100, task_gs); err == nil {
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
