package main

import "fmt"

func fibonacci() func() int {
	fi := make([]int, 0)
	return func() int {
		if len(fi) == 0 {
			fi = append(fi, 0)
			return 0
		} else if len(fi) == 1 {
			fi = append(fi, 1)
			return 1
		}
		fi = append(fi, fi[len(fi) - 1] + fi[len(fi) - 2])
		return fi[len(fi) - 1]
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
