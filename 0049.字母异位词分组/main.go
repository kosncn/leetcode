package main

import (
	"fmt"
)

func main() {
	// 示例 1
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(strs)) // Output: [["bat"], ["nat", "tan"], ["ate", "eat", "tea"]]

	// 示例 2
	strs = []string{""}
	fmt.Println(groupAnagrams(strs)) // Output: [[""]]

	// 示例 3
	strs = []string{"a"}
	fmt.Println(groupAnagrams(strs)) // Output: [["a"]]
}

func groupAnagrams(strs []string) [][]string {
	// 解法一：排序 + 哈希
	//strMap := make(map[string][]string)
	//for _, str := range strs {
	//	strList := []byte(str)
	//	sort.Slice(strList, func(i, j int) bool {
	//		return strList[i] < strList[j]
	//	})
	//	strSort := string(strList)
	//	strMap[strSort] = append(strMap[strSort], str)
	//}
	//result := make([][]string, 0, len(strMap))
	//for _, v := range strMap {
	//	result = append(result, v)
	//}
	//return result

	// 解法二：计数 + 哈希
	strMap := make(map[[26]int][]string)
	for _, str := range strs {
		var count [26]int
		for _, s := range str {
			count[s-'a']++
		}
		strMap[count] = append(strMap[count], str)
	}
	result := make([][]string, 0, len(strMap))
	for _, v := range strMap {
		result = append(result, v)
	}
	return result
}
