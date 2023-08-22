package main

import "fmt"

func main() {
	data := longestCommonPrefix([]string{"flower", "flow", "flight"})
	fmt.Println(data)
}

func longestCommonPrefix(strs []string) string {
	for i := 0; i < len(strs[0]); i++ {
		for _, s := range strs {
			if i >= len(s) || s[i] != strs[0][i] {
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}
