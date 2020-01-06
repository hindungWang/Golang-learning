package main

import "fmt"

// 堆实现最小优先队列，即最小堆
// 数据结构是数组构成的二叉树（完全二叉树）
// 树中某个节点的值总是不大于或不小于其孩子节点的值；
// 树中每个节点的子树都是堆树
//        1  2  3  4   5  6   7   8  9  10
// a[] = {4, 1, 3, 2, 16, 9, 10, 14, 8, 7}
//                  4
//                /  \
//               1    3
//              / \  / \
//             2  16 9  10
//            /\  /
//           14 8 7
// 若取a[1]为堆顶最大或最小，a[i]的左子节点为a[2i],右子节点为a[2i+1],父节点为a[i/2]
// 若取a[0]为堆顶最大或最小，a[i]的左子节点为a[2i+1],右子节点为a[2i+2],父节点为a[(i-1)/2]

type Priority_Heap struct {
	arr []int
}

func (p *Priority_Heap)push(pri int)  {
	p.arr = append(p.arr, pri)
	if len(p.arr) == 2 {
		return
	}
	// fixup 比父节点小的就上浮
	end := len(p.arr) - 1
	father := end / 2
	for father >= 1 {
		if p.arr[father] <= p.arr[end] {
			break
		}
		p.arr[father], p.arr[end] = p.arr[end], p.arr[father]
		end = father
		father = end / 2
	}
}

func (p *Priority_Heap)pop() int  {
	length := len(p.arr)
	if length <= 1 {
		return -1;
	}
	res := p.arr[1]
	if length == 2 {
		p.arr = []int{0}
		return res
	}
	p.arr[1] = p.arr[length - 1]
	p.arr = p.arr[:length - 1]
	length = len(p.arr)
	// fix down 比左、右节点小的越小越往哪边下沉
	farther := 1
	min := 2*farther
	for min < length {
		if min + 1 < length {// 右节点存在？
			if p.arr[min] > p.arr[min+1] { // 左节点比右节点大？
				min++                      // 那就用右节点，否则往左子节点下沉
			}
		}
		if p.arr[farther] <= p.arr[min] {
			break
		}
		p.arr[farther], p.arr[min] = p.arr[min], p.arr[farther]
		farther = min
		min = 2*farther
	}
	return res
}

func initHeap() *Priority_Heap  {
	p := &Priority_Heap{
		arr: []int{0},
	}
	return p
}

func main()  {
	p := initHeap()
	p.push(9)
	p.push(4)
	p.push(2)
	p.push(4)
	p.push(6)
	p.push(7)
	p.push(8)
	p.push(1)

	fmt.Println(p)

	p.pop()
	fmt.Println(p)
	p.pop()
	fmt.Println(p)
	p.pop()
	fmt.Println(p)
	p.pop()
	fmt.Println(p)
	p.pop()
	fmt.Println(p)
	p.pop()
	fmt.Println(p)
	p.pop()
	fmt.Println(p)
	p.pop()
	fmt.Println(p)
	p.pop()
	fmt.Println(p)
	p.pop()
	fmt.Println(p)


}