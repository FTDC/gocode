package main

var (
	sum int
)

func eat(n int) int {
	for i := 0; i <= n; i++ {
		eat(sum / 2)
	}

	return sum
}

func main() {
	// 每天吃一半

}
