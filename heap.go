package heapSlice

import (
	"errors"
	"fmt"
)

type Node struct {
	Priority int
	Message  string
}

type Heap struct {
	HeapList []Node
}

func NewHeap() *Heap {
	return &Heap{
		HeapList: []Node{},
	}
}

func (h *Heap) Insert(p int, m string) error {
	return h.insert(Node{
		Priority: p,
		Message:  m,
	})
}

func (h *Heap) insert(n Node) error {
	if n.Priority < 0 || n.Priority > 4 {
		return errors.New("Priority has to be 0-4")
	}
	h.HeapList = append(h.HeapList, n)
	// fmt.Println("TEST : >> ", len(h.HeapList)-1)
	// fmt.Printf("%+v\n", h.HeapList)
	return h.insertSwap(len(h.HeapList) - 1)
}

func (h *Heap) insertSwap(i int) error {

	var current int
	var parent int
	for i > 0 {
		current = i
		// fmt.Println(i)
		if i%2 == 0 {
			parent = (i - 2) / 2
		} else {
			parent = (i - 1) / 2
		}
		if h.HeapList[current].Priority < h.HeapList[parent].Priority {
			temp := h.HeapList[current]
			h.HeapList[current] = h.HeapList[parent]
			h.HeapList[parent] = temp
			i = parent
			// h.Test()
		} else {
			// fmt.Println("After swap")
			// fmt.Printf("%+v\n", h.HeapList)
			break
		}
	}

	return nil
}

func (h *Heap) Delete() (Node, error) {
	if len(h.HeapList) == 0 {
		return Node{}, errors.New("Empty Heap")
	}

	res := h.HeapList[0]
	if len(h.HeapList) == 1 {
		h.HeapList = []Node{}
		return res, nil
	}
	// h.HeapList = append(h.HeapList[len(h.HeapList)-1], h.HeapList[1:])
	lastNode := h.HeapList[len(h.HeapList)-1]
	h.HeapList = append([]Node{lastNode}, h.HeapList[1:len(h.HeapList)-1]...)
	h.deleteSwap(0)
	return res, nil
}

func (h *Heap) deleteSwap(i int) error {

	// var current, parent
	//TODO: implement delete swap procedure
	var current int
	var child1, child2, min int
	for {
		current = i
		if (current*2 + 1) > len(h.HeapList)-2 {
			break
		}
		child1 = current*2 + 1
		child2 = current*2 + 2
		min = current
		if h.HeapList[child1].Priority < h.HeapList[current].Priority {
			min = child1
		}
		if (current*2 + 2) <= len(h.HeapList)-1 {
			if h.HeapList[child2].Priority < h.HeapList[min].Priority {
				min = child2
			}
		}
		if min == current {
			break
		}
		temp := h.HeapList[current]
		h.HeapList[current] = h.HeapList[min]
		h.HeapList[min] = temp

		i = min
	}
	return nil
}

func (h *Heap) Test() {
	// fmt.Printf("Heap : %+v\n", h.HeapList)
	fmt.Println("Heap Content :  ")
	// fmt.Printf("%+v\n", h.HeapList)
	for _, node := range h.HeapList {
		fmt.Printf("%+v\n", node)
	}
	// fmt.Printf("Heap : %+v\n", h.HeapList[1:len(h.HeapList)-1])
	// last := h.HeapList[len(h.HeapList)-1]
	// fmt.Printf("Heap : %+v\n", append([]Node{last}, h.HeapList[1:len(h.HeapList)-1]...))
	// append(h.HeapList[len(h.HeapList)-1], h.HeapList[1:len(h.HeapList)-1])
}
