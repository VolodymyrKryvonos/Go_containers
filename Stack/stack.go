package Stack

import (
	"fmt"
	"reflect"
)

type (
	Stack struct {
		top       *node
		len       int
		stackType reflect.Kind
	}
	node struct {
		value    interface{}
		previous *node
	}
)

//Unsafe be careful
func New() *Stack {
	return &Stack{nil, 0, reflect.Interface}
}

func NewIntStack() *Stack {
	return &Stack{nil, 0, reflect.Int}
}

func NewStringStack() *Stack {
	return &Stack{nil, 0, reflect.String}
}

func NewFloat64Stack() *Stack {
	return &Stack{nil, 0, reflect.Float64}
}

func NewComplexStack() *Stack {
	return &Stack{nil, 0, reflect.Complex128}
}

//Be careful with interfaces
//Use explicit type conversion
//For example when you try to add 1 use 1.0 or Float64(1)
func (stack *Stack) Add(x interface{}) {
	if stack.stackType != reflect.TypeOf(x).Kind() && stack.stackType != reflect.Interface {
		panic(fmt.Sprintf("expected %s but argument is of type %s", stack.stackType.String(), reflect.TypeOf(x).Kind().String()))
	}
	stack.top = &node{x, stack.top}
	stack.len++
}

func (stack *Stack) Clear() {
	stack.len = 0
	stack.top = nil
}

func (stack *Stack) Pop() {
	if stack.len == 0 {
		return
	}
	stack.top = stack.top.previous
	stack.len--
}

func (stack Stack) IsEmpty() bool {
	return stack.len == 0
}

func (stack Stack) Len() int {
	return stack.len
}

func (stack Stack) Top() interface{} {
	if stack.len == 0 {
		return nil
	} else {
		return stack.top.value
	}
}

func (stack Stack) String() string {
	if stack.top == nil {
		return "[<nil>]"
	}
	s := fmt.Sprintf("%d [", stack.len)
	for stack.len > 1 {
		s += fmt.Sprintf("%v, ", stack.top.value)
		stack.Pop()
	}
	s += fmt.Sprintf("%v]", stack.top.value)
	return s
}
