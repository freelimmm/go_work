package main

import (
	"Go_project/infra"
	"fmt"
)

func main() {
retriever := infra.Retriever{}
fmt.Println(retriever.Get("http://www.freebuf.com"))
}