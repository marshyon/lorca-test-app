# lorca-test-app

A test application using Lorca to add a web browser based GUI in Chrome to a go program. This test application has been built and tested on Windows 10 using Powershell and Bash environments but it should likeise work on other supported platforms such as Linux and Mac.

## Prerequisites

Go

a recent version of go

see the go [downloads](https://golang.org/dl/) page if you do not already have this

Chrome

a recent version of the Chrome Browser

see the chrome [download](https://www.google.com/chrome/) page if you do not already have this

## Build / Run

Powershell :
```
go run .\cmd\generate.go ; go build -o lorca-test-app.exe main.go assets.go ; .\lorca-test-app.exe
```

Bash : 
```
go run cmd/generate.go && go build -o lorca-test-app.exe main.go assets.go && ./lorca-test-app.exe
```