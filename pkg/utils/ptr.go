package utils

func BoolPtr(b bool) *bool { return &b }

func IntPtr(i int) *int { return &i }

func StringPtr(s string) *string { return &s }

func Assert(v bool, msg string) {
	if !v {
		panic(msg)
	}
}
