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

const (
	TypeNone int8 = iota
	TypeSome
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
func (o *Option[T]) rawValue() T {
	return *o.value
}

// Type returns the type of the option as an integer that can be matched against TypeNone and TypeSome.
func (o Option[T]) Type() int8 {
	if o.IsNone() {
		return TypeNone
	}

	return TypeSome
}

// IsNone returns true if the option is None.
func (o *Option[T]) IsNone() bool {
	return o.value == nil
}

// IsSome returns true if the option is not None.
func (o *Option[T]) IsSome() bool {
	return o.value != nil
}

// Unwrap returns the value of the option, or the given default value if the option is None.
// This makes sure you are never dealing with a nil value.
func (o *Option[T]) Unwrap(defaultValue T) T {
	if o.IsNone() {
		return defaultValue
	}
	return o.rawValue()
}

// UnwrapOrElse returns the value of the option, or the result of the given function if the option is None.
func (o *Option[T]) UnwrapOrElse(f func() T) T {
	if o.IsNone() {
		return f()
	}
	return o.rawValue()
}

// DangerouslyUnwrapPtr returns the value of the option as a pointer - of course, it is also aptly named as it is dangerous.
func (o *Option[T]) DangerouslyUnwrapPtr() *T {
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
	return json.Marshal(o.rawValue())
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

func (o *Option[T]) StringWithDefault(alt string) string {
	if o.IsNone() {
		return alt
	}
	return fmt.Sprintf("%v", o.rawValue())
}

func (o *Option[T]) RawString() string {
	return fmt.Sprintf("Option[%v]", o.rawValue())
}
