package main

import (
	"fmt"
	"log"

	heapSlice "github.com/Brotchu/heap-slice"
)

func main() {
	pHeap := heapSlice.NewHeap()

	// err := pHeap.Insert(heapSlice.Node{
	// 	Priority: 1,
	// 	Message:  "First one",
	// })
	// logErr(err)
	msgs := []int{2, 3, 2, 1, 4, 2, 0, 1}

	for _, m := range msgs {
		logErr(pHeap.Insert(m, "Hello"))
	}
	// logErr(pHeap.Insert(2, "Hello"))
	pHeap.Test()

	// res, err := pHeap.Delete()
	// logErr(err)
	// fmt.Printf("Top Item deleted : %+v\n", res)
	// res, err = pHeap.Delete()
	// logErr(err)
	// fmt.Printf("Top Item deleted : %+v\n", res)
	fmt.Println("deleting one by one : ")
	for range msgs {
		res, err := pHeap.Delete()
		logErr(err)
		fmt.Printf("Top Node : %+v\n", res)
		pHeap.Test()
	}

	// pHeap.Test()
}

func logErr(err error) {
	if err != nil {
		log.Fatal("Error : ", err)
	}
}
