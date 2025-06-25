package fpm


import "fmt"

type Option[T any] struct {
	hasValue bool
	value    T
}

func Some[T any](v T) Option[T] {
	return Option[T]{hasValue: true, value: v}
}

func None[T any]() Option[T] {
	var zero T
	return Option[T]{hasValue: false, value: zero}
}

func (o Option[T]) IsSome() bool {
	return o.hasValue
}

func (o Option[T]) IsNone() bool {
	return !o.hasValue
}

func (o Option[T]) Unwrap() T {
	if !o.hasValue {
		var zero T
		fmt.Println("Unwrap() called on None")
		return zero
	}
	return o.value
}

func (o Option[T]) OrElse(defaultValue T) T {
	if o.hasValue {
		return o.value
	}
	return defaultValue
}

func Map[T any, U any](o Option[T], f func(T) U) Option[U] {
	if o.IsSome() {
		return Some(f(o.value))
	}
	return None[U]()
}

func FlatMap[T any, U any](o Option[T], f func(T) Option[U]) Option[U] {
	if o.IsSome() {
		return f(o.value)
	}
	return None[U]()
}

func (o Option[T]) Match(someFn func(T), noneFn func()) {
	if o.hasValue {
		someFn(o.value)
	} else {
		noneFn()
	}
}

func (o Option[T]) String() string {
	if o.hasValue {
		return fmt.Sprintf("Some(%v)", o.value)
	}
	return "None"
}
