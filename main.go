package main

import (
	"Go_project/infra"
	"fmt"
)

func getRetriever() retriever{
	return infra.Retriever{}
}

type retriever interface {
	Get(string) string
}

func main() {
	var r  = getRetriever()
	fmt.Println(r.Get("http://www.freebuf.com"))
}