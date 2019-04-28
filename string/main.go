package main

import (
	"fmt"
)

func main() {
	fmt.Println(`
	床前明月光，疑是地上霜。
	举头望明月，我是郭德纲。`)

	b := 'b'

	fmt.Println(string(b))
	fmt.Printf("%c\n", b)

	fmt.Println(reverse("Hello world!"))

}

func reverse(str string) string {
	var severs string
	for i := 0; i < len(str); i++ {
		severs += fmt.Sprintf("%c", str[len(str)-i-1])
	}
	return severs
}
