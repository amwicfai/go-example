package main

type trace struct{}

func main() {
	slice := make([]string, 2, 4)

	var t trace
	t.Example(slice, "hello", 10)
}

func (t *trace) Example(slice []string, str string, i int) {
	panic("Want stack trace")
}
