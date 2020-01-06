package main

import "fmt"

type Item interface {

}

type Priority struct {
	pri int
	item Item
}

type Priority_Queue_List struct {
	priority Priority
	next *Priority_Queue_List
}

func (p *Priority_Queue_List)size() int {
	if p.isEmpty() {
		return 0
	}
	count := 0
	for p != nil {
		count++
		p = p.next
	}
	return count
}

func (p *Priority_Queue_List)isEmpty() bool {
	if p == nil {
		return true
	}
	return false
}

func (p *Priority_Queue_List)push(priority Priority) *Priority_Queue_List  {
	if p.isEmpty() {
		return nil
	}
	newP := &Priority_Queue_List{
		priority:priority,
		next:nil,
	}
	if priority.pri > p.priority.pri {
		newP.next = p
		return newP
	}
	pre := p
	cnt := p.next
	for (priority.pri < cnt.priority.pri && cnt.next != nil) {
		pre = pre.next
		cnt = cnt.next
	}
	if cnt.next == nil {
		cnt.next = newP
	} else {
		pre.next = newP
		newP.next = cnt
	}
	return p
}



func Pprint(p *Priority_Queue_List) {
	for p != nil {
		fmt.Println(p.priority)
		p = p.next
	}
}

func (p *Priority_Queue_List)pop() (*Priority, *Priority_Queue_List) {
	if p.isEmpty() {
		return nil, nil
	}
	q := p.next
	return &p.priority, q
}

func main()  {
	p := &Priority_Queue_List{
		priority: struct {
			pri  int
			item Item
		}{pri: 4, item: "gdf"},
	}
	p = p.push(struct {
		pri  int
		item Item
	}{pri: 24, item: "sr"})

	p = p.push(struct {
		pri  int
		item Item
	}{pri: 4, item: "oi"})

	p = p.push(struct {
		pri  int
		item Item
	}{pri: 1, item: "oi"})

	Pprint(p)

	k , p := p.pop()

	fmt.Println(k)

	Pprint(p)
}
