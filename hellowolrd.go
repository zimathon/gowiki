package main

import (
    "fmt"
    "net"
    "os"
    "time"
)

func main() {
    fmt.Println("welcome to hte playground!")
    fmt.Println(time.Now())
    fmt.Println(os.Open("filename"))
    fmt.Println(net.Dial("tcp", "google.com"))
}
