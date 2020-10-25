package main

import "fmt"

func smallChange(small, change string) (string, error) {
	// some logic to play with small and change
	if small == "" {
		return "", fmt.Errorf("smallChange: arg1 is empty")
	}
	return small + " " + change, nil
}

func main() {
	fmt.Println("small change")
	fmt.Println(smallChange("practice", "yourself"))
}
