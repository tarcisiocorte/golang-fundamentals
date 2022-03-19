## Compiling (go build) and Running Go Applications (go run)

1. go run: it compiles and runs the application. it doesn't produce an executable
		go run file.go compiles and immediately runs Go programs

2. go build: it just compiles the application. it produces an executable
		**go build file.go** compiles a bunch of go source files. It compiles packages and dependencies.
		If you run **go build** it will compile the files in the current directory and will produce an executable file with the name of the current working directory.
		**go build -o app** will produce an executable file called app

## Compiling (go build) and Running Go Applications (go run)

 - Compiling for Windows: ***GOOS=windows GOARCH=amd64 go build -o
   winapp.exe*** 
 - Compiling for Linux: ***GOOS=linux GOARCH=amd64 go build -o linuxapp***
 - Compiling for Mac: ***GOOS=darwin GOARCH=amd64 go build -o macapp***
