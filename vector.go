package go_list

import (
	"errors"
	"reflect"
)

var (
	ErrEmpty    = errors.New("vector is empty")
	ErrIndex    = errors.New("index out of range")
	ErrNotFound = errors.New("element not found")
)

type Vector[T comparable] []T

func New[T comparable]() Vector[T] {
	return Vector[T]{}
}

func (v *Vector[T]) Add(e T) {
	*v = append(*v, e)
}

func (v *Vector[T]) AddAll(e []T) {
	*v = append(*v, e...)
}

func (v Vector[T]) Capacity() int {
	return cap(v)
}

func (v Vector[T]) Chunk(size int) []Vector[T] {
	var v2 []Vector[T]

	if size <= 0 {
		return v2
	}

	if size >= len(v) {
		v2 = append(v2, v)
		return v2
	}

	for i := 0; i < len(v); i += size {
		v2 = append(v2, v[i:v.min(i+size, len(v))])

	}

	return v2
}

func (v Vector[T]) Contains(e T) bool {
	return v.index(e) != -1
}

func (v Vector[T]) ContainsAll(e []T) bool {
	for _, el := range e {
		if v.index(el) == -1 {
			return false
		}
	}
	return true
}

func (v *Vector[T]) Clear() {
	*v = Vector[T]{}
}

func (v Vector[T]) Empty() bool {
	return len(v) == 0
}

func (v Vector[T]) Filter(f func(e T) bool) Vector[T] {
	v2 := Vector[T]{}
	for _, e := range v {
		if f(e) {
			v2.Add(e)
		}
	}
	return v2
}

func (v Vector[T]) First() T {
	return v[0]
}

func (v Vector[T]) Get(i int) (T, error) {
	err := v.validateIndex(i)
	if err != nil {
		return v.zero(), err
	}
	return v[i], nil
}

func (v Vector[T]) Index(e T) int {
	return v.index(e)
}

func (v *Vector[T]) Insert(i int, e T) error {
	err := v.validateIndex(i)
	if err != nil {
		return err
	}
	*v = append((*v)[:i], append([]T{e}, (*v)[i:]...)...)
	return nil
}

func (v Vector[T]) Last() T {
	return v[len(v)-1]
}

func (v Vector[T]) Len() int {
	return len(v)
}

func (v Vector[T]) Map(f func(e T) T) Vector[T] {
	v2 := Vector[T]{}
	for _, e := range v {
		v2.Add(f(e))
	}
	return v2
}

func (v *Vector[T]) PopFirst() (T, error) {
	if len(*v) == 0 {
		return v.zero(), ErrEmpty
	}
	e := (*v)[0]
	*v = (*v)[1:]
	return e, nil
}

func (v *Vector[T]) PopLast() (T, error) {
	if len(*v) == 0 {
		return v.zero(), ErrEmpty
	}
	e := (*v)[len(*v)-1]
	*v = (*v)[:len(*v)-1]
	return e, nil
}

func (v *Vector[T]) Remove(e T) bool {
	i := v.index(e)
	if i == -1 {
		return false
	}
	*v = append((*v)[:i], (*v)[i+1:]...)
	return true
}

func (v *Vector[T]) RemoveIndex(i int) (T, error) {
	err := v.validateIndex(i)
	if err != nil {
		return v.zero(), err
	}
	e := (*v)[i]
	*v = append((*v)[:i], (*v)[i+1:]...)
	return e, nil
}

func (v *Vector[T]) Replace(e1 T, e2 T) error {
	i := v.index(e1)
	if i == -1 {
		return ErrNotFound
	}
	(*v)[i] = e2
	return nil
}

func (v *Vector[T]) ReplaceIndex(i int, e T) error {
	err := v.validateIndex(i)
	if err != nil {
		return err
	}
	(*v)[i] = e
	return nil
}

func Reduce[T comparable, A any](v Vector[T], acc A, f func(acc A, e T) A) A {
	a := acc
	for _, e := range v {
		a = f(a, e)
	}
	return a
}

// ---------------------------------------------------

func (v Vector[T]) min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func (v Vector[T]) zero() T {
	var e T
	return e
}

func (v Vector[T]) zeroReflect() T {
	return reflect.Zero(reflect.TypeOf(v).Elem()).Interface().(T)
}

func (v Vector[T]) validateIndex(i int) error {
	if len(v) == 0 {
		return ErrEmpty
	}
	if i < 0 || i >= len(v) {
		return ErrIndex
	}
	return nil
}

func (v Vector[T]) index(e T) int {
	for i, el := range v {
		if e == el {
			return i
		}
	}
	return -1
}
