// 堆，完全二叉树，节点的值大于左右子树的值（大顶堆）
// 用数组存储，浪费第一位，节点i的左右子节点分别为2i，2i+1;i的父节点为i/2
package heap

type Heap struct {
	arr []int // 从下标1开始存储数据
	count int // 已存储元素数
}

func Init(capacity int) *Heap {
	a := make([]int, capacity + 1)
	h := &Heap{
		arr: a,
		count: 0,
	}

	return h
}

func (h *Heap) insert(data int) {
	h.count++
	h.arr[h.count] = data

	// 从下而上堆化
	i := h.count
	for ;i/2 > 0 && h.arr[i] > h.arr[i/2]; {
		h.arr[i], h.arr[i/2] = h.arr[i/2], h.arr[i]
		i = i/2
	}
}

func (h *Heap) delete(data int) {
	if h.count == 0 {
		return
	}

	h.arr[1] = h.arr[h.count]
	h.count--

	// 从上而下堆化
	i := 1
	for {
		maxPos := i
		if i*2 < h.count && h.arr[i] < h.arr[2*i] {
			maxPos = 2*i
		}
		if i*2 + 1 < h.count && h.arr[i] < h.arr[2*i+1] {
			maxPos = 2*i+1
		}
		if maxPos == i {
			break
		}
		h.arr[i], h.arr[maxPos] = h.arr[maxPos], h.arr[i]
		i = maxPos
	}
}

func buildHeap(arr []int) []int {
	return nil
}

