package main

import (
	"Go_project/infra"
	"fmt"
)

func getRetriever() infra.Retriever{
	return infra.Retriever{}
}

func main() {
	var retriever infra.Retriever = getRetriever()
	fmt.Println(retriever.Get("http://www.freebuf.com"))
}