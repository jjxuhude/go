package main

func main() {
	m := map[int]int{}
	go func() {
		for {
			m[1] = 1
		}
	}()

	go func() {
		for {
			_ = m[1]
		}
	}()

	var i int
	for {
		i++
	}

}
