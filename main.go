package main

import (
    "url-shortener/router"
    "fmt"
)

// Main function
func main() {
    fmt.Println("Ready on port 8080")
    a := router.App{}
    a.Initialise()
    a.Run(":8080")
}
