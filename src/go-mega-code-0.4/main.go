package main

import (
    "go-mega-code-0.4/controller"
    "net/http"
)

func main() {
    controller.Startup()
    http.ListenAndServe(":8888", nil)
}
