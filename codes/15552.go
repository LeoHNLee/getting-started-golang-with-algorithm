package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<14)
	writer := bufoi.NewWriterSize(os.Stdout, 1<<14)
	defer writer.Flush()
	reader.ReadLine()
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		ret := calc(line)
		writer.WriteString(strconv.Itoa(ret))
		writer.WriteByte('\n')
	}
}

func calc(line []byte) int {
	tmp := 0
	ans := 0
	for _, v := range line {
		if v == ' ' {
			ans = tmp
			tmp = 0
			continue
		}
		tmp *= 10
		tmp += int(v - '0')
	}
	return tmp + ans
}
