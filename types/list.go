package types

import (
	"fmt"
	"reflect"
)

type List[T any] struct {
	Data []T
}

// make new list for  any type data
func NewList[T any]() *List[T] {
	return &List[T]{
		Data: []T{},
	}
}

// insert data in list
func (l *List[T]) Insert(val T) {
	l.Data = append(l.Data, val)
}

// clear all data from list
func (l *List[T]) Clear() {
	l.Data = []T{}
}

// get data by index
func (l *List[T]) Get(index int) T {
	if index > len(l.Data) {
		err, _ := fmt.Printf("given index %d is higher than max index %d ", index, len(l.Data))
		panic(err)

	}
	return l.Data[index]
}

// get index by val of list
func (l *List[T]) GetIndex(val T) int {
	for i := 0; i < len(l.Data); i++ {
		if reflect.DeepEqual(val, l.Data[i]) {
			return i
		}
	}
	return -1
}

// get length of list
func (l *List[T]) Len() int {
	return len(l.Data)
}

// remove val from list if don't know index
func (l *List[T]) Remove(value T) {
	index := l.GetIndex(value)
	l.Pop(index)
}

// pop element from list by index
func (l *List[T]) Pop(index int) {
	l.Data = append(l.Data[:index], l.Data[index+1:]...)
}

// check weather element are contain or nor in list
func (l *List[T]) Contain(value T) bool {
	for i := 0; i < len(l.Data); i++ {
		if reflect.DeepEqual(value, l.Data[i]) {
			return true
		}
	}
	return false
}
