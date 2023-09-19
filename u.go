package u

// Bool is a `bool` that can be unset
type Bool = Var[bool]

// String is a `string` that can be unset
type String = Var[string]

// Float64 is a `float64` that can be unset
type Float64 = Var[float64]

// Float32 is a `float32` that can be unset
type Float32 = Var[float32]

// Uint is a `uint` that can be unset
type Uint = Var[uint]

// Uint8 is a `uint8` that can be unset
type Uint8 = Var[uint8]

// Uint16 is a `uint16` that can be unset
type Uint16 = Var[uint16]

// Uint32 is a `uint32` that can be unset
type Uint32 = Var[uint32]

// Uint64 is a `uint64` that can be unset
type Uint64 = Var[uint64]

// Byte is a `byte` that can be unset
type Byte = Var[byte]

// Rune is a `rune` that can be unset
type Rune = Var[rune]

// Complex64 is a `complex64` that can be unset
type Complex64 = Var[complex64]

// Complex128 is a `complex128` that can be unset
type Complex128 = Var[complex128]

// Int is an `int` that can be unset
type Int = Var[int]

// Int8 is an `int8` that can be unset
type Int8 = Var[int8]

// Int16 is an `int16` that can be unset
type Int16 = Var[int16]

// Int32 is an `int32` that can be unset
type Int32 = Var[int32]

// Int64 is an `int64` that can be unset
type Int64 = Var[int64]

// Var is a variable that can be set, unset and queried for its state.
type Var[T any] struct {
	val T
	set bool
}

// Get the value of the variable
// Returns the zero value if unset
func (v *Var[T]) Get() T {
	return v.val
}

// Set the value of the variable
func (v *Var[T]) Set(val T) {
	v.val = val
	v.set = true
}

// IsSet returns true when a value has been set
func (v *Var[T]) IsSet() bool {
	return v.set
}

// Unset resets the value to the zero value and sets the variable to "unset"
func (v *Var[T]) Unset() {
	v.set = false
	var temp T
	v.val = temp
}
