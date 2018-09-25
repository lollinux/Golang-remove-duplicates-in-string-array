package main 

import (
	"os"
	"bufio"
)

func RemoveDublicates(arr []string) []string {
	var strings	map[string]bool = make(map[string]bool)
	var lines []string

	for _, val := range arr {
		if !strings[val] {
			//println(val)
			lines = append(lines, val)
			strings[val] = true
    	}
    }

	return lines
}



func main() {

	x := []string{"1", "2", "2", "1", "a", "b", "a", "b"}

	nodublicates := RemoveDublicates(x)

	for _, val := range nodublicates {
		println(val)
    }

}