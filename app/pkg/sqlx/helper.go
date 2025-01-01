package sqlx

import "time"

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

// Now returns the current time formatted as a string in the "YYYY-MM-DD HH:MM:SS" format.
//
// The function uses the time.Now() function to get the current time and then formats it using the time.Format() function.
// The format string "2006-01-02 15:04:05" is used to ensure the output matches the specified format.
//
// Example usage:
//
//	currentTime := Now()
//	fmt.Println(currentTime) // Output: 2023-01-01 12:34:56
//
func Now() string {
	now := time.Now()
	return now.Format("2006-01-02 15:04:05")
}
