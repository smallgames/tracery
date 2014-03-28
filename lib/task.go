package lib

import (
	"fmt"
	"os"
)

func init() {
	fmt.Println("Initial task model")
}

type Runable interface {
	Run()
}

type Task struct {
	Name string

	ppid int
	pid  int
	gid  int
	sig  chan int
}

func NewTask(n string) (*Task, error) {
	return &Task{Name: n, ppid: os.Getppid(), pid: os.Getpid(), gid: os.Getgid(), sig: make(chan int)}, nil
}
