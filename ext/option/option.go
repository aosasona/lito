package option

import "encoding/json"

type Option[T any] struct {
	value *T
}

type (
	OptionalInt    = Option[int]
	OptionalString = Option[string]
	OptionalBool   = Option[bool]
)

func Some[T any](value T) Option[T] {
	return Option[T]{&value}
}

func None[T any]() Option[T] {
	return Option[T]{nil}
}

func (o Option[T]) Value() T {
	return *o.value
}

func (o Option[T]) IsNone() bool {
	return o.value == nil
}

func (o Option[T]) MarshalJSON() ([]byte, error) {
	if o.IsNone() {
		return []byte("null"), nil
	}
	return json.Marshal(o.Value())
}

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
