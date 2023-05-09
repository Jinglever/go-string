package jgstr

// GetStringByPointer get string by pointer
// return "" if pointer is nil
func GetStringByPointer(pointer *string) string {
	if pointer == nil {
		return ""
	}
	return *pointer
}

// CompareStringPointer compare string pointer
// return true if a and b are both nil or a == b
func CompareStringPointer(a, b *string) bool {
	if a == nil && b == nil {
		return true
	}
	if a != nil && b == nil {
		return false
	}
	if a == nil && b != nil {
		return false
	}
	return *a == *b
}
