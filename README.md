# devops-project-1

We have created a simple api application using golang. Port exposed is 8080.
code is in main.go file.

We have used github.com/gorilla/mux library to create a router.

Below basic commands are required to run this code.

go mod init main
go mod tidy
go build

now to run code execute main file created by above commands.
command to start execution ./main


We have created a dockerfile to containerise this go program.
commands to build docker container

login to docker if you have not logged in.

`docker build -t <useraccount/image-name> .`

`docker push <useraccount/image-name>`

