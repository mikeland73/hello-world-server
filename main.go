package main
import (
    "fmt"
    "net/http"
)
func hello(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(w, "hello\n")
}

func main() {
    http.HandleFunc("/hello", hello)

    fmt.Println("Serving requests at port 8080\n")
    http.ListenAndServe(":8080", nil)
}