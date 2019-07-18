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

func countOfCharMap(str string)  {
	ans := map[int32]int{}

	for _, i := range str {
		ans[i]++
	}

	for key, value := range ans {
		fmt.Println(string(key), value)
	}

}

func main() {
	dat, err := ioutil.ReadFile("text.txt")
	if err != nil {
		fmt.Println("Sorry")
	}
	//fmt.Println(string(dat))

	s := strings.Split(string(dat), "\n")
	for _, i := range s {
		go countOfCharMap(i)
	}
	fmt.Scanln()
	fmt.Println("done")
}
