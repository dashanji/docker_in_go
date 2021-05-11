package docker_in_go

import (
"log"
"os"
"os/exec"
"syscall"
)
func main(){
	cmd := exec.Command("sh")

	cmd.SysProcAttr = &syscall.SysProcAttr{
		//uts: 用来隔离主机名
		Cloneflags: syscall.CLONE_NEWUTS,
		//ipc：用来隔离 System V IPC 和 POSIX message queues
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err!=nil{
		log.Fatal(err)
	}
}
