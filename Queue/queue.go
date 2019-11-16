package Queue

import (
	"fmt"
	"reflect"
)

type (
	Queue struct {
		first, last *node
		len         int
		stackType   reflect.Kind
	}
	node struct {
		next  *node
		value interface{}
	}
)

//Unsafe be careful
func New() *Queue {
	return &Queue{nil, nil, 0, reflect.Interface}
}

func NewIntQueue() *Queue {
	return &Queue{nil, nil, 0, reflect.Int}
}

func NewStringQueue() *Queue {
	return &Queue{nil, nil, 0, reflect.String}
}

func NewFloat64Queue() *Queue {
	return &Queue{nil, nil, 0, reflect.Float64}
}

func NewComplexQueue() *Queue {
	return &Queue{nil, nil, 0, reflect.Complex128}
}

//Be careful with interfaces
//Use explicit type conversion
//For example when you try to add 1 use 1.0 or Float64(1)
func (queue *Queue) Add(x interface{}) {
	if queue.stackType != reflect.TypeOf(x).Kind() && queue.stackType != reflect.Interface {
		panic(fmt.Sprintf("expected %s but argument is of type %s", queue.stackType.String(), reflect.TypeOf(x).Kind().String()))
	}
	n := &node{next: nil, value: x}
	if queue.len == 0 {
		queue.first = n
		queue.last = n
	} else {
		queue.last.next = n
		queue.last = n
	}
	queue.len++
}

func (queue Queue) Front() interface{} {
	if queue.len == 0 {
		return nil
	}
	return queue.first.value
}
func (queue Queue) Back() interface{} {
	if queue.len == 0 {
		return nil
	}
	return queue.last.value
}
func (queue *Queue) Pop() {
	if queue.len == 0 {
		return
	}
	queue.first = queue.first.next
	queue.len--
}
func (queue Queue) Len() int {
	return queue.len
}

func (queue Queue) Empty() bool {
	return queue.len == 0
}

func (queue Queue) Clear() {
	queue.first = nil
	queue.last = nil
	queue.len = 0
}

func (queue Queue) String() string {
	if queue.first == nil {
		return "[<nil>]"
	}
	s := fmt.Sprintf("%d [", queue.len)
	for queue.len > 1 {
		s += fmt.Sprintf("%v, ", queue.first.value)
		queue.Pop()
	}
	s += fmt.Sprintf("%v]", queue.first.value)
	return s
}
