### Begineers guide to GO

If you want to run a simple GO program, just create a main.go file like this.

```go
package main

func main() {
	println("Hello world")
}
```

After creating this file type the following command in the terminal and you will see the Hello World text being printed out.

```console
foo@bar:~$ go run main.go
Hello world
```

This 'run' command of the go-cli, just does an extra step for us, it compiles the file to a binary and executes it.

We can do it like this.

```console
foo@bar:~$ go build -o main main.go
foo@bar:~$ ./main
Hello World
```
