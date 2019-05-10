package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Score int    `json:"score"`
}

func main() {
	stu := Student{
		Name:  "stu01",
		Age:   18,
		Score: 90,
	}

	data, err := json.Marshal(stu)
	if err != nil {
		fmt.Println("Json encode stu failed,err:", err)
		return
	}
	fmt.Println(string(data))
}
