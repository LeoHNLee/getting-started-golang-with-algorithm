package bojStarter

import "fmt"

func ExampleMain() {
	var hour int8 = 0
	var minute int8 = 0
	hour, minute = NewAlarm(hour, minute)
	fmt.Printf("%d %d\n", hour, minute)

	hour = 23
	minute = 59
	hour, minute = NewAlarm(hour, minute)
	fmt.Printf("%d %d\n", hour, minute)

	hour = 0
	minute = 45
	hour, minute = NewAlarm(hour, minute)
	fmt.Printf("%d %d", hour, minute)

	// Output:
	// 23 15
	// 23 14
	// 0 0
}
