package queue

import (
	"errors"
	"guia2/stack"
)

type Queue[T any] struct {
	items []T
}

func NewQueue[T any](longitud int) *Queue[T] {
	queue := new(Queue[T])
	queue.items = make([]T, 0, longitud)
	return queue
}

func (q *Queue[T]) Enqueue(v T) {
	q.items = append(q.items, v)
}

func (q *Queue[T]) Dequeue() (v T, err error) {
	if q.IsEmpty() {
		return v, errors.New("La cola esta vacia")
	}

	output := q.items[0]
	q.items = q.items[1:len(q.items)]
	return output, nil
}

func (q *Queue[T]) IsEmpty() bool {
	return len((*q).items) == 0
}

func (q *Queue[T]) Front() (T, error) {
	output := q.items[0]
	return output, nil
}

type QueueS[T any] struct {
	p1, p2 stack.Stack[T]
}

func (q *QueueS[T]) Enqueue(v T) {
	q.p1.Push(v)
}

func NewQueueS[T any](longitud int) *QueueS[T] {
	queues := new(QueueS[T])
	queues.p1 = *stack.NewStack[T](longitud)
	queues.p2 = *stack.NewStack[T](longitud)
	return queues
}

func (q *QueueS[T]) Dequeue() (T, error) {
	if q.IsEmpty() {
		var t T
		return t, errors.New("La cola esta vacia")
	}
	if q.p2.IsEmpty() {
		for !(q.p1.IsEmpty()) {
			aux, _ := q.p1.Pop()
			q.p2.Push(aux)
		}
	}
	return q.p2.Pop()
}

func (q *QueueS[T]) IsEmpty() bool {
	return q.p1.IsEmpty() && q.p2.IsEmpty()
}

func (q *QueueS[T]) Front() (T, error) {
	if q.IsEmpty() {
		var t T
		return t, errors.New("La cola esta vacia")
	}
	if q.p2.IsEmpty() {
		for !(q.p1.IsEmpty()) {
			aux, _ := q.p1.Pop()
			q.p2.Push(aux)
		}
	}
	return q.p2.Front()
}
