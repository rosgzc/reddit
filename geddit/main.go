package main 

import (
    "log"
    "fmt"
    "github.com/rossgzc/reddit"
)

func main() {
    items, err := reddit.Get("dachshund")
    if err != nil  {
        log.Fatal(err)
    }
    for _, item := range items {
        fmt.Println(item)
    }
}
