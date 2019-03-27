package hashtable

import "fmt"

func topKFrequent(nums []int, k int) []int {
	var ret []int
	maps := make(map[int]int, 0)
	for _, v := range nums {
		maps[v]++
	}

	for i:= k-1; i >= 0; i-- {
		key := topFromMap(maps)
		ret = append(ret, key)
		delete(maps, key)
	}
	return ret
}

func topFromMap(maps map[int]int) int {
	max := -1
	key := 0
	for k, v := range maps {
		if v > max {
			max = v
			key = k
		}
	}
	return key
}

func main() {
	arr := []int{1,1,3,4,2,3,1,5}
	ret := topKFrequent(arr, 2)
	fmt.Println(ret)
}