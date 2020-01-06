package main

import (
	"fmt"
)

type Item interface {
}

type Queue struct {
	Items Item
	Pre   *Queue
	Next  *Queue
}

type IQueue interface {
	Push_back(item Item)
	Push_front(item Item) *Queue
	Pop_back() (Item, *Queue)
	Pop_front() (Item, *Queue)
	Size() int
	IsEmpty() bool
}

func (q *Queue) Push_back(item Item) {
	nextQ := &Queue{
		Items: item,
		Pre:   nil,
		Next:  nil,
	}
	for q.Next != nil {
		q = q.Next
	}
	q.Next = nextQ
	nextQ.Pre = q
}

func (q *Queue) Push_front(item Item) *Queue {
	nextQ := &Queue{
		Items: item,
		Pre:   nil,
		Next:  q,
	}
	q = nextQ
	return q
}

func (q *Queue) Pop_back() (Item, *Queue) {
	if q.IsEmpty() {
		return nil, nil
	}
	if q.Next == nil {
		return q.Items, nil
	}
	pre := q
	cnt := q.Next
	for cnt.Next != nil {
		cnt = cnt.Next
		pre = pre.Next
	}
	pre.Next = nil
	return cnt.Items, q
}

func (q *Queue) Pop_front() (Item, *Queue) {
	if q.IsEmpty() {
		return nil, nil
	}
	if q.Next == nil {
		return q.Items, nil
	}
	p := q.Next
	return q.Items, p
}

func (q *Queue) Size() int {
	count := 0
	for q != nil {
		count++
		q = q.Next
	}
	return count
}

func (q *Queue) IsEmpty() bool {
	if q == nil {
		return true
	}
	return false
}

func print(queue *Queue) {
	for queue != nil {
		fmt.Println(queue.Items)
		queue = queue.Next
	}
}

func main() {
	q := &Queue{
		Items: 4,
		Pre:   nil,
		Next:  nil,
	}
	q.Push_back(7)
	q.Push_back(6)
	q = q.Push_front(9)
	q = q.Push_front(8)
	q.Push_back("hkhgjg")
	print(q)
	it, q := q.Pop_back()
	fmt.Println("pop_back: ", it)
	print(q)
	it, q = q.Pop_front()
	fmt.Println("pop_front: ", it)
	print(q)
	fmt.Println("size:", q.Size())
}
