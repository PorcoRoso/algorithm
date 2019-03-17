package main

import (
	"testing"
	"fmt"
	"math/rand"
)

func TestBubbleSort(t *testing.T){
	var arr = []int{1}
	arr1 := bubbleSort(arr)
	fmt.Print(arr1)
	t.Log(arr1)
}

func TestInsertSort(t *testing.T){
	var arr = []int{3,1,7,2,5}
	arr1 := insertSort(arr)
	fmt.Print(arr1)
	t.Log(arr1)
}

func TestSelectSort(t *testing.T){
	var arr = []int{3,1,7,2,5}
	arr1 := selectSort(arr)
	fmt.Print(arr1)
	t.Log(arr1)
}

// benchmark测试命令 ：go test -bench=. -run=^$
// 冒泡和插入的benchmark测试结果，差一个数量级;
// BenchmarkBubbleSort-4                  1        10897614558 ns/op
// BenchmarkInsertSort-4                  1        1354924389 ns/op
func BenchmarkBubbleSort(b *testing.B) {
	var a = [][]int{}

	for i := 0; i < 100; i++ {
		a = append(a, []int{})
	}
	for i := 0; i < 100; i++ {
		for j := 0; j < 10000; j++ {
			a[i] = append(a[i], rand.Intn(100))
		}
	}
	for i := 0; i < 100; i++ {
		bubbleSort(a[i])
	}
}

func BenchmarkInsertSort(b *testing.B) {
	var a = [][]int{}

	for i := 0; i < 100; i++ {
		a = append(a, []int{})
	}
	for i := 0; i < 100; i++ {
		for j := 0; j < 10000; j++ {
			a[i] = append(a[i], rand.Intn(100))
		}
	}
	for i := 0; i < 100; i++ {
		insertSort(a[i])
	}
}

func TestMergeSort(t *testing.T){
	var arr = []int{3,1,7,2,5}
	arr1 := mergeSort(arr)
	fmt.Print(arr1)
	t.Log(arr1)
}

func TestQuickSort(t *testing.T){
	var arr = []int{3,1,7,2,5}
	arr1 := quickSort(arr)
	fmt.Print(arr1)
	t.Log(arr1)
}