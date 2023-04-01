package main

import (
	"fmt"
	"guia2/queue"
)

func main() {
	q := *queue.NewQueue[int](10)

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)

	v, err := q.Dequeue()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(v)
	}

	v, err = q.Dequeue()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(v)
	}

	v, err = q.Dequeue()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(v)
	}

	v, err = q.Dequeue()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(v)
	}

	q2 := *queue.NewQueueS[int](10)

	q2.Enqueue(1)
	q2.Enqueue(2)
	q2.Enqueue(3)
	q2.Enqueue(4)

	v, err = q2.Dequeue()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(v)
	}

	v, err = q2.Front()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(v)
	}

}
