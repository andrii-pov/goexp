package main

func main() {
	go func() {
		// This is an anonymous deferred call.
		// When it fully exits, the panic 2 will spread
		// to the entry function call of the new
		// goroutine, and replace the panic 0. The
		// panic 2 will never be recovered.
		defer func() {
			// As explained in the last example,
			// panic 2 will replace panic 1.
			defer panic(2)

			// When the anonymous function call fully
			// exits, panic 1 will spread to (and
			// associate with) the nesting anonymous
			// deferred call.
			func() {
				// Once the panic 1 occurs, there will
				// be two unrecovered panics coexisting
				// in the new goroutine. One (panic 0)
				// associates with the entry function
				// call of the new goroutine, the other
				// (panic 1) associates with the
				// current anonymous function call.
				panic(1)
			}()
		}()
		panic(0)
	}()

	select {}
}
