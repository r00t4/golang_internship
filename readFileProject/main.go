package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func countOfChar(str string) [26]int {
	ans := [26]int{}
	str = strings.ToLower(str)

	for _, i := range str {
		ans[i-97]++
	}

	return ans
}

func countOfCharMap(str string, j int)  {
	ans := map[int32]int{}

	for _, i := range str {
		ans[i]++
	}

	for key, value := range ans {
		fmt.Printf("goroutine %d - '%s' : %d\n", j, string(key), value)
	}

}

func main() {
	dat, err := ioutil.ReadFile("text.txt")
	if err != nil {
		fmt.Println("Sorry")
	}
	//fmt.Println(string(dat))

	// TODO with channels

	L : for {
		fmt.Println("asd")
		if true {
			break L
		}
	}

	s := strings.Split(string(dat), "\n")
	for j, i := range s {
		go countOfCharMap(i, j)
	}
	fmt.Scanln()
	fmt.Println("done")
}
