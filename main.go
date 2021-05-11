package main

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
		Cloneflags: syscall.CLONE_NEWUTS |
					//pid:用来隔离进程 PID 号
					syscall.CLONE_NEWPID |
					//mount:用来隔离各个进程看到的挂载点视图
					syscall.CLONE_NEWNS  |
					//user:用来隔离用户
					syscall.CLONE_NEWUSER|
					//network:用来隔离网络
					syscall.CLONE_NEWNET |
					//ipc:用来隔离 System V IPC 和 POSIX message queues
					syscall.CLONE_NEWIPC,
		//设置容器的UID
		UidMappings: []syscall.SysProcIDMap{
			{
				//容器UID
				ContainerID: 1,
				//宿主机UID
				HostID: 0,
				Size: 1,
			},
		},
		//设置容器的GID
		GidMappings: []syscall.SysProcIDMap{
			{
				//容器GID
				ContainerID: 1,
				//宿主机GID
				HostID: 0,
				Size: 1,
			},
		},

	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err!=nil{
		log.Fatal(err)
	}
}
