package cast

func ToPointer[T any](value T) *T {
	return &value
}

func ToPointerDeref[T any](ptr *T, value T) T {
	if ptr == nil {
		return value
	}

	return *ptr
}
