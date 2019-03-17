package main

import "fmt"

// 第一类
// BubbleSort，让最大值在最后一位
// 时间复杂度O(n*n)，空间复杂度O(1),稳定
func BubbleSort(arr []int) []int {
	for j := len(arr)-1; j >= 0; j-- {
		for i:=0; i<j;i++ {
			if arr[i] > arr[i+1] {
				tmp := arr[i]
				arr[i] = arr[i+1]
				arr[i+1] = tmp
			}
		}
	}
	return arr
}

// 第一类
// InsertSort 数据插入左边有序集合
// 时间复杂度O(n*n)，空间复杂度O(1),稳定
// 插入排序的操作比冒泡少，冒泡有3次赋值，插入就一次
func InsertSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		tmp := arr[i]
		j := i -1
		for ; j >= 0; j-- {
			if arr[j] > tmp {
				arr[j+1] = arr[j]
			} else {
				break
			}
		}
		arr[j+1] = tmp
	}
	return arr
}

// 第一类
// SelectSort 选择最小值放入左边有序集合
// 时间复杂度O(n*n)，空间复杂度O(1),不稳定（5，1，5，2，4）2选择放入左边后，与5互换，则原来两个5的先后位置变了
func SelectSort(arr []int) []int {
	for i := 0; i < len(arr) - 1; i++ {
		j := i + 1
		minIndex := j
		min := arr[j]
		for j++; j < len(arr); j++ {
			if min > arr[j] {
				minIndex = j
				min = arr[j]
			}
		}
		arr[minIndex] = arr[i]
		arr[i] = min
	}
	return arr
}

// 第二类
// MergeSort 分治思想，递归分，再合并
// 时间复杂度O(nlogn)，空间复杂度O(n)，稳定
func MergeSort(arr []int) []int {
	mergeSortCore(arr, 0, len(arr)-1)
	return arr
}

func mergeSortCore(arr []int, p, r int) {
	if p >= r {
		return
	}

	q := (p + r) / 2
	mergeSortCore(arr, p, q)
	mergeSortCore(arr, q+1, r)
	merge(arr, p, q, r)
}

func merge(arr []int, p, q, r int) {
	var tmp = []int{}
	i, j := p, q+1

	for ; i <= q && j <= r ;  {
		if arr[i] < arr[j] {
			tmp = append(tmp, arr[i])
			i++
		} else {
			tmp = append(tmp,arr[j])
			j++
		}

	}

	for ; i <= q; i++ {
		tmp = append(tmp, arr[i])
	}

	for ; j <= r; j++ {
		tmp = append(tmp,arr[j])
	}

	for i,j := p, 0; i <= r ; i++ {
		arr[i] = tmp[j]
		j++
	}
}

// 第二类
// QuickSort 分治思想，partition分区函数
// 时间复杂度O(nlogn)，空间复杂度O(1)，稳定
func QuickSort(arr []int) []int {
	quickSortCore(arr, 0, len(arr) - 1)
	return arr
}

func quickSortCore(arr []int, p, r int) {
	if p >= r {
		return
	}
	q := partition(arr, p, r)
	quickSortCore(arr, p, q-1)
	quickSortCore(arr, q+1, r)
}

// 选r为pivot，遍历数组将数组一分为二，大于pivot和小于pivot的，巧妙实现一分为二，返回pivot的index
// 分区的思想很值得学习！
func partition(arr []int, p, r int) int {
	pivot := arr[r]
	i := p
	for j := p; j < r; j++ {
		if arr[j] < pivot {
			arr[j], arr[i] = arr[i], arr[j]
			i++
		}
	}
	arr[r], arr[i] = arr[i], arr[r]
	return i
}

// 第二类
// 堆排序，堆是一个完全二叉树，每个节点的值必须大于其左右子树每个节点的值（大顶堆）
func HeapSort(arr []int) []int {
	return nil
}

// 第三类
// 桶排序，数据均匀分到桶里，用快排分别排序，桶是有序递增的
// 时间复杂度O(n)，空间复杂度O(n)，稳定？
// 场景要求：数据有特点，如年龄都是在0-100范围内，如何给金额在0-100之间的100万条订单排序？适合外部排序，如10G数据，分成100个文件桶，每个100M，再放入内存处理；数据范围不大
func BucketSort(arr []int) []int {
	// 桶数，根据所排序数据选择合适的桶数量m
	m := len(arr)

	// 数组最大值k
	k := getMaxInArr(arr)

	// 二维切片，桶
	buckets := make([][]int, m)

	// 分配数组元素入桶，尽量均匀到各个桶里
	index := 0
	for i :=0; i < m; i++ {
		index = arr[i] * (m-1) / k // 获取arr[i]元素所在的桶index
		buckets[index] = append(buckets[index], arr[i])
	}

	// 桶内排序，用快排
	tmpPos := 0
	for i := 0; i < m; i++ {
		bucketLen := len(buckets[i])
		if bucketLen > 0 {
			QuickSort(buckets[i])
			copy(arr[tmpPos:], buckets[i])
			tmpPos += bucketLen
		}
	}
	return arr
}

func getMaxInArr(arr []int) int{
	max := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > max{ max = arr[i]}
	}
	return max
}

// 第三类
// CountingSort 计数排序，有一个累计计数数组c，通过c和逆序的arr找到index
// 时间复杂度O(n)，空间复杂度O(n),稳定
// 适合数据量大50万的重复数据很多的数组排序，数据范围不大
func CountingSort(arr []int) []int {
	if len(arr) < 1 {
		return arr
	}

	// 查找数组中数据的范围
	max := getMaxInArr(arr)

	// 申请技术数组c，大小为max+1，即【0，max+1】
	c := make([]int, max + 1)
	// 计算每个元素个数，存入c
	for i := 0; i < len(arr); i++ {
		c[arr[i]]++
	}

	// c数组元素依次累加
	for i := 1; i < max + 1; i++ {
		c[i] = c[i-1] + c[i]
	}

	// 临时数组r存储排序后的结果
	r := make([]int, len(arr))
	for i := len(arr) - 1; i >=0; i-- {
		index := c[arr[i]] - 1
		r[index] = arr[i]
		c[arr[i]]--
	}

	// 将结果拷贝给arr数组
	copy(arr, r)
	return arr
}

// 第三类
// RadixSort 基数排序
// 场景：给11位手机号排序，思路：先按照最后一位排序，再按照第二位重新排序，直至第一位，经过11次排序，就都有序了；每一位排序用桶或者基数排序，0(11*n)
func RadixSort(arr []int) []int {
	return nil
}

// 第四类
// 工程类排序算法，golang的sort算法sort/sort.go

func main() {
	var arr = []int{2,1,7,2,5,6,7}
	arr1 := CountingSort(arr)
	fmt.Print(arr1)
}


