package main

func main() {
	defer func() {
		print("WOR! Panic is still not recovered")
		recover() // ONLY THIS take effect
	}()

	defer func() {
		defer func() {
			recover() // no-op
		}()
	}()
	defer func() {
		func() {
			recover() // no-op
		}()
	}()
	func() {
		defer func() {
			recover() // no-op
		}()
	}()
	func() {
		defer recover() // no-op
	}()
	func() {
		recover() // no-op
	}()
	recover()       // no-op
	defer recover() // no-op
	panic("bye")
}
