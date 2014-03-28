package lib

import (
	"fmt"
	"os"
)

func init() {
	fmt.Println("Initial task model")
}

type Handle func()

type Task struct {
	ppid   int
	pid    int
	gid    int
	sig    chan int
	handle Handle
}

func NewTask(h Handle) (*Task, error) {
	return &Task{ppid: os.Getppid(), pid: os.Getpid(), gid: os.Getgid(), sig: make(chan int), handle: h}, nil
}

func Run(t *Task) {
	go t.handle()
}
