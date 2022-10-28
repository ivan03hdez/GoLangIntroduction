package pointers

// *int means you *must* pass a *int (pointer to int), NOT just an int!

func PointerFunc(x *int) {
	*x = 2 // Whatever variable caller passed in will now be 2
	y := 7
	x = &y // has no impact on the caller because we overwrote the pointer value!
}
