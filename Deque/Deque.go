package Deque

import (
	"fmt"
	"reflect"
)

type (
	Deque struct {
		first, last *node
		len         int
		stackType   reflect.Kind
	}
	node struct {
		next     *node
		previous *node
		value    interface{}
	}
)

//Unsafe be careful
func New() *Deque {
	return &Deque{nil, nil, 0, reflect.Interface}
}

func NewIntDeque() *Deque {
	return &Deque{nil, nil, 0, reflect.Int}
}

func NewStringDeque() *Deque {
	return &Deque{nil, nil, 0, reflect.String}
}

func NewFloat64Deque() *Deque {
	return &Deque{nil, nil, 0, reflect.Float64}
}

func NewComplexDeque() *Deque {
	return &Deque{nil, nil, 0, reflect.Complex128}
}

//Be careful with interfaces
//Use explicit type conversion
//For example when you try to add 1 use 1.0 or Float64(1)
func (queue *Deque) PushBack(x interface{}) {
	if queue.stackType != reflect.TypeOf(x).Kind() && queue.stackType != reflect.Interface {
		panic(fmt.Sprintf("expected %s but argument is of type %s", queue.stackType.String(), reflect.TypeOf(x).Kind().String()))
	}
	n := &node{next: nil, value: x, previous: nil}
	if queue.len == 0 {
		queue.first = n
		queue.last = n
	} else {
		n.previous = queue.last
		queue.last.next = n
		queue.last = n
	}
	queue.len++
}

//Be careful with interfaces
//Use explicit type conversion
//For example when you try to add 1 use 1.0 or Float64(1)
func (queue *Deque) PushFront(x interface{}) {
	if queue.stackType != reflect.TypeOf(x).Kind() && queue.stackType != reflect.Interface {
		panic(fmt.Sprintf("expected %s but argument is of type %s", queue.stackType.String(), reflect.TypeOf(x).Kind().String()))
	}
	n := &node{next: nil, value: x, previous: nil}
	if queue.len == 0 {
		queue.first = n
		queue.last = n
	} else {
		n.next = queue.first
		queue.first.previous = n
		queue.first = n
	}
	queue.len++
}

func (queue Deque) Front() interface{} {
	if queue.len == 0 {
		return nil
	}
	return queue.first.value
}
func (queue Deque) Back() interface{} {
	if queue.len == 0 {
		return nil
	}
	return queue.last.value
}
func (queue *Deque) PopFront() {
	if queue.len == 0 {
		return
	}
	queue.first = queue.first.next
	queue.len--
}
func (queue *Deque) PopBack() {
	if queue.len == 0 {
		return
	}
	if queue.len > 1 {
		queue.last = queue.last.previous
		queue.len--
		return
	}
	queue.last = nil
	queue.first = nil
	queue.len = 0
}
func (queue Deque) Len() int {
	return queue.len
}
func (queue Deque) Empty() bool {
	return queue.len == 0
}

func (queue Deque) Clear() {
	queue.first = nil
	queue.last = nil
	queue.len = 0
}

func (queue Deque) String() string {
	if queue.first == nil {
		return "[<nil>]"
	}
	s := fmt.Sprintf("%d [", queue.len)
	for queue.len > 1 {
		s += fmt.Sprintf("%v, ", queue.first.value)
		queue.PopFront()
	}
	s += fmt.Sprintf("%v]", queue.first.value)
	return s
}
