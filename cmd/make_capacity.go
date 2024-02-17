package main

import "fmt"

func main() {
	ss := make([]byte, 250) // reserve 250 bytes: 0-249
	prn(&ss, "ss - original")

	ss = ss[:len(ss)-100] // drop last 100 elements of array
	prn(&ss, "ss - reduced ?")
}

func prn(slice *[]byte, name string) {
	fmt.Printf("Lenght of slice [%s] = %d\n", name, len(*slice))
	fmt.Printf("    Capacity of [%s] = %d\n", name, cap(*slice))
	fmt.Printf("\n")
}

/**
Lenght of slice [ss - original] = 250
    Capacity of [ss - original] = 250

Lenght of slice [ss - reduced ?] = 150
    Capacity of [ss - reduced ?] = 250
*/
