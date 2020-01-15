package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	r, _, _ := reader.ReadLine()
	N, _ := strconv.Atoi(string(r))

	nums := make(map[string]int, 10000)

	var num string
	for i := 0; i < N; i++ {
		r, _, _ := reader.ReadLine()
		num := string(r)
		nums[num] = nums[num] + 1
	}
	var v int
	for i := 1; i < 10001; i++ {
		num = strconv.Itoa(i)
		v = nums[num]
		if v > 0 {
			for i := 0; i < v; i++ {
				writer.WriteString(num)
				writer.WriteByte('\n')
			}
		}
	}
}
