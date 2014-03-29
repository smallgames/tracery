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
	//b1, _ := gs.NewMem("b1", gs.BYTES_MEM)
	//b2, _ := gs.NewMem("b2", gs.BYTES_MEM)
	//b3, _ := gs.NewMem("b3", gs.BYTES_MEM)
	//b1.Append(b2)
	//b1.Append(b3)

	//var (
	//	v1, v2, v3 *gs.BytesMem
	//)

	//v1, _ = b1.(*gs.BytesMem)
	//v2, _ = b2.(*gs.BytesMem)
	//v3, _ = b3.(*gs.BytesMem)

	//fmt.Println("\n===================")
	//fmt.Println(v1.Name)
	//if v1.Head != nil {
	//	fmt.Println("h=", v1.Head.Name)
	//}
	//if v1.Pre != nil {
	//	fmt.Println("p=", v1.Pre.Name)
	//}
	//if v1.Next != nil {
	//	fmt.Println("n=", v1.Next.Name)
	//}

	//fmt.Println("\n===================")
	//fmt.Println(v2.Name)
	//if v2.Head != nil {
	//	fmt.Println("h=", v2.Head.Name)
	//}
	//if v2.Pre != nil {
	//	fmt.Println("p=", v2.Pre.Name)
	//}
	//if v2.Next != nil {
	//	fmt.Println("n=", v2.Next.Name)
	//}

	//fmt.Println("\n===================")
	//fmt.Println(v3.Name)
	//if v3.Head != nil {
	//	fmt.Println("h=", v3.Head.Name)
	//}
	//if v3.Pre != nil {
	//	fmt.Println("p=", v3.Pre.Name)
	//}
	//if v3.Next != nil {
	//	fmt.Println("n=", v3.Next.Name)
	//}

	//os.Exit(0)

	if len(strings.TrimSpace(GS_HOME)) < 1 {
		GS_HOME, _ = os.Getwd()
	}
	fmt.Println("GS_HOME:", GS_HOME)

	log, err := lib.NewLog("app", GS_HOME+"/log/app.log", lib.LOG_DEBUG)
	if err != nil {
		fmt.Println("c log e:", err)
	}

	(*log).Info("start load Configure...")

	conf, err := lib.NewConf(GS_HOME + CONF_PATH)
	if err != nil {
		fmt.Println(err)
	}

	task_gs, err := lib.NewTask("iGS")
	if err != nil {
		fmt.Errorf("Fork thread failed %v \n", err)
		os.Exit(0)
	}

	(*log).Info("start load gs...")

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
			log.Info("Server receives the signal(%v) will now exit.", sig)
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
