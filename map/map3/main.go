package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	c := strings.Split(s, " ")
	data := make(map[string]int)

	for i := 0; i < len(c); i++ {
		_, tkit := data[c[i]]
		if tkit {
			data[c[i]]++
		} else {
			data[c[i]] = 1
		}
	}

	return data
}

func WordCountWithRange(s string) map[string]int {
	s = strings.ToLower(s)
	data := make(map[string]int)
	w := strings.Fields(s)

	for _, word := range w {
		data[word]++
	}

	return data
}

func main() {
	sentence := "Kilo Good Morning New Yorkers New Kilo Good New Morning New New Yorkers"
	r := WordCount(sentence)
	rs := WordCountWithRange(sentence)
	fmt.Println(rs)
	fmt.Println(r)
}
