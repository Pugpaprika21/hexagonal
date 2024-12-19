package sqlx

// Nil is a generic function that checks if a pointer to a comparable type is nil or points to its zero value.
// If the input pointer is nil or points to its zero value, Nil returns nil. Otherwise, it returns the input pointer unchanged.
//
// The function supports any comparable type T, including basic types (int, float64, string, etc.) and custom types that implement the comparable interface.
//
// Example usage:
//
//	var i *int
//	fmt.Println(Nil(i)) // Output: <nil>
//
//	var j *int = new(int)
//	fmt.Println(Nil(j)) // Output: 0x... (memory address of j)
//
//	var k *string = new(string)
//	fmt.Println(Nil(k)) // Output: 0x... (memory address of k)
//
func Nil[T comparable](value *T) *T {
	if value == nil {
		return nil
	}

	var zero T
	if *value == zero {
		return nil
	}
	return value
}
