package option

import (
	"encoding/json"
	"fmt"
)

type Option[T any] struct {
	value *T
}

type (
	Int    = Option[int]
	String = Option[string]
	Bool   = Option[bool]
)

func IntValue(v int) Int            { return Some(v) }
func StringValue(v string) String   { return Some(v) }
func BoolValue(v bool) Bool         { return Some(v) }
func List[T any](v []T) Option[[]T] { return Some(v) }

type Optionable interface {
	IsNone() bool
}

// Some returns an option with the given value.
func Some[T any](value T) Option[T] {
	return Option[T]{&value}
}

// None returns an option with no value.
func None[T any]() Option[T] {
	return Option[T]{nil}
}

// Value returns the value of the option - ensure that the option is not None before calling this.
func (o Option[T]) Value() T {
	return *o.value
}

// IsNone returns true if the option is None.
func (o Option[T]) IsNone() bool {
	return o.value == nil
}

// IsSome returns true if the option is not None.
func (o Option[T]) IsSome() bool {
	return o.value != nil
}

// Unwrap returns the value of the option, or panics if the option is None.
// This is useful for tests, but should be avoided in production code.
func (o Option[T]) Unwrap() T {
	if o.IsNone() {
		panic("unwrap called on `None` value")
	}
	return o.Value()
}

// UnwrapOr returns the value of the option, or the given default value if the option is None.
func (o Option[T]) UnwrapOr(defaultValue T) T {
	if o.IsNone() {
		return defaultValue
	}
	return o.Value()
}

// UnwrapOrElse returns the value of the option, or the result of the given function if the option is None.
func (o Option[T]) UnwrapOrElse(f func() T) T {
	if o.IsNone() {
		return f()
	}
	return o.Value()
}

func (o Option[T]) UnwrapPtr() *T {
	if o.IsNone() {
		return nil
	}
	return o.value
}

// Map returns a new option with the result of the given function applied to the value of the option.
func (o Option[T]) MarshalJSON() ([]byte, error) {
	if o.IsNone() {
		return []byte("null"), nil
	}
	return json.Marshal(o.Value())
}

// Map returns a new option with the result of the given function applied to the value of the option.
func (o *Option[T]) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}

	var value T
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	*o = Some[T](value)

	return nil
}

func (o *Option[T]) String() string {
	return fmt.Sprintf("%v", o.Value())
}

func (o *Option[T]) RawString() string {
	return fmt.Sprintf("Option[%v]", o.Value())
}
