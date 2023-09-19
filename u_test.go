package u_test

import (
	"testing"
	"u"
)

func TestVar_String(t *testing.T) {
	var v u.String
	if v.Get() != "" {
		t.Errorf("Get() = %Var, want \"\"", v.Get())
	}

	v.Set("hello")
	if v.Get() != "hello" {
		t.Errorf("Get() = %Var, want \"hello\"", v.Get())
	}
	if v.IsSet() != true {
		t.Errorf("IsSet() = %Var, want true", v.IsSet())
	}

	v.Unset()
	if v.Get() != "" {
		t.Errorf("Get() = %Var, want \"\"", v.Get())
	}
	if v.IsSet() != false {
		t.Errorf("IsSet() = %Var, want false", v.IsSet())
	}
}

func TestVar_Bool(t *testing.T) {

	var v u.Bool
	if v.Get() != false {
		t.Errorf("Get() = %Var, want false", v.Get())
	}

	v.Set(true)
	if v.Get() != true {
		t.Errorf("Get() = %Var, want true", v.Get())
	}

	if v.IsSet() != true {
		t.Errorf("IsSet() = %Var, want true", v.IsSet())
	}

	v.Unset()
	if v.Get() != false {
		t.Errorf("Get() = %Var, want false", v.Get())
	}
	if v.IsSet() != false {
		t.Errorf("IsSet() = %Var, want false", v.IsSet())
	}
}

func TestVar_Float64(t *testing.T) {

	var v u.Float64
	if v.Get() != 0.0 {
		t.Errorf("Get() = %Var, want 0.0", v.Get())
	}

	v.Set(42.0)
	if v.Get() != 42.0 {
		t.Errorf("Get() = %Var, want 42.0", v.Get())
	}
	if v.IsSet() != true {
		t.Errorf("IsSet() = %Var, want true", v.IsSet())
	}

	v.Unset()
	if v.Get() != 0.0 {
		t.Errorf("Get() = %Var, want 0.0", v.Get())
	}
	if v.IsSet() != false {
		t.Errorf("IsSet() = %Var, want false", v.IsSet())
	}
}

func TestVar_Float32(t *testing.T) {

	var v u.Float32
	if v.Get() != 0.0 {
		t.Errorf("Get() = %Var, want 0.0", v.Get())
	}

	v.Set(42.0)
	if v.Get() != 42.0 {
		t.Errorf("Get() = %Var, want 42.0", v.Get())
	}
	if v.IsSet() != true {
		t.Errorf("IsSet() = %Var, want true", v.IsSet())
	}

	v.Unset()
	if v.Get() != 0.0 {
		t.Errorf("Get() = %Var, want 0.0", v.Get())
	}
	if v.IsSet() != false {
		t.Errorf("IsSet() = %Var, want false", v.IsSet())
	}
}

func TestVar_Byte(t *testing.T) {

	var v u.Byte
	if v.Get() != 0 {
		t.Errorf("Get() = %Var, want 0", v.Get())
	}

	v.Set(42)
	if v.Get() != 42 {
		t.Errorf("Get() = %Var, want 42", v.Get())
	}
	if v.IsSet() != true {
		t.Errorf("IsSet() = %Var, want true", v.IsSet())
	}

	v.Unset()
	if v.Get() != 0 {
		t.Errorf("Get() = %Var, want 0", v.Get())
	}
	if v.IsSet() != false {
		t.Errorf("IsSet() = %Var, want false", v.IsSet())
	}
}

func TestVar_Rune(t *testing.T) {

	var v u.Rune
	if v.Get() != 0 {
		t.Errorf("Get() = %Var, want 0", v.Get())
	}

	v.Set('x')
	if v.Get() != 'x' {
		t.Errorf("Get() = %Var, want 'x'", v.Get())
	}
	if v.IsSet() != true {
		t.Errorf("IsSet() = %Var, want true", v.IsSet())
	}

	v.Unset()
	if v.Get() != 0 {
		t.Errorf("Get() = %Var, want 0", v.Get())
	}
	if v.IsSet() != false {
		t.Errorf("IsSet() = %Var, want false", v.IsSet())
	}
}

func TestVar_Complex64(t *testing.T) {

	var v u.Complex64
	if v.Get() != 0+0i {
		t.Errorf("Get() = %Var, want 0+0i", v.Get())
	}

	v.Set(42 + 42i)
	if v.Get() != 42+42i {
		t.Errorf("Get() = %Var, want 42+42i", v.Get())
	}
	if v.IsSet() != true {
		t.Errorf("IsSet() = %Var, want true", v.IsSet())
	}

	v.Unset()
	if v.Get() != 0+0i {
		t.Errorf("Get() = %Var, want 0+0i", v.Get())
	}
	if v.IsSet() != false {
		t.Errorf("IsSet() = %Var, want false", v.IsSet())
	}
}

func TestVar_Complex128(t *testing.T) {

	var v u.Complex128
	if v.Get() != 0+0i {
		t.Errorf("Get() = %Var, want 0+0i", v.Get())
	}

	v.Set(42 + 42i)
	if v.Get() != 42+42i {
		t.Errorf("Get() = %Var, want 42+42i", v.Get())
	}
	if v.IsSet() != true {
		t.Errorf("IsSet() = %Var, want true", v.IsSet())
	}

	v.Unset()
	if v.Get() != 0+0i {
		t.Errorf("Get() = %Var, want 0+0i", v.Get())
	}
	if v.IsSet() != false {
		t.Errorf("IsSet() = %Var, want false", v.IsSet())
	}
}

func TestVar_Int(t *testing.T) {

	var v u.Int
	if v.Get() != 0 {
		t.Errorf("Get() = %Var, want 0", v.Get())
	}

	v.Set(42)
	if v.Get() != 42 {
		t.Errorf("Get() = %Var, want 42", v.Get())
	}
	if v.IsSet() != true {
		t.Errorf("IsSet() = %Var, want true", v.IsSet())
	}

	v.Unset()
	if v.Get() != 0 {
		t.Errorf("Get() = %Var, want 0", v.Get())
	}
	if v.IsSet() != false {
		t.Errorf("IsSet() = %Var, want false", v.IsSet())
	}
}

func TestVar_Int8(t *testing.T) {

	var v u.Int8
	if v.Get() != 0 {
		t.Errorf("Get() = %Var, want 0", v.Get())
	}

	v.Set(42)
	if v.Get() != 42 {
		t.Errorf("Get() = %Var, want 42", v.Get())
	}
	if v.IsSet() != true {
		t.Errorf("IsSet() = %Var, want true", v.IsSet())
	}

	v.Unset()
	if v.Get() != 0 {
		t.Errorf("Get() = %Var, want 0", v.Get())
	}
	if v.IsSet() != false {
		t.Errorf("IsSet() = %Var, want false", v.IsSet())
	}
}

func TestVar_Int16(t *testing.T) {

	var v u.Int16
	if v.Get() != 0 {
		t.Errorf("Get() = %Var, want 0", v.Get())
	}

	v.Set(42)
	if v.Get() != 42 {
		t.Errorf("Get() = %Var, want 42", v.Get())
	}
	if v.IsSet() != true {
		t.Errorf("IsSet() = %Var, want true", v.IsSet())
	}

	v.Unset()
	if v.Get() != 0 {
		t.Errorf("Get() = %Var, want 0", v.Get())
	}
	if v.IsSet() != false {
		t.Errorf("IsSet() = %Var, want false", v.IsSet())
	}
}

func TestVar_Int32(t *testing.T) {

	var v u.Int32
	if v.Get() != 0 {
		t.Errorf("Get() = %Var, want 0", v.Get())
	}

	v.Set(42)
	if v.Get() != 42 {
		t.Errorf("Get() = %Var, want 42", v.Get())
	}
	if v.IsSet() != true {
		t.Errorf("IsSet() = %Var, want true", v.IsSet())
	}

	v.Unset()
	if v.Get() != 0 {
		t.Errorf("Get() = %Var, want 0", v.Get())
	}
	if v.IsSet() != false {
		t.Errorf("IsSet() = %Var, want false", v.IsSet())
	}
}

func TestVar_Int64(t *testing.T) {

	var v u.Int64
	if v.Get() != 0 {
		t.Errorf("Get() = %Var, want 0", v.Get())
	}

	v.Set(42)
	if v.Get() != 42 {
		t.Errorf("Get() = %Var, want 42", v.Get())
	}
	if v.IsSet() != true {
		t.Errorf("IsSet() = %Var, want true", v.IsSet())
	}

	v.Unset()
	if v.Get() != 0 {
		t.Errorf("Get() = %Var, want 0", v.Get())
	}
	if v.IsSet() != false {
		t.Errorf("IsSet() = %Var, want false", v.IsSet())
	}
}

func TestVar_Uint(t *testing.T) {

	var v u.Uint
	if v.Get() != 0 {
		t.Errorf("Get() = %Var, want 0", v.Get())
	}

	v.Set(42)
	if v.Get() != 42 {
		t.Errorf("Get() = %Var, want 42", v.Get())
	}
	if v.IsSet() != true {
		t.Errorf("IsSet() = %Var, want true", v.IsSet())
	}

	v.Unset()
	if v.Get() != 0 {
		t.Errorf("Get() = %Var, want 0", v.Get())
	}
	if v.IsSet() != false {
		t.Errorf("IsSet() = %Var, want false", v.IsSet())
	}
}

func TestVar_Uint8(t *testing.T) {

	var v u.Uint8
	if v.Get() != 0 {
		t.Errorf("Get() = %Var, want 0", v.Get())
	}

	v.Set(42)
	if v.Get() != 42 {
		t.Errorf("Get() = %Var, want 42", v.Get())
	}
	if v.IsSet() != true {
		t.Errorf("IsSet() = %Var, want true", v.IsSet())
	}

	v.Unset()
	if v.Get() != 0 {
		t.Errorf("Get() = %Var, want 0", v.Get())
	}
	if v.IsSet() != false {
		t.Errorf("IsSet() = %Var, want false", v.IsSet())
	}
}

func TestVar_Uint16(t *testing.T) {

	var v u.Uint16
	if v.Get() != 0 {
		t.Errorf("Get() = %Var, want 0", v.Get())
	}

	v.Set(42)
	if v.Get() != 42 {
		t.Errorf("Get() = %Var, want 42", v.Get())
	}
	if v.IsSet() != true {
		t.Errorf("IsSet() = %Var, want true", v.IsSet())
	}

	v.Unset()
	if v.Get() != 0 {
		t.Errorf("Get() = %Var, want 0", v.Get())
	}
	if v.IsSet() != false {
		t.Errorf("IsSet() = %Var, want false", v.IsSet())
	}
}

func TestVar_Uint32(t *testing.T) {

	var v u.Uint32
	if v.Get() != 0 {
		t.Errorf("Get() = %Var, want 0", v.Get())
	}

	v.Set(42)
	if v.Get() != 42 {
		t.Errorf("Get() = %Var, want 42", v.Get())
	}
	if v.IsSet() != true {
		t.Errorf("IsSet() = %Var, want true", v.IsSet())
	}

	v.Unset()
	if v.Get() != 0 {
		t.Errorf("Get() = %Var, want 0", v.Get())
	}
	if v.IsSet() != false {
		t.Errorf("IsSet() = %Var, want false", v.IsSet())
	}
}

func TestVar_Uint64(t *testing.T) {

	var v u.Uint64
	if v.Get() != 0 {
		t.Errorf("Get() = %Var, want 0", v.Get())
	}

	v.Set(42)
	if v.Get() != 42 {
		t.Errorf("Get() = %Var, want 42", v.Get())
	}
	if v.IsSet() != true {
		t.Errorf("IsSet() = %Var, want true", v.IsSet())
	}

	v.Unset()
	if v.Get() != 0 {
		t.Errorf("Get() = %Var, want 0", v.Get())
	}
	if v.IsSet() != false {
		t.Errorf("IsSet() = %Var, want false", v.IsSet())
	}
}
