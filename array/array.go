/**
* @program: lemo
*
* @description:
*
* @author: lemo
*
* @create: 2020-01-07 21:38
**/

package array

import (
	"github.com/lemonyxk/utils/constraints"
	"sort"
)

func Ordered[T constraints.Ordered](src []T) Order[T] {
	return Order[T]{Compare: Compare[T]{Any[T]{src: src}}}
}

type Order[T constraints.Ordered] struct {
	Compare[T]
}

func (a Order[T]) Sum() T {
	var sum T
	for i := 0; i < len(a.src); i++ {
		sum += a.src[i]
	}
	return sum
}

func (a Order[T]) Max() T {

	if len(a.src) == 0 {
		panic("array is empty")
	}

	var max = a.src[0]

	for i := 0; i < len(a.src); i++ {
		if a.src[i] > max {
			max = a.src[i]
		}
	}

	return max
}

func (a Order[T]) Min() T {

	if len(a.src) == 0 {
		panic("array is empty")
	}

	var min = a.src[0]

	for i := 0; i < len(a.src); i++ {
		if a.src[i] < min {
			min = a.src[i]
		}
	}

	return min
}

func (a Order[T]) Asc() {
	for i := 0; i < len(a.src)-1; i++ {
		for j := i + 1; j < len(a.src); j++ {
			if a.src[i] > a.src[j] {
				a.src[i], a.src[j] = a.src[j], a.src[i]
			}
		}
	}
}

func (a Order[T]) Desc() {
	for i := 0; i < len(a.src)-1; i++ {
		for j := i + 1; j < len(a.src); j++ {
			if a.src[i] < a.src[j] {
				a.src[i], a.src[j] = a.src[j], a.src[i]
			}
		}
	}
}

func Comparable[T comparable](src []T) Compare[T] {
	return Compare[T]{Any[T]{src: src}}
}

type Compare[T comparable] struct {
	Any[T]
}

func (a Compare[T]) Has(s T) bool {
	for i := 0; i < len(a.src); i++ {
		if a.src[i] == s {
			return true
		}
	}
	return false
}

func (a Compare[T]) Index(s T) int {
	for i := 0; i < len(a.src); i++ {
		if a.src[i] == s {
			return i
		}
	}
	return -1
}

func (a Compare[T]) Count(s T) int {
	var count int
	for i := 0; i < len(a.src); i++ {
		if a.src[i] == s {
			count++
		}
	}
	return count
}

func (a Compare[T]) Unique() []T {
	var res []T
	var mapSet = make(map[T]bool)
	for i := 0; i < len(a.src); i++ {
		if _, ok := mapSet[a.src[i]]; !ok {
			mapSet[a.src[i]] = true
			res = append(res, a.src[i])
		}
	}
	return res
}

func (a Compare[T]) Diff(s []T) []T {
	var res []T
	var mapSet = make(map[T]bool)
	for i := 0; i < len(s); i++ {
		mapSet[s[i]] = true
	}
	for i := 0; i < len(a.src); i++ {
		if _, ok := mapSet[a.src[i]]; !ok {
			res = append(res, a.src[i])
		}
	}
	return res
}

func (a Compare[T]) Intersect(s []T) []T {
	var res []T
	var mapSet = make(map[T]bool)
	for i := 0; i < len(s); i++ {
		mapSet[s[i]] = true
	}
	for i := 0; i < len(a.src); i++ {
		if _, ok := mapSet[a.src[i]]; ok {
			res = append(res, a.src[i])
		}
	}
	return res
}

func (a Compare[T]) Union(s []T) []T {
	var res []T
	var mapSet = make(map[T]bool)
	for i := 0; i < len(a.src); i++ {
		mapSet[a.src[i]] = true
	}
	for i := 0; i < len(s); i++ {
		mapSet[s[i]] = true
	}
	for k := range mapSet {
		res = append(res, k)
	}
	return res
}

func From[T any](src []T) Any[T] {
	return Any[T]{src}
}

type Any[T any] struct {
	src []T
}

func (a Any[T]) Slice(start, end int) []T {

	var res []T

	if start < 0 {
		start = len(a.src) + start
	}

	if end < 0 {
		end = len(a.src) + end
	}

	if start > end {
		panic("start must less than end")
	}

	for i := start; i < end; i++ {
		res = append(res, a.src[i])
	}

	return res
}

func (a Any[T]) Splice(start int, count int, elem ...T) []T {

	if start < 0 {
		panic("start must be greater than 0")
	}

	if start > len(a.src) {
		panic("start must be less than length of array")
	}

	if start+count > len(a.src) {
		count = len(a.src) - start
	}

	var p1 = a.src[:start]
	var p2 = a.src[start+count:]

	var p3 = a.src[start : start+count]

	a.src = a.src[0:0]

	a.src = append(a.src, p1...)
	a.src = append(a.src, elem...)
	a.src = append(a.src, p2...)

	return p3
}

func (a Any[T]) Insert(start int, elem ...T) {

	if start < 0 {
		panic("start must be greater than 0")
	}

	if start > len(a.src) {
		panic("start must be less than length of array")
	}

	var p1 = a.src[:start]
	var p2 = a.src[start:]

	a.src = a.src[0:0]

	a.src = append(a.src, p1...)
	a.src = append(a.src, elem...)
	a.src = append(a.src, p2...)
}

func (a Any[T]) Delete(start int, count int) {

	if start < 0 {
		panic("start must be greater than 0")
	}

	if start > len(a.src) {
		panic("start must be less than length of array")
	}

	if start+count > len(a.src) {
		count = len(a.src) - start
	}

	var p1 = a.src[:start]
	var p2 = a.src[start+count:]

	a.src = a.src[0:0]

	a.src = append(a.src, p1...)
	a.src = append(a.src, p2...)
}

func (a Any[T]) Push(elem ...T) {
	a.src = append(a.src, elem...)
}

func (a Any[T]) Pop() T {
	var elem = a.src[len(a.src)-1]
	a.src = a.src[:len(a.src)-1]
	return elem
}

func (a Any[T]) Shift() T {
	var elem = a.src[0]
	a.src = a.src[1:]
	return elem
}

func (a Any[T]) UnShift(elem ...T) {
	a.src = append(elem, a.src...)
}

func (a Any[T]) Concat(src ...[]T) []T {
	var res = a.src[:]
	for i := 0; i < len(src); i++ {
		res = append(res, src[i]...)
	}
	return res
}

func (a Any[T]) Reverse() {
	for i := 0; i < len(a.src)/2; i++ {
		a.src[i], a.src[len(a.src)-1-i] = a.src[len(a.src)-1-i], a.src[i]
	}
}

func (a Any[T]) ForEach(fn func(elem T, index int)) {
	for i := 0; i < len(a.src); i++ {
		fn(a.src[i], i)
	}
}

func (a Any[T]) Map(fn func(elem T, index int) T) []T {
	var res []T
	for i := 0; i < len(a.src); i++ {
		res = append(res, fn(a.src[i], i))
	}
	return res
}

func (a Any[T]) Filter(fn func(elem T, index int) bool) []T {
	var res []T
	for i := 0; i < len(a.src); i++ {
		if fn(a.src[i], i) {
			res = append(res, a.src[i])
		}
	}
	return res
}

func (a Any[T]) Reduce(fn func(prev T, curr T, index int) T, init T) T {
	var res = init
	for i := 0; i < len(a.src); i++ {
		res = fn(res, a.src[i], i)
	}
	return res
}

func (a Any[T]) Sort(fn func(a T, b T) bool) {
	sort.Slice(a.src, func(i, j int) bool {
		return fn(a.src[i], a.src[j])
	})
}
