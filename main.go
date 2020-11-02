package main

import (
	"github.com/linhongzhao321/tokens/http"
)

func main() {
	r := http.NewRouter()
	err := r.Run(":8080")
	panic(err)
}
