// learn project main.go
package main

import (
	"fmt"
	"learn/lib"
	"os"
	"strings"
)

var (
	GS_HOME = os.Getenv("GS_HOME")
	APP_LOG *lib.Log
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
	(*APP_LOG).Debug("debug hahaha")
	(*APP_LOG).Info("debug hahaha")
	(*APP_LOG).Warn("debug hahaha")
	(*APP_LOG).Error("debug hahaha")

}
