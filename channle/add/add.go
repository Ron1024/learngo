package add

func Add(pipe chan int) int {
	a := <-pipe
	b := <-pipe

	return a + b
}

func Clac(a int, b int) (int, int) {
	s := a + b
	d := a / b
	return s, d
}
