package lib

import (
	"fmt"
	"os"
)

func init() {
	fmt.Println("Load task model")
}

type Handle func()

type runable interface {
	run(*Handle)
}

type Task struct {
	ppid int
	pid  int
	gid  int
	//live int64

	sig chan int
}

func NewTask() (*Task, error) {
	return &Task{ppid: os.Getppid(), pid: os.Getpid(), gid: os.Getgid(), sig: make(chan int)}, nil
}

func (self *Task) run(h Handle) {
	//defer func() {
	//	if r := recover(); r != nil {
	//		fmt.Printf("pid:%v\n recover", self.pid)
	//	}
	//}()
	go h()
}

func Fork(t *Task, h Handle) {
	t.run(h)
}
