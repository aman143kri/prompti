package main

import (
    "fmt"
    "os/exec"
)

func main() {
    cmd := exec.Command("touch", "hello.txt")
    err := cmd.Run()
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println("hello.txt created")
    }
}
