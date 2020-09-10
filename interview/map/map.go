package main

import (
	"fmt"
	"sort"
)

func main() {
	SortMap()
}

/*
问题：根据map中value的值有序遍历map
思路：
1. map为无序的，但是数组是有序的
2. 先遍历map，将key组成一个数组k
3. 用sort.Slice方法将数组进行排序，排序判断条件可以是用对应的key取map中的value进行比较
4. 经过第三步，k数组的顺序已经排序，然后遍历k，通过k中元素去取map中的值即有序
*/
// SortMap 有序遍历map
func SortMap() {
	m := map[string]int{
		"a": 0,
		"b": 1,
		"c": 2,
		"d": 3,
		"e": 4,
		"f": 5,
		"g": 6,
		"h": 7,
		"i": 8,
		"j": 9,
		"k": 10,
		"l": 11,
		"m": 12,
	}
	// for k, v := range m {
	// 	fmt.Println("m-----------", k, v)
	// }
	keys := []string{}
	for k := range m {
		keys = append(keys, k)
	}
	fmt.Println("keys---1---------", keys)
	sort.Slice(keys, func(i, j int) bool {
		if m[keys[i]] > m[keys[j]] {
			return true
		}
		return false
	})
	fmt.Println("keys---2---------", keys)

	for _, key := range keys {
		fmt.Println("m-value---------------", m[key])
	}
}
