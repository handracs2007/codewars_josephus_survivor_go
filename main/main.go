package main

import "fmt"

func JosephusSurvivor(n, k int) int {
	if n == 1 {
		return 1
	} else {
		return (JosephusSurvivor(n-1, k)+k-1)%n + 1
	}
}

func main() {
	fmt.Println(JosephusSurvivor(300, 300))
}
