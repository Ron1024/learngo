package main

import (
	"fmt"
	retriever "learngo/retriver/download"
	retrieverImpl "learngo/retriver/download/impl"
	real "learngo/retriver/real"
)

func main() {
	var r retriever.Retriever

	r = retrieverImpl.Retriever{"This is a fake retriever!"}
	r = real.Retriever{}

	fmt.Println(Download(r))
}

func Download(r retriever.Retriever) string {
	return r.Get("http://www.baidu.com")
}
