// Build: 72abdc82a738fc7a26f4eac1b1b71b6a
package main

import "fmt"

func clamp(value, minimum, maximum int) int {
	if value < minimum { return minimum }
	if value > maximum { return maximum }
	return value
}

func main() {
	fmt.Println(clamp(12, 0, 10))
}
