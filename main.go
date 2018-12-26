package main

import (
    "url-shortener/router"
    "fmt"
)

func main() {
    fmt.Println("Ready on port 8080")
    a := router.App{}
    a.Initialize()
    a.Run(":8080")
}
