package main

import "fmt"

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println("arr[2:6] = ", arr[2:6])
	fmt.Println("arr[2:6] = ", arr[2:])
	fmt.Println("arr[2:6] = ", arr[:6])
	fmt.Println("arr[2:6] = ", arr[:])

	//Reslice
	fmt.Println("Resilce")
	s2 := arr[:]
	fmt.Println(s2)
	s2 = s2[:5]
	fmt.Println(s2)
	s2 = s2[2:]
	fmt.Println(s2)

	//Extending Slice

	fmt.Println("Extending slice")
	s3 := arr[2:6]
	s4 := s3[3:5]
	fmt.Println(s2)
	fmt.Println(s4)

	fmt.Println("Append slice ")
	s5 := append(s4, 10)
	s6 := append(s5, 11)
	s7 := append(s6, 12)
	fmt.Println("s3,s4,s5 = ", s5, s6, s7)
	//s6 and s7 no longer view arr
	fmt.Println("arr=", arr)

}
