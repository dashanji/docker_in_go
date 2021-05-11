# docker_in_go
A simple implementation of docker in Golang.
The project is mostly based on https://github.com/pibigstar/go-docker
How to run ?
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build ./main.go
After this command,you need to move <main> to Linux
1、Test hostname
![img.png](img.png)
2、Test PID
![img_2.png](img_2.png)
3、Test Mount
![img_4.png](img_4.png)
![img_5.png](img_5.png)
4、Test User
![img_6.png](img_6.png)
5、Test Network
![img_3.png](img_3.png)
6、Test IPC
![img_1.png](img_1.png)