package main 

import (
	"os"
	"bufio"
)

func RemoveDuplicates(arr []string) []string {
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

	noduplicates := RemoveDuplicates(x)

	for _, val := range noduplicates {
		println(val)
    }

}