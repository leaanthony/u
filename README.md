# u

u provides a simple way to create variables that are "unset" by default.

u is dependency free and has 100% test coverage.

## Why?

Go's default values are great most of the time, but sometimes you want to know if a value has been set or not. This is especially true when you want to know if a value has been set to its zero value or not.

For example, let's say you had a preferences struct like this:

```go
type Preferences struct {
    UseFeatureX bool
    Threshold int
}

var prefs Preferences

if prefs.UseFeatureX {
    // Uh oh...was this explicitly set?
}
```

There is no real way to know if `UseFeatureX` or `Threshold` have been set or not.

Using this library we can do this:

```go
type Preferences struct {
    UseFeatureX u.Bool
    Threshold u.Int
}

var prefs Preferences

if prefs.UseFeatureX.IsSet() {
    // Do something
}

if prefs.Threshold.IsSet() {
    // Do something
}
```

## Why should any of this matter? 

It matters when you are working with third party libraries or frameworks that have default values that may not be the same as the Zero Go values. 
Perhaps the default value for a preference is `true` and you only want to set it if an explicit value has been specified.

## Usage

### Basic Usage

```go
var myVar u.Int

// Set the value
myVar.Set(10)

// Get the value
fmt.Println(myVar.Get()) // 10

// Check if the value has been set
fmt.Println(myVar.IsSet()) // true

// Unset the value
myVar.Unset()

// Check if the value has been set
fmt.Println(myVar.IsSet()) // false
```

## Supported types

- `u.Bool`
- `u.Int`
- `u.Int8`
- `u.Int16`
- `u.Int32`
- `u.Int64`
- `u.Uint`
- `u.Uint8`
- `u.Uint16`
- `u.Uint32`
- `u.Uint64`
- `u.Float32`
- `u.Float64`
- `u.Complex64`
- `u.Complex128`
- `u.String`
- `u.Byte`
- `u.Rune`

## Custom types

Any type can be used with this library.

```go

type MyCustomType struct {
    value string
}

var myVar u.Var[MyCustomType]

// Set the value
myVar.Set(MyCustomType{value: "hello"})
```

## License

MIT