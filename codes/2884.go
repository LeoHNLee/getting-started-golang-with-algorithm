package bojStarter

import (
	"fmt"
	"io/ioutil"
	"os"
)

// Tune the alarm
func NewAlarm(hour int8, minute int8) (int8, int8) {
	// Tune the alarm before 45minute
	switch {
	case minute >= 45:
		minute = minute - 45
	case hour > 0:
		minute = minute + 15
		hour = hour - 1
	default:
		minute = minute + 15
		hour = 23
	}
	return hour, minute
}

// Read Ints
// https://stackoverflow.com/questions/3751429/reading-an-integer-from-standard-input/22887685#22887685?newreg=d27857c9e9634627b57cd3f4a3876c35
func readInts() []int8 {
	var buf []byte
	buf, _ = ioutil.ReadAll(os.Stdin)
	var ints []int8
	num := int8(0)
	found := false

	for _, v := range buf {
		if '0' <= v && v <= '9' {
			num = 10*num + int8(v-'0')
			found = true
		} else if found {
			ints = append(ints, num)
			found = false
			num = 0
		}
	}
	if found {
		ints = append(ints, num)
	}
	return ints
}

func main() {
	hm := readInts()
	hour, minute := NewAlarm(hm[0], hm[1])
	fmt.Printf("%d %d", hour, minute)
}
