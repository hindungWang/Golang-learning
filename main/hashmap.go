package main

import "fmt"

var (
	BUCKETCOUNT int32  =  16
)

type hashEntry struct {
	key string
	value string
	next *hashEntry
}

type hashTable struct {
	bucket [16]hashEntry
}

func (h *hashTable) keyToIndex(str string) int32  {
	var index int32 = -1
	if str == "" {
		return index
	}
	len := len(str)
	index = int32(str[0])
	for i := 1; i < len; i++ {
		index *= 1103515245 + int32(str[i])
	}
	index >>= 27
	index &= (BUCKETCOUNT - 1)
	return index
}

func (h *hashTable) insertEntry(key string, val string) int32  {
	m := new(hashEntry)
	m.key = key
	m.value = val
	if key == "" || val == "" {
		return -1
	}
	index := h.keyToIndex(key)

	if h.bucket[index].key == "" {
		h.bucket[index].key = key
		h.bucket[index].value = val
	} else {
		if h.bucket[index].key == key {
			h.bucket[index].value = val
			return 0
		}
		iter := &h.bucket[index]
		for iter.next != nil {
			if key == iter.key {
				iter.value = val
				return 0
			}
			iter = iter.next
		}
		iter.next = m
	}
	return 0
}

func (h *hashTable) getVal(key string) string  {
	if key == "" {
		return ""
	}
	index := h.keyToIndex(key)
	if h.bucket[index].key == "" {
		return ""
	}
	iter := &h.bucket[index]
	for iter != nil {
		if iter.key == key {
			return iter.value
		}
		iter = iter.next
	}
	return ""
}

func (h *hashTable) remove(key string) int32 {
	if key == "" {
		return -1
	}
	index := h.keyToIndex(key)
	if h.bucket[index].key == "" {
		return -1
	} else {
		if h.bucket[index].key == key {
			if h.bucket[index].next == nil {
				h.bucket[index].key = ""
				h.bucket[index].value = ""
				return 0
			}
			h.bucket[index] = *h.bucket[index].next
			return 0
		}
		bar := &h.bucket[index]
		foo := bar.next
		for foo != nil {
			if foo.key == key {
				bar.next = foo.next
				return 0
			}
			bar = bar.next
			foo = foo.next
		}
	}
	return -1
}

func (h *hashTable)printT()  {
	for i,b := range h.bucket {
		iter := &b
		fmt.Println("bucket = ", i)
		for iter != nil {
			fmt.Printf("key = %s, vel = %s\n", iter.key, iter.value)
			iter = iter.next
		}
	}
}

func main()  {
	h := new(hashTable)
	//h.printT()
	h.insertEntry("哈哈哈", "dsaf")
	h.insertEntry("dg", "dsaf")
	h.insertEntry("dsge", "eraw")
	h.insertEntry("dg", "fsdjfdsdssajfha")
	h.printT()
}
