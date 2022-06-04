package main

import (
    "hello/sub" // subは'package'で指定したパッケージ名
    "fmt"
)

func main() {
    fmt.Printf("hello, world\n")
    sub.Hello()
}
