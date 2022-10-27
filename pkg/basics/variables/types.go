package variables

var number int
var unsignedNumber uint
var floatingNumber32 float32
var floatingNumber64 float64
var complexNumber64 complex64
var complexNumber128 complex128

var boolean bool
var cadena string

var err error

var mapOfStrings map[string]string
var mapOfAnything map[interface{}]interface{}

var slice []string  // slice --> size dynamically grows
var array [5]string // array --> size already known

type Object struct {
	property1 string
	property2 map[string]string
}

type CallbackSyntax func(string) string

var callback CallbackSyntax = func(name string) string {
	return name
}
