package ref

func Ref[T any](value T) *T {
	return &value
}

func Deref[T any](value *T, fallback T) T {
	if value == nil {
		return fallback
	}

	return *value
}

func Assert(v bool, msg string) {
	if !v {
		panic(msg)
	}
}
