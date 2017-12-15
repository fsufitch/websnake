package log

// CustomLogger is a custom logger interface meant to provide a nice
// abstraction over multiple log destinations
type CustomLogger interface {
	Print(...interface{})
	Printf(...interface{})
	GetEnabled() bool
	SetEnabled(bool)
}
