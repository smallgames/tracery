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

type A struct {
}

func main() {

	//amqp begin
	//mqconn, err := amqp.Dial("amqp://guest:guest@192.168.56.101:5672/")
	//if err != nil {
	//	fmt.Errorf("Dial: %s", err)
	//}
	//defer mqconn.Close()

	//channel, err := mqconn.Channel()
	//if err != nil {
	//	fmt.Errorf("Channel: %s", err)
	//}

	//if err := channel.ExchangeDeclare(
	//	"yate_ex1", // name
	//	"direct",   // type direct|fanout|topic|x-custom
	//	true,       // durable
	//	false,      // auto-deleted
	//	false,      // internal
	//	false,      // noWait
	//	nil,        // arguments
	//); err != nil {
	//	fmt.Errorf("Exchange Declare: %s", err)
	//}

	//if err := channel.Confirm(false); err != nil {
	//	fmt.Errorf("Channel could not be put into confirm mode: %s", err)
	//}
	//ack, nack := channel.NotifyConfirm(make(chan uint64, 1), make(chan uint64, 1))

	//for {
	//	if err = channel.Publish(
	//		"yate_ex1", // publish to an exchange
	//		"yate_rk",  // routing to 0 or more queues
	//		false,      // mandatory
	//		false,      // immediate
	//		amqp.Publishing{
	//			Headers:         amqp.Table{},
	//			ContentType:     "text/plain",
	//			ContentEncoding: "",
	//			Body:            []byte("yate_test_mq"),
	//			DeliveryMode:    amqp.Persistent, // 1=non-persistent, 2=persistent
	//			Priority:        0,               // 0-9
	//			// a bunch of application/implementation-specific fields
	//		},
	//	); err != nil {
	//		fmt.Errorf("Exchange Publish: %s", err)
	//	}
	//	select {
	//	case tag := <-ack:
	//		log.Printf("confirmed delivery with delivery tag: %d", tag)
	//	case tag := <-nack:
	//		log.Printf("failed delivery of delivery tag: %d", tag)
	//	}
	//	time.Sleep(5 * time.Second)
	//}
	//os.Exit(0)

	if len(strings.TrimSpace(GS_HOME)) < 1 {
		GS_HOME, _ = os.Getwd()
	}
	fmt.Println("GS_HOME:", GS_HOME)

	//log, err := lib.NewLog("app", GS_HOME+"/log/app.log", lib.LOG_DEBUG)
	//if err != nil {
	//	fmt.Println("c log e:", err)
	//}

	conf, err := lib.NewConf(GS_HOME + CONF_PATH)
	if err != nil {
		fmt.Println(err)
	}

	task_gs, err := lib.NewTask("iGS")
	if err != nil {
		fmt.Errorf("Fork thread failed %v \n", err)
		os.Exit(0)
	}

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

			switch sig {
			case syscall.SIGHUP:
				fmt.Println("get sighup\n") //Utils.LogInfo是我自己封装的输出信息函数
			case syscall.SIGINT:
				os.Exit(1)
			case syscall.SIGQUIT:
				fmt.Println("收到退出信号，准备退出")
			case syscall.SIGUSR1:
				fmt.Println("usr1\n")
			case syscall.SIGUSR2:
				fmt.Println("usr2\n")
			default:
				fmt.Printf("get sig=%v\n", sig)
			}
		}

	}
}
