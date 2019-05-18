package main

import (
	"log"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	// 用来指定被fork出来的新进程内的初始命令, 默认使用sh来执行.
	cmd := exec.Command("sh")
	cmd.SysProcAttr = &syscall.SysProcAttr{
		// 使用clone_newuts来创建一个新的 UTS Namespace
		Cloneflags: syscall.CLONE_NEWUTS,
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
