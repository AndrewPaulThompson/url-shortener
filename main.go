package main

import (
    "url-shortener/app"
    "fmt"
)

// Main function
func main() {
    fmt.Println("Ready on port 8080")
    a := app.App{}
    a.Initialise()
    a.Run(":8080")
}
